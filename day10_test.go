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

func TestDay10(t *testing.T) {
	fDemo1 := strings.NewReader("16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4\n")
	fDemo2 := strings.NewReader("28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3\n")
	fReal, err := os.Open("testdata/day10_task1.txt")
	if err != nil {
		t.Fatalf("got error: %v\n", err)
	}

	tests := []struct {
		reader       io.Reader
		jolt1        int
		jolt3        int
		arrangements int
	}{
		{
			reader:       fDemo1,
			jolt1:        7,
			jolt3:        5,
			arrangements: 8,
		},
		{
			reader:       fDemo2,
			jolt1:        22,
			jolt3:        10,
			arrangements: 19208,
		},
		{
			reader:       fReal,
			jolt1:        65,
			jolt3:        38,
			arrangements: 1973822685184,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("[%d] %d*%d=%d", i, tc.jolt1, tc.jolt3, tc.jolt1*tc.jolt3), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := advent2020.StreamInts(ctx, tc.reader)
			jolt1, _, jolt3, arrangements := advent2020.Jolts(c)
			assert.Equal(t, tc.jolt1, jolt1)
			assert.Equal(t, tc.jolt3, jolt3)
			assert.Equal(t, tc.arrangements, arrangements)
		})
	}
}
