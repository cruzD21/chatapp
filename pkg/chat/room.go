package chat

import (
	"log"
	"sync"
)

type Room struct {
	ID         string
	Hub        *Hub
	Members    map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
}

var (
	Rooms     map[string]*Room
	RoomsLock sync.RWMutex
)

func (r *Room) initRoom() {
	for {
		select {
		case client := <-r.Register:
			r.Members[client] = true
			log.Printf("client %v joined room %s", client, r.ID)

		case client := <-r.Unregister:
			if _, ok := r.Members[client]; ok {
				delete(r.Members, client)
				if len(r.Members) == 0 {
					r.Hub.Unregister <- r
				}
			}
		case message := <-r.Broadcast:
			for client := range r.Members {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(r.Members, client)
				}
			}

		}
	}
}

func CreateOrGetRoom(uuid string) (string, *Room) {

	if room := Rooms[uuid]; room != nil {
		return uuid, room
	}

	RoomsLock.Lock()
	defer RoomsLock.Unlock()

	//room doesn't exist
	newRoom := &Room{
		ID:         uuid,
		Hub:        Hubs,
		Members:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		Unregister: make(chan *Client),
		Register:   make(chan *Client),
	}
	Hubs.Register <- newRoom
	go newRoom.initRoom()

	return uuid, newRoom
}
