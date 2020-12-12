package advent2020

import (
	"fmt"
	"strconv"
)

func ManhattanDistance(c []string) int {
	heading := 90
	east, north := 0, 0
	for _, line := range c {
		instr := line[0]
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch instr {
		case 'F':
			east, north = moveForward(heading, val, east, north)
		case 'N':
			east, north = moveForward(0, val, east, north)
		case 'E':
			east, north = moveForward(90, val, east, north)
		case 'S':
			east, north = moveForward(180, val, east, north)
		case 'W':
			east, north = moveForward(270, val, east, north)
		case 'R':
			heading = (heading + val) % 360
		case 'L':
			heading = (heading + 360 - val) % 360
		}
	}
	return abs(east) + abs(north)
}

func WaypointMove(c []string) int {
	east, north := 0, 0
	wpE, wpN := 10, 1
	for _, line := range c {
		instr := line[0]
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch instr {
		case 'F':
			east, north = east+val*wpE, north+val*wpN
		case 'N':
			wpE, wpN = moveForward(0, val, wpE, wpN)
		case 'E':
			wpE, wpN = moveForward(90, val, wpE, wpN)
		case 'S':
			wpE, wpN = moveForward(180, val, wpE, wpN)
		case 'W':
			wpE, wpN = moveForward(270, val, wpE, wpN)
		case 'R':
			wpE, wpN = rotate(val, wpE, wpN)
		case 'L':
			wpE, wpN = rotate((360-val)%360, wpE, wpN)
		}
	}
	return abs(east) + abs(north)
}

func rotate(degrees, east, north int) (int, int) {
	switch degrees {
	case 0:
		break
	case 90:
		east, north = north, -east
	case 180:
		east, north = -east, -north
	case 270:
		east, north = -north, east
	default:
		panic(fmt.Errorf("unknown rotation %d", degrees))
	}
	return east, north
}

func moveForward(heading, n, east, north int) (int, int) {
	switch heading {
	case 0:
		north += n
	case 90:
		east += n
	case 180:
		north -= n
	case 270:
		east -= n
	default:
		panic(fmt.Errorf("unknown heading %d", heading))
	}
	return east, north
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}
