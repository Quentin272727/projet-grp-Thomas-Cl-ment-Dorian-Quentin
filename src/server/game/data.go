package game

type RowStruct struct {
	Player   int
	IsPlaced bool
}

type GameStruct struct {
	Title        string
	PlayerStatus string
	Rows         [][]RowStruct
}

var Data = GameStruct{
	Title:        "Game",
	PlayerStatus: "",
	Rows:         make([][]RowStruct, 6),
}
var IsGameStarted = false
var PlayerToPlay = 1
