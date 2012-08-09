package sfml

//#include <SFML/Window/Export.h>
//#include <SFML/Window/Types.h>
//#include <SFML/Window/Mouse.h>
//#include <SFML/System/Vector2.h>
import "C"

type MouseButton uint

const (
	MouseLeft        = iota // The left mouse button
	MouseRight              // The right mouse button
	MouseMiddle             // The middle (wheel) mouse button
	MouseXButton1           // The first extra mouse button
	MouseXButton2           // The second extra mouse button
	MouseButtonCount        // Keep last -- the total number of mouse buttons
)

// Mouse buttons
// Check if a mouse button is pressed
// \param button Button to check
// \return sfTrue if the button is pressed, sfFalse otherwise
// sfBool sfMouse_isButtonPressed(sfMouseButton button);
func (self MouseButton) IsButtonPressed() bool {
	return C.sfMouse_isButtonPressed(C.sfMouseButton(self)) == 1
}

// Get the current position of the mouse
// This function returns the current position of the mouse
// cursor relative to the given window, or desktop if NULL is passed.
// \param relativeTo Reference window
// \return Position of the mouse cursor, relative to the given window
// sfVector2i sfMouse_getPosition(const sfWindow* relativeTo);
func MousePosition(win Window) (int, int) {
	v := Vector2i{C.sfMouse_getPosition(win.Cref)}
	return v.X(), v.Y()
}

// Set the current position of the mouse
// This function sets the current position of the mouse
// cursor relative to the given window, or desktop if NULL is passed.
// \param position   New position of the mouse
// \param relativeTo Reference window
// void sfMouse_setPosition(sfVector2i position, const sfWindow* relativeTo);
func SetMousePosition(x, y int, win Window) {
	v := NewVector2i(x, y)
	C.sfMouse_setPosition(v.Cref, win.Cref)
}
