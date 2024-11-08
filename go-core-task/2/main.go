package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateSlice() []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, 10)

	for i := range slice {
		slice[i] = r.Intn(100) + 1
	}
	return slice
}

func SliceExample(slice []int) []int {
	var newSlice []int
	for _, i := range slice {
		if i%2 == 0 {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}

func AddElements(slice []int, elem int) []int {
	return append(slice, elem)
}

func CopySlice(slice []int) []int {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return copySlice
}

func RemoveElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		fmt.Println("индекс выходит за пределы слайса")
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
func main() {
	slice := GenerateSlice()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	elem := r.Intn(9)
	fmt.Println("slice", slice)
	fmt.Println("chetnyi slice", SliceExample(slice))
	fmt.Printf("dobavim element %v v konec slica: %v\n", elem, AddElements(slice, elem))
	fmt.Printf("copy slice %v\n", CopySlice(slice))
	fmt.Printf("remove element %d: %v\n", elem, RemoveElement(slice, elem))
}
