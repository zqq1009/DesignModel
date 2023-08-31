package main

import (
	"fmt"
	"testing"
)

func TestNewPokerGame(t *testing.T) {
	game1 := NewPokerGame()
	game2 := NewPokerGame()
	fmt.Println(game1.Cards[1] == game2.Cards[1])
}
