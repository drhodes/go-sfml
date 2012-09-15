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


/*
 #cgo LDFLAGS:-lcsfml-window
 #include <SFML/Window/Window.h>
*/
import "C"

//
// Structure defining the window's creation settings
type ContextSettings struct {
	Cref *C.sfContextSettings
	// DepthBits         uint //< Bits of the depth buffer
	// StencilBits       uint //< Bits of the stencil buffer
	// AntialiasingLevel uint //< Level of antialiasing
	// MajorVersion      uint //< Major number of the context version to create
	// MinorVersion      uint //< Minor number of the context version to create
}

func NewContextSettings(depthbits, stencilbits, antialiasinglevel,
	majVersion, minVersion uint) ContextSettings {

	ref := C.sfContextSettings{
		C.uint(depthbits),
		C.uint(stencilbits),
		C.uint(antialiasinglevel),
		C.uint(majVersion),
		C.uint(minVersion),
	}

	return ContextSettings{&ref}
	// todo: How is this memory collected?
}

func (self ContextSettings) Nil() bool {
	return self.Cref == nil
}
func (self ContextSettings) DepthBits() uint {
	return uint(self.Cref.depthBits)
}
func (self ContextSettings) StencilBits() uint {
	return uint(self.Cref.stencilBits)
}
func (self ContextSettings) AntialiasingLevel() uint {
	return uint(self.Cref.antialiasingLevel)
}
func (self ContextSettings) MajorVersion() uint {
	return uint(self.Cref.majorVersion)
}
func (self ContextSettings) MinorVersion() uint {
	return uint(self.Cref.minorVersion)
}
