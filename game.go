package ebiten_stardew_valley

import (
	"fmt" 

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

type Game struct {
}

const mapPath = "../resources/data/map.tmx" 

func NewGame(debugFlag bool) *Game {

	ebiten_extended.GameManager().SetIsDebug(debugFlag)


	LoadAnimationSets()

	player := NewPlayer(math2D.NewVector2D(200, 200))
	overlay := NewOverlay(player) 

	ebiten_extended.GameManager().World().AddNode(player)
	ebiten_extended.GameManager().World().AddUINode(overlay) 
	ebiten_extended.GameManager().World().Camera().SetTransformToFollow(player)
	
	return &Game{}
}

func (g *Game)  Init() error{
	g.LoadTiledMap()
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

func (g *Game) LoadTiledMap()   {
	// Parse .tmx file.
	gameMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		  fmt.Printf("error parsing map: %s", err.Error())
	}
 

	tileMap := NewTileMap(gameMap)


	err = tileMap.RenderVisibleLayers()
	if err != nil {
		fmt.Printf("error rendering tile map: %s", err.Error())
 	}

	img := tileMap.Result
	tileMap.Clear()
	sprite := ebiten_extended.NewSprite("TileMap", ebiten.NewImageFromImage(img), GROUND_LAYER, false)
	sprite.Node2D.SetPosition(math2D.NewVector2D(0, 0))
	ebiten_extended.GameManager().World().AddNode(sprite)

	water_animation := ebiten_extended.ResourceManager().GetAnimation("Water")
	water_animation.SetTimePerFrame(15)
	var water_positions []math2D.Vector2D
	for _, layer := range gameMap.Layers {
		if layer.Name == "Water" {
			water_positions = tileMap.GetTilesPositionInLayer(layer)
		}
	}
				
	for _, position := range water_positions {
		waterSprite := ebiten_extended.NewAnimationSprite("Water", water_animation, WATER_LAYER, true)
			waterSprite.SetPosition(position) 
			waterSprite.SetPivot(math2D.NewVector2D(0,0))
			ebiten_extended.GameManager().World().AddNode(waterSprite)
	}
 

	for _, objectGroup := range gameMap.ObjectGroups {
		if objectGroup.Name == "Trees" {
			for _, object := range objectGroup.Objects {
				var image *ebiten.Image
			if object.Name == "Large" {
				 image = ebiten_extended.ResourceManager().GetImage(Objects_Tree_Medium)
			}
			if object.Name == "Small" {
				image = ebiten_extended.ResourceManager().GetImage(Objects_Tree_Small)
			}
			sprite := ebiten_extended.NewSprite("Tree", image, MAIN_LAYER, false)
			sprite.SetPivot(math2D.NewVector2D(0, float64(image.Bounds().Dy())))
			tree := NewTree(sprite)
			 
		tree.SetPosition(math2D.NewVector2D(object.X, object.Y))
		ebiten_extended.GameManager().World().AddNode(tree) 
		}
	}
}

for _, objectGroup := range gameMap.ObjectGroups {
		if objectGroup.Name == "Decoration" {
			for _, object := range objectGroup.Objects {

			sprite := ebiten_extended.NewSprite("Decoration", ebiten_extended.ResourceManager().GetImage(Objects_Flower), MAIN_LAYER, false)
			sprite.SetPivot(math2D.NewVector2D(0, float64(sprite.GetTexture().Bounds().Dy())))
			 

			sprite.SetPosition(math2D.NewVector2D(object.X, object.Y))
			ebiten_extended.GameManager().World().AddNode(sprite)
		}
	}
}





}


	// water_frames = import_folder('../graphics/water')
	// 	for x, y, surf in tmx_data.get_layer_by_name('Water').tiles():
	// 		Water((x * TILE_SIZE,y * TILE_SIZE), water_frames, self.all_sprites)

// class Water(Generic):
// 	def __init__(self, pos, frames, groups):

// 		#animation setup
// 		self.frames = frames
// 		self.frame_index = 0

// 		# sprite setup
// 		super().__init__(
// 				pos = pos, 
// 				surf = self.frames[self.frame_index], 
// 				groups = groups, 
// 				z = LAYERS['water']) 












// 	house_bottom_layers := []string {"HouseFurnitureTop"}

// 	for _, layerName := range house_bottom_layers {
		
// 		 tileMap.RenderLayerByName(layerName)
// 		 image := tileMap.Result


// 		 file, err := os.Create("output.png")
// if err != nil {
//     panic(err)
// }
// defer file.Close()

// err = png.Encode(file, image)
// if err != nil {
//     panic(err)
// }

// 		if image != nil {
// 			sprite := ebiten_extended.NewSprite(layerName, ebiten.NewImageFromImage(image), HOUSE_BOTTOM_LAYER, false)
// 			sprite.Node2D.SetPosition(math2D.NewVector2D(0, 0))
// 			ebiten_extended.GameManager().World().AddNode(sprite)
// 		} 
// 		tileMap.Clear()
// 	}

	//house_top_layers := []string {"HouseWalls", "HouseFurnitureTop"}


//  house 
// 		for layer in ['HouseFloor', 'HouseFurnitureBottom']:
// 			for x, y, surf in tmx_data.get_layer_by_name(layer).tiles():
// 				Generic((x * TILE_SIZE,y * TILE_SIZE), surf, self.all_sprites, LAYERS['house bottom'])

// 		for layer in ['HouseWalls', 'HouseFurnitureTop']:
// 			for x, y, surf in tmx_data.get_layer_by_name(layer).tiles():
// 				Generic((x * TILE_SIZE,y * TILE_SIZE), surf, self.all_sprites)

// 		# Fence
// 		for x, y, surf in tmx_data.get_layer_by_name('Fence').tiles():
// 			Generic((x * TILE_SIZE,y * TILE_SIZE), surf, [self.all_sprites, self.collision_sprites])

// 		# water 
// 		water_frames = import_folder('../graphics/water')
// 		for x, y, surf in tmx_data.get_layer_by_name('Water').tiles():
// 			Water((x * TILE_SIZE,y * TILE_SIZE), water_frames, self.all_sprites)

// 		# trees 
// 		for obj in tmx_data.get_layer_by_name('Trees'):
// 			Tree((obj.x, obj.y), obj.image, [self.all_sprites, self.collision_sprites], obj.name)

// 		# wildflowers 
// 		for obj in tmx_data.get_layer_by_name('Decoration'):
// 			WildFlower((obj.x, obj.y), obj.image, [self.all_sprites, self.collision_sprites])

// 		# collion tiles
// 		for x, y, surf in tmx_data.get_layer_by_name('Collision').tiles():
// 			Generic((x * TILE_SIZE, y * TILE_SIZE), pygame.Surface((TILE_SIZE, TILE_SIZE)), self.collision_sprites)

// 		# Player 
// 		for obj in tmx_data.get_layer_by_name('Player'):
// 			if obj.name == 'Start':
// 				self.player = Player((obj.x,obj.y), self.all_sprites, self.collision_sprites)
// 		Generic(
// 			pos = (0,0),
// 			surf = pygame.image.load('../graphics/world/ground.png').convert_alpha(),
// 			groups = self.all_sprites,
// 			z = LAYERS['ground'])



// class Generic(pygame.sprite.Sprite):
// 	def __init__(self, pos, surf, groups, z = LAYERS['main']):
// 		super().__init__(groups)
// 		self.image = surf
// 		self.rect = self.image.get_rect(topleft = pos)
// 		self.z = z
// 		self.hitbox = self.rect.copy().inflate(-self.rect.width * 0.2, -self.rect.height * 0.75)

// class Water(Generic):
// 	def __init__(self, pos, frames, groups):

// 		#animation setup
// 		self.frames = frames
// 		self.frame_index = 0

// 		# sprite setup
// 		super().__init__(
// 				pos = pos, 
// 				surf = self.frames[self.frame_index], 
// 				groups = groups, 
// 				z = LAYERS['water']) 

// 	def animate(self,dt):
// 		self.frame_index += 5 * dt
// 		if self.frame_index >= len(self.frames):
// 			self.frame_index = 0
// 		self.image = self.frames[int(self.frame_index)]

// 	def update(self,dt):
// 		self.animate(dt)

// class WildFlower(Generic):
// 	def __init__(self, pos, surf, groups):
// 		super().__init__(pos, surf, groups)
// 		self.hitbox = self.rect.copy().inflate(-20,-self.rect.height * 0.9)

// class Tree(Generic):
// 	def __init__(self, pos, surf, groups, name):
// 		super().__init__(pos, surf, groups)