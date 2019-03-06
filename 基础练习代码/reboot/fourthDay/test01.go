package fourthDay

import (
	"crypto/md5"
	"fmt"
)

func T01() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a)
}

func T02() {
	data := []byte("hello World")
	md5sum := md5.Sum(data)
	fmt.Printf("%x\n", md5sum)
}
