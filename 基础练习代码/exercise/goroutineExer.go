package exercise

import (
	"fmt"
	"time"
)

func GoroutineExer() string {
	msg := make(chan string)
	
	go func() {
		msg <- "Hello World!"
	}()
	
	str := <- msg
	
	return str
}

var msg = make(chan int, 3)
func example1(msg chan int) {
	msg <- 1
	msg <- 2
	msg <- 3
}

func example2(msg chan int) {
	time.Sleep(time.Second)
	str := <- msg
	str += 10
	msg <- str
	
	close(msg)
	fmt.Println("111: ", <- msg)
}

func GoroutineExer2() {
	go example1(msg)
	go example2(msg)
	time.Sleep(time.Second)
	
	for str := range msg {
		fmt.Println("loop: ", str)
	}
}
