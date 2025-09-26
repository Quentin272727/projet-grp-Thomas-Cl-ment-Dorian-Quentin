package game

import (
	"html/template"
	"net/http"
)

func Start(w http.ResponseWriter, r *http.Request) {
	Data.PlayerToPlay = PlayerToPlay
	if IsGameStarted {
		LoadPage(w, Data)
	} else {
		for i := 1; i <= 6; i++ {
			dataToAdd := []RowStruct{}
			for y := 1; y <= 7; y++ {
				dataToAdd = append(dataToAdd, RowStruct{Player: 0, IsPlaced: false})
			}

			Data.Rows[i-1] = append(Data.Rows[i-1], dataToAdd...)
		}

		LoadPage(w, Data)
		IsGameStarted = true
	}
}

func LoadPage(w http.ResponseWriter, data GameStruct) {
	tmpl, err := template.ParseFiles("src/client/game/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, Data)
}

func PlaceCoinLine(w http.ResponseWriter, r *http.Request, line int) {
	if !appendCoinInsideRow(line) {
		switch PlayerToPlay {
		case 1:
			PlayerToPlay = 2
			return
		case 2:
			PlayerToPlay = 1
			return
		}
	}
}

func appendCoinInsideRow(l int) bool {
	// This function return true if the entire line is full and no coin can slide into it
	for i := 6; i >= 1; i-- {
		if !Data.Rows[i-1][l-1].IsPlaced {
			Data.Rows[i-1][l-1] = RowStruct{
				Player:   PlayerToPlay,
				IsPlaced: true,
			}

			if i == 1 {
				Data.IsLineFull[l-1] = true
			}
			return false
		}
	}

	return true
}
