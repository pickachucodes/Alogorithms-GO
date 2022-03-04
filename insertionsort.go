package main

import "fmt"

func insertionsort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[j] > array[i] {
				array[j], array[i] = array[i], array[j]
			}

		}
	}

	return array
}
func main() {
	k := []int{2, 1, 3, 8, 9, 11, 5}
	m := insertionsort(k)
	fmt.Println(m)
}
