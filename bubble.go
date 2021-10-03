package main

import (
	"fmt"
)

func BubbleSort(a []int) []int {
	for i := 0; i < len(a) - 1; i++ {
		for j := 0; j < len(a) - i - 1; j++ {
			if a[j] > a[j + 1] {
				a[j], a[j + 1] = a[j + 1], a[j]
			}
		}
	}
	
	return a
}

func main() {
	a := []int{3, 1, 5, 10, 2, 9, 7, 0}

	fmt.Println(BubbleSort(a))
}