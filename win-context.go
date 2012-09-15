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
// #include <SFML/Window/Context.h>
// #include <SFML/Window/Export.h>
// #include <SFML/Window/Types.h>
import "C"

import "errors"

type Context struct {
	Cref *C.sfContext
}

// Create a new context
// This function activates the new context.
// \return New sfContext object
// sfContext* sfContext_create(void);
func NewContext() (Context, error) {
	ctx := C.sfContext_create()
	if ctx == nil {
		return Context{nil}, errors.New("Couldn't make a new context")
	}
	return Context{ctx}, nil
}

//  Destroy a context
// \param context Context to destroy
// void sfContext_destroy(sfContext* context);
func (self Context) Destroy() {
	C.sfContext_destroy(self.Cref)
}

//  Activate or deactivate explicitely a context
//
// \param context Context object
// \param active  sfTrue to activate, sfFalse to deactivate
// void sfContext_setActive(sfContext* context, sfBool active);
func (self Context) SetActive(active bool) {
	if active {
		C.sfContext_setActive(self.Cref, C.sfBool(1))
	} else {
		C.sfContext_setActive(self.Cref, C.sfBool(0))
	}
}
