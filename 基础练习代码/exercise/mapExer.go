package exercise

import (
	"fmt"
)

func MapExer() map[int]string {
	info := make(map[int]string, 0)
	info[1] = "amesy"
	info[2] = "yangbin"
	info[3] = "jack"
	// delete(info, 3)
	return info
}

func MapExer2() {
	for k, v := range MapExer() {
		fmt.Printf("k = %d, v = %s\n", k, v)
	}
}

func MapExer3() string {
	value, is := MapExer()[3]
	if is {
		return value
	} else {
		return "nothing."
	}
}

func MapExer4() string {
	maps := map[string]string{"name": "amesy", "sex": "male"}
	if value, ok := maps["name"]; ok {
		return value
	} else {
		return "nothing"
	}
}
