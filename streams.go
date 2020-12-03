package advent2020

import (
	"bufio"
	"context"
	"io"
	"log"
	"strconv"
)

func StreamInts(ctx context.Context, r io.Reader) <-chan int {
	channel := make(chan int, 1024)

	go func() {
		streamLine(ctx, r, func(token string) {
			number, err := strconv.Atoi(token)
			if err != nil {
				log.Println("got non-numeric line: ", err)
				return
			}

			channel <- number
		})
		close(channel)
	}()

	return channel
}

func StreamStrings(ctx context.Context, r io.Reader) <-chan string {
	channel := make(chan string, 1024)

	go func() {
		streamLine(ctx, r, func(token string) {
			channel <- token
		})
		close(channel)
	}()

	return channel
}

func streamLine(ctx context.Context, r io.Reader, outFn func(string)) {
	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanLines)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if read := scan.Scan(); !read {
				if scan.Err() != nil {
					log.Printf("got error: %v\n", scan.Err())
				}
				return
			}

			outFn(scan.Text())
		}
	}
}
