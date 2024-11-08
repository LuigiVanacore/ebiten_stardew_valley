package ebiten_stardew_valley

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	ebiten_extended.Node2D
	direction math2D.Vector2D
	speed     int
}

func NewPlayer(pos math2D.Vector2D) *Player {
	player := &Player{
		Node2D: *ebiten_extended.NewNode2D("player"),
		speed: 20,
	}
	animationPlayer := ebiten_extended.NewAnimationPlayer()
	animationPlayer.AddAnimation(animationSets[Character_Down], Character_Down)
	// player_sprite := ebiten_extended.NewSprite("player_sprite", ebiten_extended.ResourceManager().GetTexture(PLAYER_SPRITE) , true)
	// player.AddChildren(player_sprite)
	// player.SetPosition(pos.X(), pos.Y())
	player.AddChildren(animationPlayer)
	return player
}

func (p *Player) Input() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		p.direction.SetY(-1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		p.direction.SetY(1)
	} else {
		p.direction.SetY(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		p.direction.SetX(1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		p.direction.SetX(-1)
	} else {
		p.direction.SetX(0)
	}
}

func (p *Player) Move() {
	// normalizing a vector
	if p.direction.Magnitude() > 0 {
		p.direction = p.direction.Normalize()
	}

	// horizontal movement
	horizontal_position := math2D.AddVectors(p.GetPosition(), p.direction.MultiplyScalar(float64(p.speed)))

	// vertical movement
	vertical_position := math2D.AddVectors(p.GetPosition(), p.direction.MultiplyScalar(float64(p.speed)))

	p.SetPosition(horizontal_position.X(), vertical_position.Y())
}

func (p *Player) Update() {
	p.Input()
	p.Move()
}	


// class Player(pygame.sprite.Sprite):
// 	def __init__(self, pos, group):
// 		super().__init__(group)

// 		# general setup
// 		self.image = pygame.Surface((32,64))
// 		self.image.fill('green')
// 		self.rect = self.image.get_rect(center = pos)

// 		# movement attributes
// 		self.direction = pygame.math.Vector2()
// 		self.pos = pygame.math.Vector2(self.rect.center)
// 		self.speed = 200


// 	def move(self,dt):

// 		# normalizing a vector
// 		if self.direction.magnitude() > 0:
// 			self.direction = self.direction.normalize()

// 		# horizontal movement
// 		self.pos.x += self.direction.x * self.speed * dt
// 		self.rect.centerx = self.pos.x

// 		# vertical movement
// 		self.pos.y += self.direction.y * self.speed * dt
// 		self.rect.centery = self.pos.y

// 	def update(self, dt):
// 		self.input()
// 		self.move(dt)