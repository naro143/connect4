package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_Turn(t *testing.T) {
	g := NewGame()

	// empty board
	assert.Equal(t, [6][7]int{}, g.Board.Board, "empty board")

	g.Turn()

	assert.Equal(t, 1, g.Board.Board[0][2], "drop marker")
}
