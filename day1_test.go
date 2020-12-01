package advent2020_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sakjur/advent2020"
	"io"
	"os"
	"testing"
)

func TestDay1_Task1(t *testing.T) {
	f, err := os.Open("testdata/day1_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct{
		target int
		reader io.Reader
		expectedFirst int
		expectedSecond int
	}{
		{
			target:         2020,
			reader:         bytes.NewBufferString("1721\n979\n366\n299\n675\n1456"),
			expectedFirst:  1721,
			expectedSecond: 299,
		},
		{
			target: 2020,
			reader: f,
			expectedFirst: 473,
			expectedSecond: 1547,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d = %d + %d", tc.target, tc.expectedFirst, tc.expectedSecond), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := advent2020.StreamInts(ctx, tc.reader)
			a, b := advent2020.FindAddPair(c, tc.target)
			cancel()

			if a + b != tc.target {
				t.Errorf("a = %d, b = %d. Expected a+b=%d, got %d", a, b, tc.target, a+b)
			}
			elementOf(t, []int{tc.expectedFirst, tc.expectedSecond}, a, b)

			t.Logf("%d = %d * %d", a*b, a, b)
		})
	}

}

func TestDay1_Task2(t *testing.T) {
	f, err := os.Open("testdata/day1_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		target         int
		reader         io.Reader
		expectedFirst  int
		expectedSecond int
		expectedThird  int
	}{
		{
			target:         2020,
			reader:         bytes.NewBufferString("1721\n979\n366\n299\n675\n1456"),
			expectedFirst:  979,
			expectedSecond: 366,
			expectedThird:  675,
		},
		{
			target:         2020,
			reader:         f,
			expectedFirst:  1433,
			expectedSecond: 365,
			expectedThird:  222,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d = %d + %d + %d", tc.target, tc.expectedFirst, tc.expectedSecond, tc.expectedThird), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			channel := advent2020.StreamInts(ctx, tc.reader)
			a, b, c := advent2020.FindAddTrio(channel, tc.target)
			cancel()

			if a+b+c != tc.target {
				t.Errorf("a = %d, b = %d, c = %d. Expected a+b+c=%d, got %d", a, b, c, tc.target, a+b+c)
			}
			elementOf(t, []int{tc.expectedFirst, tc.expectedSecond, tc.expectedThird}, a, b, c)

			t.Logf("%d = %d * %d * %d", a*b*c, a, b, c)
		})
	}
}