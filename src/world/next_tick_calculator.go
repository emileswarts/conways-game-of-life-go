package world

import "fmt"

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
