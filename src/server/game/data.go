package game

type RowStruct struct {
	Player   int
	IsPlaced bool
}

type WinStruct struct {
	Winner int
	IsWin  bool
}

type LeaderboardScores struct {
	Player   int
	IsWinner bool
}

type GameStruct struct {
	Title        string
	PlayerToPlay int
	Leaderboard  []LeaderboardScores
	Win          WinStruct
	Rows         [][]RowStruct
	IsLineFull   []bool
}

var Data = GameStruct{
	Title:        "Game",
	PlayerToPlay: PlayerToPlay,
	Leaderboard:  []LeaderboardScores{},
	Win: WinStruct{
		Winner: 0,
		IsWin:  false,
	},
	Rows:       make([][]RowStruct, 6),
	IsLineFull: make([]bool, 7),
}
var IsGameStarted = false
var PlayerToPlay = 1
