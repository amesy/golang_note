package exercise

import (
	"fmt"
)

func DeferExer() {
	defer fmt.Println("Hello World!")
	for i := 0; i <= 5; i++ {
		if i == 2 {
			panic("程序遇见2，故退出!")
		}
		defer fmt.Println(i)
	}
}


