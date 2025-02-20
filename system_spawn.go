package main

import "github.com/mlange-42/arche/ecs"
type SpawnSystem struct {
	Spawner *Spawner
}

func (s *SpawnSystem) Initialize(w *ecs.World) { s.Spawner.SpawnPlayer(w) }
func (s *SpawnSystem) Update(w *ecs.World)     {}
func (s *SpawnSystem) Finalize(w *ecs.World)   {}
