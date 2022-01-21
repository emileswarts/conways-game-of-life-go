package main

/*
 */

func calculate_next_tick_cell_state(cell Cell, live_neighbours []Cell) Cell {
	live_neighbour_count := len(live_neighbours)

	if live_neighbour_count < 2 || live_neighbour_count > 3 {
		cell.Alive = false
	} else if live_neighbour_count == 3 {
		cell.Alive = true
	}

	return cell
}

func calculate_neighbours(cell Cell, cells [25]Cell) []Cell {
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

func calculate(cells [25]Cell) [25]Cell {
	var new_cells [25]Cell
	var new_cell Cell

	for i, cell := range cells {
		new_cell = calculate_next_tick_cell_state(cell, live_neighbours(calculate_neighbours(cell, cells)))
		new_cells[i] = new_cell
	}

	return new_cells
}
