package main

import (
	"fmt"
)

func main() {
	var array []int

	array = []int{23, 65, 78, 90, 12, 4, 90, 3, 1}

	for i := 0; i < len(array)-1; i++ {
		minValue := i
		for j := i + 1; j < len(array); j++ {
			if array[minValue] > array[j] {
				minValue = j
			}
			if minValue != i {
				swap := array[i]
				array[i] = array[minValue]
				array[minValue] = swap
			}
		}
		fmt.Println(array)
	}

}
