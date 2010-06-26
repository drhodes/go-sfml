// Created by cgo - DO NOT EDIT
package gfx

import "unsafe"

type _ unsafe.Pointer

type _Cstruct___0 struct {
	r	_C_sfUint8
	g	_C_sfUint8
	b	_C_sfUint8
	a	_C_sfUint8
}
type _C_unsignedchar uint8
type _C_sfColor _Cstruct___0
type _C_sfUint8 _C_unsignedchar
type _C_void [0]byte

func _C_sfColor_FromRGB(_C_sfUint8, _C_sfUint8, _C_sfUint8) _C_sfColor
func _C_sfColor_Modulate(_C_sfColor, _C_sfColor) _C_sfColor
func _C_sfColor_FromRGBA(_C_sfUint8, _C_sfUint8, _C_sfUint8, _C_sfUint8) _C_sfColor
func _C_sfColor_Add(_C_sfColor, _C_sfColor) _C_sfColor
