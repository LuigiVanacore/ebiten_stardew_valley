package main

import (
	"log"

	"github.com/LuigiVanacore/ebiten_game_example"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)


func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sprite Example")

	
	
	game := ebiten_game_example.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}