// Created by cgo - DO NOT EDIT
//line win.go:1
package win

// #include <SFML/Window/Context.h>
// #include <SFML/Window/Input.h>
// #include <SFML/Window/VideoMode.h>
// #include <SFML/Window/Window.h>


import (
	"runtime"
	"unsafe"
)

// utility functions must be from in this package to make happy cgo.
func SfBool2GoBool(cVal _C_sfBool) bool {
	return int(cVal) == 1
}

func GoBool2SfBool(goVal bool) _C_sfBool {
	if goVal {
		return _C_sfBool(1)
	}
	return _C_sfBool(0)
}

/*
func NewColor(val C.sfColor) Color {
	return Color{ val }
}

type Color struct {
	Cref C.sfColor
}
*/
func NewVideoMode(val _C_sfVideoMode) VideoMode {
	return VideoMode{val}
}

type VideoMode struct {
	Cref _C_sfVideoMode
}

//-------------------------------------
func NewContext(val *[0]byte) Context {
	tmp := Context{val}
	runtime.SetFinalizer(&tmp, (*Context).Destroy)
	return tmp
}

type Context struct {
	Cref *[0]byte
}
//-------------------------------------
func NewInput(val *[0]byte) Input {
	tmp := Input{val}
	//runtime.SetFinalizer(&tmp, (*Input).Destroy)
	return tmp
}

type Input struct {
	Cref *[0]byte
}


//-------------------------------------
func NewWindow(val *[0]byte) Window {
	tmp := Window{val}
	runtime.SetFinalizer(&tmp, (*Window).Destroy)
	return tmp
}


type Event struct {
	Cref *_C_sfEvent
}


type Window struct {
	Cref *[0]byte
}


func NewWindowSettings(val _C_sfWindowSettings) WindowSettings {
	return WindowSettings{val}
}

type WindowSettings struct {
	Cref _C_sfWindowSettings
}

type WindowHandle struct {
	Cref _C_sfWindowHandle
}

type KeyCode struct {
	Cref _C_sfKeyCode
}

type JoyAxis struct {
	Cref _C_sfJoyAxis
}

type MouseButton struct {
	Cref _C_sfMouseButton
}


// _Context_
// -------------------------------------------------------------------------------
/// Construct a new context
/// return New context
// sfContext* sfContext_Create();
func CreateContext() Context {
	return NewContext(_C_sfContext_Create())
}
/// Destroy an existing context
/// param Context : Context to destroy
// void sfContext_Destroy(sfContext* Context);
func (self *Context) Destroy() {
	_C_sfContext_Destroy(self.Cref)
	self.Cref = nil
}

/// Activate or deactivate a context
/// param Context : Context to activate or deactivate
/// param Active : sfTrue to activate, sfFalse to deactivate
// void sfContext_SetActive(sfContext* Context, sfBool Active);
func (self *Context) SetActive(active bool) {
	_C_sfContext_SetActive(self.Cref, GoBool2SfBool(active))
}


