package advent2020_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/sakjur/advent2020"
)

func TestDay2(t *testing.T) {
	f, err := os.Open("testdata/day2_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader      io.Reader
		validLength int
		validIndex  int
	}{
		{
			reader:      bytes.NewBufferString("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc"),
			validLength: 2,
			validIndex:  1,
		},
		{
			reader:      f,
			validLength: 660,
			validIndex:  530,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d", i, tc.validLength), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := advent2020.StreamStrings(ctx, tc.reader)
			length, indexes := advent2020.ValidateLength(c)
			cancel()

			if length != tc.validLength {
				t.Errorf("Expected %d valid lines on length, got %d", tc.validLength, length)
			}
			if indexes != tc.validIndex {
				t.Errorf("Expected %d valid lines on indexes, got %d", tc.validIndex, indexes)
			}
		})
	}
}
