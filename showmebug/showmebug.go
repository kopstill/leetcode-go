package showmebug

const (
	EAST  = "EAST"
	WEST  = "WEST"
	NORTH = "NORTH"
	SOUTH = "SOUTH"
)

func dirReduce(directions []string) []string {
	if len(directions) == 0 || len(directions) == 1 {
		return directions
	}

	var result []string
	for i := 0; i < len(directions); i++ {
		direction := directions[i]

		if i == 0 {
			nextDirection := directions[i+1]
			if (direction == EAST && nextDirection != WEST) ||
				(direction == WEST && nextDirection != EAST) ||
				(direction == SOUTH && nextDirection != NORTH) ||
				(direction == NORTH && nextDirection != SOUTH) {
				result = append(result, direction)
			}
			continue
		}

		if i == len(directions)-1 {
			preDirection := directions[i-1]
			if (direction == EAST && preDirection != WEST) ||
				(direction == WEST && preDirection != EAST) ||
				(direction == SOUTH && preDirection != NORTH) ||
				(direction == NORTH && preDirection != SOUTH) {
				result = append(result, direction)
			}
			break
		}

		if nextDirection, preDirection := directions[i+1], directions[i-1]; (direction == EAST && nextDirection != WEST && preDirection != WEST) ||
			(direction == WEST && nextDirection != EAST && preDirection != EAST) ||
			(direction == SOUTH && nextDirection != NORTH && preDirection != NORTH) ||
			(direction == NORTH && nextDirection != SOUTH && preDirection != SOUTH) {
			result = append(result, direction)
		}
	}

	return result
}
