package advent2020

func FindAddPair(numbers <-chan int, target int) (int, int) {
	existing := map[int]struct{}{}
	for number := range numbers {
		if _, exists := existing[target-number]; exists {
			return number, target - number
		}

		existing[number] = struct{}{}
	}
	return 0, 0
}

func FindAddTrio(numbers <-chan int, target int) (int, int, int) {
	type pair struct {
		a int
		b int
	}

	existing := []int{}
	pairs := map[int]pair{}
	for number := range numbers {
		innerTarget := target - number
		if pair, exists := pairs[innerTarget]; exists {
			return pair.a, pair.b, number
		}

		for _, n := range existing {
			sum := n + number
			if sum > target {
				continue
			}
			pairs[sum] = pair{a: n, b: number}
		}

		existing = append(existing, number)
	}

	return 0, 0, 0
}
