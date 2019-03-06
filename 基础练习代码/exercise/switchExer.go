package exercise

func SwitchExer() int {
	name := 20
	switch name {
	case 20:
		return 20
		fallthrough
	default:
		return name
	}
}
