package advent2020

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
	fReal, err := os.Open("testdata/day6_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader   io.Reader
		countAny int
		countAll int
	}{
		{
			reader:   strings.NewReader("abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"),
			countAny: 11,
			countAll: 6,
		},
		{
			reader:   fReal,
			countAny: 6430,
			countAll: 3125,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.countAny), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := StreamStrings(ctx, tc.reader)
			forms := CustomForms(c)
			cancel()

			countAny, countAll := 0, 0
			for _, form := range forms {
				countAny += form.Any()
				countAll += form.All()
			}
			if countAny != tc.countAny {
				t.Errorf("expected total countAny to be '%d', got '%d'", tc.countAny, countAny)
			}
			if countAll != tc.countAll {
				t.Errorf("expected total countAny to be '%d', got '%d'", tc.countAll, countAll)
			}
		})
	}
}
