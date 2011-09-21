// Created by cgo - DO NOT EDIT

//line win.go:1
package win

import (
	"runtime"
	"unsafe"
)

func SfBool2GoBool(cVal _Ctypedef_sfBool) bool {
	return int(cVal) == 1
}

func GoBool2SfBool(goVal bool) _Ctypedef_sfBool {
	if goVal {
		return _Ctypedef_sfBool(1)
	}
	return _Ctypedef_sfBool(0)
}

func NewVideoMode(val _Ctypedef_sfVideoMode) VideoMode {
	return VideoMode{val}
}

type VideoMode struct {
	Cref _Ctypedef_sfVideoMode
}

func NewContext(val *[0]byte) Context {
	tmp := Context{val}
	runtime.SetFinalizer(&tmp, (*Context).Destroy)
	return tmp
}

type Context struct {
	Cref *[0]byte
}

func NewInput(val *[0]byte) Input {
	tmp := Input{val}

	return tmp
}

type Input struct {
	Cref *[0]byte
}

func NewWindow(val *[0]byte) Window {
	tmp := Window{val}
	runtime.SetFinalizer(&tmp, (*Window).Destroy)
	return tmp
}

type Event struct {
	Cref *_Ctypedef_sfEvent
}

type Window struct {
	Cref *[0]byte
}

func NewWindowSettings(val _Ctypedef_sfWindowSettings) WindowSettings {
	return WindowSettings{val}
}

type WindowSettings struct {
	Cref _Ctypedef_sfWindowSettings
}

type WindowHandle struct {
	Cref _Ctypedef_sfWindowHandle
}

type KeyCode struct {
	Cref _Ctypedef_sfKeyCode
}

type JoyAxis struct {
	Cref _Ctypedef_sfJoyAxis
}

type MouseButton struct {
	Cref _Ctypedef_sfMouseButton
}

func CreateContext() Context {
	return NewContext(_Cfunc_sfContext_Create())
}

func (self *Context) Destroy() {
	_Cfunc_sfContext_Destroy(self.Cref)
	self.Cref = nil
}

func (self *Context) SetActive(active bool) {
	_Cfunc_sfContext_SetActive(self.Cref, GoBool2SfBool(active))
}

func (self *Input) IsKeyDown(keyCode KeyCode) bool {
	return SfBool2GoBool(_Cfunc_sfInput_IsKeyDown(self.Cref, keyCode.Cref))
}

func (self *Input) IsMouseButtonDown(button MouseButton) bool {
	return SfBool2GoBool(_Cfunc_sfInput_IsMouseButtonDown(self.Cref, button.Cref))
}

func (self *Input) IsJoystickButtonDown(joyId uint, button uint) bool {
	return SfBool2GoBool(_Cfunc_sfInput_IsJoystickButtonDown(self.Cref, _Ctype_uint(joyId), _Ctype_uint(button)))
}

func (self *Input) GetMouseX() int {
	return int(_Cfunc_sfInput_GetMouseX(self.Cref))
}

func (self *Input) GetMouseY() int {
	return int(_Cfunc_sfInput_GetMouseY(self.Cref))
}

func (self *Input) GetJoystickAxis(joyId uint, axis JoyAxis) float32 {
	return float32(_Cfunc_sfInput_GetJoystickAxis(self.Cref, _Ctype_uint(joyId), axis.Cref))
}

func GetDesktopMode() VideoMode {
	return NewVideoMode(_Cfunc_sfVideoMode_GetDesktopMode())
}

func (self *VideoMode) SetWidth(width uint) {
	self.Cref.Width = _Ctype_uint(width)
}

func (self *VideoMode) SetHeight(height uint) {
	self.Cref.Height = _Ctype_uint(height)
}

func (self *VideoMode) SetBitsPerPixel(bitsPerPixel uint) {
	self.Cref.BitsPerPixel = _Ctype_uint(bitsPerPixel)
}

func CreateVideoMode(height, width, bpp uint) VideoMode {
	tmp := GetDesktopMode()
	tmp.SetWidth(width)
	tmp.SetHeight(height)
	tmp.SetBitsPerPixel(bpp)
	return tmp
}

func GetMode(index uint) VideoMode {
	return NewVideoMode(_Cfunc_sfVideoMode_GetMode(_Ctypedef_size_t(index)))
}

