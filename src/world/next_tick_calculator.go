package world

import "fmt"

/*
Any live cell with fewer than two live neighbours dies (referred to as underpopulation or exposure[1]).
Any live cell with more than three live neighbours dies (referred to as overpopulation or overcrowding).
Any live cell with two or three live neighbours lives, unchanged, to the next generation.
Any dead cell with exactly three live neighbours will come to life.
*/

func calculate_neighbours(cell Cell, cells [16]Cell) []Cell {
	var neighbours []Cell

	for _, v := range cells {
		if (v.X == cell.X || v.X == cell.X-1 || v.X == cell.X+1) && (v.Y == cell.Y || v.Y == cell.Y-1 || v.Y == cell.Y+1) {
			if v != cell {
				neighbours = append(neighbours, v)
			}
		}
	}

	return neighbours
}

func live_neighbours(neighbours []Cell) []Cell {
	var live_cells []Cell

	for _, cell := range neighbours {
		if cell.Alive == true {
			live_cells = append(live_cells, cell)
		}
	}

	return live_cells
}

func calculate(cells [16]Cell) [16]Cell {
	for _, cell := range cells {
		fmt.Println(cell)
		fmt.Println(live_neighbours(calculate_neighbours(cell, cells)))
	}

	return cells
}
