package exercise

import (
	"fmt"
)

func ArrayExer() [2][3]int {
	nums := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	return nums
}

func ArrayExer2() {
	for _, val := range ArrayExer() {
		fmt.Println(val)
	}
}

func ArrayExer3() int {
	var arrayNum []int
	arrayNum = make([]int, 0)
	for i := 0; i <= 10; i++ {
		arrayNum = append(arrayNum, i)
	}
	
	p := arrayNum[2:4:6]
	
	return cap(p)
}
