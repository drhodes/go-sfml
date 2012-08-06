package sfml

//#include <SFML/Window/Export.h>
//#include <SFML/Window/Types.h>
//#include <SFML/Window/Mouse.h>
//#include <SFML/System/Vector2.h>
import "C"

type MouseButton uint

const (
	sfMouseLeft        = iota // The left mouse button
	sfMouseRight              // The right mouse button
	sfMouseMiddle             // The middle (wheel) mouse button
	sfMouseXButton1           // The first extra mouse button
	sfMouseXButton2           // The second extra mouse button
	sfMouseButtonCount        // Keep last -- the total number of mouse buttons 
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
func (self MouseButton) GetPosition(win Window) Vector2i {
	cv2 := C.sfMouse_getPosition(win.Cref)
	x := int(cv2.x)
	y := int(cv2.y)
	return NewVector2i(x, y)
}

// Set the current position of the mouse
// This function sets the current position of the mouse
// cursor relative to the given window, or desktop if NULL is passed.
// \param position   New position of the mouse
// \param relativeTo Reference window
// void sfMouse_setPosition(sfVector2i position, const sfWindow* relativeTo);
func (self MouseButton) SetPosition(pos Vector2i, win Window) {
	// TODO: this is really inefficient. figure out how to fix it
	// later these constant factors are going to add up in 
	// The issue is that 
	x := pos.X()
	y := pos.Y()
	C.sfMouse_setPosition(C.sfVector2i{C.int(x), C.int(y)}, win.Cref)
}
