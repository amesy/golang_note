/*
给定一个正整数N 
要求：
返回一个数组，长度为N，元素个数为N，元素相加的和为0. 
用任意语言或伪代码实现。
*/

package main

import (
	"fmt"
)

type info struct {
	arr []int
	num int
}

func arr(num_slice, arr []int) []int {
	for _, v := range num_slice {
		arr = append(arr, v, -v)
	}
	return arr
}

func main() {
	res := info{
		arr: make([]int, 0),
		num: 5,
	}

	num_slice := make([]int, 0)
	num_half := (res.num / 2)

	if res.num % 2 != 0 {
		for i := 0; i < (num_half + 1); i++ {
			num_slice = append(num_slice, i)
		}

		L := arr(num_slice, res.arr)
		L = L[1:]
		fmt.Println(L)
		return
	}

	for i := 0; i < num_half; i++ {
		num_slice = append(num_slice, i)
	}

	fmt.Println(arr(num_slice, res.arr))
	return

}
