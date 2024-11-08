package ebiten_stardew_valley

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_stardew_valley/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func NewGame(debugFlag bool) *Game {
	ebiten_extended.GameManager().SetIsDebug(debugFlag)
	ebiten_extended.ResourceManager().AddImage(PLAYER_SPRITE, resources.Character_Down_0)
	LoadAnimationSets()
	player := NewPlayer(math2D.NewVector2D(0, 0))
	gameLayer := ebiten_extended.NewLayer(2, 2, "GameLayer")
	gameLayer.AddNode(player)
	ebiten_extended.GameManager().World().AddLayer(gameLayer)
	return &Game{}
}

func (g *Game)  Init() error{
	return nil
}

func (g *Game) Update() error {
	ebiten_extended.GameManager().Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	ebiten_extended.GameManager().Draw(screen, op)
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}