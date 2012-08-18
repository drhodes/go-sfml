package sfml

// #cgo LDFLAGS:-lcsfml-window
// #include <SFML/Window/Export.h>
// #include <SFML/Window/Event.h>
// #include <SFML/Window/VideoMode.h>
// #include <SFML/Window/Window.h>
// #include <SFML/Window/WindowHandle.h>
// #include <SFML/Window/Types.h>
// #include <SFML/System/Vector2.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"unsafe"
)

// Enumeration of window creation styles
const (
	// No border / title bar (this flag and all others are mutually exclusive)
	StyleNone = 0
	// Title bar + fixed border
	StyleTitlebar = 1 << 0
	// Titlebar + resizable border + maximize button
	StyleResize = 1 << 1
	// Titlebar + close button
	StyleClose = 1 << 2
	// Fullscreen mode (this flag and all others are mutually exclusive)
	StyleFullscreen = 1 << 3
	// Default window style
	StyleDefaultStyle = StyleTitlebar | StyleResize | StyleClose
)

type CWindowHandle C.sfWindowHandle
type WindowHandle struct {
	Cref C.sfWindowHandle
}

// need to export this C type.
type CWindow *C.sfWindow

type Window struct {
	Cref CWindow
}

// Construct a new window
// This function creates the window with the size and pixel
// depth defined in \a mode. An optional style can be passed to
// customize the look and behaviour of the window (borders,
// title bar, resizable, closable, ...). If \a style contains
// sfFullscreen, then \a mode must be a valid video mode.
// The fourth parameter is a pointer to a structure specifying
// advanced OpenGL context settings such as antialiasing,
// depth-buffer bits, etc.
// param mode     Video mode to use (defines the width, height and depth 
// of the rendering area of the window)
// param title    Title of the window
// param style    Window style
// param settings Additional settings for the underlying OpenGL context
// return A new Window.
// sfWindow* sfWindow_create(sfVideoMode mode, const char* title,
// sfUint32 style, const sfContextSettings* settings);
func NewWindow(mode VideoMode, title string, style int32, settings ContextSettings) (Window, error) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	if mode.Nil() {
		return Window{nil}, errors.New("NewWindow can't take mode with nil Cref")
	}

	win := C.sfWindow_create(*mode.Cref, t, C.sfUint32(style), settings.Cref)
	if win == nil {
		return Window{nil}, errors.New("Couldn't make a new window")
	}
	return Window{win}, nil
}

// Construct a window from an existing control
// Use this constructor if you want to create an OpenGL
// rendering area into an already existing control.
// The second parameter is a pointer to a structure specifying
// advanced OpenGL context settings such as antialiasing,
// depth-buffer bits, etc.
// \param handle   Platform-specific handle of the control
// \param settings Additional settings for the underlying OpenGL context
// \return A new sfWindow object
// sfWindow* sfWindow_createFromHandle(sfWindowHandle handle, const sfContextSettings* settings);
// func (self Window) Createfromhandle(settings ContextSettings) *Window {
// 	return C.sfWindow_createFromHandle();
// }

// Destroy a window
// void sfWindow_destroy(sfWindow* window);
func (self Window) Destroy() {
	C.sfWindow_destroy(self.Cref)
}

// Close a window and destroy all the attached resources
// After calling this function, the sfWindow object remains
// valid, you must call sfWindow_destroy to actually delete it.
// All other functions such as sfWindow_pollEvent or sfWindow_display
// will still work (i.e. you don't have to test sfWindow_isOpen
// every time), and will have no effect on closed windows.
// \param window Window object
// void sfWindow_close(sfWindow* window);
func (self Window) Close() {
	C.sfWindow_close(self.Cref)
}

// Tell whether or not a window is opened
// This function returns whether or not the window exists.
// Note that a hidden window (sfWindow_SetVisible(sfFalse)) will return
// sfTrue.
// \param window Window object
// \return sfTrue if the window is opened, sfFalse if it has been closed
// sfBool sfWindow_isOpen(const sfWindow* window);
func (self Window) IsOpen() bool {
	b := C.sfWindow_isOpen(self.Cref)
	return b == 1
}

// Get the settings of the OpenGL context of a window
// Note that these settings may be different from what was
// passed to the sfWindow_create function,
// if one or more settings were not supported. In this case,
// SFML chose the closest match.
// \param window Window object
// \return Structure containing the OpenGL context settings
// sfContextSettings sfWindow_getSettings(const sfWindow* window);
func (self Window) Settings() ContextSettings {
	sets := C.sfWindow_getSettings(self.Cref)
	return ContextSettings{&sets}
}

