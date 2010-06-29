// Created by cgo - DO NOT EDIT
//line util.go:1
package cgo

// #include <SFML/Graphics/Color.h>


func SfBool2GoBool(cVal _C_sfBool) bool {
	return int(cVal) == 1
}

func GoBool2SfBool(goVal bool) _C_sfBool {
	if goVal {
		return _C_sfBool(1)
	}
	return _C_sfBool(0)
}
