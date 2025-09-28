package data

type PlayerData struct {
	Slot string
}

type LeaderboardScores struct {
	Player   int
	IsWinner bool
}

type WinStruct struct {
	Winner int
	IsWin  bool
	IsDraw bool
}

type RowStruct struct {
	Player   int
	IsPlaced bool
}

type ServerStruct struct {
	Title               string
	PlayerSelectedIndex int
	Players             []PlayerData
	PlayerToPlay        int
	Leaderboard         []LeaderboardScores
	Win                 WinStruct
	Rows                [][]RowStruct
	IsLineFull          []bool
	AvailableSlotColors []string
}

var ServerData = ServerStruct{
	Title:               "Power4Web",
	PlayerToPlay:        PlayerToPlay,
	PlayerSelectedIndex: 0,
	Players: []PlayerData{
		{
			Slot: "red",
		},
		{
			Slot: "yellow",
		},
	},
	Leaderboard: []LeaderboardScores{},
	Win: WinStruct{
		Winner: 0,
		IsWin:  false,
		IsDraw: false,
	},
	Rows:       make([][]RowStruct, 6),
	IsLineFull: make([]bool, 7),
	AvailableSlotColors: []string{
		"red", "yellow", "green", "blue", "pink", "purple",
	},
}
var IsGameStarted = false
var PlayerToPlay = 1
var NumOfPlayers = 2

func IsValidColor(color string) bool {
	for _, c := range ServerData.AvailableSlotColors {
		if color == c {
			return true
		}
	}

	return false
}

func IsColorTaken(color string, playerNumber int) bool {
	for i, p := range ServerData.Players {
		if p.Slot == color && i != playerNumber-1 {
			return true
		}
	}

	return false
}
