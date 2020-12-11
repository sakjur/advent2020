package advent2020_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sakjur/advent2020"
)

func TestDay11(t *testing.T) {
	fDemo, err := os.Open("testdata/day11_demo.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}
	fReal, err := os.Open("testdata/day11_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader       io.Reader
		occupiedAdj  int
		occupiedSeen int
	}{
		{
			reader:       fDemo,
			occupiedAdj:  37,
			occupiedSeen: 26,
		},
		{
			reader:       fReal,
			occupiedAdj:  2299,
			occupiedSeen: 2047,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.occupiedAdj), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := advent2020.StreamStrings(ctx, tc.reader)
			f := advent2020.ScanSeats(c)
			occupiedAdj := advent2020.SimulateAdj(f)
			assert.Equal(t, tc.occupiedAdj, occupiedAdj)

			occupiedSeen := advent2020.SimulateSeen(f)
			assert.Equal(t, tc.occupiedSeen, occupiedSeen)
		})
	}
}
