package handlers

import (
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
	http.Redirect(w, r, url, http.StatusMovedPermanently)
	return
}

// TODO: fix this name
func Roomf(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]

	p := RoomPage{RoomID: roomID}
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
