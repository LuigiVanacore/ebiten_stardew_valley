package ebiten_stardew_valley

import (
	"fmt"
	"time"

	"github.com/LuigiVanacore/ebiten_extended"
	fsm "github.com/LuigiVanacore/ebiten_extended/fsm"
	inputv3 "github.com/LuigiVanacore/ebiten_extended/input_v3"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)


type orientation int

const (
	UP orientation = iota
	DOWN
	LEFT
	RIGHT
)

type selected_tool int

const (
	AXE selected_tool = iota
	HOE
	WATERING_CAN
)

type selected_seed int

const (
	CORN selected_seed = iota
	TOMATO
)



const (
    MOVE_UP inputv3.ActionID = iota
    MOVE_DOWN
    MOVE_LEFT
    MOVE_RIGHT
    USE_TOOL
	SWITCH_TOOL
	USE_SEED
	SWITCH_SEED
)

const (
	MOVE_STATE fsm.StateID = iota
	USE_TOOL_STATE
	IDLE_STATE
)

type Player struct {
	ebiten_extended.Node2D
	animationPlayer *ebiten_extended.AnimationPlayer
	tool_timer      *ebiten_extended.Timer
	switch_tool_timer *ebiten_extended.Timer
	seed_use_timer *ebiten_extended.Timer
	switch_seed_timer *ebiten_extended.Timer
	stateMachine    *fsm.StateMachine
	direction       math2D.Vector2D
	orientation     orientation
	speed          int
	selected_tool  selected_tool
	selected_seed  selected_seed
}



