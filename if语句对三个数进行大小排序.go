package main

import "fmt"

func main() {
	a := 10
	b := 1
	c := 25

	if a > b {
		if b > c {
			fmt.Println( a, b, c)
		} else {
			if a > c {
				fmt.Println(a, c, b)
			} else {
				fmt.Println(b, a, c)
			}
		}
	} else if b > c {
		if a > c {
			fmt.Println(b, a, c)
		} else {
			fmt.Println(b, c, a)
		}
	} else {
		fmt.Println(c, b, a)
	}
}
