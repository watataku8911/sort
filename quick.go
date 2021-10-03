package main
import "fmt"

func quick_sort(ar []int) []int {
    var left, right []int
    if len(ar) > 1 {
        x := ar[0]
        for _, a := range ar[1:] {
            if a < x {
                left  = append(left , a)
            } else {
                right = append(right, a)
            }
        }
        left  = quick_sort(left)
        right = quick_sort(right)
        ar = append(append(left, x), right...)
    }
    return ar
}

func main() {
    ar := quick_sort([]int{3, 1, 5, 10, 2, 9, 7, 0})
    fmt.Println(ar)
}