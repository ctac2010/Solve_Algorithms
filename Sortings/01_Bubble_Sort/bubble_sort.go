package main

import "fmt"

func BubbleSort(array []int) []int {
	ln := len(array)
	for i := 0; i < ln; i++ {
		var isSwapped bool = false
		for j := 0; j < ln-i-1; j++ {
			if array[j] > array[j+1] {
				isSwapped = true
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
		if !isSwapped {
			break
		}
	}
	return array
}

func main() {
	a := []int{4, 3, 2, 1, 1}

	fmt.Println("Array before sotring", a)

	a = BubbleSort(a)

	fmt.Println("Array after sotring", a)
}
