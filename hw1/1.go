package main

import "fmt"

func multiply(a int, b int) int {
	answer := 0
	for i := 0; i < b; i++ {
		answer += a
	}
	return answer
}

func main() {
	fmt.Println(multiply(3,7))
}