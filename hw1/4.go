package main

import "fmt"

func bubble_sort (b []int) {
	k := 1
	n := len(b)
	for i := n; i >= 0; i-- {
		if k > 0 {
			k = 0
		} else {
			break
		}
		for j := 0; j < i-1; j++ {
			if b[j] > b[j+1] {
				buf := b[j]
				b[j] = b[j+1]
				b[j+1] = buf
				k++
			}
		}
	}

}

func main() {
	a := [10]int{34,6,23,7,45,34,74,23,2,25}
	bubble_sort(a[:])
	fmt.Println(a)
}