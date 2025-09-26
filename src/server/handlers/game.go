package handlers

import (
	"net/http"
	game "power4web/src/server/game"
	"strconv"
)

func GameHandler(w http.ResponseWriter, r *http.Request) {
	game.Start(w, r)
}

func GameLinePlayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	line := r.FormValue("line_btn")

	if line == "" {
		http.Error(w, "Data sent is empty or edited", http.StatusMethodNotAllowed)
		return
	}

	lineInt, err := strconv.Atoi(line)
	if err != nil {
		http.Error(w, "Data sent is malformed", http.StatusMethodNotAllowed)
		return
	}

	game.PlaceCoinLine(w, r, lineInt)

	http.Redirect(w, r, "/play", http.StatusSeeOther)
}
