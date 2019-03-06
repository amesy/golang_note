package exercise

import (
	"fmt"
)

type name string

func (n name) length() int {
	return len(n)
}

func StructExer1() {
	var myname name = "amesy"
	line := []byte(myname)
	fmt.Println("the length of []byte is: ", len(line))
	fmt.Println("调用对象方法后的长度: ", myname.length())
}

type person struct {
	name string
	age int
}

func (info person) info() []interface{} {
	lst := make([]interface{}, 0)
	lst = append(lst, info.name, info.age)
	return lst
}

func compareOfAge(info1, info2 person) (person, int) {
	if info1.age > info2.age {
		return info1, info1.age
	}
	
	return info2, info2.age
}

func StructExer2() {
	var amesy person
	amesy.name, amesy.age = "Amesy", 24
	
	bob := person{
		"Bob",
		23,
	}
	
	tom := person{
		name: "Tom",
		age: 22,
	}
	
	ab_Older, ab_diff := compareOfAge(amesy, bob)
	at_Older, at_diff := compareOfAge(amesy, tom)
	bt_Older, bt_diff := compareOfAge(bob, tom)
	
	fmt.Printf("Of %s and %s, %s is Older by %d years\n",
		amesy.name, bob.name, ab_Older.name, ab_diff)
	
	fmt.Printf("Of %s and %s, %s is Older by %d years\n",
		amesy.name, tom.name, at_Older.name, at_diff)
	
	fmt.Printf("Of %s and %s, %s is Older by %d years\n",
		bob.name, tom.name, bt_Older.name, bt_diff)
	
	
	fmt.Println(amesy.info())
}

type Human struct {
	name string
	age, weight int
}

type Student struct {
	Human
	speciality string
}

func StructExer3() {
	var lst []interface{}
	myinfo := Student{
		Human{
			"amesy",
			24,
			120,
		},
		"Computer Science",
	}
	
	lst = append(lst, myinfo.name, myinfo.age, myinfo.weight, myinfo.speciality)
	
	myinfo.name = "tom"
	myinfo.age = 21
	myinfo.weight = 122
	myinfo.speciality = "English"
	
	lst = append(lst, myinfo.name, myinfo.age, myinfo.weight, myinfo.speciality)
	
	fmt.Println("res is: ", lst)
}