// _Input_
// -------------------------------------------------------------------------------
// Get the state of a key
// param Input : Input object
// param KeyCode : Key to check
// return sfTrue if key is down, sfFalse if key is up
// sfBool sfInput_IsKeyDown(sfInput* Input, sfKeyCode KeyCode);
func (self *Input) IsKeyDown(keyCode KeyCode) bool {
	return SfBool2GoBool(_C_sfInput_IsKeyDown(self.Cref, keyCode.Cref))
}
// Get the state of a mouse button
// param Input : Input object
// param Button : Button to check
// return sfTrue if button is down, sfFalse if button is up
// sfBool sfInput_IsMouseButtonDown(sfInput* Input, sfMouseButton Button);
func (self *Input) IsMouseButtonDown(button MouseButton) bool {
	return SfBool2GoBool(_C_sfInput_IsMouseButtonDown(self.Cref, button.Cref))
}
// Get the state of a joystick button
// param Input : Input object
// param JoyId : Identifier of the joystick to check (0 or 1)
// param Button : Button to check
// return sfTrue if button is down, sfFalse if button is up
// sfBool sfInput_IsJoystickButtonDown(sfInput* Input, unsigned int JoyId, unsigned int Button);
func (self *Input) IsJoystickButtonDown(joyId uint, button uint) bool {
	return SfBool2GoBool(_C_sfInput_IsJoystickButtonDown(self.Cref, _C_uint(joyId), _C_uint(button)))
}
// Get the mouse X position
// param Input : Input object
// return Current mouse left position, relative to owner window
// int sfInput_GetMouseX(sfInput* Input);
func (self *Input) GetMouseX() int {
	return int(_C_sfInput_GetMouseX(self.Cref))
}
// Get the mouse Y position
// param Input : Input object
// return Current mouse top position, relative to owner window
// int sfInput_GetMouseY(sfInput* Input);
func (self *Input) GetMouseY() int {
	return int(_C_sfInput_GetMouseY(self.Cref))
}
// Get the joystick position on a given axis
// param Input : Input object
// param JoyId : Identifier of the joystick to check (0 or 1)
// param Axis : Identifier of the axis to read
// return Current joystick position, in the range [-100, 100]
// float sfInput_GetJoystickAxis(sfInput* Input, unsigned int JoyId, sfJoyAxis Axis);
func (self *Input) GetJoystickAxis(joyId uint, axis JoyAxis) float {
	return float(_C_sfInput_GetJoystickAxis(self.Cref, _C_uint(joyId), axis.Cref))
}


// _VideoMode_
// -------------------------------------------------------------------------------
// Get the current desktop video mode
// return Current desktop video mode
// sfVideoMode sfVideoMode_GetDesktopMode();
func GetDesktopMode() VideoMode {
	return NewVideoMode(_C_sfVideoMode_GetDesktopMode())
}

func (self *VideoMode) SetWidth(width uint) {
	self.Cref.Width = _C_uint(width)
}

func (self *VideoMode) SetHeight(height uint) {
	self.Cref.Height = _C_uint(height)
}

func (self *VideoMode) SetBitsPerPixel(bitsPerPixel uint) {
	self.Cref.BitsPerPixel = _C_uint(bitsPerPixel)
}

// Create A VideoMode  --  ?? Consider adding error handling to this
// param: height, with, bits-per-pixel are uints
// return: VideoMode
func CreateVideoMode(height, width, bpp uint) VideoMode {
	tmp := GetDesktopMode()
	tmp.SetWidth(width)
	tmp.SetHeight(height)
	tmp.SetBitsPerPixel(bpp)
	return tmp
}

// Get a valid video mode
// Index must be in range [0, GetModesCount()[
// Modes are sorted from best to worst
// param Index : Index of video mode to get
// return Corresponding video mode (invalid mode if index is out of range)
// sfVideoMode sfVideoMode_GetMode(size_t Index);
func GetMode(index uint) VideoMode {
	return NewVideoMode(_C_sfVideoMode_GetMode(_C_size_t(index)))
}

// Get valid video modes count
// return Number of valid video modes available
// size_t sfVideoMode_GetModesCount();
func GetModesCount() int {
	return int(_C_size_t(_C_sfVideoMode_GetModesCount()))
}

// Tell whether or not a video mode is supported
// param Mode : Video mode to check
// return True if video mode is supported, false otherwise
// sfBool sfVideoMode_IsValid(sfVideoMode Mode);
func (self *VideoMode) IsValid() bool {
	return SfBool2GoBool(_C_sfVideoMode_IsValid(self.Cref))
}


// _Window_
// -------------------------------------------------------------------------------
// Structure defining the window's creation settings
// Construct a new window
// param Mode : Video mode to use
// param Title : Title of the window
// param Style : Window style
// param Params : Creation settings
// sfWindow* sfWindow_Create(sfVideoMode Mode, const char* Title, unsigned long Style, sfWindowSettings Params);
func CreateWindow(mode VideoMode, title string, style uint64, params WindowSettings) Window {
	return NewWindow(_C_sfWindow_Create(mode.Cref, _C_CString(title), _C_ulong(style), params.Cref))
}

