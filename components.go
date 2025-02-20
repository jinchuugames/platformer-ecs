package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

type Movement struct {
	Position Vector2.XY
	Velocity Vector2.XY
	Speed    Float.X
}
type GodotNode struct {
	Node Node.Instance
}
