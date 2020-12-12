package advent2020_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sakjur/advent2020"
)

func TestDay12(t *testing.T) {
	fDemo := strings.NewReader("F10\nN3\nF7\nR90\nF11\n")
	fReal, err := os.Open("testdata/day12_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader            io.Reader
		manhattanDistance int
		waypointDistance  int
	}{
		{
			reader:            fDemo,
			manhattanDistance: 25,
			waypointDistance:  286,
		},
		{
			reader:            fReal,
			manhattanDistance: 1565,
			waypointDistance:  78883,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.manhattanDistance), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := advent2020.StreamStrings(ctx, tc.reader)
			instr := []string{}
			for line := range c {
				instr = append(instr, line)
			}

			distance := advent2020.ManhattanDistance(instr)
			assert.Equal(t, tc.manhattanDistance, distance)

			distance = advent2020.WaypointMove(instr)
			assert.Equal(t, tc.waypointDistance, distance)
		})
	}
}
