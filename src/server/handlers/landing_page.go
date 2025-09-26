package handlers

import (
	"log"
	"net/http"
	"os"
)

func LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("src/client/index.html")
	if err != nil {
		log.Fatal(err)
	}

	w.Write(data)
}
