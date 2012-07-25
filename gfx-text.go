package sfml

// #cgo LDFLAGS:-lcsfml-graphics -ljpeg
// #include <SFML/Graphics/Text.h>
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/System/Vector2.h>
// #include <stddef.h>
import "C"

type Text struct {
	Cref *C.sfText
}

/*
// sfText styles
// Create a new text
// \return A new sfText object, or NULL if it failed
// sfText* sfText_create(void);
func (self Text) NewText() Text { 
    return C.sfText_create(self.Cref);
}
            
// Copy an existing text
// \param text Text to copy
// \return Copied object
// sfText* sfText_copy(sfText* text);
func (self Text) Copy() *Text { 
    return C.sfText_copy(self.Cref);
}
            
// Destroy an existing text
// \param text Text to delete
// void sfText_destroy(sfText* text);
func (self Text) Destroy() void { 
    return C.sfText_destroy(self.Cref);
}
            
// Set the position of a text
// This function completely overwrites the previous position.
// See sfText_move to apply an offset based on the previous position instead.
// The default position of a text Text object is (0, 0).
// \param text     Text object
// \param position New position
// void sfText_setPosition(sfText* text, sfVector2f position);
func (self Text) Setposition(position Vector2f) void { 
    return C.sfText_setPosition(self.Cref, sfVector2f(position));
}
            
// Set the orientation of a text
// This function completely overwrites the previous rotation.
// See sfText_rotate to add an angle based on the previous rotation instead.
// The default rotation of a text Text object is 0.
// \param text  Text object
// \param angle New rotation, in degrees
// void sfText_setRotation(sfText* text, float angle);
func (self Text) Setrotation(angle float) void { 
    return C.sfText_setRotation(self.Cref, sfFloat(angle));
}
            
// Set the scale factors of a text
// This function completely overwrites the previous scale.
// See sfText_scale to add a factor based on the previous scale instead.
// The default scale of a text Text object is (1, 1).
// \param text  Text object
// \param scale New scale factors
// void sfText_setScale(sfText* text, sfVector2f scale);
func (self Text) Setscale(scale Vector2f) void { 
    return C.sfText_setScale(self.Cref, sfVector2f(scale));
}
            
// Set the local origin of a text
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a text object is (0, 0).
// \param text   Text object
// \param origin New origin
// void sfText_setOrigin(sfText* text, sfVector2f origin);
func (self Text) Setorigin(origin Vector2f) void { 
    return C.sfText_setOrigin(self.Cref, sfVector2f(origin));
}
            
// Get the position of a text
// \param text Text object
// \return Current position
// sfVector2f sfText_getPosition(const sfText* text);
func (self Text) Getposition() Vector2f { 
    return C.sfText_getPosition(self.Cref);
}
            
// Get the orientation of a text
// The rotation is always in the range [0, 360].
// \param text Text object
// \return Current rotation, in degrees
// float sfText_getRotation(const sfText* text);
func (self Text) Getrotation() float { 
    return C.sfText_getRotation(self.Cref);
}
            
// Get the current scale of a text
// \param text Text object
// \return Current scale factors
// sfVector2f sfText_getScale(const sfText* text);
func (self Text) Getscale() Vector2f { 
    return C.sfText_getScale(self.Cref);
}
            
// Get the local origin of a text
// \param text Text object
// \return Current origin
// sfVector2f sfText_getOrigin(const sfText* text);
func (self Text) Getorigin() Vector2f { 
    return C.sfText_getOrigin(self.Cref);
}
            
// Move a text by a given offset
// This function adds to the current position of the object,
// unlike sfText_setPosition which overwrites it.
// \param text   Text object
// \param offset Offset
// void sfText_move(sfText* text, sfVector2f offset);
func (self Text) Move(offset Vector2f) void { 
    return C.sfText_move(self.Cref, sfVector2f(offset));
}
            
// Rotate a text
// This function adds to the current rotation of the object,
// unlike sfText_setRotation which overwrites it.
// \param text  Text object
// \param angle Angle of rotation, in degrees
// void sfText_rotate(sfText* text, float angle);
func (self Text) Rotate(angle float) void { 
    return C.sfText_rotate(self.Cref, sfFloat(angle));
}
            
// Scale a text
// This function multiplies the current scale of the object,
// unlike sfText_setScale which overwrites it.
// \param text    Text object
// \param factors Scale factors
// void sfText_scale(sfText* text, sfVector2f factors);
func (self Text) Scale(factors Vector2f) void { 
    return C.sfText_scale(self.Cref, sfVector2f(factors));
}
            
// Get the combined transform of a text
// \param text Text object
// \return Transform combining the position/rotation/scale/origin of the object
// const sfTransform* sfText_getTransform(const sfText* text);
func (self *Transform) *Transform(Text_getTransform)  { 
    return C.sf*Transform(self.Cref, sfVector2f(factors));
}
            
// Get the inverse of the combined transform of a text
// \param text Text object
// \return Inverse of the combined transformations applied to the object
// const sfTransform* sfText_getInverseTransform(const sfText* text);
func (self *Transform) *Transform(Text_getInverseTransform)  { 
    return C.sf*Transform(self.Cref, sfVector2f(factors));
}
            
// Set the string of a text (from an ANSI string)
// A text's string is empty by default.
// \param text   Text object
// \param string New string
// void sfText_setString(sfText* text, const char* string);
func (self Text) Setstring(string *char ) void { 
    return C.sfText_setString(self.Cref, sf(*char));
}
            
// Set the string of a text (from a unicode string)
// \param text   Text object
// \param string New string
// void sfText_setUnicodeString(sfText* text, const sfUint32* string);
func (self Text) Setunicodestring(string *Uint32 ) void { 
    return C.sfText_setUnicodeString(self.Cref, sf(*Uint32));
}
            
// Set the font of a text
// The \a font argument refers to a texture that must
// exist as long as the text uses it. Indeed, the text
// doesn't store its own copy of the font, but rather keeps
// a pointer to the one that you passed to this function.
// If the font is destroyed and the text tries to
// use it, the behaviour is undefined.
// Texts have a valid font by default, which the built-in
// sfFont_getDefaultFont().
// \param text Text object
// \param font New font
// void sfText_setFont(sfText* text, const sfFont* font);
func (self Text) Setfont(font *Font ) void { 
    return C.sfText_setFont(self.Cref, sf(*Font));
}
            
// Set the character size of a text
// The default size is 30.
// \param text Text object
// \param size New character size, in pixels
// void sfText_setCharacterSize(sfText* text, unsigned int size);
func (self Text) Setcharactersize(size int ) void { 
    return C.sfText_setCharacterSize(self.Cref, sf(int));
}
            
// Set the style of a text
// You can pass a combination of one or more styles, for
// example sfTextBold | sfTextItalic.
// The default style is sfTextRegular.
// \param text  Text object
// \param style New style
// void sfText_setStyle(sfText* text, sfUint32 style);
func (self Text) Setstyle(style Uint32) void { 
    return C.sfText_setStyle(self.Cref, sfUint32(style));
}
            
// Set the global color of a text
// By default, the text's color is opaque white.
// \param text  Text object
// \param color New color of the text
// void sfText_setColor(sfText* text, sfColor color);
func (self Text) Setcolor(color Color) void { 
    return C.sfText_setColor(self.Cref, sfColor(color));
}
            
// Get the text of a text (returns an ANSI string)
// \param text Text object
// \return String an a locale-dependant ANSI string
// const char* sfText_getString(const sfText* text);
func (self *char) *char(Text_getString)  { 
    return C.sf*char(self.Cref, sfColor(color));
}
            
// Get the string of a text (returns a unicode string)
// \param text Text object
// \return String as UTF-32
// const sfUint32* sfText_getUnicodeString(const sfText* text);
func (self *Uint32) *Uint32(Text_getUnicodeString)  { 
    return C.sf*Uint32(self.Cref, sfColor(color));
}
            
// Get the font used by a text
// \param text Text object
// \return Pointer to the font
// const sfFont* sfText_getFont(const sfText* text);
func (self *Font) *Font(Text_getFont)  { 
    return C.sf*Font(self.Cref, sfColor(color));
}
            
// Get the size of the characters of a text
// \param text Text object
// \return Size of the characters
// unsigned int sfText_getCharacterSize(const sfText* text);
func (self int) int(Text_getCharacterSize)  { 
    return C.sfint(self.Cref, sfColor(color));
}
            
// Get the style of a text
// \param text Text object
// \return Current string style (see sfTextStyle enum)
// sfUint32 sfText_getStyle(const sfText* text);
func (self Text) Getstyle() Uint32 { 
    return C.sfText_getStyle(self.Cref);
}
            
// Get the global color of a text
// \param text Text object
// \return Global color of the text
// sfColor sfText_getColor(const sfText* text);
func (self Text) Getcolor() Color { 
    return C.sfText_getColor(self.Cref);
}
            
// Return the position of the \a index-th character in a text
// This function computes the visual position of a character
// from its index in the string. The returned position is
// in global coordinates (translation, rotation, scale and
// origin are applied).
// If \a index is out of range, the position of the end of
// the string is returned.
// \param text  Text object
// \param index Index of the character
// \return Position of the character
// sfVector2f sfText_findCharacterPos(const sfText* text, size_t index);
func (self Text) Findcharacterpos(index size_t) Vector2f { 
    return C.sfText_findCharacterPos(self.Cref, sfSize_t(index));
}
            
// Get the local bounding rectangle of a text
// The returned rectangle is in local coordinates, which means
// that it ignores the transformations (translation, rotation,
// scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// entity in the entity's coordinate system.
// \param text Text object
// \return Local bounding rectangle of the entity
// sfFloatRect sfText_getLocalBounds(const sfText* text);
func (self Text) Getlocalbounds() FloatRect { 
    return C.sfText_getLocalBounds(self.Cref);
}
            
// Get the global bounding rectangle of a text
// The returned rectangle is in global coordinates, which means
// that it takes in account the transformations (translation,
// rotation, scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// text in the global 2D world's coordinate system.
// \param text Text object
// \return Global bounding rectangle of the entity
// sfFloatRect sfText_getGlobalBounds(const sfText* text);
func (self Text) Getglobalbounds() FloatRect { 
    return C.sfText_getGlobalBounds(self.Cref);
}
            
*/