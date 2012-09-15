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


// #cgo LDFLAGS:-lcsfml-graphics -ljpeg
// #include <SFML/Graphics/Color.h>
// #include <SFML/Window/Context.h>
// #include <SFML/Window/Export.h>
// #include <SFML/Window/Types.h>
// #include <SFML/Graphics/Export.h>
import "C"

import "fmt"

// Utility class for manpulating RGBA colors
type Color struct {
	Cref C.sfColor
}

// Construct a color from its 3 RGB components
//
// \param red   Red component   (0 .. 255)
// \param green Green component (0 .. 255)
// \param blue  Blue component  (0 .. 255)
//
// \return sfColor constructed from the components
// sfColor sfColor_fromRGB(sfUint8 red, sfUint8 green, sfUint8 blue);
func FromRGB(red uint8, green uint8, blue uint8) Color {
	return Color{C.sfColor_fromRGB(
		C.sfUint8(red),
		C.sfUint8(green),
		C.sfUint8(blue),
	)}
}

// Construct a color from its 4 RGBA components
//
// \param red   Red component   (0 .. 255)
// \param green Green component (0 .. 255)
// \param blue  Blue component  (0 .. 255)
// \param alpha Alpha component (0 .. 255)
//
// \return sfColor constructed from the components
//
// sfColor sfColor_fromRGBA(sfUint8 red, sfUint8 green, sfUint8 blue, sfUint8 alpha);
func FromRGBA(red uint8, green uint8, blue uint8, alpha uint8) Color {
	return Color{C.sfColor_fromRGBA(
		C.sfUint8(red),
		C.sfUint8(green),
		C.sfUint8(blue),
		C.sfUint8(alpha),
	)}
}

// Add two colors
//
// \param color2 Second color
//
// \return Component-wise saturated addition of the two colors
//
// sfColor sfColor_add(sfColor color1, sfColor color2);
func (self Color) Add(color2 Color) Color {
	return Color{C.sfColor_add(self.Cref, color2.Cref)}
}

// Modulate two colors
//
// \param color1 First color
// \param color2 Second color
//
// \return Component-wise multiplication of the two colors
//
// sfColor sfColor_modulate(sfColor color1, sfColor color2);
func (self Color) Modulate(color2 Color) Color {
	return Color{C.sfColor_modulate(self.Cref, color2.Cref)}
}

// Return the red component of a color
func (self Color) Red() uint8 {
	return uint8(self.Cref.r)
}

// Return the green component of a color
func (self Color) Green() uint8 {
	return uint8(self.Cref.g)
}

// Return the blue component of a color
func (self Color) Blue() uint8 {
	return uint8(self.Cref.b)
}

// Return the alpha component of a color
func (self Color) Alpha() uint8 {
	return uint8(self.Cref.a)
}

// Return the components of a color
func (self Color) Components() (uint8, uint8, uint8, uint8) {
	return self.Red(), self.Green(), self.Blue(), self.Alpha()
}

// Return a string describing the color
func (self Color) String() string {
	return fmt.Sprintf("Color{r=%d, g=%d, b=%d, a=%d}", self.Red(), self.Green(), self.Blue(), self.Alpha())
}