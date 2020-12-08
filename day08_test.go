package advent2020_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/sakjur/advent2020"
)

func TestDay8(t *testing.T) {
	fDemo, err := os.Open("testdata/day8_demo.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}
	fReal, err := os.Open("testdata/day8_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader   io.Reader
		crashAcc int
		fixedAcc int
	}{
		{
			reader:   fDemo,
			crashAcc: 5,
			fixedAcc: 8,
		},
		{
			reader:   fReal,
			crashAcc: 1087,
			fixedAcc: 0,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.crashAcc), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := advent2020.StreamStrings(ctx, tc.reader)
			ops, err := advent2020.GameConsoleOps(c)
			cancel()
			if err != nil {
				t.Fatalf("failed to read bags: %v", err)
			}

			acc, ok := advent2020.GameConsoleDetectLoop(ops)
			if ok {
				t.Errorf("[unexpected result] application finished successfully")
			}
			if acc != tc.crashAcc {
				t.Errorf("expected accumulator on loop entry to be %d, got %d", tc.crashAcc, acc)
			}

			acc, ok = advent2020.GameConsoleAutoPatcher(ops)
			if !ok {
				t.Error("could not fix game boot code")
			}
			if acc != tc.fixedAcc {
				t.Errorf("expected accumulator for fixed application to be %d, got %d", tc.fixedAcc, acc)
			}
		})
	}
}
