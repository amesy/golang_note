package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "China中国"
	fmt.Println(s)
	fmt.Printf("%X", []byte(s))
	fmt.Println()  // 空一行.

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {  // ch is a rune. rune is a alise of int32 and it's integer of four byte.
		fmt.Printf("(%d %X) ", i, ch)  // 将字符串进行utf8解码，然后转unicode,最后再放到rune中。
	}
	fmt.Println()

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeLastRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)  // 出现乱码.
	}
	fmt.Println()

	for i, ch := range []rune(s) {  // 每个rune转出来是4个字节.
		fmt.Printf("(%d %c )", i, ch)  // 该结果是重新开出来的rune数组保存的，内存地址发生了改变。
	}
}