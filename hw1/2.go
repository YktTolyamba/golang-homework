package main

import "fmt"

func fibonacci1 (n int) int {
	var a1, a2 = 0, 1
	if n == 1 {
		return a1
	}
	if n == 2 {
		return a2
	}
	var buf int
	for i := 2; i < n; i++ {
		buf = a1 + a2
		a1 = a2
		a2 = buf
	}
	return buf
}

func main() {
	fmt.Println(fibonacci1(10))
}