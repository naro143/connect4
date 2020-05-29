package models

type Player struct {
	Number  int
	Markers int
}

func NewPlayer(number int) *Player {
	return &Player{
		Number:  number,
		Markers: 21,
	}
}

// コマを置く（座標を指定）
func (p *Player) Play() int {
	p.Markers -= 1
	// input column
	column := 2
	return column
}
