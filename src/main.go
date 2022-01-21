package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Cell struct {
	X     int
	Y     int
	Alive bool
}

func main() {
	var world [25]Cell

	start_world := [25]Cell{
		{X: 0, Y: 0, Alive: false},
		{X: 0, Y: 1, Alive: false},
		{X: 0, Y: 2, Alive: false},
		{X: 0, Y: 3, Alive: true},
		{X: 0, Y: 4, Alive: true},
		{X: 1, Y: 0, Alive: true},
		{X: 1, Y: 1, Alive: true},
		{X: 1, Y: 2, Alive: true},
		{X: 1, Y: 3, Alive: true},
		{X: 1, Y: 4, Alive: true},
		{X: 2, Y: 0, Alive: false},
		{X: 2, Y: 1, Alive: false},
		{X: 2, Y: 2, Alive: false},
		{X: 2, Y: 3, Alive: false},
		{X: 2, Y: 4, Alive: false},
		{X: 3, Y: 0, Alive: false},
		{X: 3, Y: 1, Alive: false},
		{X: 3, Y: 2, Alive: false},
		{X: 3, Y: 3, Alive: false},
		{X: 3, Y: 4, Alive: false},
		{X: 4, Y: 0, Alive: false},
		{X: 4, Y: 1, Alive: false},
		{X: 4, Y: 2, Alive: false},
		{X: 4, Y: 3, Alive: false},
		{X: 4, Y: 4, Alive: false},
	}

	world = start_world
	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println(render(world))
		world = calculate(world)
		time.Sleep(1 * time.Second)
	}
}
