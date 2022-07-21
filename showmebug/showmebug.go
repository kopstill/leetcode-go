/*
	https://www.showmebug.com/
*/
package showmebug

type DirType string

const (
	EAST    DirType = "EAST"
	WEST    DirType = "WEST"
	NORTH   DirType = "NORTH"
	SOUTH   DirType = "SOUTH"
	UNKNOWN DirType = "UNKNOWN"
)

func dirReduce(directions []string) []string {
	if len(directions) == 0 || len(directions) == 1 {
		return directions
	}

	for i := 0; i < len(directions)-1; i++ {
		cuDirection := directions[i]
		reDirection := getReDirection(cuDirection)
		neDirection := directions[i+1]

		if reDirection == neDirection {
			return dirReduce(append(directions[:i], directions[i+2:]...))
		}
	}

	return directions
}

func getReDirection[T string](direction T) T {
	switch direction {
	case T(EAST):
		return T(WEST)
	case T(WEST):
		return T(EAST)
	case T(SOUTH):
		return T(NORTH)
	case T(NORTH):
		return T(SOUTH)
	default:
		return T(UNKNOWN)
	}
}
