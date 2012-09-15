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

type VertexArray struct {
	Cref *C.sfVertexArray
}



// \brief Create a new vertex array
// \return A new sfVertexArray object
// sfVertexArray* sfVertexArray_create(void);
func NewVertexArray() VertexArray { 
    return VertexArray{C.sfVertexArray_create()}
}

// \brief Copy an existing vertex array
// \param vertexArray Vertex array to copy
// \return Copied object
// sfVertexArray* sfVertexArray_copy(sfVertexArray* vertexArray);
func (self VertexArray) Copy() VertexArray { 
    return VertexArray{C.sfVertexArray_copy(self.Cref)}
}
           
// \brief Destroy an existing vertex array
// \param vertexArray Vertex array to delete
// void sfVertexArray_destroy(sfVertexArray* vertexArray);
func (self VertexArray) Destroy() { 
    C.sfVertexArray_destroy(self.Cref)
}
           
// \brief Return the vertex count of a vertex array
// \param vertexArray Vertex array object
// \return Number of vertices in the array
// unsigned int sfVertexArray_getVertexCount(const sfVertexArray* vertexArray);
func (self VertexArray) GetVertexCount() uint {
    return uint(C.sfVertexArray_getVertexCount(self.Cref))
}
           
// \brief Get access to a vertex by its index
// This function doesn't check \a index, it must be in range
// [0, vertex count - 1]. The behaviour is undefined
// otherwise.
// \param vertexArray Vertex array object
// \param index       Index of the vertex to get
// \return Pointer to the index-th vertex
// sfVertex* sfVertexArray_getVertex(sfVertexArray* vertexArray, unsigned int index);
func (self VertexArray) GetVertex(index uint) Vertex { 
	return Vertex{ C.sfVertexArray_getVertex(self.Cref, C.uint(index)) }
}

           
// \brief Clear a vertex array
// This function removes all the vertices from the array.
// It doesn't deallocate the corresponding memory, so that
// adding new vertices after clearing doesn't involve
// reallocating all the memory.
// \param vertexArray Vertex array object
// void sfVertexArray_clear(sfVertexArray* vertexArray);
func (self VertexArray) Clear() { 
    C.sfVertexArray_clear(self.Cref)
}
            
// \brief Resize the vertex array
// If \a vertexCount is greater than the current size, the previous
// vertices are kept and new (default-constructed) vertices are
// added.
// If \a vertexCount is less than the current size, existing vertices
// are removed from the array.
// \param vertexArray Vertex array objet
// \param vertexCount New size of the array (number of vertices)
// void sfVertexArray_resize(sfVertexArray* vertexArray, unsigned int vertexCount);
func (self VertexArray) Resize(vertexCount uint) { 
    C.sfVertexArray_resize(self.Cref, C.uint(vertexCount))
}
            
// \brief Add a vertex to a vertex array array
// \param vertexArray Vertex array objet
// \param vertex      Vertex to add
// void sfVertexArray_append(sfVertexArray* vertexArray, sfVertex vertex);
func (self VertexArray) Append(vertex Vertex) { 
    C.sfVertexArray_append(self.Cref, *vertex.Cref)
}

// \brief Set the type of primitives of a vertex array
// This function defines how the vertices must be interpreted
// when it's time to draw them:
// \li As points
// \li As lines
// \li As triangles
// \li As quads
// The default primitive type is sfPoints.
// \param vertexArray Vertex array objet
// \param type        Type of primitive
// void sfVertexArray_setPrimitiveType(sfVertexArray* vertexArray, sfPrimitiveType type);
func (self VertexArray) SetPrimitiveType(ptype PrimitiveType) { 
    C.sfVertexArray_setPrimitiveType(self.Cref, C.sfPrimitiveType(ptype))
}

// \brief Get the type of primitives drawn by a vertex array
// \param vertexArray Vertex array objet
// \return Primitive type
// sfPrimitiveType sfVertexArray_getPrimitiveType(sfVertexArray* vertexArray);
func (self VertexArray) Getprimitivetype() PrimitiveType { 
    return PrimitiveType(C.sfVertexArray_getPrimitiveType(self.Cref))
}
            
// \brief Compute the bounding rectangle of a vertex array
// This function returns the axis-aligned rectangle that
// contains all the vertices of the array.
// \param vertexArray Vertex array objet
// \return Bounding rectangle of the vertex array
// sfFloatRect sfVertexArray_getBounds(sfVertexArray* vertexArray);
func (self VertexArray) Getbounds() FloatRect { 
	r := C.sfVertexArray_getBounds(self.Cref)
	return FloatRect{&r}
}
            