// Pop the event on top of events stack, if any, and return it
// This function is not blocking: if there's no pending event then
// it will return false and leave \a event unmodified.
// Note that more than one event may be present in the events stack,
// thus you should always call this function in a loop
// to make sure that you process every pending event
// in: window Window object
// in: event  Event to be returned
// returns true if an event was returned, or false if the events stack was empty
// sfBool sfWindow_pollEvent(sfWindow* window, sfEvent* event);
func (self Window) PollEvent() (interface{}, bool) {
	// ok if got event.	
	e := NewEvent()
	ok := C.sfWindow_pollEvent(self.Cref, e.Cref) == 1
	if ok {
		// look at the first byte, it's the event type
		et := EventType((*e.Cref)[0])
		switch et {
		case EvtClosed:
		case EvtResized:
		case EvtLostFocus:
		case EvtGainedFocus:
		case EvtTextEntered:
			return e.ToTextEvent(), true
		case EvtKeyPressed, EvtKeyReleased:
			return e.ToKeyEvent(), true
		case EvtMouseWheelMoved:
			return e.ToMouseWheelEvent(), true
		case EvtMouseButtonPressed, EvtMouseButtonReleased:
			return e.ToMouseButtonEvent(), true
		case EvtMouseMoved, EvtMouseEntered, EvtMouseLeft:
			return e.ToMouseMoveEvent(), true
		case EvtJoystickButtonPressed, EvtJoystickButtonReleased, EvtJoystickMoved:
			return e.ToJoystickMoveEvent(), true
		case EvtJoystickConnected:
		case EvtJoystickDisconnected:
			return e.ToJoystickConnectEvent(), true
		case EvtNone:
			return NullEvent(0), false
		}
	}
	return NullEvent(0), false
}

// Wait for an event and return it
// This function is blocking: if there's no pending event then
// it will wait until an event is received.
// After this function returns (and no error occured),
// the \a event object is always valid and filled properly.
// This function is typically used when you have a thread that
// is dedicated to events handling: you want to make this thread
// sleep as long as no new event is received.
// \param window Window object
// \param event  Event to be returned
// \return sfFalse if any error occured
// sfBool sfWindow_waitEvent(sfWindow* window, sfEvent* event);
// func (self Window) Waitevent(event *Event) Bool {
// 	return C.sfWindow_waitEvent();
// }

// Get the position of a window
// \param window Window object
// \return (x, y)
// sfVector2i sfWindow_getPosition(const sfWindow* window);
// TODO evalutate which is better [ multiple return | vector type ]
func (self Window) Position() (x, y int) {
	pos := C.sfWindow_getPosition(self.Cref)
	return int(pos.x), int(pos.y)
}

// Change the position of a window on screen
// This function only works for top-level windows
// (i.e. it will be ignored for windows created from
// the handle of a child window/control).
// \param window   Window object
// \param position New position of the window, in pixels
// void sfWindow_setPosition(sfWindow* window, sfVector2i position);
// TODO evalutate which is better [ multiple return | vector type ]
func (self Window) SetPosition(x, y int) {
	v := C.sfVector2i{C.int(x), C.int(y)}
	C.sfWindow_setPosition(self.Cref, v)
}

// Get the size of the rendering region of a window
// The size doesn't include the titlebar and borders
// of the window.
// \param window Window object
// \return Size in pixels
// sfVector2u sfWindow_getSize(const sfWindow* window);
func (self Window) GetSize() (x, y uint) {
	pos := C.sfWindow_getSize(self.Cref)
	return uint(pos.x), uint(pos.y)
}

// TODO WHY THIS NO WORK.
// Change the size of the rendering region of a window
// \param window Window object
// \param size   New size, in pixels
// void sfWindow_setSize(sfWindow* window, sfVector2u size);
func (self Window) SetSize(x, y uint) {
	size := NewVector2u(x, y)
	C.sfWindow_setSize(self.Cref, size.Cref)
}

// Change the title of a window
// \param window Window object
// \param title  New title
// void sfWindow_setTitle(sfWindow* window, const char* title);
func (self Window) SetTitle(title string) {
	C.sfWindow_setTitle(self.Cref, C.CString(title))
}

// Change a window's icon
// \a pixels must be an array of \a width x \a height pixels
// in 32-bits RGBA format.
// \param window Window object
// \param width  Icon's width, in pixels
// \param height Icon's height, in pixels
// \param pixels Pointer to the array of pixels in memory
// void sfWindow_setIcon(sfWindow* window, unsigned int width, unsigned int height, const sfUint8* pixels);
func (self Window) SetIcon(width, height int, pixels []uint8) {
	ptr := unsafe.Pointer(&pixels[0])
	p := (*C.sfUint8)(ptr)
	C.sfWindow_setIcon(self.Cref, C.uint(width), C.uint(height), p)
}

