package game

type RowStruct struct {
	Player   int
	IsPlaced bool
}

type GameStruct struct {
	Title        string
	PlayerToPlay int
	Rows         [][]RowStruct
	IsLineFull   []bool
}

var Data = GameStruct{
	Title:        "Game",
	PlayerToPlay: PlayerToPlay,
	Rows:         make([][]RowStruct, 6),
	IsLineFull:   make([]bool, 7),
}
var IsGameStarted = false
var PlayerToPlay = 1
