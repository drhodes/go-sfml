package win

// #include <SFML/Window/Context.h>
// #include <SFML/Window/Input.h>
// #include <SFML/Window/VideoMode.h>
// #include <SFML/Window/Window.h>
import "C"

import(
	"runtime"
	"unsafe"
)

// utility functions must be from in this package to make happy cgo.
func SfBool2GoBool(cVal C.sfBool) bool {
	return int(cVal) == 1
}

func GoBool2SfBool(goVal bool) C.sfBool {
	if goVal {
		return C.sfBool(1)
	}
	return C.sfBool(0)
}

/*
func NewColor(val C.sfColor) Color {
	return Color{ val }
}

type Color struct { 
	Cref C.sfColor 
}
*/
func NewVideoMode(val C.sfVideoMode) VideoMode {
	return VideoMode{ val }
}

type VideoMode struct {
	Cref C.sfVideoMode
} 

//-------------------------------------
func NewContext(val *C.sfContext) Context {
    tmp := Context{ val }
    runtime.SetFinalizer(&tmp, (*Context).Destroy)
    return tmp
}

type Context struct {
	Cref *C.sfContext
}
//-------------------------------------
func NewInput(val *C.sfInput) Input {
    tmp := Input{ val }
    //runtime.SetFinalizer(&tmp, (*Input).Destroy)
    return tmp
}

type Input struct {
	Cref *C.sfInput
}




//-------------------------------------
func NewWindow(val *C.sfWindow) Window {
    tmp := Window{ val }
    runtime.SetFinalizer(&tmp, (*Window).Destroy)
    return tmp
}


type Event struct {
	Cref *C.sfEvent
}


type Window struct {
	Cref *C.sfWindow
}


func NewWindowSettings(val C.sfWindowSettings) WindowSettings {
	return WindowSettings{ val }
}

type WindowSettings struct {
	Cref C.sfWindowSettings
}

type WindowHandle struct {
	Cref C.sfWindowHandle
}

type KeyCode struct {
	Cref C.sfKeyCode
}

type JoyAxis struct {
	Cref C.sfJoyAxis
}

type MouseButton struct {
	Cref C.sfMouseButton
}




// _Context_
// -------------------------------------------------------------------------------
/// Construct a new context
/// return New context
// sfContext* sfContext_Create();
func CreateContext() Context {
    return NewContext( C.sfContext_Create() )
}
/// Destroy an existing context
/// param Context : Context to destroy
// void sfContext_Destroy(sfContext* Context);
func (self *Context) Destroy() {
    C.sfContext_Destroy(self.Cref)
	self.Cref = nil
}

/// Activate or deactivate a context
/// param Context : Context to activate or deactivate
/// param Active : sfTrue to activate, sfFalse to deactivate
// void sfContext_SetActive(sfContext* Context, sfBool Active);
func (self *Context) SetActive(active bool) {
    C.sfContext_SetActive(self.Cref, GoBool2SfBool(active))
}




// _Input_
// -------------------------------------------------------------------------------
// Get the state of a key
// param Input : Input object
// param KeyCode : Key to check
// return sfTrue if key is down, sfFalse if key is up
// sfBool sfInput_IsKeyDown(sfInput* Input, sfKeyCode KeyCode);
func (self *Input) IsKeyDown(keyCode KeyCode) bool {
    return SfBool2GoBool( C.sfInput_IsKeyDown(self.Cref, keyCode.Cref) )
}
// Get the state of a mouse button
// param Input : Input object
// param Button : Button to check
// return sfTrue if button is down, sfFalse if button is up
// sfBool sfInput_IsMouseButtonDown(sfInput* Input, sfMouseButton Button);
func (self *Input) IsMouseButtonDown(button MouseButton) bool {
    return  SfBool2GoBool( C.sfInput_IsMouseButtonDown(self.Cref, button.Cref) )
}
// Get the state of a joystick button
// param Input : Input object
// param JoyId : Identifier of the joystick to check (0 or 1)
// param Button : Button to check
// return sfTrue if button is down, sfFalse if button is up
// sfBool sfInput_IsJoystickButtonDown(sfInput* Input, unsigned int JoyId, unsigned int Button);
func (self *Input) IsJoystickButtonDown(joyId uint, button uint) bool {
    return SfBool2GoBool( C.sfInput_IsJoystickButtonDown(self.Cref, C.uint(joyId), C.uint(button)) )
}
// Get the mouse X position
// param Input : Input object
// return Current mouse left position, relative to owner window
// int sfInput_GetMouseX(sfInput* Input);
func (self *Input) GetMouseX() int {
    return int( C.sfInput_GetMouseX(self.Cref) )
}
// Get the mouse Y position
// param Input : Input object
// return Current mouse top position, relative to owner window
// int sfInput_GetMouseY(sfInput* Input);
func (self *Input) GetMouseY() int {
    return int( C.sfInput_GetMouseY(self.Cref) )
}
// Get the joystick position on a given axis
// param Input : Input object
// param JoyId : Identifier of the joystick to check (0 or 1)
// param Axis : Identifier of the axis to read
// return Current joystick position, in the range [-100, 100]
// float sfInput_GetJoystickAxis(sfInput* Input, unsigned int JoyId, sfJoyAxis Axis);
func (self *Input) GetJoystickAxis(joyId uint, axis JoyAxis) float32 {
    return float32( C.sfInput_GetJoystickAxis(self.Cref, C.uint(joyId), axis.Cref) )
}




