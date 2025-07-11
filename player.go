package ebiten_stardew_valley

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	ebiten_extended.Node2D
	animationPlayer *ebiten_extended.AnimationPlayer
	direction math2D.Vector2D
	speed     int
}

func NewPlayer(pos math2D.Vector2D) *Player {
	animationPlayer := ebiten_extended.NewAnimationPlayer()
	player := &Player{
		Node2D: *ebiten_extended.NewNode2D("player"),
		animationPlayer: animationPlayer,
		speed: PLAYER_SPEED,
		direction: math2D.NewVector2D(0, 0),
	}

	player.loadAnimations()

	player.AddChild(animationPlayer)
	player.animationPlayer.SetCurrentAnimation(Character_Down_Idle)
	player.animationPlayer.Start()

	return player
}

func (p *Player) loadAnimations() {
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Down), Character_Down)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Up), Character_Up)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Left), Character_Left)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Right), Character_Right)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Down_Idle), Character_Down_Idle)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Up_Idle), Character_Up_Idle)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Left_Idle), Character_Left_Idle)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Right_Idle), Character_Right_Idle)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Down_Axe), Character_Down_Axe)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Up_Axe), Character_Up_Axe)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Left_Axe), Character_Left_Axe)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Right_Axe), Character_Right_Axe)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Down_Hoe), Character_Down_Hoe)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Up_Hoe), Character_Up_Hoe)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Left_Hoe), Character_Left_Hoe)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Right_Hoe), Character_Right_Hoe)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Down_Water), Character_Down_Water)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Up_Water), Character_Up_Water)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Left_Water), Character_Left_Water)
	p.animationPlayer.AddAnimation(ebiten_extended.ResourceManager().GetAnimation(Character_Right_Water), Character_Right_Water)
}

func (p *Player) Input() {
	key_pressed := inpututil.AppendJustPressedKeys()
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		p.direction.SetY(-1)
		p.animationPlayer.SetCurrentAnimation(Character_Up)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		p.direction.SetY(1)
	} else {
		p.direction.SetY(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		p.direction.SetX(1)
		p.animationPlayer.SetCurrentAnimation(Character_Right)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		p.direction.SetX(-1)
		p.animationPlayer.SetCurrentAnimation(Character_Left)
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