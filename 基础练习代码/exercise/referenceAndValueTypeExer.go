package exercise

import (
	"fmt"
)

func value_type(n []int) {
	m := n
	m[2] = 10
	fmt.Println(n, m)
}

func reference_type(n [5]int) {
	m := n
	m[2] = 10
	fmt.Println(n, m)
}

func RandValueType() {
	b := []int{1, 2, 3, 4, 5}
	value_type(b)
	
	a := [5]int{1, 2, 3, 4, 5}
	reference_type(a)
}
