package exercise

import (
	"fmt"
)

func FuncExer(args ...int) []int {
	var sets []int
	for _, v := range args {
		sets = append(sets, v)
	}
	
	return sets
}

func FuncExer2(args ...interface{}) []interface{} {
	var sets []interface{}
	for _, v := range args {
		sets = append(sets, v)
	}
	
	return sets
}

func FuncExer3() {
	for _, v := range FuncExer2(1, "amesy", 2.999, true) {
		switch v.(type) {
		case int:
			fmt.Println("int: ", v)
		case string:
			fmt.Println("string: ", v)
		case float64:
			fmt.Println("float: ", v)
		case bool:
			fmt.Println("bool: ", v)
		default:
			fmt.Println("nothing")
		}
	}
}

var sets []interface{}

var FuncExer4 = func(args ...interface{}) (sets []interface{}){
	for _, v := range args {
		sets = append(sets, v)
	}
	
	return
}

var FuncExer5 = func(args ...interface{}) (sets []interface{}){
	for _, v := range args {
		sets = append(sets, v)
	}
	
	return
} ("English", "math", 100, 99.999)

var FuncExer6 = func(a int) (func()) {
	b := 20
	return func() {
		fmt.Println("a * b is: ", a * b)
	}
}

type testFunc func(int) bool

func Istest(a int) bool {
	if a % 2 == 0 {
		return false
	}
	
	return true
}

func Iseven(b int) bool {
	if b % 2 == 0 {
		return true
	}
	
	return false
}

func FuncFilter(slice []int, f testFunc) []int {
	var result []int
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

var FuncRes1 = func() (func(int) int) {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAddr func(int) (int, iAddr)

func FuncAddr2(base int) iAddr {
	return func(v int) (int, iAddr) {
		return base +v, FuncAddr2(base + v)
	}
}






