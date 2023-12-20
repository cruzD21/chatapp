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

func Run() error {
	flag.Parse()
	r := mux.NewRouter()
	//hub := chat.NewHub()
	go chat.Hubs.Run()
	r.Use(noCacheMiddleware)
	r.HandleFunc("/", handlers.ServeHome)
	r.HandleFunc("/room/create", handlers.RoomCreate)
	r.HandleFunc("/room/{room_id}", handlers.Room)
	r.HandleFunc("/room/{room_id}/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(chat.Hubs, w, r)
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

func noCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers to prevent caching
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
