// Created by cgo - DO NOT EDIT

package sys

import "unsafe"

import "os"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *os.Error, x int) { *dst = os.Errno(x) }
type _Ctype_float float32
type _Ctype_void [0]byte

func _Cfunc_sfClock_GetTime(*[0]byte) _Ctype_float
func _Cfunc_sfClock_Reset(*[0]byte)
func _Cfunc_sfClock_Create() *[0]byte
func _Cfunc_sfSleep(_Ctype_float)
func _Cfunc_sfClock_Destroy(*[0]byte)
