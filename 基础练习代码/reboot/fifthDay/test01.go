package fifthDay

// IInstance
// Instance

type IInstance interface {
	Instance() float64
}

type Put struct {}

func (p Put) Cal() IInstance {
	
}


func main() {
	var i IInstance
	p := Put{}
	
	i = p.Cal()
}
