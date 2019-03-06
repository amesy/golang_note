package exercise

import (
	"fmt"
)

var  lastname, firstname string

func IoExer() string {
	fmt.Printf("input name: ")
	n, _ := fmt.Scanln(&lastname, &firstname)
	
	fmt.Println(n)
	
	return "you name is：" + lastname + " " + firstname
}

func IoExer2() (string, string, int){
	var (
		name, speciality string
		age int
	)
	
	myinfo := "amesy computer 23"
	fmt.Sscanln(myinfo, &name, &speciality, &age)
	
	return name, speciality, age
}


