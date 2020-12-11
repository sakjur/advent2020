package advent2020

import (
	"fmt"
	"sort"
)

func Jolts(values <-chan int) (jolt1 int, jolt2 int, jolt3 int, arrangements int) {
	adapters := []int{}
	for v := range values {
		adapters = append(adapters, v)
	}

	sort.Ints(adapters)

	current := 0
	for _, a := range adapters {
		switch a - current {
		case 1:
			jolt1++
		case 2:
			jolt2++
		case 3:
			jolt3++
		default:
			panic(fmt.Errorf("unexpected difference of %d", a-current))
		}
		current = a
	}
	jolt3++

	arrangementsTo := map[int]int{}

	for i, a := range adapters {
		paths := 0
		if a <= 3 { // can be reached from start
			paths++
		}

		start := i - 3
		if start < 0 {
			start = 0
		}
		for _, ancestor := range adapters[start:i] {
			if a-ancestor <= 3 {
				paths += arrangementsTo[ancestor]
			}
		}
		arrangementsTo[a] = paths
	}

	arrangements = arrangementsTo[adapters[len(adapters)-1]]

	return
}