// _VideoMode_
// -------------------------------------------------------------------------------
// Get the current desktop video mode
// return Current desktop video mode
// sfVideoMode sfVideoMode_GetDesktopMode();
func GetDesktopMode() VideoMode {
    return NewVideoMode( C.sfVideoMode_GetDesktopMode() )
}

func (self *VideoMode) SetWidth (width uint) {
	self.Cref.Width = C.uint(width)
}

func (self *VideoMode) SetHeight (height uint) {
	self.Cref.Height = C.uint(height)
}

func (self *VideoMode) SetBitsPerPixel (bitsPerPixel uint) {
	self.Cref.BitsPerPixel = C.uint(bitsPerPixel)
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
    return NewVideoMode( C.sfVideoMode_GetMode(C.size_t(index)) )
}

// Get valid video modes count
// return Number of valid video modes available
// size_t sfVideoMode_GetModesCount();
func GetModesCount() int  {
    return int( C.size_t(C.sfVideoMode_GetModesCount()) )
}

// Tell whether or not a video mode is supported
// param Mode : Video mode to check
// return True if video mode is supported, false otherwise
// sfBool sfVideoMode_IsValid(sfVideoMode Mode);
func (self *VideoMode) IsValid() bool {
    return SfBool2GoBool( C.sfVideoMode_IsValid(self.Cref) )
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
    return NewWindow( C.sfWindow_Create(mode.Cref, C.CString(title), C.ulong(style), params.Cref) )
}

// Construct a window from an existing control
// param Handle : Platform-specific handle of the control
// param Params : Creation settings
// sfWindow* sfWindow_CreateFromHandle(sfWindowHandle Handle, sfWindowSettings Params);
func (self *Window) CreateFromHandle(handle WindowHandle, params WindowSettings) Window {
    return NewWindow( C.sfWindow_CreateFromHandle(handle.Cref, params.Cref) )
}

// Destroy an existing window
// param Window : Window to destroy
// void sfWindow_Destroy(sfWindow* Window);
func (self *Window) Destroy() {
    C.sfWindow_Destroy(self.Cref)
}

// Close a window (but doesn't destroy the internal data)
// param Window : Window to close
// void sfWindow_Close(sfWindow* Window);
func (self *Window) Close() {
    C.sfWindow_Close(self.Cref)
}

// Tell whether or not a window is opened
// param Window : Window object
// sfbool sfWindow_IsOpened(sfWindow* Window);
func (self *Window) IsOpened() bool {
    return SfBool2GoBool( C.sfWindow_IsOpened(self.Cref) )
}

// Get the width of the rendering region of a window
// param Window : Window object
// return Width in pixels
// unsigned int sfWindow_GetWidth(sfWindow* Window);
func (self *Window) GetWidth() uint {
	return uint( C.sfWindow_GetWidth(self.Cref) )
}

// Get the height of the rendering region of a window
// param Window : Window object
// return Height in pixels
// unsigned int sfWindow_GetHeight(sfWindow* Window);
func (self *Window) GetHeight() uint {
	return uint( C.sfWindow_GetHeight(self.Cref) )
}


// Get the creation settings of a window
// param Window : Window object
// return Settings used to create the window
// sfWindowSettings sfWindow_GetSettings(sfWindow* Window);
func (self *Window) GetSettings() WindowSettings {
    return NewWindowSettings( C.sfWindow_GetSettings(self.Cref) )
}

func CreateWindowSettings(depthBits, stencilBits, antialiasingLevel uint) WindowSettings {
	return NewWindowSettings( C.sfWindowSettings{ C.uint(depthBits), C.uint(stencilBits), C.uint(antialiasingLevel)} )
}

// Get the event on top of events stack of a window, if any, and pop it
// param Window : Window object
// param Event : Event to fill, if any
// return sfTrue if an event was returned, sfFalse if events stack was empty
// sfbool sfWindow_GetEvent(sfWindow* Window, sfEvent* Event);
func (self *Window) GetEvent(event Event) bool {
    return SfBool2GoBool( C.sfWindow_GetEvent(self.Cref, event.Cref) )
}

// Enable / disable vertical synchronization on a window
// param Window : Window object
// param Enabled : sfTrue to enable v-sync, sfFalse to deactivate
// void sfWindow_UseVerticalSync(sfWindow* Window, sfbool Enabled);
func (self *Window) UseVerticalSync(enabled bool) {
    C.sfWindow_UseVerticalSync(self.Cref, GoBool2SfBool(enabled))
}

