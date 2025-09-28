package handlers

import (
	"net/http"
	"power4web/src/server/data"
	"power4web/src/server/landing"
	"strconv"
)

func LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	landing.LoadLanding(w, r)
}

func ChangeSlotHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	selectedSlotColor := r.FormValue("change_slot")

	if selectedSlotColor == "" {
		http.Error(w, "Data sent is malformed", http.StatusMethodNotAllowed)
		return
	}

	if data.IsValidColor(selectedSlotColor) {
		if !data.IsColorTaken(selectedSlotColor, data.ServerData.PlayerSelectedIndex+1) {
			data.ServerData.Players[data.ServerData.PlayerSelectedIndex].Slot = selectedSlotColor

			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Error(w, "This color is already taken by another player", http.StatusMethodNotAllowed)
			return
		}
	} else {
		http.Error(w, "Invalid color", http.StatusMethodNotAllowed)
		return
	}
}

func SelectPlayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	player := r.FormValue("select-player_btn")
	playerC, err := strconv.Atoi(player)
	if err != nil {
		http.Error(w, "Data sent is malformed", http.StatusMethodNotAllowed)
		return
	}

	if playerC <= 0 || playerC > data.NumOfPlayers {
		http.Error(w, "Data sent is malformed", http.StatusMethodNotAllowed)
		return
	}

	data.ServerData.PlayerSelectedIndex = playerC - 1

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
