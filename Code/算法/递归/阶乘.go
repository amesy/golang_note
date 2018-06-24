package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return factorial(n-1) * n
}

func main() {
	result := factorial(5)
	fmt.Println(result)
}
