package main

import (
	"fmt"
)

func bsort(n []int) {
	for i:=0; i< len(n); i++ {
		for j := 1; j < len(n)-i; j++ {
			if n[j] < n[j-1] {
				n[j], n[j-1] = n[j-1], n[j]
			}
		}
	}
}

func main() {
	arr := [...]int{5, 2, 7, 4, 1}
	bsort(arr[:])

	fmt.Println(arr)
}
