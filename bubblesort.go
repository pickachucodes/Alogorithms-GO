package main

import "fmt"

func bubsort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
func main() {
	k := []int{12, 11, 7, 8, 45, 33}
	m := bubsort(k)
	fmt.Println(m)
}
