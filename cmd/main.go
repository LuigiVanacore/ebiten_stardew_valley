package main

import (
	"log"

	"github.com/LuigiVanacore/ebiten_stardew_valley"
	"github.com/hajimehoshi/ebiten/v2"
)




func main() {
	ebiten.SetWindowSize(ebiten_stardew_valley.SCREEN_WIDTH, ebiten_stardew_valley.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Sprite Example")

	
	
	game := ebiten_stardew_valley.NewGame(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}