package exercise

import (
	"fmt"
)

type Info struct {
	name string
	age int
}

func (i *Info) setName(name string) *Info {
	i.name = name
	return i
}

func (i *Info) setAge(age int) *Info {
	i.age = age
	return i
}

func (i *Info) print() {
	fmt.Println(*i)
}

func ChainOperating() {
	p := &Info{}
	p.setName("amesy").setAge(24).print()
}
