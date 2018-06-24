package main

import (
	"fmt"
)

func ssort(n []int) {
	for i := 0; i < len(n); i++ {
		var min int = i
		for j := i + 1; j < len(n); j++ {
			if n[min] > n[j] {
				min = j
			}
			n[i], n[min] = n[min], n[i]
		}
	}
}

func main() {
	arr := [...]int{5, 2, 7, 4, 1}
	ssort(arr[:])

	fmt.Println(arr)
}