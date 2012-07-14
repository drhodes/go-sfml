// Created by cgo - DO NOT EDIT

package sys

import "unsafe"

import "syscall"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *error, x int) { *dst = syscall.Errno(x) }
type _Ctype_void [0]byte

func _Cfunc_sfClock_copy(*[0]byte) *[0]byte
func _Cfunc_sfClock_create() *[0]byte
func _Cfunc_sfClock_destroy(*[0]byte)
