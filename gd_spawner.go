package main

import (
	"log/slog"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/variant/Vector2"
)

// Spawner will spawn a player in the world.
type Spawner struct {
	classdb.Extension[Spawner, Node.Instance] `gd:"NXRSpawner"`
	Player                                    PackedScene.Instance
}

// AsNode is required to show the node in the Godot editor under the parent World.
// Without this you'll get nil pointer errors.
func (s *Spawner) AsNode() Node.Instance { return s.Super().AsNode() }
func (s *Spawner) SpawnPlayer(w *ecs.World) {
	player := s.Player.Instantiate()
	s.Super().AddChild(player)
	p := &GodotNode{player}
	mvmt := &Movement{
		Position: Vector2.XY{X: 0, Y: 0},
		Velocity: Vector2.XY{X: 0, Y: 0},
		Speed:    100,
	}
	body, ok := classdb.As[CharacterBody2D.Instance](Node.Instance(player))
	if !ok {
		slog.Warn("could not get CharacterBody2D from player")
		return
	}
	body.SetVelocity(mvmt.Velocity)
	m := generic.NewMap2[GodotNode, Movement](w)
	m.NewWith(p, mvmt)
}
