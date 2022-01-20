package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderDefaultWorld(t *testing.T) {
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
		{X: 2, Y: 3, Alive: false},
		{X: 3, Y: 0, Alive: false},
		{X: 3, Y: 1, Alive: false},
		{X: 3, Y: 2, Alive: false},
		{X: 3, Y: 3, Alive: false},
	}

	expected_result := `
....
....
....
....`

	result := render(world_data)
	assert.Equal(t, result, expected_result)
}

func TestRenderWorldWithOneLiveCell(t *testing.T) {
	world_data := [16]Cell{
		{X: 0, Y: 0, Alive: false},
		{X: 0, Y: 1, Alive: true},
		{X: 0, Y: 2, Alive: false},
		{X: 0, Y: 3, Alive: false},
		{X: 1, Y: 0, Alive: false},
		{X: 1, Y: 1, Alive: false},
		{X: 1, Y: 2, Alive: false},
		{X: 1, Y: 3, Alive: false},
		{X: 2, Y: 0, Alive: false},
		{X: 2, Y: 1, Alive: false},
		{X: 2, Y: 2, Alive: false},
		{X: 2, Y: 3, Alive: false},
		{X: 3, Y: 0, Alive: true},
		{X: 3, Y: 1, Alive: false},
		{X: 3, Y: 2, Alive: false},
		{X: 3, Y: 3, Alive: false},
	}

	expected_result := `
.x..
....
....
x...`

	result := render(world_data)
	assert.Equal(t, result, expected_result)
}
