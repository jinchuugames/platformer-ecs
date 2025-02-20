package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/variant/Vector2"
)

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

		velocity := Vector2.XY{X: 0, Y: 0}
		if Input.IsActionPressed(PlayerControls.MoveRight) {
			velocity.X += 1
		}
		if Input.IsActionPressed(PlayerControls.MoveLeft) {
			velocity.X -= 1
		}
		if Input.IsActionPressed(PlayerControls.MoveDown) {
			velocity.Y += 1
		}
		if Input.IsActionPressed(PlayerControls.MoveUp) {
			velocity.Y -= 1
		}
		if Vector2.Length(velocity) > 0 {
			velocity = Vector2.MulX(Vector2.Normalized(velocity), mvmt.Speed)
		}
		body.SetVelocity(velocity)
	}
}
func (s *InputSystem) Finalize(w *ecs.World) {}
