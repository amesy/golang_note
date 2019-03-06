package main

import (
	"./exercise"
	"./reboot/fourthDay"
	"fmt"
)

func main() {
	fmt.Println(
		// exercise.Calculation(),
		exercise.Yushu(),
		exercise.JingDu(),
	)
	
	f1 := 0.0000012
	f2 := 0.0000003213
	res := exercise.Compare(f1, f2)
	fmt.Println(res)
	
	fmt.Println(exercise.StringExer())
	
	fmt.Println(exercise.StringRange1())
	fmt.Println(exercise.StringRange2())
	
	fmt.Println(exercise.ArrayExer())
	exercise.ArrayExer2()
	fmt.Println(exercise.ArrayExer3())
	
	fmt.Println(exercise.MapExer())
	exercise.MapExer2()
	fmt.Println(exercise.MapExer3())
	
	fmt.Println(exercise.PtrExer())
	
	exercise.RandValueType()
	
	fmt.Println(exercise.MapExer4())
	
	fmt.Println(exercise.SwitchExer())
	
	fmt.Println(exercise.ForExer())
	
	// exercise.GotoExer()
	
	fmt.Println(exercise.FuncExer(1, 2, 3, 4, 5))
	fmt.Println(exercise.FuncExer2("amesy", 1, 2.0, "2019"))
	
	exercise.FuncExer3()
	
	fmt.Println(exercise.FuncExer4("English", "math", 100, 99.999))
	fmt.Println(exercise.FuncExer5)
	
	a := 5
	exercise.FuncExer6(a)()
	a *= 2
	exercise.FuncExer6(a)()
	
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)  //
	odd := exercise.FuncFilter(slice, exercise.Istest)
	fmt.Println("Istest func result = ", odd)  // 1, 3, 5, 7
	even := exercise.FuncFilter(slice, exercise.Iseven)
	fmt.Println("Iseven func result = ", even)  // 2, 4
	
	p := exercise.FuncRes1()(1)
	fmt.Println(p)
	
	/*
	resv := exercise.FuncAddr2(0)
	var s int
	for i := 0; i < 10; i++ {
		s, resv = resv(i)
		fmt.Println("result: ", s)
	}
	*/
	
	fmt.Println(exercise.PtrExer2())
	exercise.PtrExer3()
	
	exercise.PtrExer4()
	exercise.PtrExer5()
	
	// fmt.Println(exercise.PtrExer6())
	
	// exercise.DeferExer()
	
	exercise.StructExer1()
	exercise.StructExer2()
	exercise.StructExer3()
	
	// fmt.Println(exercise.IoExer())
	fmt.Println(exercise.IoExer2())
	
	/*
		func Println(a ...interface{}) (n int, err error) {
		return Fprintln(os.Stdout, a...)
		}
	*/
	
	// var name string
	// fmt.Fscanln(os.Stdin, &name)
	// // fmt.Fprintf(os.Stdin, name)
	// fmt.Println(name)
	
	fmt.Println(exercise.MethodExer())
	
	exercise.MethodExer2()
	
	fmt.Println(exercise.GoroutineExer())
	
	exercise.GoroutineExer2()
	
	fourthDay.T01()
	fourthDay.T02()
	
	exercise.ChainOperating()
	exercise.StringOperating()
}