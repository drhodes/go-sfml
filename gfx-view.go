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
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/System/Vector2.h>
// #include <SFML/Graphics/View.h>
import "C"

type View struct {
	Cref *C.sfView
}

// \brief Create a default view
// This function creates a default view of (0, 0, 1000, 1000)
// \return A new sfView object
// sfView* sfView_create(void);
func (self View) NewView() View {
	return View{C.sfView_create()}
}

// \brief Construct a view from a rectangle
// \param rectangle Rectangle defining the zone to display
// \return A new sfView object
// sfView* sfView_createFromRect(sfFloatRect rectangle);
func (self View) NewViewFromRect(r FloatRect) View {
	return View{C.sfView_createFromRect(*r.Cref)}
}

// \brief Copy an existing view
// \param view View to copy
// \return Copied object
// sfView* sfView_copy(sfView* view);
func (self View) Copy() View {
	return View{C.sfView_copy(self.Cref)}
}

// \brief Destroy an existing view
// \param view View to destroy
// void sfView_destroy(sfView* view);
func (self View) Destroy() {
	C.sfView_destroy(self.Cref)
}

// \brief Set the center of a view
// \param view   View object
// \param center New center
// void sfView_setCenter(sfView* view, sfVector2f center);
func (self View) SetCenter(x, y float32) {
	C.sfView_setCenter(self.Cref, C.sfVector2f{
		C.float(x), C.float(y),
	})
}

// \brief Set the size of a view
// \param view View object
// \param size New size of the view
// void sfView_setSize(sfView* view, sfVector2f size);
func (self View) SetSize(x, y float32) {
	C.sfView_setSize(self.Cref, C.sfVector2f{
		C.float(x), C.float(y),
	})
}
            
// \brief Set the orientation of a view
// The default rotation of a view is 0 degree.
// \param view  View object
// \param angle New angle, in degrees
// void sfView_setRotation(sfView* view, float angle);
func (self View) SetRotation(angle float32) { 
	C.sfView_setRotation(self.Cref, C.float(angle))
}
            
// \brief Set the target viewport of a view
// The viewport is the rectangle into which the contents of the
// view are displayed, expressed as a factor (between 0 and 1)
// of the size of the render target to which the view is applied.
// For example, a view which takes the left side of the target would
// be defined by a rect of (0, 0, 0.5, 1).
// By default, a view has a viewport which covers the entire target.
// \param view     View object
// \param viewport New viewport rectangle
// void sfView_setViewport(sfView* view, sfFloatRect viewport);
func (self View) SetViewport(viewport FloatRect) {
	C.sfView_setViewport(self.Cref, *viewport.Cref)
}

// \brief Reset a view to the given rectangle
// Note that this function resets the rotation angle to 0.
// \param view      View object
// \param rectangle Rectangle defining the zone to display
// void sfView_reset(sfView* view, sfFloatRect rectangle);
func (self View) Reset(rectangle FloatRect) {
	C.sfView_reset(self.Cref, *rectangle.Cref)
}

// \brief Get the center of a view
// \param view View object
// \return Center of the view
// sfVector2f sfView_getCenter(const sfView* view);
func (self View) GetCenter() (x, y float32) {
	p := C.sfView_getCenter(self.Cref)
	return float32(p.x), float32(p.y)
}

// \brief Get the size of a view
// \param view View object
// \return Size of the view
// sfVector2f sfView_getSize(const sfView* view);
func (self View) GetSize() (x, y float32) {
	p := C.sfView_getSize(self.Cref)
	return float32(p.x), float32(p.y)
}

// \brief Get the current orientation of a view
// \param view View object
// \return Rotation angle of the view, in degrees
// float sfView_getRotation(const sfView* view);
func (self View) GetRotation() float32 { 
    return float32(C.sfView_getRotation(self.Cref))
}
            
// \brief Get the target viewport rectangle of a view
// \param view View object
// \return Viewport rectangle, expressed as a factor of the target size
// sfFloatRect sfView_getViewport(const sfView* view);
func (self View) GetViewport() FloatRect { 
    r := C.sfView_getViewport(self.Cref)
	return FloatRect{&r}
}

// \brief Move a view relatively to its current position
// \param view   View object
// \param offset Offset
// void sfView_move(sfView* view, sfVector2f offset);
func (self View) Move(x, y float32) { 
	C.sfView_move(self.Cref, C.sfVector2f {
		C.float(x), C.float(y),
	})
}
            
// \brief Rotate a view relatively to its current orientation
// \param view  View object
// \param angle Angle to rotate, in degrees
// void sfView_rotate(sfView* view, float angle);
func (self View) Rotate(angle float32) { 
	C.sfView_rotate(self.Cref, C.float(angle))
}

// \brief Resize a view rectangle relatively to its current size
// Resizing the view simulates a zoom, as the zone displayed on
// screen grows or shrinks.
// \a factor is a multiplier:
// \li 1 keeps the size unchanged
// \li > 1 makes the view bigger (objects appear smaller)
// \li < 1 makes the view smaller (objects appear bigger)
// \param view   View object
// \param factor Zoom factor to apply
// void sfView_zoom(sfView* view, float factor);
func (self View) Zoom(factor float32) {
	C.sfView_zoom(self.Cref, C.float(factor))
}
