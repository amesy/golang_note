package main

import "fmt"

type person struct {
	name string
	age int
}

func Older (info1, info2 person) (person, int) {
	if info1.age > info2.age {
		return info1, info1.age
	}
	return info2, info2.age
}

func main() {
	var tom person
	tom.name, tom.age = "Tom", 20

	bob := person {
		name: "Bob",
		age: 18,
	}

	pual := person{name: "Pual"}  // defalut pual' age is zero.

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, pual)
	pb_Older, pb_diff := Older(bob, pual)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, pual.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, pual.name, pb_Older.name, pb_diff)

	fmt.Printf("%d", pual.age)
}