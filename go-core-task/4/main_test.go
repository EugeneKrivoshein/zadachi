package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "уникальные слайсы",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "слайсы с одинаковыми элементами",
			slice1:   []string{"apple", "banana"},
			slice2:   []string{"apple", "banana"},
			expected: []string{},
		},
		{
			name:     "второй слайс пустой",
			slice1:   []string{"apple", "banana", "cherry"},
			slice2:   []string{},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "первый слайс пустой",
			slice1:   []string{},
			slice2:   []string{"apple", "banana", "cherry"},
			expected: []string{},
		},
		{
			name:     "оба слайса пустые",
			slice1:   []string{},
			slice2:   []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("получили %v, ожидаем %v", result, tt.expected)
			}
		})
	}
}
