package main

import (
	"fmt"
)

func MergeChannels(channels ...<-chan int) <-chan int {
	res := make(chan int)

	go func() {
		for _, ch := range channels {
			for num := range ch {
				res <- num
			}
		}
		close(res)
	}()

	return res
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	go func() {
		for i := 10; i < 15; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	merge := MergeChannels(ch1, ch2, ch3)
	for num := range merge {
		fmt.Println(num)
	}
}
