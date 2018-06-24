package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		result := fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
		return res
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
