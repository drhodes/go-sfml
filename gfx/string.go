

// _String_
//-
// Create a new string
// return A new sfString object, or NULL if it failed
// sfString* sfString_Create();
func StringCreate() String {
    return C.sfString_Create(self.Cref)
}

// Destroy an existing string
// param String : String to delete
// sfString_Destroy(sfString* String);
func (self *String) Destroy() {
    return C.sfString_Destroy(self.Cref)
}

// Set the X position of a string
// param String : String to modify
// param X : New X coordinate
// sfString_SetX(sfString* String, float X);
func (self *String) SetX(x float) {
    return C.sfString_SetX(self.Cref, x)
}

// Set the Y position of a string
// param String : String to modify
// param Y : New Y coordinate
// sfString_SetY(sfString* String, float Y);
func (self *String) SetY(y float) {
    return C.sfString_SetY(self.Cref, y)
}

// Set the position of a string
// param String : String to modify
// param Left : New left coordinate
// param Top : New top coordinate
// sfString_SetPosition(sfString* String, float Left, float Top);
func (self *String) SetPosition(left float, top float) {
    return C.sfString_SetPosition(self.Cref, left, top)
}

// Set the horizontal scale of a string
// param String : String to modify
// param Scale : New scale (must be strictly positive)
// sfString_SetScaleX(sfString* String, float Scale);
func (self *String) SetScaleX(scale float) {
    return C.sfString_SetScaleX(self.Cref, scale)
}

// Set the vertical scale of a string
// param String : String to modify
// param Scale : New scale (must be strictly positive)
// sfString_SetScaleY(sfString* String, float Scale);
func (self *String) SetScaleY(scale float) {
    return C.sfString_SetScaleY(self.Cref, scale)
}

// Set the scale of a string
// param String : String to modify
// param ScaleX : New horizontal scale (must be strictly positive)
// param ScaleY : New vertical scale (must be strictly positive)
// sfString_SetScale(sfString* String, float ScaleX, float ScaleY);
func (self *String) SetScale(scaleX float, scaleY float) {
    return C.sfString_SetScale(self.Cref, scaleX, scaleY)
}

// Set the orientation of a string
// param String : String to modify
// param Rotation : Angle of rotation, in degrees
// sfString_SetRotation(sfString* String, float Rotation);
func (self *String) SetRotation(rotation float) {
    return C.sfString_SetRotation(self.Cref, rotation)
}

// Set the center of a string, in coordinates
// relative to its left-top corner
// param String : String to modify
// param X : X coordinate of the center
// param Y : Y coordinate of the center
// sfString_SetCenter(sfString* String, float X, float Y);
func (self *String) SetCenter(x float, y float) {
    return C.sfString_SetCenter(self.Cref, x, y)
}

// Set the color of a string
// param String : String to modify
// param Color : New color
// sfString_SetColor(sfString* String, sfColor Color);
func (self *String) SetColor(color sfColor) {
    return C.sfString_SetColor(self.Cref, color)
}

// Set the blending mode for a string
// param String : String to modify
// param Mode : New blending mode
// sfString_SetBlendMode(sfString* String, sfBlendMode Mode);
func (self *String) SetBlendMode(mode sfBlendMode) {
    return C.sfString_SetBlendMode(self.Cref, mode)
}

// Get the X position of a string
// param String : String to read
// return Current X position
// float sfString_GetX(sfString* String);
func (self *String) GetX() float {
    return C.sfString_GetX(self.Cref)
}

// Get the top Y of a string
// param String : String to read
// return Current Y position
// float sfString_GetY(sfString* String);
func (self *String) GetY() float {
    return C.sfString_GetY(self.Cref)
}

// Get the horizontal scale of a string
// param String : String to read
// return Current X scale factor (always positive)
// float sfString_GetScaleX(sfString* String);
func (self *String) GetScaleX() float {
    return C.sfString_GetScaleX(self.Cref)
}

// Get the vertical scale of a string
// param String : String to read
// return Current Y scale factor (always positive)
// float sfString_GetScaleY(sfString* String);
func (self *String) GetScaleY() float {
    return C.sfString_GetScaleY(self.Cref)
}

// Get the orientation of a string
// param String : String to read
// return Current rotation, in degrees
// float sfString_GetRotation(sfString* String);
func (self *String) GetRotation() float {
    return C.sfString_GetRotation(self.Cref)
}

// Get the X position of the center a string
// param String : String to read
// return Current X center position
// float sfString_GetCenterX(sfString* String);
func (self *String) GetCenterX() float {
    return C.sfString_GetCenterX(self.Cref)
}

// Get the top Y of the center of a string
// param String : String to read
// return Current Y center position
// float sfString_GetCenterY(sfString* String);
func (self *String) GetCenterY() float {
    return C.sfString_GetCenterY(self.Cref)
}

// Get the color of a string
// param String : String to read
// return Current color
// sfColor sfString_GetColor(sfString* String);
func (self *String) GetColor() sfColor {
    return C.sfString_GetColor(self.Cref)
}

// Get the current blending mode of a string
// param String : String to read
// return Current blending mode
// sfBlendMode sfString_GetBlendMode(sfString* String);
func (self *String) GetBlendMode() sfBlendMode {
    return C.sfString_GetBlendMode(self.Cref)
}

