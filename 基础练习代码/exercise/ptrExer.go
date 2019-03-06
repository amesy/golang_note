package exercise

import (
	"fmt"
)

func PtrExer() *map[int]int {
	name := &map[int]int{3:3}
	fmt.Println("res: ", name)
	return name
}

func PtrExer2() *string {
	name := "amesy"
	return &name
}

func PtrExer3() {
	a := 10
	
	var res *int
	
	res = &a
	
	fmt.Printf("a 变量的地址是: %x\n", &a  )
	
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", res )
	
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *res )
}

func PtrExer4() {
	list := []int{10, 20, 30}
	const MAX int = 3
	
	for i := 0; i < MAX; i++ {
		fmt.Printf("list[%d] = %d\n", i, list[i])
	}
}

func PtrExer5() {
	list := []int{10, 20, 30}
	length := len(list)
	var ptrlist []*int
	
	for i := 0; i < length; i++ {
		ptrlist = append(ptrlist, &list[i])
	}
	
	fmt.Println(ptrlist)
	
	for i := 0; i < length; i++ {
		fmt.Printf("%d - %d - %x\n", i, *ptrlist[i], ptrlist[i])
	}
}

func PtrExer6() []interface{} {
	var a int = 100
	var ptr *int
	var pptr **int
	
	var lst []interface{}
	
	ptr = &a
	pptr = &ptr
	
	lst = append(lst, a, ptr, pptr)
	fmt.Println(lst)
	
	return lst
}
