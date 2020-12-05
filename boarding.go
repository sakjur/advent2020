package advent2020

import (
	"fmt"
	"math"
)

func BinaryBoard(input string, rowStart, rowEnd, seatStart, seatEnd int) (int, int, error) {
	if len(input) == 0 {
		if rowStart != rowEnd {
			return 0, 0, fmt.Errorf("unclear row: %d-%d", rowStart, rowEnd)
		}
		if seatStart != seatEnd {
			return 0, 0, fmt.Errorf("unclear seat: %d-%d", seatStart, seatEnd)
		}
		return rowStart, seatStart, nil
	}

	switch input[0] {
	case 'F':
		newRow := updatePosition(rowStart, rowEnd, math.Floor)
		return BinaryBoard(input[1:], rowStart, newRow, seatStart, seatEnd)
	case 'B':
		newRow := updatePosition(rowStart, rowEnd, math.Ceil)
		return BinaryBoard(input[1:], newRow, rowEnd, seatStart, seatEnd)
	case 'R':
		newSeat := updatePosition(seatStart, seatEnd, math.Ceil)
		return BinaryBoard(input[1:], rowStart, rowEnd, newSeat, seatEnd)
	case 'L':
		newSeat := updatePosition(seatStart, seatEnd, math.Floor)
		return BinaryBoard(input[1:], rowStart, rowEnd, seatStart, newSeat)
	}

	return 0, 0, fmt.Errorf("bad input: %s", input)
}

func updatePosition(start, end int, roundFn func(float64) float64) int {
	return start + int(roundFn(float64(end-start)/2.0))
}
