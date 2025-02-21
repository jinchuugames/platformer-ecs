package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
)

// InputSystem handles player input
// It reads the input and sets the velocity of the player character based on the speed in its Movement component.
type InputSystem struct {
	f *generic.Filter2[GodotNode, Movement]
}

func (s *InputSystem) Initialize(w *ecs.World) {}
func (s *InputSystem) Update(w *ecs.World) {
	q := s.f.Query(w)
	for q.Next() {
		node, mvmt := q.Get()
		body, ok := classdb.As[CharacterBody2D.Instance](node.Node)
		if !ok {
			continue
		}
		gravity := ecs.GetResource[Gravity](w)
		delta := ecs.GetResource[DeltaTime](w)
		velocity := body.Velocity()
		if !body.IsOnFloor() {
			velocity.Y += gravity.Value * delta.Delta
		} else {
			velocity.Y = 0
		}
		if Input.IsActionPressed(PlayerControls.MoveUp) && body.IsOnFloor() {
			velocity.Y = mvmt.JumpVelocity
		}

		velocity.X = 0
		if Input.IsActionPressed(PlayerControls.MoveRight) {
			velocity.X += mvmt.Speed
		}
		if Input.IsActionPressed(PlayerControls.MoveLeft) {
			velocity.X -= mvmt.Speed
		}
		body.SetVelocity(velocity)
	}
}
func (s *InputSystem) Finalize(w *ecs.World) {}
