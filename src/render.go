package main

func render(world_data [25]Cell) string {
	var response string

	for i, cell := range world_data {

		if i%5 == 0 {
			response = response + "\n"
		}

		if cell.Alive == true {
			response = response + " x "
		} else {
			response = response + " . "
		}
	}

	return response
}
