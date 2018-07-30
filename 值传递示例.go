package main

import "fmt"

func main() {
	i:=10
	ip:=&i
	fmt.Println("ip, &i", ip, &i)
	fmt.Printf("原始指针的内存地址是：%p\n",&ip)
	fmt.Println("---ip", *ip)
	*ip = 3
	fmt.Println("---i", i)
	modify(ip)
	fmt.Println("int值被修改了，新值为:",i)
}
func modify(ip *int){
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n",&ip)
	*ip=1
}


/*
Go语言没有引用传递，以上面的例子为例，如果在modify函数里打印出来的内存地址是不变的，也是0xc42000c028，那么就是引用传递。
*/