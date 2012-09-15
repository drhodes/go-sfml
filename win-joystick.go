// Copyright 2012.  All rights reserved. 
// Use of this source code is governed by a BSD-style  
// license that can be found in the LICENSE file.

////////////////////////////////////////////////////////////
//
// SFML - Simple and Fast Multimedia Library
// Copyright (C) 2007-2012 Laurent Gomila (laurent.gom@gmail.com)
//
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it freely,
// subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented;
//    you must not claim that you wrote the original software.
//    If you use this software in a product, an acknowledgment
//    in the product documentation would be appreciated but is not required.
//
// 2. Altered source versions must be plainly marked as such,
//    and must not be misrepresented as being the original software.
//
// 3. This notice may not be removed or altered from any source distribution.
//
////////////////////////////////////////////////////////////
package sfml


//#include <SFML/Window/Export.h>
//#include <SFML/Window/Joystick.h>
import "C"

import "log"

type Joystick uint

type JoystickAxis struct {
	Cref *C.sfJoystickAxis
}

// Check if a joystick is connected
// \param joystick Index of the joystick to check
// \return sfTrue if the joystick is connected, sfFalse otherwise
// sfBool sfJoystick_isConnected(unsigned int joystick);
func (self Joystick) Isconnected() bool {
	return C.sfJoystick_isConnected(C.uint(self)) == 1
}

// Return the number of buttons supported by a joystick
// If the joystick is not connected, this function returns 0.
// \param joystick Index of the joystick
// \return Number of buttons supported by the joystick
// unsigned int sfJoystick_getButtonCount(unsigned int joystick);
func (self Joystick) ButtonCount() uint {
	return uint(C.uint(self))
}

// Check if a joystick supports a given axis
// If the joystick is not connected, this function returns false.
// \param joystick Index of the joystick
// \param axis     Axis to check
// \return sfTrue if the joystick supports the axis, sfFalse otherwise
// sfBool sfJoystick_hasAxis(unsigned int joystick, sfJoystickAxis axis);
func (self Joystick) HasAxis(axis JoystickAxis) bool {
	if axis.Cref == nil {
		log.Panic("nil Joystick encountered in Joystick.HasAxis")
	}
	return C.sfJoystick_hasAxis(C.uint(self), *axis.Cref) == 1
}

// Check if a joystick button is pressed
// If the joystick is not connected, this function returns false.
// \param joystick Index of the joystick
// \param button   Button to check
// \return sfTrue if the button is pressed, sfFalse otherwise
// sfBool sfJoystick_isButtonPressed(unsigned int joystick, unsigned int button);
func (self Joystick) IsButtonPressed(button uint) bool {
	return C.sfJoystick_isButtonPressed(C.uint(self), C.uint(button)) == 1
}

// Get the current position of a joystick axis
// If the joystick is not connected, this function returns 0.
// \param joystick Index of the joystick
// \param axis     Axis to check
// \return Current position of the axis, in range [-100 .. 100]
// float sfJoystick_getAxisPosition(unsigned int joystick, sfJoystickAxis axis);
func (self Joystick) AxisPosition(axis JoystickAxis) float32 {
	return float32(
		C.sfJoystick_getAxisPosition(C.uint(self), *axis.Cref))
}

// Update the states of all joysticks
// This function is used internally by SFML, so you normally
// don't have to call it explicitely. However, you may need to
// call it if you have no window yet (or no window at all):
// in this case the joysticks states are not updated automatically.
// void sfJoystick_update(void);
func (self Joystick) Update() {
	C.sfJoystick_update()
}
