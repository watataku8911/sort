package main

import (
	"fmt"
)

func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i - j - 1] > a[i - j] {
				a[i - j - 1], a[i - j] = a[i - j], a[i - j - 1]
			} else {
				break
			}
		}
	}

	return a
}

func main() {
	a := []int{3, 1, 5, 10, 2, 9, 7, 0}
	fmt.Println(InsertionSort(a))
}