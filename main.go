package main

import (
	"my-game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

// Main function to run the game
func main() {
	g := game.NewGame()
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
