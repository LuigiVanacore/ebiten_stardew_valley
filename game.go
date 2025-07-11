package ebiten_stardew_valley

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func NewGame(debugFlag bool) *Game {

	ebiten_extended.GameManager().SetIsDebug(debugFlag)


	LoadAnimationSets()

	player := NewPlayer(math2D.NewVector2D(0, 0))
	overlay := NewOverlay(player)
	
	
	ebiten_extended.GameManager().World().AddNode(player)
	ebiten_extended.GameManager().World().AddNode(overlay)
	ebiten_extended.GameManager().World().Camera().SetTransformToFollow(player)
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