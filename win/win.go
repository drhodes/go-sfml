package win

// #include <SFML/Window/Context.h>
import "C"

import(
	"runtime"
	//"unsafe"
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

//-------------------------------------
func NewContext(val *C.sfContext) Context {
    tmp := Context{ val }
    runtime.SetFinalizer(&tmp, (*Context).Destroy)
    return tmp
}

type Context struct {
	Cref *C.sfContext
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
    return SfBool2GoBool( C.sfInput_IsJoystickButtonDown(self.Cref, joyId, button) )
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
func (self *Input) GetJoystickAxis(joyId uint, axis JoyAxis) float {
    return float( C.sfInput_GetJoystickAxis(self.Cref, joyId, axis.Cref) )
}

