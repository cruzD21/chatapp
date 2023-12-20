package handlers

import (
	chat "chatapp/pkg/chat"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type RoomPage struct {
	RoomID string
}

func RoomCreate(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("/room/%s", guuid.New().String())
	http.Redirect(w, r, url, http.StatusFound)
	return
}

func Room(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]
	uuid, _ := chat.CreateOrGetRoom(roomID)

	p := RoomPage{RoomID: uuid}

	t, err := template.ParseFiles("static/room.html")
	if err != nil {
		log.Println("error parsing file maybe wrong path", err)
	}

	if err := t.Execute(w, p); err != nil {
		log.Println("error executing template", err)
	}

	//log.Printf("room id: %s", roomID)
	return
}
