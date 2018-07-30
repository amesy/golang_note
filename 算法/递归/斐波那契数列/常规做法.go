package main

import "fmt"

func main() {
	a := 0
	b := 1
	for i:=1; i<11; i++ {
		a, b = b, a+b
		fmt.Println(a)
	}
}
