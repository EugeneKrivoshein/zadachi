package main

import (
	"testing"
	"time"
)

func TestGenerateInt(t *testing.T) {
	tests := []struct {
		name     string
		numCount int
	}{
		{
			name:     "generate 10",
			numCount: 10,
		},
		{
			name:     "generate 100",
			numCount: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			randChan := GenerateInt()

			timeout := time.After(time.Second * 1)
			count := 0 // чтобы считать правильное ли количество сгенерируется

			for {
				select {
				case <-timeout:
					t.Errorf("не удалось сгенерировать %d чисел за секунду", tt.numCount)
					return
				case num := <-randChan:
					count++
					if count >= tt.numCount {
						return
					}
					if num < 0 || num > 100 {
						t.Errorf("выходное число %d не входит в диапазон от 0 до 100", num)
						return
					}
				}
			}
		})
	}
}
