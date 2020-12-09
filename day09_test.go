package advent2020_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/sakjur/advent2020"
)

func TestDay9(t *testing.T) {
	fDemo := strings.NewReader("35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576\n")
	fReal, err := os.Open("testdata/day9_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader     io.Reader
		preamble   int
		firstBreak int
		weakness   int
	}{
		{
			reader:     fDemo,
			preamble:   5,
			firstBreak: 127,
			weakness:   62,
		},
		{
			reader:     fReal,
			preamble:   25,
			firstBreak: 675280050,
			weakness:   96081673,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.firstBreak), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := advent2020.StreamInts(ctx, tc.reader)
			ops, weakness, ok := advent2020.XMASCodeBreak(c, tc.preamble)
			cancel()
			if ok {
				t.Errorf("[unexpected result] application finished successfully")
			}
			if ops != tc.firstBreak {
				t.Errorf("expected first failing n to be %d, got %d", tc.firstBreak, ops)
			}
			if weakness != tc.weakness {
				t.Errorf("expected weakness to be %d, got %d", tc.weakness, weakness)
			}
		})
	}
}
