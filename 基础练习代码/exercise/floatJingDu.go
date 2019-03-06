package exercise

import (
	"reflect"
)

func JingDu() reflect.Type {
	// var val float32
	// val = 12
	
	val2 := 12.0
	
	resType := reflect.TypeOf(val2)
	return resType
}
