package main

import (
	"fmt"
)

func slesort(array []int) []int {
	for i := 0; i < len(array); i++ {
		min_index := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
			temp := array[i]
			array[i] = array[min_index]
			array[min_index] = temp
		}
	}
	return array
}
func main() {
	k := []int{3, 1, 8, 2, 8, 7}
	m := slesort(k)
	fmt.Println(m)
}
                                           
