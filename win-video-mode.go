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


// #cgo LDFLAGS:-lcsfml-window
// #include <SFML/Window/VideoMode.h>
// #include <string.h>
// 
//
import "C"
import (
	"fmt"
	"unsafe"
)

// sfVideoMode defines a video mode (width, height, bpp, frequency)
// and provides functions for getting modes supported by the display device
type VideoMode struct {
	Cref *C.sfVideoMode
}

func NewVideoMode(w, h, bpp uint) VideoMode {
	v := C.sfVideoMode{C.uint(w), C.uint(h), C.uint(bpp)}
	return VideoMode{&v}
}

func (self VideoMode) Nil() bool {
	return self.Cref == nil
}

func (self VideoMode) Width() uint {
	return uint(self.Cref.width)
}

func (self VideoMode) Height() uint {
	return uint(self.Cref.height)
}

func (self VideoMode) BitsPerPixel() uint {
	return uint(self.Cref.bitsPerPixel)
}

func (self VideoMode) String() string {
	s := `
VideoMode:
   Width:        %d
   Height:       %d
   BitsPerPixel: %d
`
	return fmt.Sprintf(s, self.Width(), self.Height(), self.BitsPerPixel())
}

// Get the current desktop video mode
// sfVideoMode sfVideoMode_getDesktopMode(void)
func DesktopMode() VideoMode {
	temp := C.sfVideoMode_getDesktopMode()
	return VideoMode{&temp}
}

// Retrieve all the video modes supported in fullscreen mode
//
// When creating a fullscreen window, the video mode is restricted
// to be compatible with what the graphics driver and monitor
// support. This function returns the complete list of all video
// modes that can be used in fullscreen mode.
// The returned slice is sorted from best to worst, so that
// the first element will always give the best mode (higher
// width, height and bits-per-pixel).
// const sfVideoMode* sfVideoMode_getFullscreenModes(size_t* Count)
func FullscreenModes() []VideoMode {
	// get the size_of a sfVideoMode
	size := unsafe.Sizeof(*VideoMode{}.Cref)

	var nmodes C.size_t
	cmodes := C.sfVideoMode_getFullscreenModes(&nmodes)

	modes := []VideoMode{}
	p := unsafe.Pointer(cmodes)
	ptr := uintptr(p)

	for nmodes > 0 {
		m := VideoMode{(*C.sfVideoMode)(p)}
		if m.IsValid() {
			modes = append(modes, m)
		}
		ptr += size
		p = unsafe.Pointer(ptr)
		nmodes--
	}
	return modes
}

// Tell whether or not a video mode is valid
// The validity of video modes is only relevant when using
// fullscreen windows otherwise any video mode can be used
// with no restriction.
// *return true if the video mode is valid for fullscreen mode
// sfBool sfVideoMode_isValid(sfVideoMode mode)
func (self VideoMode) IsValid() bool {
	return int(C.sfVideoMode_isValid(*self.Cref)) == 1
}
