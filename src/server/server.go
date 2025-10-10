package server

import (
	"fmt"
	"log"
	"net/http"
	"power4web/src/server/handlers"
)

func StartHttpServer() {
	// File server
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("src/public"))))

	// API
	http.HandleFunc("/placeLine", handlers.GameLinePlayHandler)
	http.HandleFunc("/newGame", handlers.GameNewPartyHandler)
	http.HandleFunc("/changeSlot", handlers.ChangeSlotHandler)
	http.HandleFunc("/selectPlayer", handlers.SelectPlayerHandler)

	// Pages
	http.HandleFunc("/room/", handlers.GameRoomHandler)
	http.HandleFunc("/play", handlers.GameHandler)
	http.HandleFunc("/", handlers.LandingPageHandler)

	fmt.Println("HTTP Server is listening on 127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
