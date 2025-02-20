package main

import (
	"log/slog"

	"graphics.gd/classdb"
	"graphics.gd/startup"
)

func main() {
	PlayerControls.MoveRight = "ui_right"
	PlayerControls.MoveLeft = "ui_left"
	PlayerControls.MoveDown = "ui_down"
	PlayerControls.MoveUp = "ui_up"
	classdb.Register[Spawner]()
	classdb.Register[World]()
	startup.LoadingScene()
	slog.Info("main ready")
	startup.Scene()
}
