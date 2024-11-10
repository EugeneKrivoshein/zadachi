package main

import (
	"fmt"
	"time"
)

type CustomWaitGroup struct {
	semaphore chan struct{}
	counter   int
	done      chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		semaphore: make(chan struct{}, 1),
		done:      make(chan struct{}),
	}
}

func (wg *CustomWaitGroup) Add(num int) {
	wg.semaphore <- struct{}{}
	defer func() { <-wg.semaphore }()

	wg.counter += num
	if wg.counter == 0 {
		close(wg.done)
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	<-wg.done
}

func main() {
	wg := NewCustomWaitGroup()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			fmt.Printf("горутина %d запущена\n", index)
			time.Sleep(time.Second * 1)
			fmt.Printf("горутина %d завершилась\n", index)
		}(i)
	}
	wg.Wait()
	fmt.Println("все горутины завершились")
}