// Show or hide a window
// \param window  Window object
// \param visible sfTrue to show the window, sfFalse to hide it
// void sfWindow_setVisible(sfWindow* window, sfBool visible);
func (self Window) SetVisible(visible bool) {
	if visible {
		C.sfWindow_setVisible(self.Cref, C.sfBool(1))
	} else {
		C.sfWindow_setVisible(self.Cref, C.sfBool(0))
	}
}

// Show or hide the mouse cursor
// \param window  Window object
// \param visible sfTrue to show, sfFalse to hide
// void sfWindow_setMouseCursorVisible(sfWindow* window, sfBool visible);
func (self Window) SetMouseCursorVisible(visible bool) {
	if visible {
		C.sfWindow_setMouseCursorVisible(self.Cref, C.sfBool(1))
	} else {
		C.sfWindow_setMouseCursorVisible(self.Cref, C.sfBool(0))
	}
}

// Enable or disable vertical synchronization
// Activating vertical synchronization will limit the number
// of frames displayed to the refresh rate of the monitor.
// This can avoid some visual artifacts, and limit the framerate
// to a good value (but not constant across different computers).
// \param window  Window object
// \param enabled sfTrue to enable v-sync, sfFalse to deactivate
// void sfWindow_setVerticalSyncEnabled(sfWindow* window, sfBool enabled);
func (self Window) SetVerticalSyncEnabled(enabled bool) {
	if enabled {
		C.sfWindow_setVerticalSyncEnabled(self.Cref, C.sfBool(1))
	} else {
		C.sfWindow_setVerticalSyncEnabled(self.Cref, C.sfBool(0))
	}
}

// Enable or disable automatic key-repeat
// If key repeat is enabled, you will receive repeated
// KeyPress events while keeping a key pressed. If it is disabled,
// you will only get a single event when the key is pressed.
// Key repeat is enabled by default.
// \param window  Window object
// \param enabled sfTrue to enable, sfFalse to disable
// void sfWindow_setKeyRepeatEnabled(sfWindow* window, sfBool enabled);
func (self Window) SetKeyRepeatEnabled(enabled bool) {
	if enabled {
		C.sfWindow_setKeyRepeatEnabled(self.Cref, C.sfBool(1))
	} else {
		C.sfWindow_setKeyRepeatEnabled(self.Cref, C.sfBool(0))
	}
}

// Activate or deactivate a window as the current target
//        for OpenGL rendering
// A window is active only on the current thread, if you want to
// make it active on another thread you have to deactivate it
// on the previous thread first if it was active.
// Only one window can be active on a thread at a time, thus
// the window previously active (if any) automatically gets deactivated.
// \param window Window object
// \param active sfTrue to activate, sfFalse to deactivate
// \return sfTrue if operation was successful, sfFalse otherwise
// sfBool sfWindow_setActive(sfWindow* window, sfBool active);
func (self Window) SetActive(active bool) bool {
	if active {
		return C.sfWindow_setActive(self.Cref, C.sfBool(1)) == 1
	}
	return C.sfWindow_setActive(self.Cref, C.sfBool(0)) == 1
}

// Display on screen what has been rendered to the
//        window so far
// This function is typically called after all OpenGL rendering
// has been done for the current frame, in order to show
// it on screen.
// \param window Window object
// void sfWindow_display(sfWindow* window);
func (self Window) Display() {
	C.sfWindow_display(self.Cref)
}

// Limit the framerate to a maximum fixed frequency
// If a limit is set, the window will use a small delay after
// each call to sfWindow_Display to ensure that the current frame
// lasted long enough to match the framerate limit.
// \param window Window object
// \param limit  Framerate limit, in frames per seconds (use 0 to disable limit)
// void sfWindow_setFramerateLimit(sfWindow* window, unsigned int limit);
func (self Window) SetFramerateLimit(limit uint) {
	C.sfWindow_setFramerateLimit(self.Cref, C.uint(limit))
}

// Change the joystick threshold
// The joystick threshold is the value below which
// no JoyMoved event will be generated.
// \param window    Window object
// \param threshold New threshold, in the range [0, 100]
// void sfWindow_setJoystickThreshold(sfWindow* window, float threshold);
func (self Window) SetJoystickThreshold(threshold float32) {
	C.sfWindow_setJoystickThreshold(self.Cref, C.float(threshold))
}

// Get the OS-specific handle of the window
// The type of the returned handle is sfWindowHandle,
// which is a typedef to the handle type defined by the OS.
// You shouldn't need to use this function, unless you have
// very specific stuff to implement that SFML doesn't support,
// or implement a temporary workaround until a bug is fixed.
// \param window Window object
// \return System handle of the window
// sfWindowHandle sfWindow_getSystemHandle(const sfWindow* window);
// func (self Window) GetSystemHandle() WindowHandle {
// 	return C.sfWindow_getSystemHandle();
// }