// Construct a window from an existing control
// param Handle : Platform-specific handle of the control
// param Params : Creation settings
// sfWindow* sfWindow_CreateFromHandle(sfWindowHandle Handle, sfWindowSettings Params);
func (self *Window) CreateFromHandle(handle WindowHandle, params WindowSettings) Window {
	return NewWindow(_C_sfWindow_CreateFromHandle(handle.Cref, params.Cref))
}

// Destroy an existing window
// param Window : Window to destroy
// void sfWindow_Destroy(sfWindow* Window);
func (self *Window) Destroy() {
	_C_sfWindow_Destroy(self.Cref)
}

// Close a window (but doesn't destroy the internal data)
// param Window : Window to close
// void sfWindow_Close(sfWindow* Window);
func (self *Window) Close() {
	_C_sfWindow_Close(self.Cref)
}

// Tell whether or not a window is opened
// param Window : Window object
// sfbool sfWindow_IsOpened(sfWindow* Window);
func (self *Window) IsOpened() bool {
	return SfBool2GoBool(_C_sfWindow_IsOpened(self.Cref))
}

// Get the width of the rendering region of a window
// param Window : Window object
// return Width in pixels
// unsigned int sfWindow_GetWidth(sfWindow* Window);
func (self *Window) GetWidth() uint {
	return uint(_C_sfWindow_GetWidth(self.Cref))
}

// Get the height of the rendering region of a window
// param Window : Window object
// return Height in pixels
// unsigned int sfWindow_GetHeight(sfWindow* Window);
func (self *Window) GetHeight() uint {
	return uint(_C_sfWindow_GetHeight(self.Cref))
}


// Get the creation settings of a window
// param Window : Window object
// return Settings used to create the window
// sfWindowSettings sfWindow_GetSettings(sfWindow* Window);
func (self *Window) GetSettings() WindowSettings {
	return NewWindowSettings(_C_sfWindow_GetSettings(self.Cref))
}

func CreateWindowSettings(depthBits, stencilBits, antialiasingLevel uint) WindowSettings {
	return NewWindowSettings(_C_sfWindowSettings{_C_uint(depthBits), _C_uint(stencilBits), _C_uint(antialiasingLevel)})
}

// Get the event on top of events stack of a window, if any, and pop it
// param Window : Window object
// param Event : Event to fill, if any
// return sfTrue if an event was returned, sfFalse if events stack was empty
// sfbool sfWindow_GetEvent(sfWindow* Window, sfEvent* Event);
func (self *Window) GetEvent(event Event) bool {
	return SfBool2GoBool(_C_sfWindow_GetEvent(self.Cref, event.Cref))
}

// Enable / disable vertical synchronization on a window
// param Window : Window object
// param Enabled : sfTrue to enable v-sync, sfFalse to deactivate
// void sfWindow_UseVerticalSync(sfWindow* Window, sfbool Enabled);
func (self *Window) UseVerticalSync(enabled bool) {
	_C_sfWindow_UseVerticalSync(self.Cref, GoBool2SfBool(enabled))
}

// Show or hide the mouse cursor on a window
// param Window : Window object
// param Show : sfTrue to show, sfFalse to hide
// void sfWindow_ShowMouseCursor(sfWindow* Window, sfbool Show);
func (self *Window) ShowMouseCursor(show bool) {
	_C_sfWindow_ShowMouseCursor(self.Cref, GoBool2SfBool(show))
}

// Change the position of the mouse cursor on a window
// param Window : Window object
// param Left : Left coordinate of the cursor, relative to the window
// param Top : Top coordinate of the cursor, relative to the window
// void sfWindow_SetCursorPosition(sfWindow* Window, unsigned int Left, unsigned int Top);
func (self *Window) SetCursorPosition(left uint, top uint) {
	_C_sfWindow_SetCursorPosition(self.Cref, _C_uint(left), _C_uint(top))
}

