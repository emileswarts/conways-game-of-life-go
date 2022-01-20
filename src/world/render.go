package world

type Cell struct {
	X     int
	Y     int
	Alive bool
}

func render(world_data [16]Cell) string {
	var response string

	for i, cell := range world_data {

		if i%4 == 0 {
			response = response + "\n"
		}

		if cell.Alive == true {
			response = response + "x"
		} else {
			response = response + "."
		}
	}

	return response
}
