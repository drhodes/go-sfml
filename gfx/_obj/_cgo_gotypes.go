// Created by cgo - DO NOT EDIT

package gfx

import "unsafe"

import "os"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *os.Error, x int) { *dst = os.Errno(x) }
type _Ctypedef_sfColor _Ctype_struct___2
type _Ctypedef_sfIntRect _Ctype_struct___5
type _Ctype_struct___3 struct {
	Left	_Ctype_float
	Top	_Ctype_float
	Right	_Ctype_float
	Bottom	_Ctype_float
}
type _Ctypedef_sfFloatRect _Ctype_struct___3
type _Ctype_uint uint32
type _Ctypedef_size_t _Ctype_uint
type _Ctypedef_sfBlendMode uint32
type _Ctype_unsignedchar uint8
type _Ctype_char int8
type _Ctype_struct___2 struct {
	r	_Ctypedef_sfUint8
	g	_Ctypedef_sfUint8
	b	_Ctypedef_sfUint8
	a	_Ctypedef_sfUint8
}
type _Ctypedef_sfBool _Ctype_int
type _Ctypedef_sfUint32 _Ctype_uint
type _Ctypedef_sfColor_P _Ctype_struct___4
type _Ctype_struct___5 struct {
	Left	_Ctype_int
	Top	_Ctype_int
	Right	_Ctype_int
	Bottom	_Ctype_int
}
type _Ctypedef_sfUint8 _Ctype_unsignedchar
type _Ctype_struct___1 struct {
	DepthBits		_Ctype_uint
	StencilBits		_Ctype_uint
	AntialiasingLevel	_Ctype_uint
}
type _Ctypedef_sfWindowSettings _Ctype_struct___1
type _Ctypedef_sfVideoMode _Ctype_struct___0
type _Ctype_struct___4 struct {
	R	_Ctypedef_sfUint8
	G	_Ctypedef_sfUint8
	B	_Ctypedef_sfUint8
	A	_Ctypedef_sfUint8
}
type _Ctype_ulong uint32
type _Ctype_struct___0 struct {
	Width		_Ctype_uint
	Height		_Ctype_uint
	BitsPerPixel	_Ctype_uint
}
type _Ctype_float float32
type _Ctype_int int32
type _Ctype_void [0]byte

