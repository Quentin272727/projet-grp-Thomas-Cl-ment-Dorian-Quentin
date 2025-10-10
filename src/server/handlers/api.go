package handlers

import (
	"net/http"
	"power4web/src/server/data"
	game "power4web/src/server/game"
	"strconv"
)

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

	game.PlaceCoinLine(lineInt)

	http.Redirect(w, r, "/play", http.StatusSeeOther)
}

func GameNewPartyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	btn := r.FormValue("newgame_btn")

	if btn == "" {
		game.NewParty()
		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	} else {
		http.Error(w, "Data sent is malformed", http.StatusMethodNotAllowed)
		return
	}
}
