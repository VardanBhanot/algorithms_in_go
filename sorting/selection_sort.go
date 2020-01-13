package main

import(
	"fmt"
)

func main() {
	var array []int
	//var temp int

	array = []int{23, 65, 78, 90, 12, 4,90,3,1}

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
		fmt.Println(array)
	}
	
}
