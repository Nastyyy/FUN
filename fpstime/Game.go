package main

import (
	"fmt"
	"time"
)

type Game struct {
	Players  []*Player
	tickData []tickData
}

func (game *Game) tick(epoch int) {
	fmt.Println("Running tick: ", epoch)
	for _, player := range game.Players {
		fmt.Println(player)
	}
	time.Sleep(time.Millisecond * 60)
}

func NewGame(players ...*Player) *Game {
	game := Game{Players: players}

	return &game
}