// Show or hide the mouse cursor on a window
// param Window : Window object
// param Show : sfTrue to show, sfFalse to hide
// void sfWindow_ShowMouseCursor(sfWindow* Window, sfbool Show);
func (self *Window) ShowMouseCursor(show bool) {
    C.sfWindow_ShowMouseCursor(self.Cref, GoBool2SfBool(show))
}

// Change the position of the mouse cursor on a window
// param Window : Window object
// param Left : Left coordinate of the cursor, relative to the window
// param Top : Top coordinate of the cursor, relative to the window
// void sfWindow_SetCursorPosition(sfWindow* Window, unsigned int Left, unsigned int Top);
func (self *Window) SetCursorPosition(left uint, top uint) {
    C.sfWindow_SetCursorPosition(self.Cref, C.uint(left), C.uint(top))
}

// Change the position of a window on screen.
// Only works for top-level windows
// param Window : Window object
// param Left : Left position
// param Top : Top position
// void sfWindow_SetPosition(sfWindow* Window, int Left, int Top);
func (self *Window) SetPosition(left int, top int) {
    C.sfWindow_SetPosition(self.Cref, C.int(left), C.int(top))
}

// Change the size of the rendering region of a window
// param Window : Window object
// param Width : New Width
// param Height : New Height
// void sfWindow_SetSize(sfWindow* Window, unsigned int Width, unsigned int Height);
func (self *Window) SetSize(width uint, height uint) {
    C.sfWindow_SetSize(self.Cref, C.uint(width), C.uint(height))
}

// Show or hide a window
// param Window : Window object
// param State : sfTrue to show, sfFalse to hide
// void sfWindow_Show(sfWindow* Window, sfbool State);
func (self *Window) Show(state bool) {
    C.sfWindow_Show(self.Cref, GoBool2SfBool(state))
}

// Enable or disable automatic key-repeat for keydown events.
// Automatic key-repeat is enabled by default
// param Window : Window object
// param Enabled : sfTrue to enable, sfFalse to disable
// void sfWindow_EnableKeyRepeat(sfWindow* Window, sfbool Enabled);
func (self *Window) EnableKeyRepeat(enabled bool) {
    C.sfWindow_EnableKeyRepeat(self.Cref, GoBool2SfBool(enabled))
}

// Change the window's icon
// param Window : Window object
// param Width : Icon's width, in pixels
// param Height : Icon's height, in pixels
// param Pixels : Pointer to the pixels in memory, format must be RGBA 32 bits
// void sfWindow_SetIcon(sfWindow* Window, unsigned int Width, unsigned int Height, sfUint8* Pixels);
func (self *Window) SetIcon(width uint, height uint, pixels *uint8) {
	pixelsPtr := (*C.sfUint8)(unsafe.Pointer(pixels))
    C.sfWindow_SetIcon(self.Cref, C.uint(width), C.uint(height), pixelsPtr)
}

// Activate or deactivate a window as the current target for rendering
// param Window : Window object
// param Active : sfTrue to activate, sfFalse to deactivate
// return True if operation was successful, false otherwise
// sfbool sfWindow_SetActive(sfWindow* Window, sfbool Active);
func (self *Window) SetActive(active bool) bool {
    return SfBool2GoBool( C.sfWindow_SetActive(self.Cref, GoBool2SfBool(active)))
}

// Display a window on screen
// param Window : Window object
// void sfWindow_Display(sfWindow* Window);
func (self *Window) Display() {
    C.sfWindow_Display(self.Cref)
}

// Get the input manager of a window
// param Window : Window object
// return Reference to the input
// sfInput* sfWindow_GetInput(sfWindow* Window);
// ?? Is this really going to work?
func (self *Window) GetInput() Input {
    return NewInput( C.sfWindow_GetInput(self.Cref) )
}

// Limit the framerate to a maximum fixed frequency for a window
// param Window : Window object
// param Limit : Framerate limit, in frames per seconds (use 0 to disable limit)
// void sfWindow_SetFramerateLimit(sfWindow* Window, unsigned int Limit);
func (self *Window) SetFramerateLimit(limit uint) {
    C.sfWindow_SetFramerateLimit(self.Cref, C.uint(limit))
}

// Get time elapsed since last frame of a window
// param Window : Window object
// return Time elapsed, in seconds
// float sfWindow_GetFrameTime(sfWindow* Window);
func (self *Window) GetFrameTime() float32 {
    return float32(C.sfWindow_GetFrameTime(self.Cref))
}

// Change the joystick threshold, ie. the value below which
// no move event will be generated
// param Window : Window object
// param Threshold : New threshold, in range [0, 100]
// void sfWindow_SetJoystickThreshold(sfWindow* Window, float Threshold);
func (self *Window) SetJoystickThreshold(threshold float32) {
    C.sfWindow_SetJoystickThreshold(self.Cref, C.float(threshold))
}
