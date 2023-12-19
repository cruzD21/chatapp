package server

import (
	"chatapp/internal/handlers"
	"chatapp/pkg/chat"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")
var Hub = chat.NewHub()

func Run() error {
	flag.Parse()
	r := mux.NewRouter()
	//hub := chat.NewHub()
	go Hub.Run()
	r.HandleFunc("/", handlers.ServeHome)
	r.HandleFunc("/room/create", handlers.RoomCreate)
	r.HandleFunc("/room/{room_id}", handlers.Room)
	r.HandleFunc("/room/{room_id}/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(Hub, w, r)
	})

	server := &http.Server{
		Addr:              *addr,
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe", err)
		return err
	}
	return nil
}
