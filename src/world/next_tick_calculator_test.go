package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateNextTick(t *testing.T) {
	world_data := [16]Cell{
		{X: 0, Y: 0, Alive: false},
		{X: 0, Y: 1, Alive: false},
		{X: 0, Y: 2, Alive: false},
		{X: 0, Y: 3, Alive: false},
		{X: 1, Y: 0, Alive: false},
		{X: 1, Y: 1, Alive: false},
		{X: 1, Y: 2, Alive: false},
		{X: 1, Y: 3, Alive: false},
		{X: 2, Y: 0, Alive: false},
		{X: 2, Y: 1, Alive: false},
		{X: 2, Y: 2, Alive: false},
		{X: 2, Y: 3, Alive: true},
		{X: 3, Y: 0, Alive: false},
		{X: 3, Y: 1, Alive: false},
		{X: 3, Y: 2, Alive: false},
		{X: 3, Y: 3, Alive: false},
	}

	next_tick_world_data := [16]Cell{
		{X: 0, Y: 0, Alive: false},
		{X: 0, Y: 1, Alive: false},
		{X: 0, Y: 2, Alive: false},
		{X: 0, Y: 3, Alive: false},
		{X: 1, Y: 0, Alive: false},
		{X: 1, Y: 1, Alive: false},
		{X: 1, Y: 2, Alive: false},
		{X: 1, Y: 3, Alive: false},
		{X: 2, Y: 0, Alive: false},
		{X: 2, Y: 1, Alive: false},
		{X: 2, Y: 2, Alive: false},
		{X: 2, Y: 3, Alive: true},
		{X: 3, Y: 0, Alive: false},
		{X: 3, Y: 1, Alive: false},
		{X: 3, Y: 2, Alive: false},
		{X: 3, Y: 3, Alive: false},
	}
	result := calculate(world_data)
	assert.Equal(t, result, next_tick_world_data)
}
