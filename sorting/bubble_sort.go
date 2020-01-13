package main

import (
	"fmt"
)

func main() {

	var array []int

	array = []int{100, 45, 21, 67, 34, 10, 8, 439, 410}

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				swap := array[j]
				array[j] = array[j+1]
				array[j+1] = swap
			}
		}
	}
	fmt.Println(array)
}
