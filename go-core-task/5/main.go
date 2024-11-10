package main

import "fmt"

func Intersection(slice1, slice2 []int) (bool, []int) {
	set := make(map[int]struct{})
	for _, val := range slice1 {
		set[val] = struct{}{}
	}

	var res []int
	for _, val := range slice2 {
		if _, ok := set[val]; ok {
			res = append(res, val)
		}
	}
	return len(res) > 0, res
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(Intersection(a, b))
}
