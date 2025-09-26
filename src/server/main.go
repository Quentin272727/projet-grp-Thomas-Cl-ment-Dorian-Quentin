package main

import (
	"log"
	"net/http"
	"power4web/src/server/handlers"
)

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("src/public"))))
	http.HandleFunc("/play", handlers.GameHandler)
	http.HandleFunc("/placeLine", handlers.GameLinePlayHandler)
	http.HandleFunc("/", handlers.LandingPageHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
