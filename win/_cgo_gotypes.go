// Created by cgo - DO NOT EDIT
package win

import "unsafe"

type _ unsafe.Pointer

type _Cstruct___0 struct {
	Width		_C_uint
	Height		_C_uint
	BitsPerPixel	_C_uint
}
type _C_sfWindowSettings _Cstruct___2
type _C_unsignedchar uint8
type _C_size_t _C_uint
type _C_ulong uint32
type _C_float float32
type _C_uint uint32
type _C_sfWindowHandle _C_ulong
type _C_sfKeyCode uint32
type _C_sfJoyAxis uint32
type _C_sfVideoMode _Cstruct___0
type _C_int int32
type _Cstruct___2 struct {
	DepthBits		_C_uint
	StencilBits		_C_uint
	AntialiasingLevel	_C_uint
}
type _C_sfBool _C_int
type _Cunion___1 [20]byte
type _C_sfUint8 _C_unsignedchar
type _C_char int8
type _C_sfEvent _Cunion___1
type _C_sfMouseButton uint32
type _C_void [0]byte

func _C_sfWindow_GetSettings(*[0]byte) _C_sfWindowSettings
func _C_sfVideoMode_GetDesktopMode() _C_sfVideoMode
func _C_sfContext_Create() *[0]byte
func _C_sfWindow_GetHeight(*[0]byte) _C_uint
func _C_CString(string) *_C_char
func _C_sfContext_SetActive(*[0]byte, _C_sfBool)
func _C_sfWindow_UseVerticalSync(*[0]byte, _C_sfBool)
func _C_sfInput_IsJoystickButtonDown(*[0]byte, _C_uint, _C_uint) _C_sfBool
func _C_sfWindow_SetPosition(*[0]byte, _C_int, _C_int)
func _C_sfVideoMode_IsValid(_C_sfVideoMode) _C_sfBool
func _C_sfWindow_GetInput(*[0]byte) *[0]byte
func _C_sfWindow_Show(*[0]byte, _C_sfBool)
func _C_sfWindow_GetFrameTime(*[0]byte) _C_float
func _C_sfWindow_Close(*[0]byte)
func _C_sfWindow_Destroy(*[0]byte)
func _C_sfInput_GetMouseY(*[0]byte) _C_int
func _C_sfWindow_CreateFromHandle(_C_sfWindowHandle, _C_sfWindowSettings) *[0]byte
func _C_sfWindow_SetCursorPosition(*[0]byte, _C_uint, _C_uint)
func _C_sfWindow_Create(_C_sfVideoMode, *_C_char, _C_ulong, _C_sfWindowSettings) *[0]byte
func _C_sfWindow_SetIcon(*[0]byte, _C_uint, _C_uint, *_C_sfUint8)
func _C_sfWindow_SetJoystickThreshold(*[0]byte, _C_float)
func _C_sfWindow_SetActive(*[0]byte, _C_sfBool) _C_sfBool
func _C_sfInput_IsKeyDown(*[0]byte, _C_sfKeyCode) _C_sfBool
func _C_sfWindow_GetEvent(*[0]byte, *_C_sfEvent) _C_sfBool
func _C_sfInput_GetJoystickAxis(*[0]byte, _C_uint, _C_sfJoyAxis) _C_float
func _C_sfWindow_Display(*[0]byte)
func _C_sfWindow_SetFramerateLimit(*[0]byte, _C_uint)
func _C_sfWindow_IsOpened(*[0]byte) _C_sfBool
func _C_sfWindow_EnableKeyRepeat(*[0]byte, _C_sfBool)
func _C_sfVideoMode_GetMode(_C_size_t) _C_sfVideoMode
func _C_sfContext_Destroy(*[0]byte)
func _C_sfInput_GetMouseX(*[0]byte) _C_int
func _C_sfWindow_GetWidth(*[0]byte) _C_uint
func _C_sfWindow_ShowMouseCursor(*[0]byte, _C_sfBool)
func _C_sfInput_IsMouseButtonDown(*[0]byte, _C_sfMouseButton) _C_sfBool
func _C_sfWindow_SetSize(*[0]byte, _C_uint, _C_uint)
func _C_sfVideoMode_GetModesCount() _C_size_t
