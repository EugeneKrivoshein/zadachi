package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name         string
		slice1       []int
		slice2       []int
		expectedBool bool
		expected     []int
	}{
		{
			name:         "intersection",
			slice1:       []int{1, 2, 3, 4, 5},
			slice2:       []int{4, 5, 6, 7, 8},
			expectedBool: true,
			expected:     []int{4, 5},
		},
		{
			name:         "no intersection",
			slice1:       []int{1, 2, 3},
			slice2:       []int{4, 5, 6},
			expectedBool: false,
			expected:     []int{},
		},
		{
			name:         "empty slice1",
			slice1:       []int{},
			slice2:       []int{1, 2, 3},
			expectedBool: false,
			expected:     []int{},
		},
		{
			name:         "empty slice2",
			slice1:       []int{1, 2, 3},
			slice2:       []int{},
			expectedBool: false,
			expected:     []int{},
		},
		{
			name:         "empty slices",
			slice1:       []int{},
			slice2:       []int{},
			expectedBool: false,
			expected:     []int{},
		},
		{
			name:         "all intersection",
			slice1:       []int{1, 2, 3},
			slice2:       []int{1, 2, 3},
			expectedBool: true,
			expected:     []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultBool, result := Intersection(tt.slice1, tt.slice2)
			if resultBool != tt.expectedBool {
				t.Errorf("ожидалось %v, получили %v", tt.expectedBool, resultBool)
			}
			if (len(result) == 0 && len(tt.expected) == 0) || reflect.DeepEqual(result, tt.expected) {
				return
			}
			t.Errorf("ожидалось %v, получили %v", tt.expected, result)
		})
	}
}
