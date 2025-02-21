# Brackey's Platformer


https://github.com/user-attachments/assets/6b77f3e8-9790-409e-831d-ae22882f06e6



Brackey's released a recent tutorial for Godot on how to build a platform.

What I tried to do differently is to use [graphics.gd](https://github.com/grow-graphics/gd) library, which lets us use Go to build extensions for Godot.

I also tried to use the ECS pattern to build the game, using mlange42's [ECS](https://github.com/mlange-42/arche) library.

At this point its more of a proof of concept that its possible.

## Spinup Instructions


```
go install graphics.gd/cmd/gd@master
go mod tidy
gd run
```
