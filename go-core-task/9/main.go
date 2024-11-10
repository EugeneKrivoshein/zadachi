package main

import "fmt"

func dataProcessingPipeline(in <-chan uint8, out chan<- float64) {
	for num := range in {
		result := float64(num) * float64(num) * float64(num)
		out <- result
	}
	close(out)
}

func main() {
	in := make(chan uint8, 10)
	out := make(chan float64, 10)

	go dataProcessingPipeline(in, out)

	for i := uint8(1); i <= 5; i++ {
		in <- i
	}
	close(in)

	for result := range out {
		fmt.Println(result)
	}
}
