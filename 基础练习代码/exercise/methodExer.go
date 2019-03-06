package exercise

import (
	"fmt"
	"math"
)

// 计算矩形和圆形的面积

type (
	rectangle struct {
		length, height float64
	}
	
	circle struct {
		randius float64
	}
)

func (val rectangle) area() float64 {
	return val.length * val.height
}

func (r circle) area() float64 {
	return r.randius * r.randius * math.Pi
}

func MethodExer() (lst []float64){
	lst = make([]float64, 0)
	rtg := rectangle{
		length: 10,
		height: 1.5,
	}
	
	cce := circle{
		randius: 1.8,
	}
	
	lst = append(lst, rtg.area(), cce.area())
	
	return
}

type Point struct {
	px, py float64
}

var lst = make([]float64, 0)

// 获取
func (p *Point) getpxy() (float64, float64) {
	return p.px, p.py
}

// 修改
func (p *Point) setpxy(x, y float64){
	p.px = x
	p.py = y
}

func MethodExer2() {
	p := Point{
		px: 1.1,
		py: 1.2,
	}
	
	xx, yy := p.getpxy()
	
	fmt.Println("before the values of modify is: ", xx, yy)
	
	pp := &Point{}
	var x, y float64 = 2.1, 2.2
	pp.setpxy(x, y)
	xx, yy = pp.getpxy()
	fmt.Println("after the values of modify is: ", xx, yy)
}