func NewPlayer(pos math2D.Vector2D) *Player {
	animationPlayer := ebiten_extended.NewAnimationPlayer("player_animationPlayer", MAIN_LAYER)
	tool_timer := ebiten_extended.NewTimer(time.Duration(3)*time.Second, false)
	switch_tool_timer := ebiten_extended.NewTimer(time.Duration(200)*time.Millisecond, false)
	seed_use_timer := ebiten_extended.NewTimer(time.Duration(350)*time.Millisecond, false)
	switch_seed_timer := ebiten_extended.NewTimer(time.Duration(200)*time.Millisecond, false)

	player := &Player{
		Node2D: *ebiten_extended.NewNode2D("player"),
		animationPlayer: animationPlayer,
		tool_timer: tool_timer,
		switch_tool_timer: switch_tool_timer,
		seed_use_timer: seed_use_timer,
		switch_seed_timer: switch_seed_timer,
		speed: PLAYER_SPEED,
		direction: math2D.NewVector2D(0, 0),
		orientation: DOWN,
	}

	player.SetPosition(pos)

	player.stateMachine = NewPlayerStateMachine(player)

	player.loadAnimations()
	player.loadActions()

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

func (p *Player) loadActions() {
	// Load actions if needed, e.g., for using tools like axe, hoe, watering can
	// This could be similar to loadAnimations but for actions
	move_up := inputv3.NewKeyAction(ebiten.KeyW, inputv3.Hold | inputv3.PressOnce)
	move_down := inputv3.NewKeyAction(ebiten.KeyS, inputv3.Hold | inputv3.PressOnce)
	move_left := inputv3.NewKeyAction(ebiten.KeyA, inputv3.Hold | inputv3.PressOnce)
	move_right := inputv3.NewKeyAction(ebiten.KeyD, inputv3.Hold | inputv3.PressOnce)
	use_tool := inputv3.NewKeyAction(ebiten.KeySpace, inputv3.PressOnce)
    switch_tool := inputv3.NewKeyAction(ebiten.KeyK, inputv3.PressOnce)
	use_seed := inputv3.NewKeyAction(ebiten.KeyControlLeft, inputv3.PressOnce)
	switch_seed := inputv3.NewKeyAction(ebiten.KeyE, inputv3.PressOnce)

	ebiten_extended.InputManager().RegisterAction(MOVE_UP, move_up)
	ebiten_extended.InputManager().RegisterAction(MOVE_DOWN, move_down)
	ebiten_extended.InputManager().RegisterAction(MOVE_LEFT, move_left)
	ebiten_extended.InputManager().RegisterAction(MOVE_RIGHT, move_right)
	ebiten_extended.InputManager().RegisterAction(USE_TOOL, use_tool)
	ebiten_extended.InputManager().RegisterAction(SWITCH_TOOL, switch_tool)
	ebiten_extended.InputManager().RegisterAction(USE_SEED, use_seed)
	ebiten_extended.InputManager().RegisterAction(SWITCH_SEED, switch_seed)
}

func (p *Player) GetCurrentTool() selected_tool {
	return p.selected_tool
}

func (p *Player) GetCurrentSeed() selected_seed {
	return p.selected_seed
}

func (p *Player) Input() {

	if ebiten_extended.InputManager().IsActionActive(MOVE_UP) {
		p.direction.SetY(-1)
	} else if ebiten_extended.InputManager().IsActionActive(MOVE_DOWN) {
		p.direction.SetY(1)
	} else {
		p.direction.SetY(0)
	}

		if ebiten_extended.InputManager().IsActionActive(MOVE_RIGHT) {
			p.direction.SetX(1)
		} else if ebiten_extended.InputManager().IsActionActive(MOVE_LEFT) {
			p.direction.SetX(-1)
		} else {
			p.direction.SetX(0)
		}

	if ebiten_extended.InputManager().IsActionActive(SWITCH_TOOL) {
		p.SwitchTool()
	}

	if ebiten_extended.InputManager().IsActionActive(SWITCH_SEED) {
		p.SwitchSeed()
	}
}

func (p *Player) SwitchSeed() {
		if p.switch_seed_timer.IsEnded() && p.tool_timer.IsEnded() {
				switch p.selected_seed {
				case CORN:
					p.selected_seed = TOMATO
				case TOMATO:
					p.selected_seed = CORN
				}
				p.switch_seed_timer.Start()
			}
		}


func (p *Player) SwitchTool() {
	if p.switch_tool_timer.IsEnded() && p.tool_timer.IsEnded() {
			switch p.selected_tool {
			case AXE:
				p.selected_tool = HOE
			case HOE:
				p.selected_tool = WATERING_CAN
			case WATERING_CAN:
				p.selected_tool = AXE
			}
			p.switch_tool_timer.Start()
		}
}

func (p *Player) Move() {
	// normalizing a vector
	if p.direction.Magnitude() > 0 {
		p.direction = p.direction.Normalize()
	
	// horizontal movement
	horizontal_position := math2D.AddVectors(p.GetPosition(), p.direction.MultiplyScalar(float64(p.speed)))

	// vertical movement
	vertical_position := math2D.AddVectors(p.GetPosition(), p.direction.MultiplyScalar(float64(p.speed)))

	p.SetPosition(math2D.NewVector2D(horizontal_position.X(), vertical_position.Y()))
		if p.direction.Y() < 0 {
			p.animationPlayer.SetCurrentAnimation(Character_Up)
		} else if p.direction.Y() > 0 {
			p.animationPlayer.SetCurrentAnimation(Character_Down)
		} else if p.direction.X() < 0 {
			p.animationPlayer.SetCurrentAnimation(Character_Left)
		} else if p.direction.X() > 0 {
			p.animationPlayer.SetCurrentAnimation(Character_Right)
		}
	} else {
		switch p.animationPlayer.GetCurrentAnimation() {
		case Character_Up:	
			p.animationPlayer.SetCurrentAnimation(Character_Up_Idle)
		case Character_Down:
			p.animationPlayer.SetCurrentAnimation(Character_Down_Idle)
		case Character_Left:
			p.animationPlayer.SetCurrentAnimation(Character_Left_Idle)
		case Character_Right:
			p.animationPlayer.SetCurrentAnimation(Character_Right_Idle)
	}
}
}

 

func (p *Player) Update() {
	p.Input()
	p.stateMachine.Update()
	fmt.Println("Current Position:", p.GetPosition())
}	


type PlayerMoveState struct {
	player *Player
}

func NewPlayerMoveState(player *Player) *PlayerMoveState {
	return &PlayerMoveState{
		player: player,
	}
}

func (s *PlayerMoveState) Enter() {
}

func (s *PlayerMoveState) Exit() {
}

func (s *PlayerMoveState) Update() {
 	// horizontal movement
	horizontal_position := math2D.AddVectors(s.player.GetPosition(), s.player.direction.MultiplyScalar(float64(s.player.speed)))

	// vertical movement
	vertical_position := math2D.AddVectors(s.player.GetPosition(), s.player.direction.MultiplyScalar(float64(s.player.speed)))

	s.player.SetPosition(math2D.NewVector2D(horizontal_position.X(), vertical_position.Y()))
		if s.player.direction.Y() < 0 {
			s.player.orientation = UP
			s.player.animationPlayer.SetCurrentAnimation(Character_Up)
		} else if s.player.direction.Y() > 0 {
			s.player.orientation = DOWN
			s.player.animationPlayer.SetCurrentAnimation(Character_Down)
		} else if s.player.direction.X() < 0 {
			s.player.orientation = LEFT
			s.player.animationPlayer.SetCurrentAnimation(Character_Left)
		} else if s.player.direction.X() > 0 {
			s.player.orientation = RIGHT
			s.player.animationPlayer.SetCurrentAnimation(Character_Right)
		}
}



type PlayerUseToolState struct {
	player *Player
}

func NewPlayerUseToolState(player *Player) *PlayerUseToolState {
	return &PlayerUseToolState{
		player: player,
	}
}

func (s *PlayerUseToolState) Enter() {
	var tool string
	switch s.player.selected_tool {
	case AXE:
		tool = "Axe"
	case HOE:
		tool = "Hoe"
	case WATERING_CAN:
		tool = "Water"
	default:
		return
	}

	if s.player.orientation == UP {
		s.player.animationPlayer.SetCurrentAnimation("Character_" + "Up_" + tool)
	} else if s.player.orientation == DOWN {
		s.player.animationPlayer.SetCurrentAnimation("Character_" + "Down_" + tool)
	} else if s.player.orientation == LEFT {
		s.player.animationPlayer.SetCurrentAnimation("Character_" + "Left_" + tool)
	} else if s.player.orientation == RIGHT {
		s.player.animationPlayer.SetCurrentAnimation("Character_" + "Right_" + tool)
	}
	s.player.tool_timer.Start()
}

func (s *PlayerUseToolState) Exit() {
}

func (s *PlayerUseToolState) Update() {
	if s.player.tool_timer.IsEnded() {
		s.Exit()
	} 
}


type PlayerIdleState struct {
	player *Player
}

func NewPlayerIdleState(player *Player) *PlayerIdleState {
	return &PlayerIdleState{
		player: player,
	}
}	

func (s *PlayerIdleState) Enter() {
switch s.player.orientation {
		case UP:
			s.player.animationPlayer.SetCurrentAnimation(Character_Up_Idle)
		case DOWN:
			s.player.animationPlayer.SetCurrentAnimation(Character_Down_Idle)
		case LEFT:
			s.player.animationPlayer.SetCurrentAnimation(Character_Left_Idle)
		case RIGHT:
			s.player.animationPlayer.SetCurrentAnimation(Character_Right_Idle)
		}
}

func (s *PlayerIdleState) Exit() {
}

func (s *PlayerIdleState) Update() {

}

func NewPlayerIdleTransitions(p *Player) []fsm.Transition {
	return []fsm.Transition{
		{
			Origin:  IDLE_STATE,
			Target:  MOVE_STATE,
			Trigger: func() bool { return p.direction.Magnitude() > 0  },
		},
		{
			Origin:  IDLE_STATE,
			Target:  USE_TOOL_STATE,
			Trigger: func() bool { return ebiten_extended.InputManager().IsActionActive(USE_TOOL) },
		},
	}
}

func NewPlayerMoveTransitions(p *Player) []fsm.Transition {
	return []fsm.Transition{
		{
			Origin:  MOVE_STATE,
			Target:  IDLE_STATE,
			Trigger: func() bool { return p.direction.Magnitude() == 0 },
		},
		{
			Origin:  MOVE_STATE,
			Target:  USE_TOOL_STATE,
			Trigger: func() bool { return ebiten_extended.InputManager().IsActionActive(USE_TOOL) },
		},
	}
}

func NewPlayerUseToolTransitions(p *Player) []fsm.Transition {
	return []fsm.Transition{
		{
			Origin:  USE_TOOL_STATE,
			Target:  IDLE_STATE,
			Trigger: func() bool { return p.tool_timer.IsEnded() },
		},
	}
}


func NewPlayerStateMachine(player *Player) *fsm.StateMachine {
	stateMachine := fsm.NewStateMachine()

	moveState := NewPlayerMoveState(player)
	useToolState := NewPlayerUseToolState(player)
	idleState := NewPlayerIdleState(player)

	stateMachine.AddState(MOVE_STATE, moveState)
	stateMachine.AddState(USE_TOOL_STATE, useToolState)
	stateMachine.AddState(IDLE_STATE, idleState)

	stateMachine.AddTransitions( NewPlayerMoveTransitions(player)...)
	stateMachine.AddTransitions( NewPlayerUseToolTransitions(player)...)
	stateMachine.AddTransitions( NewPlayerIdleTransitions(player)...)

	stateMachine.SetState(IDLE_STATE)

	return stateMachine
}
