package sfml

//#include <SFML/Window/Export.h>
//#include <SFML/Window/Types.h>
//#include <SFML/Window/Mouse.h>
//#include <SFML/System/Vector2.h>
import "C"

type MouseButton uint

// Mouse buttons
const (
	MouseLeft        = iota // The left mouse button
	MouseRight              // The right mouse button
	MouseMiddle             // The middle (wheel) mouse button
	MouseXButton1           // The first extra mouse button
	MouseXButton2           // The second extra mouse button
	MouseButtonCount        // Keep last -- the total number of mouse buttons
)

// Check if a mouse button is pressed
// Returns true if the button is pressed, false otherwise
func (self MouseButton) IsButtonPressed() bool {
	// sfBool sfMouse_isButtonPressed(sfMouseButton button)
	return C.sfMouse_isButtonPressed(C.sfMouseButton(self)) == 1
}


// Get the current position of the mouse
// This function returns the current position of the mouse
// cursor relative to the desktop 
func MouseGetPositionAbsolute() (int, int) {
	v := C.sfMouse_getPosition(nil)
	return int(v.x), int(v.y)
}

// Get the current position of the mouse
// This function returns the current position of the mouse
// cursor relative to the given window.
func MouseGetPosition(win Window) (int, int) {
	v := C.sfMouse_getPosition(win.Cref)
	return int(v.x), int(v.y)
}

// Set the current position of the mouse.
// This function sets the current position of the mouse
// cursor relative to the desktop.
func SetMousePositionAbsolute(x, y int) {
	// void sfMouse_setPosition(sfVector2i position, const sfWindow* relativeTo);
	v := C.sfVector2i{C.int(x), C.int(y)}
	C.sfMouse_setPosition(v, nil)
}

// Set the current position of the mouse
// This function sets the current position of the mouse
// cursor relative to the given window.
func SetMousePosition(x, y int, win Window) {
	// void sfMouse_setPosition(sfVector2i position, const sfWindow* relativeTo);
	v := C.sfVector2i{C.int(x), C.int(y)}
	C.sfMouse_setPosition(v, win.Cref)
} 
