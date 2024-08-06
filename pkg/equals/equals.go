package equals

import "reflect"

// Equals used to compare whether two variables have equal depth
// Can compare values of primitive types
// It also recursively compares all fields of composite types such as slices, maps, and structs
func Equals(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
