package main

import (
	"testing"
)

func TestDataProcessingPipeline(t *testing.T) {
	tests := []struct {
		name        string
		inputData   []uint8
		expectedOut []float64
	}{
		{
			name:        "тест от 1 до 5",
			inputData:   []uint8{1, 2, 3, 4, 5},
			expectedOut: []float64{1, 8, 27, 64, 125},
		},
		{
			name:        "тест с 1 значением",
			inputData:   []uint8{10},
			expectedOut: []float64{1000},
		},
		{
			name:        "пустой",
			inputData:   []uint8{},
			expectedOut: []float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan uint8, len(tt.inputData))
			out := make(chan float64, len(tt.inputData))

			go dataProcessingPipeline(in, out)

			for _, num := range tt.inputData {
				in <- num
			}
			close(in)

			var res []float64
			for result := range out {
				res = append(res, result)
			}

			if len(res) != len(tt.expectedOut) {
				t.Errorf("ожидали длину %d, получили %d", len(tt.expectedOut), len(res))
			}

			for i, expected := range tt.expectedOut {
				if res[i] != expected {
					t.Errorf("ожидали %v, получили %v по индексу %d", expected, res[i], i)
				}
			}
		})
	}
}