// Change the position of a window on screen.
// Only works for top-level windows
// param Window : Window object
// param Left : Left position
// param Top : Top position
// void sfWindow_SetPosition(sfWindow* Window, int Left, int Top);
func (self *Window) SetPosition(left int, top int) {
	_C_sfWindow_SetPosition(self.Cref, _C_int(left), _C_int(top))
}

// Change the size of the rendering region of a window
// param Window : Window object
// param Width : New Width
// param Height : New Height
// void sfWindow_SetSize(sfWindow* Window, unsigned int Width, unsigned int Height);
func (self *Window) SetSize(width uint, height uint) {
	_C_sfWindow_SetSize(self.Cref, _C_uint(width), _C_uint(height))
}

// Show or hide a window
// param Window : Window object
// param State : sfTrue to show, sfFalse to hide
// void sfWindow_Show(sfWindow* Window, sfbool State);
func (self *Window) Show(state bool) {
	_C_sfWindow_Show(self.Cref, GoBool2SfBool(state))
}

// Enable or disable automatic key-repeat for keydown events.
// Automatic key-repeat is enabled by default
// param Window : Window object
// param Enabled : sfTrue to enable, sfFalse to disable
// void sfWindow_EnableKeyRepeat(sfWindow* Window, sfbool Enabled);
func (self *Window) EnableKeyRepeat(enabled bool) {
	_C_sfWindow_EnableKeyRepeat(self.Cref, GoBool2SfBool(enabled))
}

// Change the window's icon
// param Window : Window object
// param Width : Icon's width, in pixels
// param Height : Icon's height, in pixels
// param Pixels : Pointer to the pixels in memory, format must be RGBA 32 bits
// void sfWindow_SetIcon(sfWindow* Window, unsigned int Width, unsigned int Height, sfUint8* Pixels);
func (self *Window) SetIcon(width uint, height uint, pixels *uint8) {
	pixelsPtr := (*_C_sfUint8)(unsafe.Pointer(pixels))
	_C_sfWindow_SetIcon(self.Cref, _C_uint(width), _C_uint(height), pixelsPtr)
}

// Activate or deactivate a window as the current target for rendering
// param Window : Window object
// param Active : sfTrue to activate, sfFalse to deactivate
// return True if operation was successful, false otherwise
// sfbool sfWindow_SetActive(sfWindow* Window, sfbool Active);
func (self *Window) SetActive(active bool) bool {
	return SfBool2GoBool(_C_sfWindow_SetActive(self.Cref, GoBool2SfBool(active)))
}

// Display a window on screen
// param Window : Window object
// void sfWindow_Display(sfWindow* Window);
func (self *Window) Display() {
	_C_sfWindow_Display(self.Cref)
}

// Get the input manager of a window
// param Window : Window object
// return Reference to the input
// sfInput* sfWindow_GetInput(sfWindow* Window);
// ?? Is this really going to work?
func (self *Window) GetInput() Input {
	return NewInput(_C_sfWindow_GetInput(self.Cref))
}

// Limit the framerate to a maximum fixed frequency for a window
// param Window : Window object
// param Limit : Framerate limit, in frames per seconds (use 0 to disable limit)
// void sfWindow_SetFramerateLimit(sfWindow* Window, unsigned int Limit);
func (self *Window) SetFramerateLimit(limit uint) {
	_C_sfWindow_SetFramerateLimit(self.Cref, _C_uint(limit))
}

// Get time elapsed since last frame of a window
// param Window : Window object
// return Time elapsed, in seconds
// float sfWindow_GetFrameTime(sfWindow* Window);
func (self *Window) GetFrameTime() float {
	return float(_C_sfWindow_GetFrameTime(self.Cref))
}

// Change the joystick threshold, ie. the value below which
// no move event will be generated
// param Window : Window object
// param Threshold : New threshold, in range [0, 100]
// void sfWindow_SetJoystickThreshold(sfWindow* Window, float Threshold);
func (self *Window) SetJoystickThreshold(threshold float) {
	_C_sfWindow_SetJoystickThreshold(self.Cref, _C_float(threshold))
}
