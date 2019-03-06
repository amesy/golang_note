package exercise

func ForExer() []int {
	var val []int
	for k, v := range []int{1, 2, 3, 4} {
		val = append(val, k, v, 8)
	}
	
	return val
}
