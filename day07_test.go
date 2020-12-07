package advent2020_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/sakjur/advent2020"
)

func TestDay7(t *testing.T) {
	fDemo, err := os.Open("testdata/day7_demo.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}
	fReal, err := os.Open("testdata/day7_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader     io.Reader
		parentBags int
		childBags  int
	}{
		{
			reader:     fDemo,
			parentBags: 4,
			childBags:  32,
		},
		{
			reader:     fReal,
			parentBags: 211,
			childBags:  12414,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.parentBags), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := advent2020.StreamStrings(ctx, tc.reader)
			bags, err := advent2020.AllBags(c)
			cancel()
			if err != nil {
				t.Fatalf("failed to read bags: %v", err)
			}

			parents := advent2020.ParentsOf("shiny gold bags", bags)
			if len(parents) != tc.parentBags {
				t.Errorf("expected %d parent bags to 'shiny gold bags', got %d", tc.parentBags, len(parents))
			}

			children := advent2020.ChildCount("shiny gold bags", bags)
			if children != tc.childBags {
				t.Errorf("expected %d child bags to 'shiny gold bags', got %d", tc.childBags, children)
			}
		})
	}
}
