package advent2020_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/sakjur/advent2020"
)

func TestDay3_Task1(t *testing.T) {
	fDemo, err := os.Open("testdata/day3_demo.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}
	fReal, err := os.Open("testdata/day3_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader   io.Reader
		hitTrees int
		allHits  int
	}{
		{
			reader:   fDemo,
			hitTrees: 7,
			allHits:  336,
		},
		{
			reader:   fReal,
			hitTrees: 220,
			allHits:  2138320800,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.hitTrees), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := advent2020.StreamStrings(ctx, tc.reader)
			m := advent2020.CreateMap(c)
			cancel()

			hits := advent2020.Travel(m, advent2020.XY{
				X: 3,
				Y: 1,
			})

			if tc.hitTrees != hits {
				t.Errorf("expected to hit %d trees, hit %d", tc.hitTrees, hits)
			}

			slopes := []advent2020.XY{
				{X: 1, Y: 1},
				{X: 3, Y: 1},
				{X: 5, Y: 1},
				{X: 7, Y: 1},
				{X: 1, Y: 2},
			}

			hits = 1
			for _, slope := range slopes {
				hits *= advent2020.Travel(m, slope)
			}

			if tc.allHits != hits {
				t.Errorf("expected to hit %d trees, hit %d", tc.allHits, hits)
			}
		})
	}
}
