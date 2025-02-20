package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

// Movement for any entity in the world that can move
type Movement struct {
	Position Vector2.XY
	Velocity Vector2.XY
	Speed    Float.X
}

// GodotNode is a wrapper for a Godot node
// You'll be able to cast it to the correct type in your systems
type GodotNode struct {
	Node Node.Instance
}
