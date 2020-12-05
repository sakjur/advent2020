package advent2020

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinaryBoard(t *testing.T) {
	tests := []struct {
		input string
		row   int
		seat  int
	}{
		{"BFFFBBFRRR", 70, 7},
		{"FFFBBBFRRR", 14, 7},
		{"BBFFBBFRLL", 102, 4},
		{"FBFBBFFRLR", 44, 5},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			row, seat, err := BinaryBoard(tc.input, 0, 127, 0, 7)
			require.NoError(t, err)
			assert.Equal(t, tc.row, row)
			assert.Equal(t, tc.seat, seat)
		})
	}
}

func TestDay5(t *testing.T) {
	f, err := os.Open("testdata/day5_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	expectedMaxSeatID := 947
	seats := make([]bool, expectedMaxSeatID+1)

	c := StreamStrings(ctx, f)
	maxSeatID := 0
	for pass := range c {
		row, seat, err := BinaryBoard(pass, 0, 127, 0, 7)
		require.NoError(t, err)
		seatID := row*8 + seat
		if maxSeatID < seatID {
			maxSeatID = seatID
		}
		seats[seatID] = true
	}

	assert.Equal(t, expectedMaxSeatID, maxSeatID)

	mySeat := 0
	for seat, exists := range seats {
		if !exists && seat-1 >= 0 && seats[seat-1] && seat+1 < len(seats) && seats[seat+1] {
			mySeat = seat
			break
		}
	}
	assert.Equal(t, 636, mySeat)
}
