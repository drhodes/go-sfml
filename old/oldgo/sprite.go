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
#ifndef SFML_SPRITE_H
#define SFML_SPRITE_H
// Headers
#include <SFML/Config.h>
#include <SFML/Graphics/BlendMode.h>
#include <SFML/Graphics/Color.h>
#include <SFML/Graphics/Rect.h>
#include <SFML/Graphics/Types.h>
/// Create a new sprite
/// return A new sfSprite object, or NULL if it failed
// sfSprite* sfSprite_Create();
func (self *Sprite) Create() sfSprite* {
    return C.sfSprite_Create(self.Cref)
}
/// Destroy an existing sprite
/// param Sprite : Sprite to delete
// void sfSprite_Destroy(sfSprite* Sprite);
func (self *Sprite) Destroy() void {
    return C.sfSprite_Destroy(self.Cref)
}
/// Set the X position of a sprite
/// param Sprite : Sprite to modify
/// param X : New X coordinate
// void sfSprite_SetX(sfSprite* Sprite, float X);
func (self *Sprite) SetX(x float) void {
    return C.sfSprite_SetX(self.Cref, x)
}
/// Set the T position of a sprite
/// param Sprite : Sprite to modify
/// param Y : New Y coordinate
// void sfSprite_SetY(sfSprite* Sprite, float Y);
func (self *Sprite) SetY(y float) void {
    return C.sfSprite_SetY(self.Cref, y)
}
/// Set the position of a sprite
/// param Sprite : Sprite to modify
/// param X : New X coordinate
/// param Y : New Y coordinate
// void sfSprite_SetPosition(sfSprite* Sprite, float X, float Y);
func (self *Sprite) SetPosition(x float, y float) void {
    return C.sfSprite_SetPosition(self.Cref, x, y)
}
/// Set the horizontal scale of a sprite
/// param Sprite : Sprite to modify
/// param Scale : New scale (must be strictly positive)
// void sfSprite_SetScaleX(sfSprite* Sprite, float Scale);
func (self *Sprite) SetScaleX(scale float) void {
    return C.sfSprite_SetScaleX(self.Cref, scale)
}
/// Set the vertical scale of a sprite
/// param Sprite : Sprite to modify
/// param Scale : New scale (must be strictly positive)
// void sfSprite_SetScaleY(sfSprite* Sprite, float Scale);
func (self *Sprite) SetScaleY(scale float) void {
    return C.sfSprite_SetScaleY(self.Cref, scale)
}
/// Set the scale of a sprite
/// param Sprite : Sprite to modify
/// param ScaleX : New horizontal scale (must be strictly positive)
/// param ScaleY : New vertical scale (must be strictly positive)
// void sfSprite_SetScale(sfSprite* Sprite, float ScaleX, float ScaleY);
func (self *Sprite) SetScale(scaleX float, scaleY float) void {
    return C.sfSprite_SetScale(self.Cref, scaleX, scaleY)
}
/// Set the orientation of a sprite
/// param Sprite : Sprite to modify
/// param Rotation : Angle of rotation, in degrees
// void sfSprite_SetRotation(sfSprite* Sprite, float Rotation);
func (self *Sprite) SetRotation(rotation float) void {
    return C.sfSprite_SetRotation(self.Cref, rotation)
}
/// Set the center of a sprite, in coordinates relative to
/// its left-top corner
/// param Sprite : Sprite to modify
/// param X : X coordinate of the center
/// param Y : Y coordinate of the center
// void sfSprite_SetCenter(sfSprite* Sprite, float X, float Y);
func (self *Sprite) SetCenter(x float, y float) void {
    return C.sfSprite_SetCenter(self.Cref, x, y)
}
/// Set the color of a sprite
/// param Sprite : Sprite to modify
/// param Color : New color
// void sfSprite_SetColor(sfSprite* Sprite, sfColor Color);
func (self *Sprite) SetColor(color sfColor) void {
    return C.sfSprite_SetColor(self.Cref, color)
}
/// Set the blending mode for a sprite
/// param Sprite : Sprite to modify
/// param Mode : New blending mode
// void sfSprite_SetBlendMode(sfSprite* Sprite, sfBlendMode Mode);
func (self *Sprite) SetBlendMode(mode sfBlendMode) void {
    return C.sfSprite_SetBlendMode(self.Cref, mode)
}
/// Get the X position of a sprite
/// param Sprite : Sprite to read
/// return Current X position
// float sfSprite_GetX(sfSprite* Sprite);
func (self *Sprite) GetX() float {
    return C.sfSprite_GetX(self.Cref)
}
/// Get the Y position of a sprite
/// param Sprite : Sprite to read
/// return Current Y position
// float sfSprite_GetY(sfSprite* Sprite);
func (self *Sprite) GetY() float {
    return C.sfSprite_GetY(self.Cref)
}
/// Get the horizontal scale of a sprite
/// param Sprite : Sprite to read
/// return Current X scale factor (always positive)
// float sfSprite_GetScaleX(sfSprite* Sprite);
func (self *Sprite) GetScaleX() float {
    return C.sfSprite_GetScaleX(self.Cref)
}
/// Get the vertical scale of a sprite
/// param Sprite : Sprite to read
/// return Current Y scale factor (always positive)
// float sfSprite_GetScaleY(sfSprite* Sprite);
func (self *Sprite) GetScaleY() float {
    return C.sfSprite_GetScaleY(self.Cref)
}
/// Get the orientation of a sprite
/// param Sprite : Sprite to read
/// return Current rotation, in degrees
// float sfSprite_GetRotation(sfSprite* Sprite);
func (self *Sprite) GetRotation() float {
    return C.sfSprite_GetRotation(self.Cref)
}
/// Get the X position of the center a sprite
/// param Sprite : Sprite to read
/// return Current X center
// float sfSprite_GetCenterX(sfSprite* Sprite);
func (self *Sprite) GetCenterX() float {
    return C.sfSprite_GetCenterX(self.Cref)
}
/// Get the Y position of the center a sprite
/// param Sprite : Sprite to read
/// return Current Y center
// float sfSprite_GetCenterY(sfSprite* Sprite);
func (self *Sprite) GetCenterY() float {
    return C.sfSprite_GetCenterY(self.Cref)
}
/// Get the color of a sprite
/// param Sprite : Sprite to read
/// return Current color
// sfColor sfSprite_GetColor(sfSprite* Sprite);
func (self *Sprite) GetColor() sfColor {
    return C.sfSprite_GetColor(self.Cref)
}
/// Get the current blending mode of a sprite
/// param Sprite : Sprite to read
/// return Current blending mode
// sfBlendMode sfSprite_GetBlendMode(sfSprite* Sprite);
func (self *Sprite) GetBlendMode() sfBlendMode {
    return C.sfSprite_GetBlendMode(self.Cref)
}
/// Move a sprite
/// param Sprite : Sprite to modify
/// param OffsetX : Offset on the X axis
/// param OffsetY : Offset on the Y axis
// void sfSprite_Move(sfSprite* Sprite, float OffsetX, float OffsetY);
func (self *Sprite) Move(offsetX float, offsetY float) void {
    return C.sfSprite_Move(self.Cref, offsetX, offsetY)
}
/// Scale a sprite
/// param Sprite : Sprite to modify
/// param FactorX : Horizontal scaling factor (must be strictly positive)
/// param FactorY : Vertical scaling factor (must be strictly positive)
// void sfSprite_Scale(sfSprite* Sprite, float FactorX, float FactorY);
func (self *Sprite) Scale(factorX float, factorY float) void {
    return C.sfSprite_Scale(self.Cref, factorX, factorY)
}
/// Rotate a sprite
/// param Sprite : Sprite to modify
/// param Angle : Angle of rotation, in degrees
// void sfSprite_Rotate(sfSprite* Sprite, float Angle);
func (self *Sprite) Rotate(angle float) void {
    return C.sfSprite_Rotate(self.Cref, angle)
}
/// Transform a point from global coordinates into the sprite's local coordinates
/// (ie it applies the inverse of object's center, translation, rotation and scale to the point)
/// param Sprite : Sprite object
/// param PointX : X coordinate of the point to transform
/// param PointY : Y coordinate of the point to transform
/// param X : Value to fill with the X coordinate of the converted point
/// param Y : Value to fill with the y coordinate of the converted point
// void sfSprite_TransformToLocal(sfSprite* Sprite, float PointX, float PointY, float* X, float* Y);
func (self *Sprite) TransformToLocal(pointX float, pointY float, x float*, y float*) void {
    return C.sfSprite_TransformToLocal(self.Cref, pointX, pointY, x, y)
}
/// Transform a point from the sprite's local coordinates into global coordinates
/// (ie it applies the object's center, translation, rotation and scale to the point)
/// param Sprite : Sprite object
/// param PointX : X coordinate of the point to transform
/// param PointY : Y coordinate of the point to transform
/// param X : Value to fill with the X coordinate of the converted point
/// param Y : Value to fill with the y coordinate of the converted point
// void sfSprite_TransformToGlobal(sfSprite* Sprite, float PointX, float PointY, float* X, float* Y);
func (self *Sprite) TransformToGlobal(pointX float, pointY float, x float*, y float*) void {
    return C.sfSprite_TransformToGlobal(self.Cref, pointX, pointY, x, y)
}
/// Change the image of a sprite
/// param Sprite : Sprite to modify
/// param Image : New image
// void sfSprite_SetImage(sfSprite* Sprite, sfImage* Image);
func (self *Sprite) SetImage(image sfImage*) void {
    return C.sfSprite_SetImage(self.Cref, image)
}
/// Set the sub-rectangle of a sprite inside the source image
/// param Sprite : Sprite to modify
/// param SubRect : New sub-rectangle
// void sfSprite_SetSubRect(sfSprite* Sprite, sfIntRect SubRect);
func (self *Sprite) SetSubRect(subRect sfIntRect) void {
    return C.sfSprite_SetSubRect(self.Cref, subRect)
}
/// Resize a sprite (by changing its scale factors)
/// param Sprite : Sprite to modify
/// param Width : New width (must be strictly positive)
/// param Height : New height (must be strictly positive)
// void sfSprite_Resize(sfSprite* Sprite, float Width, float Height);
func (self *Sprite) Resize(width float, height float) void {
    return C.sfSprite_Resize(self.Cref, width, height)
}
/// Flip a sprite horizontally
/// param Sprite : Sprite to modify
/// param Flipped : sfTrue to flip the sprite
// void sfSprite_FlipX(sfSprite* Sprite, sfBool Flipped);
func (self *Sprite) FlipX(flipped sfBool) void {
    return C.sfSprite_FlipX(self.Cref, flipped)
}
/// Flip a sprite vertically
/// param Sprite : Sprite to modify
/// param Flipped : sfTrue to flip the sprite
// void sfSprite_FlipY(sfSprite* Sprite, sfBool Flipped);
func (self *Sprite) FlipY(flipped sfBool) void {
    return C.sfSprite_FlipY(self.Cref, flipped)
}
/// Get the source image of a sprite
/// param Sprite : Sprite to read
/// return Pointer to the image (can be NULL)
// sfImage* sfSprite_GetImage(sfSprite* Sprite);
func (self *Sprite) GetImage() sfImage* {
    return C.sfSprite_GetImage(self.Cref)
}
/// Get the sub-rectangle of a sprite inside the source image
/// param Sprite : Sprite to read
/// return Sub-rectangle
// sfIntRect sfSprite_GetSubRect(sfSprite* Sprite);
func (self *Sprite) GetSubRect() sfIntRect {
    return C.sfSprite_GetSubRect(self.Cref)
}
/// Get a sprite width
/// param Sprite : Sprite to read
/// return Width of the sprite
// float sfSprite_GetWidth(sfSprite* Sprite);
func (self *Sprite) GetWidth() float {
    return C.sfSprite_GetWidth(self.Cref)
}
/// Get a sprite height
/// param Sprite : Sprite to read
/// return Height of the sprite
// float sfSprite_GetHeight(sfSprite* Sprite);
func (self *Sprite) GetHeight() float {
    return C.sfSprite_GetHeight(self.Cref)
}
/// Get the color of a given pixel in a sprite
/// param Sprite : Sprite to read
/// param X : X coordinate of the pixel to get
/// param Y : Y coordinate of the pixel to get
/// return Color of pixel (X, Y)
// sfColor sfSprite_GetPixel(sfSprite* Sprite, unsigned int X, unsigned int Y);
func (self *Sprite) GetPixel(x uint, y uint) sfColor {
    return C.sfSprite_GetPixel(self.Cref, x, y)
}
#endif // SFML_SPRITE_H
