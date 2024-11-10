package main

import "fmt"

func Difference(slice1, slice2 []string) []string {
	result := make([]string, 0)

	for _, v := range slice1 {
		found := false
		for _, w := range slice2 {
			if v == w {
				found = true
				break
			}
		}
		if !found {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(Difference(slice1, slice2))
}
