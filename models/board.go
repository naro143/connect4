package models

import (
	"errors"
)

type Board struct {
	Board [6][7]int
}

func NewBoard() *Board {
	return &Board{}
}

const (
	Empty = 0
)

// 終了判定

// コマの配置
func (b *Board) DropMarker(column, playerNumber int) error {
	for index, row := range b.Board {
		if row[column] == Empty {
			row[column] = playerNumber
			b.Board[index] = row
			return nil
		}
	}
	return errors.New("overflow")
}
