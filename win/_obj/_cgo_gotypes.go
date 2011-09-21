// Created by cgo - DO NOT EDIT

package win

import "unsafe"

import "os"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *os.Error, x int) { *dst = os.Errno(x) }
type _Ctype_uint uint32
type _Ctypedef_size_t _Ctype_uint
type _Ctype_unsignedchar uint8
type _Ctypedef_sfJoyAxis uint32
type _Ctype_char int8
type _Ctypedef_sfBool _Ctype_int
type _Ctypedef_sfWindowHandle _Ctype_ulong
type _Ctypedef_sfUint8 _Ctype_unsignedchar
type _Ctype_struct___1 struct {
	Width		_Ctype_uint
	Height		_Ctype_uint
	BitsPerPixel	_Ctype_uint
}
type _Ctypedef_sfWindowSettings _Ctype_struct___0
type _Ctypedef_sfMouseButton uint32
type _Ctypedef_sfVideoMode _Ctype_struct___1
type _Ctype_ulong uint32
type _Ctype_union___2 [20]byte
type _Ctype_struct___0 struct {
	DepthBits		_Ctype_uint
	StencilBits		_Ctype_uint
	AntialiasingLevel	_Ctype_uint
}
type _Ctype_float float32
type _Ctype_int int32
type _Ctypedef_sfEvent _Ctype_union___2
type _Ctypedef_sfKeyCode uint32
type _Ctype_void [0]byte

func _Cfunc_sfWindow_GetSettings(*[0]byte) _Ctypedef_sfWindowSettings
func _Cfunc_sfVideoMode_GetDesktopMode() _Ctypedef_sfVideoMode
func _Cfunc_sfContext_Create() *[0]byte
func _Cfunc_sfWindow_GetHeight(*[0]byte) _Ctype_uint
func _Cfunc_CString(string) *_Ctype_char
func _Cfunc_sfContext_SetActive(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfWindow_UseVerticalSync(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfInput_IsJoystickButtonDown(*[0]byte, _Ctype_uint, _Ctype_uint) _Ctypedef_sfBool
func _Cfunc_sfWindow_SetPosition(*[0]byte, _Ctype_int, _Ctype_int)
func _Cfunc_sfVideoMode_IsValid(_Ctypedef_sfVideoMode) _Ctypedef_sfBool
func _Cfunc_sfWindow_GetInput(*[0]byte) *[0]byte
func _Cfunc_sfWindow_Show(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfWindow_GetFrameTime(*[0]byte) _Ctype_float
func _Cfunc_sfWindow_Close(*[0]byte)
func _Cfunc_sfWindow_Destroy(*[0]byte)
func _Cfunc_sfInput_GetMouseY(*[0]byte) _Ctype_int
func _Cfunc_sfWindow_CreateFromHandle(_Ctypedef_sfWindowHandle, _Ctypedef_sfWindowSettings) *[0]byte
func _Cfunc_sfWindow_SetCursorPosition(*[0]byte, _Ctype_uint, _Ctype_uint)
func _Cfunc_sfWindow_Create(_Ctypedef_sfVideoMode, *_Ctype_char, _Ctype_ulong, _Ctypedef_sfWindowSettings) *[0]byte
func _Cfunc_sfWindow_SetIcon(*[0]byte, _Ctype_uint, _Ctype_uint, *_Ctypedef_sfUint8)
func _Cfunc_sfWindow_SetJoystickThreshold(*[0]byte, _Ctype_float)
func _Cfunc_sfWindow_SetActive(*[0]byte, _Ctypedef_sfBool) _Ctypedef_sfBool
func _Cfunc_sfInput_IsKeyDown(*[0]byte, _Ctypedef_sfKeyCode) _Ctypedef_sfBool
func _Cfunc_sfWindow_GetEvent(*[0]byte, *_Ctypedef_sfEvent) _Ctypedef_sfBool
func _Cfunc_sfInput_GetJoystickAxis(*[0]byte, _Ctype_uint, _Ctypedef_sfJoyAxis) _Ctype_float
func _Cfunc_sfWindow_Display(*[0]byte)
func _Cfunc_sfWindow_SetFramerateLimit(*[0]byte, _Ctype_uint)
func _Cfunc_sfWindow_IsOpened(*[0]byte) _Ctypedef_sfBool
func _Cfunc_sfWindow_EnableKeyRepeat(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfVideoMode_GetMode(_Ctypedef_size_t) _Ctypedef_sfVideoMode
func _Cfunc_sfContext_Destroy(*[0]byte)
func _Cfunc_sfInput_GetMouseX(*[0]byte) _Ctype_int
func _Cfunc_sfWindow_GetWidth(*[0]byte) _Ctype_uint
func _Cfunc_sfWindow_ShowMouseCursor(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfInput_IsMouseButtonDown(*[0]byte, _Ctypedef_sfMouseButton) _Ctypedef_sfBool
func _Cfunc_sfWindow_SetSize(*[0]byte, _Ctype_uint, _Ctype_uint)
func _Cfunc_sfVideoMode_GetModesCount() _Ctypedef_size_t
