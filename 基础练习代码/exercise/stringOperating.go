package exercise

import (
	"fmt"
)

func StringOperating() {
	beforeReverse := "6a7a7a8a6a7a8a1a4a3a1a"
	
	afterReverse := func (s string) (result string){
		for _, v := range s {
			result = string(v) + result  // 推导.
		}
		return
	}(beforeReverse)
	
	fmt.Println(afterReverse)  // a1a3a4a1a8a7a6a8a7a7a6
}