func GetModesCount() int {
	return int(_Ctypedef_size_t(_Cfunc_sfVideoMode_GetModesCount()))
}

func (self *VideoMode) IsValid() bool {
	return SfBool2GoBool(_Cfunc_sfVideoMode_IsValid(self.Cref))
}

func CreateWindow(mode VideoMode, title string, style uint64, params WindowSettings) Window {
	return NewWindow(_Cfunc_sfWindow_Create(mode.Cref, _Cfunc_CString(title), _Ctype_ulong(style), params.Cref))
}

func (self *Window) CreateFromHandle(handle WindowHandle, params WindowSettings) Window {
	return NewWindow(_Cfunc_sfWindow_CreateFromHandle(handle.Cref, params.Cref))
}

func (self *Window) Destroy() {
	_Cfunc_sfWindow_Destroy(self.Cref)
}

func (self *Window) Close() {
	_Cfunc_sfWindow_Close(self.Cref)
}

func (self *Window) IsOpened() bool {
	return SfBool2GoBool(_Cfunc_sfWindow_IsOpened(self.Cref))
}

func (self *Window) GetWidth() uint {
	return uint(_Cfunc_sfWindow_GetWidth(self.Cref))
}

func (self *Window) GetHeight() uint {
	return uint(_Cfunc_sfWindow_GetHeight(self.Cref))
}

func (self *Window) GetSettings() WindowSettings {
	return NewWindowSettings(_Cfunc_sfWindow_GetSettings(self.Cref))
}

func CreateWindowSettings(depthBits, stencilBits, antialiasingLevel uint) WindowSettings {
	return NewWindowSettings(_Ctypedef_sfWindowSettings{_Ctype_uint(depthBits), _Ctype_uint(stencilBits), _Ctype_uint(antialiasingLevel)})
}

func (self *Window) GetEvent(event Event) bool {
	return SfBool2GoBool(_Cfunc_sfWindow_GetEvent(self.Cref, event.Cref))
}

func (self *Window) UseVerticalSync(enabled bool) {
	_Cfunc_sfWindow_UseVerticalSync(self.Cref, GoBool2SfBool(enabled))
}

func (self *Window) ShowMouseCursor(show bool) {
	_Cfunc_sfWindow_ShowMouseCursor(self.Cref, GoBool2SfBool(show))
}

func (self *Window) SetCursorPosition(left uint, top uint) {
	_Cfunc_sfWindow_SetCursorPosition(self.Cref, _Ctype_uint(left), _Ctype_uint(top))
}

func (self *Window) SetPosition(left int, top int) {
	_Cfunc_sfWindow_SetPosition(self.Cref, _Ctype_int(left), _Ctype_int(top))
}

func (self *Window) SetSize(width uint, height uint) {
	_Cfunc_sfWindow_SetSize(self.Cref, _Ctype_uint(width), _Ctype_uint(height))
}

func (self *Window) Show(state bool) {
	_Cfunc_sfWindow_Show(self.Cref, GoBool2SfBool(state))
}

func (self *Window) EnableKeyRepeat(enabled bool) {
	_Cfunc_sfWindow_EnableKeyRepeat(self.Cref, GoBool2SfBool(enabled))
}

func (self *Window) SetIcon(width uint, height uint, pixels *uint8) {
	pixelsPtr := (*_Ctypedef_sfUint8)(unsafe.Pointer(pixels))
	_Cfunc_sfWindow_SetIcon(self.Cref, _Ctype_uint(width), _Ctype_uint(height), pixelsPtr)
}

func (self *Window) SetActive(active bool) bool {
	return SfBool2GoBool(_Cfunc_sfWindow_SetActive(self.Cref, GoBool2SfBool(active)))
}

func (self *Window) Display() {
	_Cfunc_sfWindow_Display(self.Cref)
}

func (self *Window) GetInput() Input {
	return NewInput(_Cfunc_sfWindow_GetInput(self.Cref))
}

func (self *Window) SetFramerateLimit(limit uint) {
	_Cfunc_sfWindow_SetFramerateLimit(self.Cref, _Ctype_uint(limit))
}

func (self *Window) GetFrameTime() float32 {
	return float32(_Cfunc_sfWindow_GetFrameTime(self.Cref))
}

func (self *Window) SetJoystickThreshold(threshold float32) {
	_Cfunc_sfWindow_SetJoystickThreshold(self.Cref, _Ctype_float(threshold))
}
