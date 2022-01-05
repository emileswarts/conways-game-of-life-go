package world

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateDefaultWorld(t *testing.T) {
	var expected_world = [][map]
	expected_world := [
		{ "x": 0, "y": 0, "alive": false }
		{ "x": 0, "y": 1, "alive": false }
		{ "x": 0, "y": 2, "alive": false }
		{ "x": 1, "y": 0, "alive": false }
		{ "x": 1, "y": 1, "alive": false }
		{ "x": 1, "y": 2, "alive": false }
		{ "x": 2, "y": 0, "alive": false }
		{ "x": 2, "y": 1, "alive": false }
		{ "x": 2, "y": 2, "alive": false }
	]

	result := generate()
	fmt.Println(result)
	// assert.Equal(t, result, expected_world)
}