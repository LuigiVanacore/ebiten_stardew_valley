package ebiten_stardew_valley

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

type Overlay struct {
	ebiten_extended.Node2D
	player *Player
	overlay_tool *ebiten_extended.Sprite
	overlay_seed *ebiten_extended.Sprite
	overlay_tools_images map[selected_tool]*ebiten.Image
	overlay_seed_images map[selected_seed]*ebiten.Image
}

var (
    OVERLAY_TOOL = math2D.NewVector2D(40, SCREEN_HEIGHT - 80)
    OVERLAY_SEED = math2D.NewVector2D(90, SCREEN_HEIGHT - 80)
)

func NewOverlay(player *Player) *Overlay {
	overlay := &Overlay{
		Node2D: *ebiten_extended.NewNode2D("Overlay"),
		player: player,
		overlay_tools_images: make(map[selected_tool]*ebiten.Image),
		overlay_seed_images: make(map[selected_seed]*ebiten.Image),
	}

	overlay.overlay_tools_images[AXE] = ebiten_extended.ResourceManager().GetImage(Overlay_Axe)
	overlay.overlay_tools_images[HOE] = ebiten_extended.ResourceManager().GetImage(Overlay_Hoe)
	overlay.overlay_tools_images[WATERING_CAN] = ebiten_extended.ResourceManager().GetImage(Overlay_Water)
	overlay.overlay_seed_images[TOMATO] = ebiten_extended.ResourceManager().GetImage(Overlay_Tomato)
	overlay.overlay_seed_images[CORN] = ebiten_extended.ResourceManager().GetImage(Overlay_Corn)

	overlay.overlay_tool = ebiten_extended.NewSprite("tool_overlay_sprite",overlay.overlay_tools_images[AXE], UI_LAYER, false)
	overlay.overlay_seed = ebiten_extended.NewSprite("seed_overlay_sprite", overlay.overlay_seed_images[TOMATO], UI_LAYER, false)
	overlay.overlay_tool.SetPosition(OVERLAY_TOOL)
	overlay.overlay_seed.SetPosition(OVERLAY_SEED)

	overlay.AddChild(overlay.overlay_tool)
	overlay.AddChild(overlay.overlay_seed)

	return overlay
}


func (o *Overlay) Update()  {
	o.updateCurrentTool()
	o.updateCurrentSeed()
}



func (o *Overlay) updateCurrentTool() {
	tool := o.player.GetCurrentTool()

	if image, exists := o.overlay_tools_images[tool]; exists {
		o.overlay_tool.SetTexture(image)
	}
}

func (o *Overlay) updateCurrentSeed() {
	seed := o.player.GetCurrentSeed()

	if image, exists := o.overlay_seed_images[seed]; exists {
		o.overlay_seed.SetTexture(image)
	}
}