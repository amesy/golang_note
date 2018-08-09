package main

import (
	"fmt"
	"sync/atomic"
)
var value int32

func main()  {
	fmt.Println("======old value=======")
	fmt.Println(value)
	fmt.Println("======CAS value=======")
	addValue(3)
	fmt.Println(value)
}

func addValue(delta int32){
	for {
		//v := value
		v := atomic.LoadInt32(&value)
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)){
			break
		}
	}
}