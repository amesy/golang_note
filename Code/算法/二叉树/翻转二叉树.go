package main

import "fmt"

type trees struct {
	data int
	left *tree
	right *tree
}

func create(index int, value []int) (T *tree) {
	T = &tree{}
	T.data = value[index-1]
	fmt.Printf("index %v", index)
	if index < len(value)-1 && 2*index <= len(value) && 2*index+1 <= len(value) {
		
	}

}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	create(1, arr[:])
}
