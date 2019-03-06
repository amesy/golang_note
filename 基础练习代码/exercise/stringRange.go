package exercise

// UTF-8遍历
func StringRange1() string {
	str := "Hello 世界!"
	var line = ""
	for i := 0; i < len(str); i++ {
		res := str[i]
		line += string(res)
	}
	
	return line  // Hello ä¸ç!
}

// Unicode遍历
func StringRange2() string {
	str := "Hello 世界!"
	var line = ""
	for _, val := range str {
		line += string(val)
	}
	
	return line  // Hello 世界!
}
