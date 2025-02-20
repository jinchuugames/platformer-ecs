package main

import "github.com/mlange-42/arche/ecs"

// SpawnSystem is a system that spawns the player character
// It has an embedded spawner as the system handles the lifecycle only and runs the spawner on initialization.
type SpawnSystem struct {
	Spawner *Spawner
}

func (s *SpawnSystem) Initialize(w *ecs.World) { s.Spawner.SpawnPlayer(w) }
func (s *SpawnSystem) Update(w *ecs.World)     {}
func (s *SpawnSystem) Finalize(w *ecs.World)   {}
