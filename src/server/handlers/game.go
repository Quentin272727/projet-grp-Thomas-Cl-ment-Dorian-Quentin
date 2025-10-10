package handlers

import (
	"net/http"
	game "power4web/src/server/game"
	"strings"
)

// Pages
func GameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	game.Start(w, r)
}

func GameRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var roomId string = strings.Split(r.URL.Path, "/room/")[1]

	if roomId == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Write([]byte(roomId))
}
