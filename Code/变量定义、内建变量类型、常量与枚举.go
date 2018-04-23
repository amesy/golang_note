```go 
package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var name = "yangbin"
var age = 23
var (
	sex = "male"
	work = "linux"
)

func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "amesy"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	// 字符串用双引号.
	var a, b, c, s = 3, 4, true, "amesy"
	var p = "abc"
	fmt.Println(a, b, c, s, p)
}

// 用 var 定义变量和用 := 效果一样。
func variableShorter() {
	a, b, c, s := 3, 4, true, "amesy"
	b = 10  // 第二次赋值不能加冒号.
	p := "abc"
	fmt.Println(a, b, c, s, p)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	// 定义多个常量时也可以用括号括起来。
	//变量名不用大写，因为go语言里面大写有特殊含义。
	const filename = "abc.txt"
    //未定义类型时，a和b可当做int或float.
	const a, b = 3, 4  
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		python
		java
		golang
	)
	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, java, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(name, age, sex, work)
	euler()
	triangle()
	consts()
	enums()
}

```

