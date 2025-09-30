package game

import (
	"html/template"
	"net/http"
	"power4web/src/server/data"
)

type checkForAWinnerStruct struct {
	Player        int
	IsThereWinner bool
	IsDraw        bool
}

func Start(w http.ResponseWriter, r *http.Request) {
	if data.IsGameStarted {
		LoadPage(w, data.ServerData)
	} else {
		loadRows()
		LoadPage(w, data.ServerData)
		data.IsGameStarted = true
	}
}

func NewParty() {
	if !data.IsGameStarted {
		return
	}

	// Reset variables that needs to be reset
	data.ServerData.Rows = make([][]data.RowStruct, 6)
	data.ServerData.IsLineFull = make([]bool, 7)
	data.ServerData.Win = data.WinStruct{
		Winner: 0,
		IsWin:  false,
		IsDraw: false,
	}

	loadRows()
}

func Reset() {
	// -----
}

func LoadPage(w http.ResponseWriter, d data.ServerStruct) {
	funcs := template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
	}

	tmpl, err := template.New("index.html").Funcs(funcs).ParseFiles("src/client/game/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "index.html", data.ServerData)
}

func PlaceCoinLine(line int) {
	if data.ServerData.Win.IsWin || data.ServerData.Win.IsDraw {
		return
	}

	if !appendCoinInsideRow(line) {
		win := checkForAWinner()

		if win.IsDraw {
			data.ServerData.Win.IsDraw = true
			data.ServerData.Leaderboard = append(data.ServerData.Leaderboard, data.LeaderboardScores{Player: 0, IsWinner: false})

			return
		}

		if win.IsThereWinner {
			data.ServerData.Win.IsWin = true
			data.ServerData.Win.Winner = win.Player

			data.ServerData.Leaderboard = append(data.ServerData.Leaderboard, data.LeaderboardScores{Player: win.Player, IsWinner: true})
			return
		}

		switch data.ServerData.PlayerToPlay {
		case 1:
			data.ServerData.PlayerToPlay = 2
			return
		case 2:
			data.ServerData.PlayerToPlay = 1
			return
		}
	}
}

func loadRows() {
	for i := 1; i <= 6; i++ {
		rowToAdd := []data.RowStruct{}
		for y := 1; y <= 7; y++ {
			rowToAdd = append(rowToAdd, data.RowStruct{Player: 0, IsPlaced: false})
		}
		data.ServerData.Rows[i-1] = append(data.ServerData.Rows[i-1], rowToAdd...)
	}
}

func appendCoinInsideRow(l int) bool {
	if data.ServerData.Win.IsWin || data.ServerData.Win.IsDraw {
		return true
	}

	// This function return true if the entire line is full and no coin can slide into it
	for i := 6; i >= 1; i-- {
		if !data.ServerData.Rows[i-1][l-1].IsPlaced {
			data.ServerData.Rows[i-1][l-1] = data.RowStruct{
				Player:   data.ServerData.PlayerToPlay,
				IsPlaced: true,
			}

			if i == 1 {
				data.ServerData.IsLineFull[l-1] = true
			}
			return false
		}
	}

	return true
}

func checkForAWinner() checkForAWinnerStruct {
	// Check if the match is draw
	cl := 0
	for _, l := range data.ServerData.IsLineFull {
		if l {
			cl += 1
		}

		if cl == 7 {
			return checkForAWinnerStruct{
				Player:        0,
				IsThereWinner: false,
				IsDraw:        true,
			}
		}
	}

	ps := make([]int, data.NumOfPlayers)
	lastPlayerChecked := 0

	// Check for row winner
	for _, r := range data.ServerData.Rows {
		for _, s := range r {
			if lastPlayerChecked == 0 {
				lastPlayerChecked = data.ServerData.PlayerToPlay
			}

			if !s.IsPlaced {
				ps[lastPlayerChecked-1] = 0
			} else {
				if s.IsPlaced && s.Player == lastPlayerChecked {
					ps[s.Player-1] += 1
				} else {
					ps[lastPlayerChecked-1] = 0
					ps[s.Player-1] += 1
				}

				if ps[s.Player-1] >= 4 {
					return checkForAWinnerStruct{
						Player:        s.Player,
						IsThereWinner: true,
						IsDraw:        false,
					}
				}

				lastPlayerChecked = s.Player
			}
		}
		ps = make([]int, data.NumOfPlayers)
	}

	// Check for line winner
	ps = make([]int, data.NumOfPlayers)
	lastPlayerChecked = 0
	for i := 1; i <= 7; i++ {
		for _, r := range data.ServerData.Rows {
			if lastPlayerChecked == 0 {
				lastPlayerChecked = data.ServerData.PlayerToPlay
			}

			if !r[i-1].IsPlaced {
				ps[lastPlayerChecked-1] = 0
			} else {
				if r[i-1].IsPlaced && r[i-1].Player == lastPlayerChecked {
					ps[r[i-1].Player-1] += 1
				} else {
					ps[lastPlayerChecked-1] = 0
					ps[r[i-1].Player-1] += 1
				}

				if ps[r[i-1].Player-1] >= 4 {
					return checkForAWinnerStruct{
						Player:        r[i-1].Player,
						IsThereWinner: true,
						IsDraw:        false,
					}
				}

				lastPlayerChecked = r[i-1].Player
			}
		}
		ps = make([]int, data.NumOfPlayers)
	}

	return checkForAWinnerStruct{
		Player:        0,
		IsThereWinner: false,
		IsDraw:        false,
	}
}
