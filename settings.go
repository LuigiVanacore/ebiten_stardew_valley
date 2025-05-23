package ebiten_stardew_valley

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_stardew_valley/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 720
	TILE_SIZE     = 64
)


const (
    PLAYER_SPEED = 5
)


const (
    PLAYER_SPRITE = "player"

    // Define constants for each key string
    Character_Down_0      = "Character_Down_0"
    Character_Down_1      = "Character_Down_1"
    Character_Down_2      = "Character_Down_2"
    Character_Down_3      = "Character_Down_3"
    Character_Down_Axe_0  = "Character_Down_Axe_0"
    Character_Down_Axe_1  = "Character_Down_Axe_1"
    Character_Down_Hoe_0  = "Character_Down_Hoe_0"
    Character_Down_Hoe_1  = "Character_Down_Hoe_1"
    Character_Down_Idle_0 = "Character_Down_Idle_0"
    Character_Down_Idle_1 = "Character_Down_Idle_1"
    Character_Down_Water_0 = "Character_Down_Water_0"
    Character_Down_Water_1 = "Character_Down_Water_1"
    Character_Left_0      = "Character_Left_0"
    Character_Left_1      = "Character_Left_1"
    Character_Left_2      = "Character_Left_2"
    Character_Left_3      = "Character_Left_3"
    Character_Left_Axe_0  = "Character_Left_Axe_0"
    Character_Left_Axe_1  = "Character_Left_Axe_1"
    Character_Left_Hoe_0  = "Character_Left_Hoe_0"
    Character_Left_Hoe_1  = "Character_Left_Hoe_1"
    Character_Left_Idle_0 = "Character_Left_Idle_0"
    Character_Left_Idle_1 = "Character_Left_Idle_1"
    Character_Left_Water_0 = "Character_Left_Water_0"
    Character_Left_Water_1 = "Character_Left_Water_1"
    Character_Right_0     = "Character_Right_0"
    Character_Right_1     = "Character_Right_1"
    Character_Right_2     = "Character_Right_2"
    Character_Right_3     = "Character_Right_3"
    Character_Right_Axe_0 = "Character_Right_Axe_0"
    Character_Right_Axe_1 = "Character_Right_Axe_1"
    Character_Right_Hoe_0 = "Character_Right_Hoe_0"
    Character_Right_Hoe_1 = "Character_Right_Hoe_1"
    Character_Right_Idle_0 = "Character_Right_Idle_0"
    Character_Right_Idle_1 = "Character_Right_Idle_1"
    Character_Right_Water_0 = "Character_Right_Water_0"
    Character_Right_Water_1 = "Character_Right_Water_1"
    Character_Up_0        = "Character_Up_0"
    Character_Up_1        = "Character_Up_1"
    Character_Up_2        = "Character_Up_2"
    Character_Up_3        = "Character_Up_3"
    Character_Up_Axe_0    = "Character_Up_Axe_0"
    Character_Up_Axe_1    = "Character_Up_Axe_1"
    Character_Up_Hoe_0    = "Character_Up_Hoe_0"
    Character_Up_Hoe_1    = "Character_Up_Hoe_1"
    Character_Up_Idle_0   = "Character_Up_Idle_0"
    Character_Up_Idle_1   = "Character_Up_Idle_1"
    Character_Up_Water_0  = "Character_Up_Water_0"
    Character_Up_Water_1  = "Character_Up_Water_1"
    Environment_Bridge    = "Environment_Bridge"
    Environment_Collision = "Environment_Collision"
    Environment_Fences    = "Environment_Fences"
    Environment_Grass     = "Environment_Grass"
    Environment_Hills     = "Environment_Hills"
    Environment_House_Decoration = "Environment_House_Decoration"
    Environment_House     = "Environment_House"
    Environment_Paths     = "Environment_Paths"
    Environment_Plant_Decoration = "Environment_Plant_Decoration"
    Environment_Water_Decoration = "Environment_Water_Decoration"
    Environment_Water     = "Environment_Water"
    Environment_Interaction = "Environment_Interaction"
    Fruit_Apple           = "Fruit_Apple"
    Fruit_Corn_0          = "Fruit_Corn_0"
    Fruit_Corn_1          = "Fruit_Corn_1"
    Fruit_Corn_2          = "Fruit_Corn_2"
    Fruit_Corn_3          = "Fruit_Corn_3"
    Fruit_Tomato_0        = "Fruit_Tomato_0"
    Fruit_Tomato_1        = "Fruit_Tomato_1"
    Fruit_Tomato_2        = "Fruit_Tomato_2"
    Fruit_Tomato_3        = "Fruit_Tomato_3"
    Objects_Bush          = "Objects_Bush"
    Objects_Flower        = "Objects_Flower"
    Objects_Merchant      = "Objects_Merchant"
    Objects_Mushroom      = "Objects_Mushroom"
    Objects_Mushrooms     = "Objects_Mushrooms"
    Objects_Stump_Medium  = "Objects_Stump_Medium"
    Objects_Stump_Small   = "Objects_Stump_Small"
    Objects_Sunflower     = "Objects_Sunflower"
    Objects_Tree_Medium   = "Objects_Tree_Medium"
    Objects_Tree_Small    = "Objects_Tree_Small"
    Overlay_Axe           = "Overlay_Axe"
    Overlay_Corn          = "Overlay_Corn"
    Overlay_Hoe           = "Overlay_Hoe"
    Overlay_Tomato        = "Overlay_Tomato"
    Overlay_Water         = "Overlay_Water"
    Rain_Drops_0          = "Rain_Drops_0"
    Rain_Drops_1          = "Rain_Drops_1"
    Rain_Drops_2          = "Rain_Drops_2"
    Rain_Floor_0          = "Rain_Floor_0"
    Rain_Floor_1          = "Rain_Floor_1"
    Rain_Floor_2          = "Rain_Floor_2"
    Soil_B                = "Soil_B"
    Soil_Bl               = "Soil_Bl"
    Soil_Bm               = "Soil_Bm"
    Soil_Br               = "Soil_Br"
    Soil_L                = "Soil_L"
    Soil_Lm               = "Soil_Lm"
    Soil_Lr               = "Soil_Lr"
    Soil_Lrb              = "Soil_Lrb"
    Soil_Lrt              = "Soil_Lrt"
    Soil_O                = "Soil_O"
    Soil_R                = "Soil_R"
    Soil_Rm               = "Soil_Rm"
    Soil_Soil             = "Soil_Soil"
    Soil_T                = "Soil_T"
    Soil_Tb               = "Soil_Tb"
    Soil_Tbl              = "Soil_Tbl"
    Soil_Tbr              = "Soil_Tbr"
    Soil_Tl               = "Soil_Tl"
    Soil_Tm               = "Soil_Tm"
    Soil_Tr               = "Soil_Tr"
    Soil_X                = "Soil_X"
    Soil_Water_0          = "Soil_Water_0"
    Soil_Water_1          = "Soil_Water_1"
    Soil_Water_2          = "Soil_Water_2"
    Stumps_Large          = "Stumps_Large"
    Stumps_Small          = "Stumps_Small"
    Water_0               = "Water_0"
    Water_1               = "Water_1"
    Water_2               = "Water_2"
    Water_3               = "Water_3"
    World_Ground          = "World_Ground"
)

