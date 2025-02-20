package main

import (
	"fmt"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/ecs/event"
	"github.com/mlange-42/arche/generic"
	"github.com/mlange-42/arche/listener"
	"graphics.gd/classdb"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/variant/Float"
)

type World struct {
	classdb.Extension[World, Node.Instance] `gd:"NXRWorld"`
	Player                                  PackedScene.Instance
	Spawner                                 *Spawner
	Camera                                  Camera3D.Instance
	model                                   *model.Model
}

func (w *World) Ready() {
	m := model.New()
	l := listener.NewCallback(func(w *ecs.World, ee ecs.EntityEvent) { fmt.Printf("EntityEvent: %+v", ee) }, event.All)
	d := listener.NewDispatch(&l)
	m.World.SetListener(&d)

	m.AddSystem(&SpawnSystem{w.Spawner})
	m.AddSystem(&MovementSystem{generic.NewFilter2[GodotNode, Movement]()})
	m.AddSystem(&InputSystem{generic.NewFilter2[GodotNode, Movement]()})
	m.Initialize()
	w.model = m
}

func (w *World) Process(delta Float.X) {
	w.model.Update()
}
