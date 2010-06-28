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
type _C_sfUint32 _C_uint
type _C_sfIntRect _Cstruct___2
type _C_sfBlendMode uint32
type _C_sfColor _Cstruct___0
type _C_size_t _C_uint
type _Cstruct___1 struct {
	Left	_C_float
	Top	_C_float
	Right	_C_float
	Bottom	_C_float
}
type _C_ulong uint32
type _C_float float32
type _C_uint uint32
type _C_int int32
type _Cstruct___2 struct {
	Left	_C_int
	Top	_C_int
	Right	_C_int
	Bottom	_C_int
}
type _C_sfBool _C_int
type _C_sfUint8 _C_unsignedchar
type _C_char int8
type _C_sfFloatRect _Cstruct___1
type _C_void [0]byte

func _C_sfString_SetColor(*[0]byte, _C_sfColor)
func _C_sfImage_CreateMaskFromColor(*[0]byte, _C_sfColor, _C_sfUint8)
func _C_sfColor_FromRGB(_C_sfUint8, _C_sfUint8, _C_sfUint8) _C_sfColor
func _C_sfFloatRect_Intersects(*_C_sfFloatRect, *_C_sfFloatRect, *_C_sfFloatRect) _C_sfBool
func _C_sfImage_CreateFromFile(*_C_char) *[0]byte
func _C_sfString_GetCenterX(*[0]byte) _C_float
func _C_sfImage_SetSmooth(*[0]byte, _C_sfBool)
func _C_sfString_GetCharacterPos(*[0]byte, _C_size_t, *_C_float, *_C_float)
func _C_sfString_SetScaleY(*[0]byte, _C_float)
func _C_CString(string) *_C_char
func _C_sfIntRect_Offset(*_C_sfIntRect, _C_int, _C_int)
func _C_sfString_Create() *[0]byte
func _C_sfString_SetUnicodeText(*[0]byte, *_C_sfUint32)
func _C_sfImage_CreateFromMemory(*_C_char, _C_size_t) *[0]byte
func _C_sfString_GetFont(*[0]byte) *[0]byte
func _C_sfString_GetUnicodeText(*[0]byte) *_C_sfUint32
func _C_sfString_TransformToGlobal(*[0]byte, _C_float, _C_float, *_C_float, *_C_float)
func _C_sfImage_SetPixel(*[0]byte, _C_uint, _C_uint, _C_sfColor)
func _C_sfString_GetCenterY(*[0]byte) _C_float
func _C_sfString_GetRect(*[0]byte) _C_sfFloatRect
func _C_sfString_SetScaleX(*[0]byte, _C_float)
func _C_sfColor_Modulate(_C_sfColor, _C_sfColor) _C_sfColor
func _C_sfIntRect_Contains(*_C_sfIntRect, _C_int, _C_int) _C_sfBool
func _C_sfFont_Create() *[0]byte
func _C_sfString_SetScale(*[0]byte, _C_float, _C_float)
func _C_sfFloatRect_Contains(*_C_sfFloatRect, _C_float, _C_float) _C_sfBool
func _C_sfImage_Copy(*[0]byte, *[0]byte, _C_uint, _C_uint, _C_sfIntRect)
func _C_sfString_Destroy(*[0]byte)
func _C_sfColor_FromRGBA(_C_sfUint8, _C_sfUint8, _C_sfUint8, _C_sfUint8) _C_sfColor
func _C_sfImage_Bind(*[0]byte)
func _C_sfImage_GetHeight(*[0]byte) _C_uint
func _C_sfImage_Destroy(*[0]byte)
func _C_sfString_GetScaleY(*[0]byte) _C_float
func _C_sfImage_CreateFromColor(_C_uint, _C_uint, _C_sfColor) *[0]byte
func _C_sfString_SetSize(*[0]byte, _C_float)
func _C_sfImage_IsSmooth(*[0]byte) _C_sfBool
func _C_sfImage_CopyScreen(*[0]byte, *[0]byte, _C_sfIntRect) _C_sfBool
func _C_sfString_SetPosition(*[0]byte, _C_float, _C_float)
func _C_sfFont_CreateFromFile(*_C_char, _C_uint, *_C_sfUint32) *[0]byte
func _C_sfString_SetCenter(*[0]byte, _C_float, _C_float)
func _C_sfString_GetY(*[0]byte) _C_float
func _C_sfString_SetY(*[0]byte, _C_float)
func _C_sfImage_GetWidth(*[0]byte) _C_uint
func _C_sfString_GetBlendMode(*[0]byte) _C_sfBlendMode
func _C_sfIntRect_Intersects(*_C_sfIntRect, *_C_sfIntRect, *_C_sfIntRect) _C_sfBool
func _C_sfString_GetColor(*[0]byte) _C_sfColor
func _C_sfImage_Create() *[0]byte
func _C_sfImage_CreateFromPixels(_C_uint, _C_uint, *_C_sfUint8) *[0]byte
func _C_sfFont_GetDefaultFont() *[0]byte
func _C_sfString_SetStyle(*[0]byte, _C_ulong)
func _C_sfImage_GetPixel(*[0]byte, _C_uint, _C_uint) _C_sfColor
func _C_sfFont_Destroy(*[0]byte)
func _C_sfString_Rotate(*[0]byte, _C_float)
func _C_sfString_SetText(*[0]byte, *_C_char)
func _C_sfString_GetSize(*[0]byte) _C_float
func _C_sfFloatRect_Offset(*_C_sfFloatRect, _C_float, _C_float)
func _C_sfString_GetScaleX(*[0]byte) _C_float
func _C_sfString_Scale(*[0]byte, _C_float, _C_float)
func _C_sfColor_Add(_C_sfColor, _C_sfColor) _C_sfColor
func _C_sfString_SetBlendMode(*[0]byte, _C_sfBlendMode)
func _C_sfImage_SaveToFile(*[0]byte, *_C_char) _C_sfBool
func _C_sfString_TransformToLocal(*[0]byte, _C_float, _C_float, *_C_float, *_C_float)
func _C_sfString_GetX(*[0]byte) _C_float
func _C_sfString_SetX(*[0]byte, _C_float)
func _C_sfString_GetRotation(*[0]byte) _C_float
func _C_sfString_Move(*[0]byte, _C_float, _C_float)
func _C_sfString_SetFont(*[0]byte, *[0]byte)
func _C_sfString_SetRotation(*[0]byte, _C_float)
