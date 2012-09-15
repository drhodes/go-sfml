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
// #include <SFML/System/Vector3.h>
//
// sfVector3f _NewVector3f(float x, float y, float z) {
//   sfVector3f v;
//   v.x = x; v.y=y; v.z=z;
//   return v;
// }
import "C"

type Vector3f struct {
	Cref C.sfVector3f
}

func NewVector3f(x, y, z float32) Vector3f {
	return Vector3f{C._NewVector3f(C.float(x), C.float(y), C.float(z))}
}

func (self Vector3f) X() float32 {
	return float32(self.Cref.x)
}

func (self Vector3f) Y() float32 {
	return float32(self.Cref.y)
}

func (self Vector3f) Z() float32 {
	return float32(self.Cref.z)
}

// internal
func CNewVector3f(x, y, z C.float) Vector3f {
	return Vector3f{C._NewVector3f(x, y, z)}
}
