package main

import (
	"reflect"
	"testing"
)

func TestGenerateSlice(t *testing.T) {
	slice := GenerateSlice()

	if len(slice) != 10 {
		t.Errorf("ожидалось 10, получего %d", len(slice))
	}
}

func TestSliceExample(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: []int{2, 4, 6, 8, 10}},
		{input: []int{11, 13, 15, 17}, expected: []int{}},
		{input: []int{2, 4, 6, 8}, expected: []int{2, 4, 6, 8}},
		{input: []int{}, expected: []int{}},
	}

	for _, tc := range testCases {
		res := SliceExample(tc.input)

		if len(res) != len(tc.expected) {
			t.Errorf("входной слайс %v ожидалось %v, получено %v", tc.input, tc.expected, res)
		}
		for i := range res {
			if res[i] != tc.expected[i] {
				t.Errorf("входной слайс %v ожидалось %v, получено %v", tc.input, tc.expected, res)
			}
		}
	}
}

func TestAddElements(t *testing.T) {
	testCases := []struct {
		inputSlice []int
		element    int
		expected   []int
	}{
		{inputSlice: []int{1, 2, 3}, element: 4, expected: []int{1, 2, 3, 4}},
		{inputSlice: []int{}, element: 5, expected: []int{5}},
		{inputSlice: []int{10, 20}, element: 30, expected: []int{10, 20, 30}},
	}

	for _, tc := range testCases {
		res := AddElements(tc.inputSlice, tc.element)

		if !reflect.DeepEqual(res, tc.expected) {
			t.Errorf("ожидалось %v, получено %v", tc.expected, res)
		}

		if len(tc.inputSlice) > 0 && tc.inputSlice[len(tc.inputSlice)-1] == tc.element {
			t.Errorf("исходный слайс изменен: %v", tc.inputSlice)
		}
	}
}

func TestCopySlice(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{10, 20, 30}, expected: []int{10, 20, 30}},
		{input: []int{}, expected: []int{}},
	}

	for _, tc := range testCases {
		result := CopySlice(tc.input)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("ожидалось %v, получено %v", tc.expected, result)
		}

		//проверка, что изменения в ориг слайсе не затрагивают копию
		if len(tc.input) > 0 {
			tc.input[0] = 9
			if tc.input[0] == result[0] {
				t.Errorf("копия слайса была изменена: оригинал %v, копия %v", tc.input, result)
			}
		}
	}
}

func TestRemoveElement(t *testing.T) {
	// Тестовые данные
	testCases := []struct {
		input    []int
		index    int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, index: 3, expected: []int{1, 2, 3, 5}},
		{input: []int{10, 20, 30, 40}, index: 3, expected: []int{10, 20, 30}},
		{input: []int{1, 2, 3}, index: 2, expected: []int{1, 2}},
		{input: []int{1}, index: 0, expected: []int{}},
	}

	for _, tc := range testCases {
		result := RemoveElement(tc.input, tc.index)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("ожидалось %v, получено %v", tc.expected, result)
		}
	}

	// тест выхода за пределы слайса
	invalidIndex := 10
	result := RemoveElement([]int{1, 2, 3}, invalidIndex)
	if len(result) != 3 {
		t.Errorf("для индекса %d слайс изменен, хотя индекс выходит за пределы", invalidIndex)
	}
}
