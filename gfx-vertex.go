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

// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/PrimitiveType.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/Graphics/Vertex.h>
// #include <SFML/Graphics/VertexArray.h>
import "C"

// typedef struct
// {
//     sfVector2f position;  ///< Position of the vertex
//     sfColor    color;     ///< Color of the vertex
//     sfVector2f texCoords; ///< Coordinates of the texture's pixel to map to the vertex
// } sfVertex;

type Vertex struct {
	Cref *C.sfVertex
}

func (self Vertex) GetPosition() (x, y float32) {
	return float32(self.Cref.position.x), float32(self.Cref.position.y)
}

func (self Vertex) SetPosition(x, y float32) {
	self.Cref.position.x = C.float(x)
	self.Cref.position.y = C.float(y)
}

func (self Vertex) GetTexCoords() (x, y float32) {
	return float32(self.Cref.texCoords.x), float32(self.Cref.texCoords.y)
}

func (self Vertex) SetTexCoords(x, y float32) {
	self.Cref.texCoords.x = C.float(x)
	self.Cref.texCoords.y = C.float(y)
}

func (self Vertex) GetColor() Color {
	return Color{self.Cref.color}
}

func (self Vertex) SetColor(c Color) {
	self.Cref.color = c.Cref
} 