package main

import (
	"fmt"
)

func Min(a []int) (idx, n int) {	
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}
	
	return
}

func SelectionSort(a []int) []int {
	for i, _ := range a {
		idx, _ := Min(a[i:len(a)])
		a[i], a[i + idx] = a[i + idx], a[i]
	}
	
	return a
}

func main() {
	a := []int{3, 1, 5, 10, 2, 9, 7, 0}
	fmt.Println(SelectionSort(a))
}