// LoadImages loads the embedded images into the resourceManager
func LoadImages() {
    imageVars := map[string][]byte{
        Character_Down_0:      resources.Character_Down_0,
        Character_Down_1:      resources.Character_Down_1,
        Character_Down_2:      resources.Character_Down_2,
        Character_Down_3:      resources.Character_Down_3,
        Character_Down_Axe_0:  resources.Character_Down_Axe_0,
        Character_Down_Axe_1:  resources.Character_Down_Axe_1,
        Character_Down_Hoe_0:  resources.Character_Down_Hoe_0,
        Character_Down_Hoe_1:  resources.Character_Down_Hoe_1,
        Character_Down_Idle_0: resources.Character_Down_Idle_0,
        Character_Down_Idle_1: resources.Character_Down_Idle_1,
        Character_Down_Water_0: resources.Character_Down_Water_0,
        Character_Down_Water_1: resources.Character_Down_Water_1,
        Character_Left_0:      resources.Character_Left_0,
        Character_Left_1:      resources.Character_Left_1,
        Character_Left_2:      resources.Character_Left_2,
        Character_Left_3:      resources.Character_Left_3,
        Character_Left_Axe_0:  resources.Character_Left_Axe_0,
        Character_Left_Axe_1:  resources.Character_Left_Axe_1,
        Character_Left_Hoe_0:  resources.Character_Left_Hoe_0,
        Character_Left_Hoe_1:  resources.Character_Left_Hoe_1,
        Character_Left_Idle_0: resources.Character_Left_Idle_0,
        Character_Left_Idle_1: resources.Character_Left_Idle_1,
        Character_Left_Water_0: resources.Character_Left_Water_0,
        Character_Left_Water_1: resources.Character_Left_Water_1,
        Character_Right_0:     resources.Character_Right_0,
        Character_Right_1:     resources.Character_Right_1,
        Character_Right_2:     resources.Character_Right_2,
        Character_Right_3:     resources.Character_Right_3,
        Character_Right_Axe_0: resources.Character_Right_Axe_0,
        Character_Right_Axe_1: resources.Character_Right_Axe_1,
        Character_Right_Hoe_0: resources.Character_Right_Hoe_0,
        Character_Right_Hoe_1: resources.Character_Right_Hoe_1,
        Character_Right_Idle_0: resources.Character_Right_Idle_0,
        Character_Right_Idle_1: resources.Character_Right_Idle_1,
        Character_Right_Water_0: resources.Character_Right_Water_0,
        Character_Right_Water_1: resources.Character_Right_Water_1,
        Character_Up_0:        resources.Character_Up_0,
        Character_Up_1:        resources.Character_Up_1,
        Character_Up_2:        resources.Character_Up_2,
        Character_Up_3:        resources.Character_Up_3,
        Character_Up_Axe_0:    resources.Character_Up_Axe_0,
        Character_Up_Axe_1:    resources.Character_Up_Axe_1,
        Character_Up_Hoe_0:    resources.Character_Up_Hoe_0,
        Character_Up_Hoe_1:    resources.Character_Up_Hoe_1,
        Character_Up_Idle_0:   resources.Character_Up_Idle_0,
        Character_Up_Idle_1:   resources.Character_Up_Idle_1,
        Character_Up_Water_0:  resources.Character_Up_Water_0,
        Character_Up_Water_1:  resources.Character_Up_Water_1,
        Environment_Bridge:    resources.Environment_Bridge,
        Environment_Collision: resources.Environment_Collision,
        Environment_Fences:    resources.Environment_Fences,
        Environment_Grass:     resources.Environment_Grass,
        Environment_Hills:     resources.Environment_Hills,
        Environment_House_Decoration: resources.Environment_House_Decoration,
        Environment_House:     resources.Environment_House,
        Environment_Paths:     resources.Environment_Paths,
        Environment_Plant_Decoration: resources.Environment_Plant_Decoration,
        Environment_Water_Decoration: resources.Environment_Water_Decoration,
        Environment_Water:     resources.Environment_Water,
        Environment_Interaction: resources.Environment_Interaction,
        Fruit_Apple:           resources.Fruit_Apple,
        Fruit_Corn_0:          resources.Fruit_Corn_0,
        Fruit_Corn_1:          resources.Fruit_Corn_1,
        Fruit_Corn_2:          resources.Fruit_Corn_2,
        Fruit_Corn_3:          resources.Fruit_Corn_3,
        Fruit_Tomato_0:        resources.Fruit_Tomato_0,
        Fruit_Tomato_1:        resources.Fruit_Tomato_1,
        Fruit_Tomato_2:        resources.Fruit_Tomato_2,
        Fruit_Tomato_3:        resources.Fruit_Tomato_3,
        Objects_Bush:          resources.Objects_Bush,
        Objects_Flower:        resources.Objects_Flower,
        Objects_Merchant:      resources.Objects_Merchant,
        Objects_Mushroom:      resources.Objects_Mushroom,
        Objects_Mushrooms:     resources.Objects_Mushrooms,
        Objects_Stump_Medium:  resources.Objects_Stump_Medium,
        Objects_Stump_Small:   resources.Objects_Stump_Small,
        Objects_Sunflower:     resources.Objects_Sunflower,
        Objects_Tree_Medium:   resources.Objects_Tree_Medium,
        Objects_Tree_Small:    resources.Objects_Tree_Small,
        Overlay_Axe:           resources.Overlay_Axe,
        Overlay_Corn:          resources.Overlay_Corn,
        Overlay_Hoe:           resources.Overlay_Hoe,
        Overlay_Tomato:        resources.Overlay_Tomato,
        Overlay_Water:         resources.Overlay_Water,
        Rain_Drops_0:          resources.Rain_Drops_0,
        Rain_Drops_1:          resources.Rain_Drops_1,
        Rain_Drops_2:          resources.Rain_Drops_2,
        Rain_Floor_0:          resources.Rain_Floor_0,
        Rain_Floor_1:          resources.Rain_Floor_1,
        Rain_Floor_2:          resources.Rain_Floor_2,
        Soil_B:                resources.Soil_B,
        Soil_Bl:               resources.Soil_Bl,
        Soil_Bm:               resources.Soil_Bm,
        Soil_Br:               resources.Soil_Br,
        Soil_L:                resources.Soil_L,
        Soil_Lm:               resources.Soil_Lm,
        Soil_Lr:               resources.Soil_Lr,
        Soil_Lrb:              resources.Soil_Lrb,
        Soil_Lrt:              resources.Soil_Lrt,
        Soil_O:                resources.Soil_O,
        Soil_R:                resources.Soil_R,
        Soil_Rm:               resources.Soil_Rm,
        Soil_Soil:             resources.Soil_Soil,
        Soil_T:                resources.Soil_T,
        Soil_Tb:               resources.Soil_Tb,
        Soil_Tbl:              resources.Soil_Tbl,
        Soil_Tbr:              resources.Soil_Tbr,
        Soil_Tl:               resources.Soil_Tl,
        Soil_Tm:               resources.Soil_Tm,
        Soil_Tr:               resources.Soil_Tr,
        Soil_X:                resources.Soil_X,
        Soil_Water_0:          resources.Soil_Water_0,
        Soil_Water_1:          resources.Soil_Water_1,
        Soil_Water_2:          resources.Soil_Water_2,
        Stumps_Large:          resources.Stumps_Large,
        Stumps_Small:          resources.Stumps_Small,
        Water_0:               resources.Water_0,
        Water_1:               resources.Water_1,
        Water_2:               resources.Water_2,
        Water_3:               resources.Water_3,
        World_Ground:          resources.World_Ground,
    }

    for name, data := range imageVars {
        ebiten_extended.ResourceManager().AddImage(name, data)
	}
}


