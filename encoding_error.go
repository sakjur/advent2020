package advent2020

import "math"

func XMASCodeBreak(c <-chan int, preambleSize int) (int, int, bool) {
	fullList := make([]int, 0)
	preamble := make([]int, 0, preambleSize)
	target := 0
	weakness := 0

	for n := range c {
		fullList = append(fullList, n)
		if len(preamble) != preambleSize {
			preamble = append(preamble, n)
			continue
		}

		preambleC := make(chan int)
		go func() {
			for _, num := range preamble {
				preambleC <- num
			}
			close(preambleC)
		}()
		a, b := FindAddPair(preambleC, n)

		if a == 0 && b == 0 {
			target = n
		}

		preamble = append(preamble, n)
		preamble = preamble[1:]
	}

outer:
	for i := range fullList {
		for i2 := 1; i-i2 >= 0; i2++ {
			tot, min, max := sumInts(fullList[i-i2 : i])
			if tot > target {
				break
			}
			if tot == target {
				weakness = min + max
				break outer
			}
		}
	}

	return target, weakness, target == 0
}

func sumInts(ns []int) (int, int, int) {
	total := 0
	min := math.MaxInt64
	max := math.MinInt64
	for _, n := range ns {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		total += n
	}
	return total, min, max
}
