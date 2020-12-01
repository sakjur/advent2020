package advent2020_test

import "testing"

func elementOf(t testing.TB, expected []int, nums ...int) {
	notFound := []int{}

	outer: for _, num := range nums {
		for _, expect := range expected {
			if num == expect {
				continue outer
			}
		}
		notFound = append(notFound, num)
	}

	if len(notFound) != 0 {
		t.Errorf("%v ⊄ %v.", notFound, expected)
	}
}
