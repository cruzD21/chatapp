package chat

import "log"

var Hubs = NewHub()

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Room
	Unregister chan *Room
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Room),
		Unregister: make(chan *Room),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case room := <-h.Register:
			h.Rooms[room.ID] = room
			log.Printf("registered room %s to the hub", room.ID)
		case room := <-h.Unregister:
			if _, ok := h.Rooms[room.ID]; ok {
				log.Printf("unregistering room %s from hub", room.ID)
				delete(h.Rooms, room.ID)
			}
		}
	}
}
