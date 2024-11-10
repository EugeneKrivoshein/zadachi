package main

import "testing"

func TestMergeChannels(t *testing.T) {
	tests := []struct {
		name           string
		channels       []<-chan int
		expectedOutput []int
	}{
		{
			name: "merge two channels",
			channels: []<-chan int{
				func() <-chan int {
					ch := make(chan int)
					go func() {
						defer close(ch)
						ch <- 1
						ch <- 2
					}()
					return ch
				}(),
				func() <-chan int {
					ch := make(chan int)
					go func() {
						defer close(ch)
						ch <- 3
						ch <- 4
					}()
					return ch
				}(),
			},
			expectedOutput: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := MergeChannels(tt.channels...)
			var res []int
			for num := range merged {
				res = append(res, num)
			}

			if len(res) != len(tt.expectedOutput) {
				t.Errorf("ожидали %d, получили %d", len(tt.expectedOutput), len(res))
			}

			for i := range tt.expectedOutput {
				if res[i] != tt.expectedOutput[i] {
					t.Errorf("ожидали %d, получили %d по индексу %d", tt.expectedOutput[i], res[i], i)
				}
			}
		})
	}
}