// Define constants for animation set keys
const (
    Character_Down      = "Character_Down"
    Character_Down_Axe  = "Character_Down_Axe"
    Character_Down_Hoe  = "Character_Down_Hoe"
    Character_Down_Idle = "Character_Down_Idle"
    Character_Down_Water = "Character_Down_Water"
    Character_Left      = "Character_Left"
    Character_Left_Axe  = "Character_Left_Axe"
    Character_Left_Hoe  = "Character_Left_Hoe"
    Character_Left_Idle = "Character_Left_Idle"
    Character_Left_Water = "Character_Left_Water"
    Character_Right     = "Character_Right"
    Character_Right_Axe = "Character_Right_Axe"
    Character_Right_Hoe = "Character_Right_Hoe"
    Character_Right_Idle = "Character_Right_Idle"
    Character_Right_Water = "Character_Right_Water"
    Character_Up        = "Character_Up"
    Character_Up_Axe    = "Character_Up_Axe"
    Character_Up_Hoe    = "Character_Up_Hoe"
    Character_Up_Idle   = "Character_Up_Idle"
    Character_Up_Water  = "Character_Up_Water"
)

// AnimationSet represents a set of images for animation


// CreateAnimationSets creates a map of AnimationSet with key as constant string
func LoadAnimationSets() {

	LoadImages()

	
	// Define the mapping of key strings to their corresponding image keys
	animmationSets := map[string][]string{
		Character_Down:      {Character_Down_0, Character_Down_1, Character_Down_2, Character_Down_3},
		Character_Down_Axe:  {Character_Down_Axe_0, Character_Down_Axe_1},
		Character_Down_Hoe:  {Character_Down_Hoe_0, Character_Down_Hoe_1},
		Character_Down_Idle: {Character_Down_Idle_0, Character_Down_Idle_1},
		Character_Down_Water: {Character_Down_Water_0, Character_Down_Water_1},
		Character_Left:      {Character_Left_0, Character_Left_1, Character_Left_2, Character_Left_3},
		Character_Left_Axe:  {Character_Left_Axe_0, Character_Left_Axe_1},
		Character_Left_Hoe:  {Character_Left_Hoe_0, Character_Left_Hoe_1},
		Character_Left_Idle: {Character_Left_Idle_0, Character_Left_Idle_1},
		Character_Left_Water: {Character_Left_Water_0, Character_Left_Water_1},
		Character_Right:     {Character_Right_0, Character_Right_1, Character_Right_2, Character_Right_3},
		Character_Right_Axe: {Character_Right_Axe_0, Character_Right_Axe_1},
		Character_Right_Hoe: {Character_Right_Hoe_0, Character_Right_Hoe_1},
		Character_Right_Idle: {Character_Right_Idle_0, Character_Right_Idle_1},
		Character_Right_Water: {Character_Right_Water_0, Character_Right_Water_1},
		Character_Up:        {Character_Up_0, Character_Up_1, Character_Up_2, Character_Up_3},
		Character_Up_Axe:    {Character_Up_Axe_0, Character_Up_Axe_1},
		Character_Up_Hoe:    {Character_Up_Hoe_0, Character_Up_Hoe_1},
		Character_Up_Idle:   {Character_Up_Idle_0, Character_Up_Idle_1},
		Character_Up_Water:  {Character_Up_Water_0, Character_Up_Water_1},
	}

	var spriteSheet []*ebiten.Image

	// Load images into the AnimationSet
	for key, animationKeys := range animmationSets {
		spriteSheet = nil
		for _, animationKey := range animationKeys {
			img := ebiten_extended.ResourceManager().GetImage(animationKey)
			if img != nil {
				spriteSheet = append(spriteSheet, img)
			}
		}
		ebiten_extended.ResourceManager().AddAnimation(key, ebiten_extended.NewAnimationSet(spriteSheet, getCenterImage(spriteSheet[0]), uint(len(spriteSheet)), 4 / float64(len(spriteSheet))  , true))
	}
}

func getCenterImage(img *ebiten.Image) math2D.Vector2D {
	return math2D.NewVector2D(float64(img.Bounds().Dx()/2), float64(img.Bounds().Dy()/2))
}	