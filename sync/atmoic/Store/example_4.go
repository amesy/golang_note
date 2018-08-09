package main

import (
	"fmt"
	"sync/atomic"
)
var val int32

func main()  {
	fmt.Println("======old value=======")
	fmt.Println(val)
	fmt.Println("======new value=======")
	res := addVal(3, val)
	fmt.Println(res)
}

func addVal(delta, val int32) int32 {
	atomic.StoreInt32(&val, delta)
	return val
}