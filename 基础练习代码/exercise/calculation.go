package exercise

func Calculation() (int, int) {
	const (
		a = iota
		b
		c
		d = 8
		e = 8
		f = iota
		g
	)
	
	return c, g
}
