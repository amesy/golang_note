package main

import (
	"fmt"
)

func lengthOfNonrepeatingSubstring(s string) int {
	start := 0
	Maxlength := 0
	lastOccurred := make(map[rune]int)

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > Maxlength {
			Maxlength = i - start + 1
		}

		lastOccurred[ch] = i

	}
	//fmt.Println(lastOccurred)
	return Maxlength
}

func main() {
	fmt.Println(lengthOfNonrepeatingSubstring("abcabcbb"))
	fmt.Println(lengthOfNonrepeatingSubstring("bbbbb"))
	fmt.Println(lengthOfNonrepeatingSubstring("pwwkew"))
	fmt.Println(lengthOfNonrepeatingSubstring(""))
	fmt.Println(lengthOfNonrepeatingSubstring("b"))
	fmt.Println(lengthOfNonrepeatingSubstring("asdfgh"))
	fmt.Println(lengthOfNonrepeatingSubstring("昨天今天明天"))
	fmt.Println(lengthOfNonrepeatingSubstring("一二三二一"))
}