func _Cfunc_sfSprite_SetScaleX(*[0]byte, _Ctype_float)
func _Cfunc_sfRenderWindow_Create(_Ctypedef_sfVideoMode, *_Ctype_char, _Ctype_ulong, _Ctypedef_sfWindowSettings) *[0]byte
func _Cfunc_sfShape_GetColor(*[0]byte) _Ctypedef_sfColor
func _Cfunc_sfRenderWindow_SetActive(*[0]byte, _Ctypedef_sfBool) _Ctypedef_sfBool
func _Cfunc_sfString_SetColor(*[0]byte, _Ctypedef_sfColor)
func _Cfunc_sfShape_SetCenter(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfSprite_GetCenterX(*[0]byte) _Ctype_float
func _Cfunc_sfView_Zoom(*[0]byte, _Ctype_float)
func _Cfunc_sfRenderWindow_SetView(*[0]byte, *[0]byte)
func _Cfunc_sfView_SetHalfSize(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfShape_SetX(*[0]byte, _Ctype_float)
func _Cfunc_sfRenderWindow_UseVerticalSync(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfImage_CreateMaskFromColor(*[0]byte, _Ctypedef_sfColor, _Ctypedef_sfUint8)
func _Cfunc_sfShape_CreateCircle(_Ctype_float, _Ctype_float, _Ctype_float, _Ctypedef_sfColor, _Ctype_float, _Ctypedef_sfColor) *[0]byte
func _Cfunc_sfView_GetRect(*[0]byte) _Ctypedef_sfFloatRect
func _Cfunc_sfRenderWindow_SetCursorPosition(*[0]byte, _Ctype_uint, _Ctype_uint)
func _Cfunc_sfColor_FromRGB(_Ctypedef_sfUint8, _Ctypedef_sfUint8, _Ctypedef_sfUint8) _Ctypedef_sfColor
func _Cfunc_sfShape_CreateRectangle(_Ctype_float, _Ctype_float, _Ctype_float, _Ctype_float, _Ctypedef_sfColor, _Ctype_float, _Ctypedef_sfColor) *[0]byte
func _Cfunc_sfFloatRect_Intersects(*_Ctypedef_sfFloatRect, *_Ctypedef_sfFloatRect, *_Ctypedef_sfFloatRect) _Ctypedef_sfBool
func _Cfunc_sfShape_SetPointPosition(*[0]byte, _Ctype_uint, _Ctype_float, _Ctype_float)
func _Cfunc_sfView_CreateFromRect(_Ctypedef_sfFloatRect) *[0]byte
func _Cfunc_sfUnwraP(_Ctypedef_sfColor_P) _Ctypedef_sfColor
func _Cfunc_sfRenderWindow_SetJoystickThreshold(*[0]byte, _Ctype_float)
func _Cfunc_sfRenderWindow_Destroy(*[0]byte)
func _Cfunc_sfShape_GetScaleX(*[0]byte) _Ctype_float
func _Cfunc_sfSprite_SetRotation(*[0]byte, _Ctype_float)
func _Cfunc_sfImage_CreateFromFile(*_Ctype_char) *[0]byte
func _Cfunc_sfString_GetCenterX(*[0]byte) _Ctype_float
func _Cfunc_sfSprite_FlipX(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfView_SetFromRect(*[0]byte, _Ctypedef_sfFloatRect)
func _Cfunc_sfImage_SetSmooth(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfPostFX_SetParameter4(*[0]byte, *_Ctype_char, _Ctype_float, _Ctype_float, _Ctype_float, _Ctype_float)
func _Cfunc_sfSprite_GetPixel(*[0]byte, _Ctype_uint, _Ctype_uint) _Ctypedef_sfColor
func _Cfunc_sfView_GetHalfSizeX(*[0]byte) _Ctype_float
func _Cfunc_sfString_SetScaleY(*[0]byte, _Ctype_float)
func _Cfunc_sfShape_GetY(*[0]byte) _Ctype_float
func _Cfunc_sfShape_EnableFill(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfSprite_Scale(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_CString(string) *_Ctype_char
func _Cfunc_sfIntRect_Offset(*_Ctypedef_sfIntRect, _Ctype_int, _Ctype_int)
func _Cfunc_sfSprite_GetImage(*[0]byte) *[0]byte
func _Cfunc_sfRenderWindow_DrawSprite(*[0]byte, *[0]byte)
func _Cfunc_sfRenderWindow_GetSettings(*[0]byte) _Ctypedef_sfWindowSettings
func _Cfunc_sfSprite_SetSubRect(*[0]byte, _Ctypedef_sfIntRect)
func _Cfunc_sfRenderWindow_Show(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfRenderWindow_Capture(*[0]byte) *[0]byte
func _Cfunc_sfString_Create() *[0]byte
func _Cfunc_sfShape_GetCenterX(*[0]byte) _Ctype_float
func _Cfunc_sfShape_Move(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfString_SetUnicodeText(*[0]byte, *_Ctypedef_sfUint32)
func _Cfunc_sfSprite_SetScale(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfSprite_GetColor(*[0]byte) _Ctypedef_sfColor
func _Cfunc_sfSprite_GetCenterY(*[0]byte) _Ctype_float
func _Cfunc_sfImage_CreateFromMemory(*_Ctype_char, _Ctypedef_size_t) *[0]byte
func _Cfunc_sfString_GetFont(*[0]byte) *[0]byte
func _Cfunc_sfString_GetUnicodeText(*[0]byte) *_Ctypedef_sfUint32
func _Cfunc_sfSprite_SetY(*[0]byte, _Ctype_float)
func _Cfunc_sfPostFX_CreateFromFile(*_Ctype_char) *[0]byte
func _Cfunc_sfString_TransformToGlobal(*[0]byte, _Ctype_float, _Ctype_float, *_Ctype_float, *_Ctype_float)
func _Cfunc_sfShape_GetRotation(*[0]byte) _Ctype_float
func _Cfunc_sfImage_SetPixel(*[0]byte, _Ctype_uint, _Ctype_uint, _Ctypedef_sfColor)
func _Cfunc_sfSprite_Rotate(*[0]byte, _Ctype_float)
func _Cfunc_sfShape_AddPoint(*[0]byte, _Ctype_float, _Ctype_float, _Ctypedef_sfColor, _Ctypedef_sfColor)
func _Cfunc_sfSprite_SetColor(*[0]byte, _Ctypedef_sfColor)
func _Cfunc_sfRenderWindow_GetDefaultView(*[0]byte) *[0]byte
func _Cfunc_sfShape_Scale(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfShape_GetNbPoints(*[0]byte) _Ctype_uint
func _Cfunc_sfShape_GetBlendMode(*[0]byte) _Ctypedef_sfBlendMode
func _Cfunc_sfString_GetCenterY(*[0]byte) _Ctype_float
func _Cfunc_sfSprite_FlipY(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfSprite_SetPosition(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfShape_GetPointOutlineColor(*[0]byte, _Ctype_uint) _Ctypedef_sfColor
func _Cfunc_sfSprite_GetRotation(*[0]byte) _Ctype_float
func _Cfunc_sfPostFX_CanUsePostFX() _Ctypedef_sfBool
func _Cfunc_sfString_GetRect(*[0]byte) _Ctypedef_sfFloatRect
func _Cfunc_sfString_SetScaleX(*[0]byte, _Ctype_float)
func _Cfunc_sfSprite_Resize(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfRenderWindow_GetWidth(*[0]byte) _Ctype_uint
func _Cfunc_sfRenderWindow_Close(*[0]byte)
func _Cfunc_sfShape_GetX(*[0]byte) _Ctype_float
func _Cfunc_sfPostFX_SetTexture(*[0]byte, *_Ctype_char, *[0]byte)
func _Cfunc_sfColor_Modulate(_Ctypedef_sfColor, _Ctypedef_sfColor) _Ctypedef_sfColor
func _Cfunc_sfView_Move(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfRenderWindow_SetSize(*[0]byte, _Ctype_uint, _Ctype_uint)
func _Cfunc_sfIntRect_Contains(*_Ctypedef_sfIntRect, _Ctype_int, _Ctype_int) _Ctypedef_sfBool
func _Cfunc_sfFont_Create() *[0]byte
func _Cfunc_sfPostFX_SetParameter3(*[0]byte, *_Ctype_char, _Ctype_float, _Ctype_float, _Ctype_float)
func _Cfunc_sfColor_FromRGB_P(_Ctypedef_sfUint8, _Ctypedef_sfUint8, _Ctypedef_sfUint8) _Ctypedef_sfColor_P
func _Cfunc_sfRenderWindow_GetHeight(*[0]byte) _Ctype_uint
func _Cfunc_sfShape_SetScale(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfString_SetScale(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfShape_GetCenterY(*[0]byte) _Ctype_float
func _Cfunc_sfFloatRect_Contains(*_Ctypedef_sfFloatRect, _Ctype_float, _Ctype_float) _Ctypedef_sfBool
func _Cfunc_sfShape_Rotate(*[0]byte, _Ctype_float)
func _Cfunc_sfImage_Copy(*[0]byte, *[0]byte, _Ctype_uint, _Ctype_uint, _Ctypedef_sfIntRect)
func _Cfunc_sfPostFX_Destroy(*[0]byte)
func _Cfunc_sfView_Destroy(*[0]byte)
func _Cfunc_sfShape_SetScaleY(*[0]byte, _Ctype_float)
func _Cfunc_sfShape_SetBlendMode(*[0]byte, _Ctypedef_sfBlendMode)
func _Cfunc_sfString_Destroy(*[0]byte)
func _Cfunc_sfSprite_SetCenter(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfRenderWindow_DrawShape(*[0]byte, *[0]byte)
func _Cfunc_sfSprite_SetX(*[0]byte, _Ctype_float)
func _Cfunc_sfSprite_Create() *[0]byte
func _Cfunc_sfShape_EnableOutline(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfColor_FromRGBA(_Ctypedef_sfUint8, _Ctypedef_sfUint8, _Ctypedef_sfUint8, _Ctypedef_sfUint8) _Ctypedef_sfColor
func _Cfunc_sfSprite_GetScaleY(*[0]byte) _Ctype_float
func _Cfunc_sfRenderWindow_EnableKeyRepeat(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfImage_Bind(*[0]byte)
func _Cfunc_sfImage_GetHeight(*[0]byte) _Ctype_uint
func _Cfunc_sfSprite_Move(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfShape_SetPosition(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfShape_GetOutlineWidth(*[0]byte) _Ctype_float
func _Cfunc_sfImage_Destroy(*[0]byte)
func _Cfunc_sfShape_Destroy(*[0]byte)
func _Cfunc_sfString_GetScaleY(*[0]byte) _Ctype_float
func _Cfunc_sfImage_CreateFromColor(_Ctype_uint, _Ctype_uint, _Ctypedef_sfColor) *[0]byte
func _Cfunc_sfRenderWindow_DrawString(*[0]byte, *[0]byte)
func _Cfunc_sfShape_Create() *[0]byte
func _Cfunc_sfString_SetSize(*[0]byte, _Ctype_float)
func _Cfunc_sfShape_SetPointColor(*[0]byte, _Ctype_uint, _Ctypedef_sfColor)
func _Cfunc_sfRenderWindow_IsOpened(*[0]byte) _Ctypedef_sfBool
func _Cfunc_sfSprite_SetImage(*[0]byte, *[0]byte)
func _Cfunc_sfImage_IsSmooth(*[0]byte) _Ctypedef_sfBool
func _Cfunc_sfSprite_GetY(*[0]byte) _Ctype_float
func _Cfunc_sfImage_CopyScreen(*[0]byte, *[0]byte, _Ctypedef_sfIntRect) _Ctypedef_sfBool
func _Cfunc_sfView_SetCenter(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfString_SetPosition(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfPostFX_SetParameter2(*[0]byte, *_Ctype_char, _Ctype_float, _Ctype_float)
func _Cfunc_sfFont_CreateFromFile(*_Ctype_char, _Ctype_uint, *_Ctypedef_sfUint32) *[0]byte
func _Cfunc_sfView_GetCenterX(*[0]byte) _Ctype_float
func _Cfunc_sfString_SetCenter(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfString_GetY(*[0]byte) _Ctype_float
func _Cfunc_sfRenderWindow_Clear(*[0]byte, _Ctypedef_sfColor)
func _Cfunc_sfRenderWindow_PreserveOpenGLStates(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfShape_SetOutlineWidth(*[0]byte, _Ctype_float)
func _Cfunc_sfRenderWindow_DrawPostFX(*[0]byte, *[0]byte)
func _Cfunc_sfRenderWindow_ShowMouseCursor(*[0]byte, _Ctypedef_sfBool)
func _Cfunc_sfString_SetY(*[0]byte, _Ctype_float)
func _Cfunc_sfImage_GetWidth(*[0]byte) _Ctype_uint
func _Cfunc_sfShape_SetRotation(*[0]byte, _Ctype_float)
func _Cfunc_sfShape_SetScaleX(*[0]byte, _Ctype_float)
func _Cfunc_sfString_GetBlendMode(*[0]byte) _Ctypedef_sfBlendMode
func _Cfunc_sfIntRect_Intersects(*_Ctypedef_sfIntRect, *_Ctypedef_sfIntRect, *_Ctypedef_sfIntRect) _Ctypedef_sfBool
func _Cfunc_sfString_GetColor(*[0]byte) _Ctypedef_sfColor
func _Cfunc_sfImage_Create() *[0]byte
func _Cfunc_sfImage_CreateFromPixels(_Ctype_uint, _Ctype_uint, *_Ctypedef_sfUint8) *[0]byte
func _Cfunc_sfFont_GetDefaultFont() *[0]byte
func _Cfunc_sfRenderWindow_SetPosition(*[0]byte, _Ctype_int, _Ctype_int)
func _Cfunc_sfSprite_SetScaleY(*[0]byte, _Ctype_float)
func _Cfunc_sfString_SetStyle(*[0]byte, _Ctype_ulong)
func _Cfunc_sfShape_SetPointOutlineColor(*[0]byte, _Ctype_uint, _Ctypedef_sfColor)
func _Cfunc_sfView_Create() *[0]byte
func _Cfunc_sfImage_GetPixel(*[0]byte, _Ctype_uint, _Ctype_uint) _Ctypedef_sfColor
func _Cfunc_sfShape_CreateLine(_Ctype_float, _Ctype_float, _Ctype_float, _Ctype_float, _Ctype_float, _Ctypedef_sfColor, _Ctype_float, _Ctypedef_sfColor) *[0]byte
func _Cfunc_sfSprite_GetScaleX(*[0]byte) _Ctype_float
func _Cfunc_sfFont_Destroy(*[0]byte)
func _Cfunc_sfShape_SetY(*[0]byte, _Ctype_float)
func _Cfunc_sfString_Rotate(*[0]byte, _Ctype_float)
func _Cfunc_sfSprite_GetBlendMode(*[0]byte) _Ctypedef_sfBlendMode
func _Cfunc_sfShape_GetPointColor(*[0]byte, _Ctype_uint) _Ctypedef_sfColor
func _Cfunc_sfString_SetText(*[0]byte, *_Ctype_char)
func _Cfunc_sfString_GetSize(*[0]byte) _Ctype_float
func _Cfunc_sfFloatRect_Offset(*_Ctypedef_sfFloatRect, _Ctype_float, _Ctype_float)
func _Cfunc_sfString_GetScaleX(*[0]byte) _Ctype_float
func _Cfunc_sfShape_GetScaleY(*[0]byte) _Ctype_float
func _Cfunc_sfString_Scale(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfSprite_GetHeight(*[0]byte) _Ctype_float
func _Cfunc_sfColor_Add(_Ctypedef_sfColor, _Ctypedef_sfColor) _Ctypedef_sfColor
func _Cfunc_sfString_SetBlendMode(*[0]byte, _Ctypedef_sfBlendMode)
func _Cfunc_sfImage_SaveToFile(*[0]byte, *_Ctype_char) _Ctypedef_sfBool
func _Cfunc_sfSprite_GetX(*[0]byte) _Ctype_float
func _Cfunc_sfRenderWindow_Display(*[0]byte)
func _Cfunc_sfString_TransformToLocal(*[0]byte, _Ctype_float, _Ctype_float, *_Ctype_float, *_Ctype_float)
func _Cfunc_sfRenderWindow_SetFramerateLimit(*[0]byte, _Ctype_uint)
func _Cfunc_sfView_GetCenterY(*[0]byte) _Ctype_float
func _Cfunc_sfString_GetX(*[0]byte) _Ctype_float
func _Cfunc_sfView_GetHalfSizeY(*[0]byte) _Ctype_float
func _Cfunc_sfSprite_SetBlendMode(*[0]byte, _Ctypedef_sfBlendMode)
func _Cfunc_sfSprite_Destroy(*[0]byte)
func _Cfunc_sfString_SetX(*[0]byte, _Ctype_float)
func _Cfunc_sfShape_SetColor(*[0]byte, _Ctypedef_sfColor)
func _Cfunc_sfString_GetRotation(*[0]byte) _Ctype_float
func _Cfunc_sfPostFX_SetParameter1(*[0]byte, *_Ctype_char, _Ctype_float)
func _Cfunc_sfString_Move(*[0]byte, _Ctype_float, _Ctype_float)
func _Cfunc_sfSprite_GetWidth(*[0]byte) _Ctype_float
func _Cfunc_sfSprite_GetSubRect(*[0]byte) _Ctypedef_sfIntRect
func _Cfunc_sfString_SetFont(*[0]byte, *[0]byte)
func _Cfunc_sfString_SetRotation(*[0]byte, _Ctype_float)
