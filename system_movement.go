package main

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
)

type MovementSystem struct {
	f *generic.Filter2[GodotNode, Movement]
}

func (s *MovementSystem) Initialize(w *ecs.World) {}
func (s *MovementSystem) Update(w *ecs.World) {
	q := s.f.Query(w)
	for q.Next() {
		node, _ := q.Get()
		body, ok := classdb.As[CharacterBody2D.Instance](node.Node)
		if !ok {
			continue
		}
		body.MoveAndSlide()
	}
}
func (s *MovementSystem) Finalize(w *ecs.World) {}
