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


// #cgo LDFLAGS:-lcsfml-graphics
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Rect.h>
import "C"

import "fmt"

// sfFloatRect and sfIntRect are utility classes for
// manipulating rectangles.
// \brief Check if a point is inside a rectangle's area
//
// \param rect Rectangle to test
// \param x    X coordinate of the point to test
// \param y    Y coordinate of the point to test
//
// \return sfTrue if the point is inside
//
// sfBool sfFloatRect_contains(const sfFloatRect* rect, float x, float y);

type IntRect struct {
	Cref *C.sfIntRect
}

func NewIntRect(left, top, width, height int32) IntRect {
	fr := C.sfIntRect{
		C.int(left),
		C.int(top),
		C.int(width),
		C.int(height),
	}
	return IntRect{&fr}
}

// brief Check intersection between two rectangles
// param rect1        First rectangle to test
// param rect2        Second rectangle to test
// param intersection Rectangle to be filled with overlapping rect (can be NULL)
// return sfTrue if rectangles overlap
// sfBool sfIntRect_intersects(const sfIntRect* rect1, const sfIntRect* rect2, sfIntRect* // intersection);
func (self IntRect) Intersects(rect IntRect) (*IntRect, bool) {
	intersect := new(C.sfIntRect)
	b := C.sfIntRect_intersects(self.Cref, rect.Cref, intersect) == 1
	return &IntRect{intersect}, b
}

func (self IntRect) Left() int32 {
	return int32(self.Cref.left)
}

func (self IntRect) Top() int32 {
	return int32(self.Cref.top)
}

func (self IntRect) Width() int32 {
	return int32(self.Cref.width)
}

func (self IntRect) Height() int32 {
	return int32(self.Cref.height)
}

func (self IntRect) Contains(x, y int32) bool {
	return C.sfIntRect_contains(self.Cref, C.int(x), C.int(y)) == 1
}

func (self IntRect) Show() string {
	return fmt.Sprintf("<left: %d, top %d, width %d, height %d>",
		self.Cref.left, self.Cref.top, self.Cref.width, self.Cref.height)
}

// ------------------------------------------------------------------
type FloatRect struct {
	Cref *C.sfFloatRect
}

func NewFloatRect(left, top, width, height float32) FloatRect {
	fr := C.sfFloatRect{
		C.float(left),
		C.float(top),
		C.float(width),
		C.float(height),
	}

	//func SetFinalizer(x, f interface{})
	return FloatRect{&fr}
}

func (self FloatRect) Left() float32 {
	return float32(self.Cref.left)
}

func (self FloatRect) Top() float32 {
	return float32(self.Cref.top)
}

func (self FloatRect) Width() float32 {
	return float32(self.Cref.width)
}

func (self FloatRect) Height() float32 {
	return float32(self.Cref.height)
}

func (self FloatRect) Contains(x, y float32) bool {
	return C.sfFloatRect_contains(self.Cref, C.float(x), C.float(y)) == 1
}

func (self FloatRect) Show() string {
	return fmt.Sprintf("<left: %f, top %f, width %f, height %f>",
		self.Cref.left, self.Cref.top, self.Cref.width, self.Cref.height)
}

// brief Check intersection between two rectangles
// param rect1        First rectangle to test
// param rect2        Second rectangle to test
// param intersection Rectangle to be filled with overlapping rect (can be NULL)
// return sfTrue if rectangles overlap
// sfBool sfFloatRect_intersects(const sfFloatRect* rect1, const sfFloatRect* rect2, sfFloatRect* // intersection);
func (self FloatRect) Intersects(rect FloatRect) (*FloatRect, bool) {
	intersect := new(C.sfFloatRect)
	b := C.sfFloatRect_intersects(self.Cref, rect.Cref, intersect) == 1
	return &FloatRect{intersect}, b
}
