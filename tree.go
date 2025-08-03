package ebiten_stardew_valley

import "github.com/LuigiVanacore/ebiten_extended"

type Tree struct {
	ebiten_extended.Node2D
	sprite *ebiten_extended.Sprite
}


func NewTree(sprite *ebiten_extended.Sprite) *Tree {
	tree := &Tree{
		Node2D: *ebiten_extended.NewNode2D("Tree"),
		sprite: sprite,
	}
	tree.AddChild(tree.sprite)
	return tree
}



// ef __init__(self, pos, surf, groups, name, player_add):
// 		super().__init__(pos, surf, groups)

// 		# tree attributes
// 		self.health = 5
// 		self.alive = True
// 		stump_path = f'../graphics/stumps/{"small" if name == "Small" else "large"}.png'
// 		self.stump_surf = pygame.image.load(stump_path).convert_alpha()

// 		# apples
// 		self.apple_surf = pygame.image.load('../graphics/fruit/apple.png')
// 		self.apple_pos = APPLE_POS[name]
// 		self.apple_sprites = pygame.sprite.Group()
// 		self.create_fruit()

// 		self.player_add = player_add

// 		# sounds
// 		self.axe_sound = pygame.mixer.Sound('../audio/axe.mp3')

// 	def damage(self):
		
// 		# damaging the tree
// 		self.health -= 1

// 		# play sound
// 		self.axe_sound.play()

// 		# remove an apple
// 		if len(self.apple_sprites.sprites()) > 0:
// 			random_apple = choice(self.apple_sprites.sprites())
// 			Particle(
// 				pos = random_apple.rect.topleft,
// 				surf = random_apple.image, 
// 				groups = self.groups()[0], 
// 				z = LAYERS['fruit'])
// 			self.player_add('apple')
// 			random_apple.kill()

// 	def check_death(self):
// 		if self.health <= 0:
// 			Particle(self.rect.topleft, self.image, self.groups()[0], LAYERS['fruit'], 300)
// 			self.image = self.stump_surf
// 			self.rect = self.image.get_rect(midbottom = self.rect.midbottom)
// 			self.hitbox = self.rect.copy().inflate(-10,-self.rect.height * 0.6)
// 			self.alive = False
// 			self.player_add('wood')

// 	def update(self,dt):
// 		if self.alive:
// 			self.check_death()

// 	def create_fruit(self):
// 		for pos in self.apple_pos:
// 			if randint(0,10) < 2:
// 				x = pos[0] + self.rect.left
// 				y = pos[1] + self.rect.top
// 				Generic(
// 					pos = (x,y), 
// 					surf = self.apple_surf, 
// 					groups = [self.apple_sprites,self.groups()[0]],
// 					z = LAYERS['fruit'])