package main

type position struct {
    X int
    Y int
}

type World struct {
    Pos position
}

func createWorld() *World {
    var world *World = new(World)
    world.Pos.X = 0
    world.Pos.Y = 0

    return world;
}

func (world *World) Update(cmd string) interface{} {
    if(cmd == "l") {
            world.Pos.X -= 10;
        } else if(cmd == "r") {
            world.Pos.X += 10;
        } else if(cmd == "u") {
            world.Pos.Y -= 10;
        } else {
            world.Pos.Y += 10;
        }

    return world.Pos
}