// Move a string
// param String : String to modify
// param OffsetX : Offset on the X axis
// param OffsetY : Offset on the Y axis
// sfString_Move(sfString* String, float OffsetX, float OffsetY);
func (self *String) Move(offsetX float, offsetY float) {
    return C.sfString_Move(self.Cref, offsetX, offsetY)
}

// Scale a string
// param String : String to modify
// param FactorX : Horizontal scaling factor (must be strictly positive)
// param FactorY : Vertical scaling factor (must be strictly positive)
// sfString_Scale(sfString* String, float FactorX, float FactorY);
func (self *String) Scale(factorX float, factorY float) {
    return C.sfString_Scale(self.Cref, factorX, factorY)
}

// Rotate a string
// param String : String to modify
// param Angle : Angle of rotation, in degrees
// sfString_Rotate(sfString* String, float Angle);
func (self *String) Rotate(angle float) {
    return C.sfString_Rotate(self.Cref, angle)
}

// Transform a point from global coordinates into the string's local coordinates
// (ie it applies the inverse of object's center, translation, rotation and scale to the point)
// param String : String object
// param PointX : X coordinate of the point to transform
// param PointY : Y coordinate of the point to transform
// param X : Value to fill with the X coordinate of the converted point
// param Y : Value to fill with the y coordinate of the converted point
// sfString_TransformToLocal(sfString* String, float PointX, float PointY, float* X, float* Y);
func (self *String) TransformToLocal(pointX float, pointY float, x float*, y float*) {
    return C.sfString_TransformToLocal(self.Cref, pointX, pointY, x, y)
}

// Transform a point from the string's local coordinates into global coordinates
// (ie it applies the object's center, translation, rotation and scale to the point)
// param String : String object
// param PointX : X coordinate of the point to transform
// param PointY : Y coordinate of the point to transform
// param X : Value to fill with the X coordinate of the converted point
// param Y : Value to fill with the y coordinate of the converted point
// sfString_TransformToGlobal(sfString* String, float PointX, float PointY, float* X, float* Y);
func (self *String) TransformToGlobal(pointX float, pointY float, x float*, y float*) {
    return C.sfString_TransformToGlobal(self.Cref, pointX, pointY, x, y)
}

// Set the text of a string (from a multibyte string)
// param String : String to modify
// param Text : New text
// sfString_SetText(sfString* String, const char* Text);
func (self *String) SetText(text char*) {
    return C.sfString_SetText(self.Cref, text)
}

// Set the text of a string (from a unicode string)
// param String : String to modify
// param Text : New text
// sfString_SetUnicodeText(sfString* String, const sfUint32* Text);
func (self *String) SetUnicodeText(text sfUint32*) {
    return C.sfString_SetUnicodeText(self.Cref, text)
}

// Set the font of a string
// param String : String to modify
// param Font : Font to use
// sfString_SetFont(sfString* String, sfFont* Font);
func (self *String) SetFont(font sfFont*) {
    return C.sfString_SetFont(self.Cref, font)
}

// Set the size of a string
// param String : String to modify
// param Size : New size, in pixels
// sfString_SetSize(sfString* String, float Size);
func (self *String) SetSize(size float) {
    return C.sfString_SetSize(self.Cref, size)
}

// Set the style of a string
// param String : String to modify
// param Size : New style (see sfStringStyle enum)
// sfString_SetStyle(sfString* String, unsigned long Style);
func (self *String) SetStyle(style ulong) {
    return C.sfString_SetStyle(self.Cref, style)
}

// Get the text of a string (returns a unicode string)
// param String : String to read
// return Text as UTF-32
// const sfUint32* sfString_GetUnicodeText(sfString* String);
// Get the text of a string (returns an ANSI string)
// param String : String to read
// return Text an a locale-dependant ANSI string
// const char* sfString_GetText(sfString* String);
// Get the font used by a string
// param String : String to read
// return Pointer to the font
// sfFont* sfString_GetFont(sfString* String);
func (self *String) GetFont() sfFont* {
    return C.sfString_GetFont(self.Cref)
}

// Get the size of the characters of a string
// param String : String to read
// return Size of the characters
// float sfString_GetSize(sfString* String);
func (self *String) GetSize() float {
    return C.sfString_GetSize(self.Cref)
}

// Get the style of a string
// param String : String to read
// return Current string style (see sfStringStyle enum)
// unsigned long sfString_GetStyle(sfString* String);
// Return the visual position of the Index-th character of the string,
// in coordinates relative to the string
// (note : translation, center, rotation and scale are not applied)
// param String : String to read
// param Index : Index of the character
// param X : Value to fill with the X coordinate of the position
// param Y : Value to fill with the y coordinate of the position
// sfString_GetCharacterPos(sfString* String, size_t Index, float* X, float* Y);
func (self *String) GetCharacterPos(index size_t, x float*, y float*) {
    return C.sfString_GetCharacterPos(self.Cref, index, x, y)
}

// Get the bounding rectangle of a string on screen
// param String : String to read
// return Rectangle contaning the string in screen coordinates
// sfFloatRect sfString_GetRect(sfString* String);
func (self *String) GetRect() sfFloatRect {
    return C.sfString_GetRect(self.Cref)
}

#endif // SFML_STRING_H
