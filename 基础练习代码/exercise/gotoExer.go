package exercise

import (
	"fmt"
)

func GotoExer() {
	i := 0
	HERE:
		fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}
