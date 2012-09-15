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


// #cgo LDFLAGS:-lcsfml-system
// #include <SFML/System.h>
// #include <SFML/System/Vector2.h>
//
// sfVector2u _NewVector2u(unsigned int x, unsigned int y) {
//   sfVector2u v;
//   v.x = x; v.y=y;
//   return v;
// }
// sfVector2i _NewVector2i(int x, int y) {
//   sfVector2i v;
//   v.x = x; v.y=y;
//   return v;
// }
// sfVector2f _NewVector2f(float x, float y) {
//   sfVector2f v;
//   v.x = x; v.y=y;
//   return v;
// }
import "C"

type (
	Vector2u struct {
		Cref C.sfVector2u
	}
	Vector2i struct {
		Cref C.sfVector2i
	}
	Vector2f struct {
		Cref C.sfVector2f
	}
)

func NewVector2u(x, y uint) Vector2u {
	return Vector2u{C._NewVector2u(C.uint(x), C.uint(y))}
}

func (self Vector2u) X() uint {
	return uint(self.Cref.x)
}

func (self Vector2u) Y() uint {
	return uint(self.Cref.y)
}

func NewVector2i(x, y int) Vector2i {
	return Vector2i{C._NewVector2i(C.int(x), C.int(y))}
}
func (self Vector2i) X() int {
	return int(self.Cref.x)
}
func (self Vector2i) Y() int {
	return int(self.Cref.y)
}

func NewVector2f(x, y float32) Vector2f {
	return Vector2f{C._NewVector2f(C.float(x), C.float(y))}
}

func (self Vector2f) X() float64 {
	return float64(self.Cref.x)
}

func (self Vector2f) Y() float64 {
	return float64(self.Cref.y)
}

// internal
func CNewVector2u(x, y C.uint) Vector2u {
	return Vector2u{C._NewVector2u(x, y)}
}

// internal
func CNewVector2i(x, y C.int) Vector2i {
	return Vector2i{C._NewVector2i(x, y)}
}

// internal
func CNewVector2f(x, y C.float) Vector2f {
	return Vector2f{C._NewVector2f(x, y)}
}
