package advent2020

import (
	"bufio"
	"context"
	"io"
	"log"
	"strconv"
)

func FindAddPair(numbers <-chan int, target int) (int, int) {
	existing := map[int]struct{}{}
	for number := range numbers {
		if _, exists := existing[target - number]; exists {
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
			sum := n+number
			if sum > target {
				continue
			}
			pairs[sum] = pair{a: n, b: number}
		}

		existing = append(existing, number)
	}

	return 0, 0, 0
}

func StreamInts(ctx context.Context, r io.Reader) (<- chan int) {
	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanLines)
	channel := make(chan int, 1024)

	go func() {
		loop: for {
			select {
			case <-ctx.Done():
				break loop
			default:
				if read := scan.Scan(); !read {
					if scan.Err() != nil {
						log.Printf("got error: %v\n", scan.Err())
					}
					break loop
				}

				token := scan.Text()
				number, err := strconv.Atoi(token)
				if err != nil {
					log.Println("got non-numeric line: ", err)
					continue
				}

				channel <- number
			}
		}
		close(channel)
	}()

	return channel
}
