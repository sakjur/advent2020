package advent2020

import (
	"bytes"
	"reflect"
)

const (
	ferryOcean = '~'
	ferryFloor = '.'
	ferryEmpty = 'L'
	ferryTaken = '#'
)

type Ferry [][]byte

func (f Ferry) Equal(other Ferry) bool {
	return reflect.DeepEqual(f, other)
}

func (f Ferry) Tick(adj func(row, seat int) int, takenLimit int) Ferry {
	neue := make(Ferry, len(f))
	for i1, row := range f {
		neue[i1] = make([]byte, len(row))
		for i2, seat := range row {
			neue[i1][i2] = seat
			switch seat {
			case ferryEmpty:
				if adj(i1, i2) == 0 {
					neue[i1][i2] = ferryTaken
				}
			case ferryTaken:
				if adj(i1, i2) >= takenLimit {
					neue[i1][i2] = ferryEmpty
				}
			case ferryFloor:
				continue
			}
		}
	}
	return neue
}

func (f Ferry) Adjacents(row, seat int) int {
	adj := 0
	checks := map[int][]int{
		row - 1: {seat - 1, seat, seat + 1},
		row:     {seat - 1, seat + 1},
		row + 1: {seat - 1, seat, seat + 1},
	}
	for r, seats := range checks {
		for _, s := range seats {
			if f.IsOccupied(r, s) {
				adj++
			}
		}
	}
	return adj
}

func (f Ferry) Seen(row, seat int) int {
	seen := 0
	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, dir := range dirs {
	loop:
		for i := 1; ; i++ {
			switch f.Seat(row+i*dir[0], seat+i*dir[1]) {
			case ferryFloor:
				continue
			case ferryTaken:
				seen++
				break loop
			case ferryEmpty, ferryOcean:
				break loop
			}
		}
	}
	return seen
}

func (f Ferry) Seat(row, seat int) byte {
	if row < 0 || row >= len(f) {
		return ferryOcean
	}
	if seat < 0 || seat >= len(f[row]) {
		return ferryOcean
	}
	return f[row][seat]
}

func (f Ferry) IsOccupied(row, seat int) bool {
	return f.Seat(row, seat) == ferryTaken
}

func (f Ferry) OccupiedSeats() int {
	occ := 0
	for _, seats := range f {
		for _, seat := range seats {
			if seat == ferryTaken {
				occ++
			}
		}
	}
	return occ
}

func (f Ferry) String() string {
	return string(bytes.Join(f, []byte("\n")))
}

func SimulateAdj(f Ferry) int {
	var prev Ferry
	for !f.Equal(prev) {
		prev = f
		f = f.Tick(f.Adjacents, 4)
	}
	return f.OccupiedSeats()
}

func SimulateSeen(f Ferry) int {
	var prev Ferry
	for !f.Equal(prev) {
		prev = f
		f = f.Tick(f.Seen, 5)
	}
	return f.OccupiedSeats()
}

func ScanSeats(lines <-chan string) Ferry {
	f := Ferry{}
	for line := range lines {
		f = append(f, []byte(line))
	}
	return f
}
