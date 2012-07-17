//
// SFML - Simple and Fast Multimedia Library
// Copyright (C) 2007-2009 Laurent Gomila (laurent.gom@gmail.com)
//
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it freely,
// subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented;
// you must not claim that you wrote the original software.
// If you use this software in a product, an acknowledgment
// in the product documentation would be appreciated but is not required.
//
// 2. Altered source versions must be plainly marked as such,
// and must not be misrepresented as being the original software.
//
// 3. This notice may not be removed or altered from any source distribution.
//
#ifndef SFML_VIEW_H
#define SFML_VIEW_H
// Headers
#include <SFML/Config.h>
#include <SFML/Graphics/Rect.h>
#include <SFML/Graphics/Types.h>
/// Construct a default view (1000x1000)
// sfView* sfView_Create();
func (self *View) Create() sfView* {
    return C.sfView_Create(self.Cref)
}
/// Construct a view from a rectangle
/// param Rect : Rectangle defining the bounds of the view
// sfView* sfView_CreateFromRect(sfFloatRect Rect);
func (self *View) CreateFromRect() sfView* {
    return C.sfView_CreateFromRect(self.Cref)
}
/// Destroy an existing view
/// param View : View to destroy
// void sfView_Destroy(sfView* View);
func (self *View) Destroy() void {
    return C.sfView_Destroy(self.Cref)
}
/// Change the center of a view
/// param View : View to modify
/// param X : X coordinate of the new center
/// param Y : Y coordinate of the new center
// void sfView_SetCenter(sfView* View, float X, float Y);
func (self *View) SetCenter(x float, y float) void {
    return C.sfView_SetCenter(self.Cref, x, y)
}
/// Change the half-size of a view
/// param View : View to modify
/// param HalfWidth : New half-width
/// param HalfHeight : New half-height
// void sfView_SetHalfSize(sfView* View, float HalfWidth, float HalfHeight);
func (self *View) SetHalfSize(halfWidth float, halfHeight float) void {
    return C.sfView_SetHalfSize(self.Cref, halfWidth, halfHeight)
}
/// Rebuild a view from a rectangle
/// param View : View to modify
/// param ViewRect : Rectangle defining the position and size of the view
// void sfView_SetFromRect(sfView* View, sfFloatRect ViewRect);
func (self *View) SetFromRect(viewRect sfFloatRect) void {
    return C.sfView_SetFromRect(self.Cref, viewRect)
}
/// Get the X coordinate of the center of a view
/// param View : View to read
/// return X coordinate of the center of the view
// float sfView_GetCenterX(sfView* View);
func (self *View) GetCenterX() float {
    return C.sfView_GetCenterX(self.Cref)
}
/// Get the Y coordinate of the center of a view
/// param View : View to read
/// return Y coordinate of the center of the view
// float sfView_GetCenterY(sfView* View);
func (self *View) GetCenterY() float {
    return C.sfView_GetCenterY(self.Cref)
}
/// Get the half-width of the view
/// param View : View to read
/// return Half-width of the view
// float sfView_GetHalfSizeX(sfView* View);
func (self *View) GetHalfSizeX() float {
    return C.sfView_GetHalfSizeX(self.Cref)
}
/// Get the half-height of the view
/// param View : View to read
/// return Half-height of the view
// float sfView_GetHalfSizeY(sfView* View);
func (self *View) GetHalfSizeY() float {
    return C.sfView_GetHalfSizeY(self.Cref)
}
/// Get the bounding rectangle of a view
/// param View : View to read
/// return Bounding rectangle of the view
// sfFloatRect sfView_GetRect(sfView* View);
func (self *View) GetRect() sfFloatRect {
    return C.sfView_GetRect(self.Cref)
}
/// Move a view
/// param View : View to move
/// param OffsetX : Offset to move the view, on X axis
/// param OffsetY : Offset to move the view, on Y axis
// void sfView_Move(sfView* View, float OffsetX, float OffsetY);
func (self *View) Move(offsetX float, offsetY float) void {
    return C.sfView_Move(self.Cref, offsetX, offsetY)
}
/// Resize a view rectangle to simulate a zoom / unzoom effect
/// param View : View to zoom
/// param Factor : Zoom factor to apply, relative to the current zoom
// void sfView_Zoom(sfView* View, float Factor);
func (self *View) Zoom(factor float) void {
    return C.sfView_Zoom(self.Cref, factor)
}
#endif // SFML_VIEW_H
