package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateInt() <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ch := make(chan int)
	go func() {
		for {
			ch <- r.Intn(100)
		}
	}()
	return ch
}

func main() {
	randChan := GenerateInt()

	for i := 0; i < 10; i++ {
		fmt.Println(<-randChan)
	}
}
