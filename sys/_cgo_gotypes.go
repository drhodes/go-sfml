// Created by cgo - DO NOT EDIT
package sys

import "unsafe"

type _ unsafe.Pointer

type _C_float float32
type _C_void [0]byte

func _C_sfClock_GetTime(*[0]byte) _C_float
func _C_sfClock_Reset(*[0]byte)
func _C_sfClock_Create() *[0]byte
func _C_sfSleep(_C_float)
func _C_sfClock_Destroy(*[0]byte)
