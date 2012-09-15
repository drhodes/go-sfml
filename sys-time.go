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
import "C"

import "fmt"

type Time struct {
	Cref C.sfTime
}

func (self Time) AsSeconds() float32 {
	return float32(C.sfTime_asSeconds(self.Cref))
}

func (self Time) AsMilliseconds() int32 {
	return int32(C.sfTime_asMilliseconds(self.Cref))
}

func (self Time) AsMicroseconds() int64 {
	return int64(C.sfTime_asMicroseconds(self.Cref))
}

func (self Time) String() string {
	return fmt.Sprintf("%d.%ds", self.AsSeconds(), self.AsMilliseconds())
}

func Seconds(amount float32) Time {
	return Time{C.sfSeconds(C.float(amount))}
}

func Milliseconds(amount int32) Time {
	return Time{C.sfMilliseconds(C.sfInt32(amount))}
}

func Microseconds(amount int64) Time {
	return Time{C.sfMicroseconds(C.sfInt64(amount))}
}

var Zero Time = Seconds(0)
