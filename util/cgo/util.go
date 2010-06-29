package cgo

// #include <SFML/Graphics/Color.h>
import "C"

func SfBool2GoBool(cVal C.sfBool) bool {
	return int(cVal) == 1
}

func GoBool2SfBool(goVal bool) C.sfBool {
	if goVal {
		return C.sfBool(1)
	}
	return C.sfBool(0)
}
