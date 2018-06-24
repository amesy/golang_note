package main

import "fmt"

func isort(n []int) {
	for i := 1; i < len(n); i++ {
		for j := i; j > 0; j-- {
			if n[j] > n[j-1] {
				break
			}
			n[j], n[j-1] = n[j-1], n[j]
		}
	}
}

func main() {
	arr := [...]int{5, 2, 7, 4, 1}
	isort(arr[:])
	fmt.Println(arr)
}
