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
	"graphics.gd/classdb/TileMapLayer"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/String"
	"graphics.gd/variant/Vector2"
)

// World is the ECS world for the game
type World struct {
	classdb.Extension[World, Node.Instance] `gd:"NXRWorld"`
	Map                                     PackedScene.Instance
	Player                                  PackedScene.Instance
	Spawner                                 *Spawner
	Camera                                  Camera3D.Instance
	model                                   *model.Model
}

func (w *World) OnPositionUpdate() Callable.Function {
	return Callable.New(func(id String.Readable, pos Vector2.XY) {
		fmt.Printf("PositionUpdate: %s %v", id, pos)
	})
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
	gravity := &Gravity{Value: 1500}
	levelMap := w.Map.Instantiate()
	tileMap, ok := classdb.As[TileMapLayer.Instance](Node.Instance(levelMap))
	if !ok {
		panic("could not get TileMapLayer from map")
	}
	w.Super().AddChild(tileMap.AsNode())
	ecs.AddResource(&m.World, &Map{tileMap})
	ecs.AddResource(&m.World, gravity)
	ecs.AddResource(&m.World, &DeltaTime{Delta: 0})
	w.model = m
}

type Map struct {
	TileMapLayer.Instance
}

type DeltaTime struct {
	Delta Float.X
}

type Gravity struct {
	Value Float.X
}

func (w *World) Process(delta Float.X) {
	deltaResource := ecs.GetResource[DeltaTime](&w.model.World)
	deltaResource.Delta = delta
	w.model.Update()
}
