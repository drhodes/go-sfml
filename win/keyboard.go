package win

// #cgo LDFLAGS:-lcsfml-window
// #include <SFML/Window/Keyboard.h>
// #include <SFML/Window/Export.h>
import "C"
import (
	//"unsafe"
	//"fmt"
)

// Check if a key is pressed
// key Key to check
// return true if the key is pressed, false otherwise
// sfBool sfKeyboard_isKeyPressed(sfKeyCode key);
// func (self Keyboard) IsKeyPressed(key KeyCode) bool { 
// 	return false
//     //return C.sfKeyboard_isKeyPressed();
// }
            
