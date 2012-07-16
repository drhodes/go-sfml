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
#ifndef SFML_Shape_H
#define SFML_Shape_H
// Headers
#include <SFML/Config.h>
#include <SFML/Graphics/BlendMode.h>
#include <SFML/Graphics/Color.h>
#include <SFML/Graphics/Types.h>
/// Create a new shape
/// return A new sfShape object, or NULL if it failed
// sfShape* sfShape_Create();
func (self *Shape) Create() sfShape* {
    return C.sfShape_Create(self.Cref)
}
/// Create a new shape made of a single line
/// param P1X, P1Y : Position of the first point
/// param P2X, P2Y : Position second point
/// param Thickness : Line thickness
/// param Col : Color used to draw the line
/// param Outline : Outline width
/// param OutlineCol : Color used to draw the outline
// sfShape* sfShape_CreateLine(float P1X, float P1Y, float P2X, float P2Y, float Thickness, sfColor Col, float Outline, sfColor OutlineCol);
func (self *Shape) CreateLine(p1Y float, p2X float, p2Y float, thickness float, col sfColor, outline float, outlineCol sfColor) sfShape* {
    return C.sfShape_CreateLine(self.Cref, p1Y, p2X, p2Y, thickness, col, outline, outlineCol)
}
/// Create a new shape made of a single rectangle
/// param P1X, P1Y : Position of the first point
/// param P2X, P2Y : Position second point
/// param Col : Color used to fill the rectangle
/// param Outline : Outline width
/// param OutlineCol : Color used to draw the outline
// sfShape* sfShape_CreateRectangle(float P1X, float P1Y, float P2X, float P2Y, sfColor Col, float Outline, sfColor OutlineCol);
func (self *Shape) CreateRectangle(p1Y float, p2X float, p2Y float, col sfColor, outline float, outlineCol sfColor) sfShape* {
    return C.sfShape_CreateRectangle(self.Cref, p1Y, p2X, p2Y, col, outline, outlineCol)
}
/// Create a new shape made of a single circle
/// param X, Y : Position of the center
/// param Radius : Radius
/// param Col : Color used to fill the circle
/// param Outline : Outline width
/// param OutlineCol : Color used to draw the outline
// sfShape* sfShape_CreateCircle(float X, float Y, float Radius, sfColor Col, float Outline, sfColor OutlineCol);
func (self *Shape) CreateCircle(y float, radius float, col sfColor, outline float, outlineCol sfColor) sfShape* {
    return C.sfShape_CreateCircle(self.Cref, y, radius, col, outline, outlineCol)
}
/// Destroy an existing Shape
/// param Shape : Shape to delete
// void sfShape_Destroy(sfShape* Shape);
func (self *Shape) Destroy() void {
    return C.sfShape_Destroy(self.Cref)
}
/// Set the X position of a shape
/// param Shape : Shape to modify
/// param X : New X coordinate
// void sfShape_SetX(sfShape* Shape, float X);
func (self *Shape) SetX(x float) void {
    return C.sfShape_SetX(self.Cref, x)
}
/// Set the Y position of a shape
/// param Shape : Shape to modify
/// param Y : New Y coordinate
// void sfShape_SetY(sfShape* Shape, float Y);
func (self *Shape) SetY(y float) void {
    return C.sfShape_SetY(self.Cref, y)
}
/// Set the position of a shape
/// param Shape : Shape to modify
/// param X : New X coordinate
/// param Y : New Y coordinate
// void sfShape_SetPosition(sfShape* Shape, float X, float Y);
func (self *Shape) SetPosition(x float, y float) void {
    return C.sfShape_SetPosition(self.Cref, x, y)
}
/// Set the horizontal scale of a shape
/// param Shape : Shape to modify
/// param Scale : New scale (must be strictly positive)
// void sfShape_SetScaleX(sfShape* Shape, float Scale);
func (self *Shape) SetScaleX(scale float) void {
    return C.sfShape_SetScaleX(self.Cref, scale)
}
/// Set the vertical scale of a shape
/// param Shape : Shape to modify
/// param Scale : New scale (must be strictly positive)
// void sfShape_SetScaleY(sfShape* Shape, float Scale);
func (self *Shape) SetScaleY(scale float) void {
    return C.sfShape_SetScaleY(self.Cref, scale)
}
/// Set the scale of a shape
/// param Shape : Shape to modify
/// param ScaleX : New horizontal scale (must be strictly positive)
/// param ScaleY : New vertical scale (must be strictly positive)
// void sfShape_SetScale(sfShape* Shape, float ScaleX, float ScaleY);
func (self *Shape) SetScale(scaleX float, scaleY float) void {
    return C.sfShape_SetScale(self.Cref, scaleX, scaleY)
}
/// Set the orientation of a shape
/// param Shape : Shape to modify
/// param Rotation : Angle of rotation, in degrees
// void sfShape_SetRotation(sfShape* Shape, float Rotation);
func (self *Shape) SetRotation(rotation float) void {
    return C.sfShape_SetRotation(self.Cref, rotation)
}
/// Set the center of a shape, in coordinates relative to
/// its left-top corner
/// param Shape : Shape to modify
/// param X : X coordinate of the center
/// param Y : Y coordinate of the center
// void sfShape_SetCenter(sfShape* Shape, float X, float Y);
func (self *Shape) SetCenter(x float, y float) void {
    return C.sfShape_SetCenter(self.Cref, x, y)
}
/// Set the color of a shape
/// param Shape : Shape to modify
/// param Color : New color
// void sfShape_SetColor(sfShape* Shape, sfColor Color);
func (self *Shape) SetColor(color sfColor) void {
    return C.sfShape_SetColor(self.Cref, color)
}
/// Set the blending mode for a shape
/// param Shape : Shape to modify
/// param Mode : New blending mode
// void sfShape_SetBlendMode(sfShape* Shape, sfBlendMode Mode);
func (self *Shape) SetBlendMode(mode sfBlendMode) void {
    return C.sfShape_SetBlendMode(self.Cref, mode)
}
/// Get the X position of a shape
/// param Shape : Shape to read
/// return Current X position
// float sfShape_GetX(sfShape* Shape);
func (self *Shape) GetX() float {
    return C.sfShape_GetX(self.Cref)
}
/// Get the Y position of a shape
/// param Shape : Shape to read
/// return Current Y position
// float sfShape_GetY(sfShape* Shape);
func (self *Shape) GetY() float {
    return C.sfShape_GetY(self.Cref)
}
/// Get the horizontal scale of a shape
/// param Shape : Shape to read
/// return Current X scale factor (always positive)
// float sfShape_GetScaleX(sfShape* Shape);
func (self *Shape) GetScaleX() float {
    return C.sfShape_GetScaleX(self.Cref)
}
/// Get the vertical scale of a shape
/// param Shape : Shape to read
/// return Current Y scale factor (always positive)
// float sfShape_GetScaleY(sfShape* Shape);
func (self *Shape) GetScaleY() float {
    return C.sfShape_GetScaleY(self.Cref)
}
/// Get the orientation of a shape
/// param Shape : Shape to read
/// return Current rotation, in degrees
// float sfShape_GetRotation(sfShape* Shape);
func (self *Shape) GetRotation() float {
    return C.sfShape_GetRotation(self.Cref)
}
/// Get the X position of the center a shape
/// param Shape : Shape to read
/// return Current X center
// float sfShape_GetCenterX(sfShape* Shape);
func (self *Shape) GetCenterX() float {
    return C.sfShape_GetCenterX(self.Cref)
}
/// Get the Y position of the center a shape
/// param Shape : Shape to read
/// return Current Y center
// float sfShape_GetCenterY(sfShape* Shape);
func (self *Shape) GetCenterY() float {
    return C.sfShape_GetCenterY(self.Cref)
}
/// Get the color of a shape
/// param Shape : Shape to read
/// return Current color
// sfColor sfShape_GetColor(sfShape* Shape);
func (self *Shape) GetColor() sfColor {
    return C.sfShape_GetColor(self.Cref)
}
/// Get the current blending mode of a shape
/// param Shape : Shape to read
/// return Current blending mode
// sfBlendMode sfShape_GetBlendMode(sfShape* Shape);
func (self *Shape) GetBlendMode() sfBlendMode {
    return C.sfShape_GetBlendMode(self.Cref)
}
/// Move a shape
/// param Shape : Shape to modify
/// param OffsetX : Offset on the X axis
/// param OffsetY : Offset on the Y axis
// void sfShape_Move(sfShape* Shape, float OffsetX, float OffsetY);
func (self *Shape) Move(offsetX float, offsetY float) void {
    return C.sfShape_Move(self.Cref, offsetX, offsetY)
}
/// Scale a shape
/// param Shape : Shape to modify
/// param FactorX : Horizontal scaling factor (must be strictly positive)
/// param FactorY : Vertical scaling factor (must be strictly positive)
// void sfShape_Scale(sfShape* Shape, float FactorX, float FactorY);
func (self *Shape) Scale(factorX float, factorY float) void {
    return C.sfShape_Scale(self.Cref, factorX, factorY)
}
/// Rotate a shape
/// param Shape : Shape to modify
/// param Angle : Angle of rotation, in degrees
// void sfShape_Rotate(sfShape* Shape, float Angle);
func (self *Shape) Rotate(angle float) void {
    return C.sfShape_Rotate(self.Cref, angle)
}
/// Transform a point from global coordinates into the shape's local coordinates
/// (ie it applies the inverse of object's center, translation, rotation and scale to the point)
/// param Shape : Shape object
/// param PointX : X coordinate of the point to transform
/// param PointY : Y coordinate of the point to transform
/// param X : Value to fill with the X coordinate of the converted point
/// param Y : Value to fill with the y coordinate of the converted point
// void sfShape_TransformToLocal(sfShape* Shape, float PointX, float PointY, float* X, float* Y);
func (self *Shape) TransformToLocal(pointX float, pointY float, x float*, y float*) void {
    return C.sfShape_TransformToLocal(self.Cref, pointX, pointY, x, y)
}
/// Transform a point from the shape's local coordinates into global coordinates
/// (ie it applies the object's center, translation, rotation and scale to the point)
/// param Shape : Shape object
/// param PointX : X coordinate of the point to transform
/// param PointY : Y coordinate of the point to transform
/// param X : Value to fill with the X coordinate of the converted point
/// param Y : Value to fill with the y coordinate of the converted point
// void sfShape_TransformToGlobal(sfShape* Shape, float PointX, float PointY, float* X, float* Y);
func (self *Shape) TransformToGlobal(pointX float, pointY float, x float*, y float*) void {
    return C.sfShape_TransformToGlobal(self.Cref, pointX, pointY, x, y)
}
/// Add a point to a shape
/// param Shape : Shape to modify
/// param X, Y : Position of the point
/// param Col : Color of the point
/// param OutlineCol : Outline color of the point
// void sfShape_AddPoint(sfShape* Shape, float X, float Y, sfColor Col, sfColor OutlineCol);
func (self *Shape) AddPoint(x float, y float, col sfColor, outlineCol sfColor) void {
    return C.sfShape_AddPoint(self.Cref, x, y, col, outlineCol)
}
/// Enable or disable filling a shape.
/// Fill is enabled by default
/// param Enable : True to enable, false to disable
// void sfShape_EnableFill(sfShape* Shape, sfBool Enable);
func (self *Shape) EnableFill(enable sfBool) void {
    return C.sfShape_EnableFill(self.Cref, enable)
}
/// Enable or disable drawing a shape outline.
/// Outline is enabled by default
/// param Shape : Shape to modify
/// param Enable : True to enable, false to disable
// void sfShape_EnableOutline(sfShape* Shape, sfBool Enable);
func (self *Shape) EnableOutline(enable sfBool) void {
    return C.sfShape_EnableOutline(self.Cref, enable)
}
/// Change the width of a shape outline
/// param Shape : Shape to modify
/// param Width : New width
// void sfShape_SetOutlineWidth(sfShape* Shape, float Width);
func (self *Shape) SetOutlineWidth(width float) void {
    return C.sfShape_SetOutlineWidth(self.Cref, width)
}
/// Get the width of a shape outline
/// param Shape : Shape to read
/// param return Current outline width
// float sfShape_GetOutlineWidth(sfShape* Shape);
func (self *Shape) GetOutlineWidth() float {
    return C.sfShape_GetOutlineWidth(self.Cref)
}
/// Get the number of points composing a shape
/// param Shape : Shape to read
/// return Total number of points
// unsigned int sfShape_GetNbPoints(sfShape* Shape);
/// Get a the position of a shape's point
/// param Shape : Shape to read
/// param Index : Index of the point to get
/// param X : Variable to fill with the X coordinate of the point
/// param Y : Variable to fill with the Y coordinate of the point
// void sfShape_GetPointPosition(sfShape* Shape, unsigned int Index, float* X, float* Y);
func (self *Shape) GetPointPosition(index uint, x float*, y float*) void {
    return C.sfShape_GetPointPosition(self.Cref, index, x, y)
}
/// Get a the color of a shape's point
/// param Shape : Shape to read
/// param Index : Index of the point to get
/// return Color of the point
// sfColor sfShape_GetPointColor(sfShape* Shape, unsigned int Index);
func (self *Shape) GetPointColor(index uint) sfColor {
    return C.sfShape_GetPointColor(self.Cref, index)
}
/// Get a the outline color of a shape's point
/// param Shape : Shape to read
/// param Index : Index of the point to get
/// return Outline color of the point
// sfColor sfShape_GetPointOutlineColor(sfShape* Shape, unsigned int Index);
func (self *Shape) GetPointOutlineColor(index uint) sfColor {
    return C.sfShape_GetPointOutlineColor(self.Cref, index)
}
/// Set a the position of a shape's point
/// param Shape : Shape to read
/// param Index : Index of the point to get
/// param X : X coordinate of the point
/// param Y : Y coordinate of the point
// void sfShape_SetPointPosition(sfShape* Shape, unsigned int Index, float X, float Y);
func (self *Shape) SetPointPosition(index uint, x float, y float) void {
    return C.sfShape_SetPointPosition(self.Cref, index, x, y)
}
/// Set a the color of a shape's point
/// param Shape : Shape to read
/// param Index : Index of the point to get
/// param Color : Color of the point
// void sfShape_SetPointColor(sfShape* Shape, unsigned int Index, sfColor Color);
func (self *Shape) SetPointColor(index uint, color sfColor) void {
    return C.sfShape_SetPointColor(self.Cref, index, color)
}
/// Set a the outline color of a shape's point
/// param Shape : Shape to read
/// param Index : Index of the point to get
/// param Color : Outline color of the point
// void sfShape_SetPointOutlineColor(sfShape* Shape, unsigned int Index, sfColor Color);
func (self *Shape) SetPointOutlineColor(index uint, color sfColor) void {
    return C.sfShape_SetPointOutlineColor(self.Cref, index, color)
}
#endif // SFML_Shape_H
