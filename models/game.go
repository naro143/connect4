package models

type Game struct {
	Board      *Board
	Player1    *Player
	Player2    *Player
	nextPlayer *Player
}

func NewGame() *Game {
	p1 := NewPlayer(1)
	return &Game{
		Board:      NewBoard(),
		Player1:    p1,
		Player2:    NewPlayer(2),
		nextPlayer: p1,
	}
}

// 勝利判定

// ターンを回す
func (g *Game) Turn() {
	column := g.nextPlayer.Play()
	g.Board.DropMarker(column, g.nextPlayer.Number)
}
