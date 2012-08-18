package sfml

// #cgo LDFLAGS:-lcsfml-graphics -ljpeg
// #include <SFML/Graphics/Text.h>
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/System/Vector2.h>
// #include <stddef.h>
// #include <stdio.h>
// extern int _UnicodeStringSize(sfUint32 *s) {
//   int i = 0;
//   for (i = 0; s[i] != 0; i++)
//     ;
//   return i;
// }
import "C"

import (
	"errors"
	"unsafe"
	"unicode/utf8"
)

type Text struct {
	Cref *C.sfText
}

// sfText styles
// Create a new text
// \return A new sfText object, or NULL if it failed
// sfText* sfText_create(void);
func NewText() (Text, error) {
	txt := C.sfText_create()
	if txt == nil {
		return Text{nil}, errors.New("Couldn't make a text")
	}
	return Text{txt}, nil
}

// Copy an existing text
// \param text Text to copy
// \return Copied object
// sfText* sfText_copy(sfText* text);
func (self Text) Copy() Text {
	return Text{C.sfText_copy(self.Cref)}
}

// Destroy an existing text
// \param text Text to delete
// void sfText_destroy(sfText* text);
func (self Text) Destroy() {
	C.sfText_destroy(self.Cref)
}

// Set the position of a text
// This function completely overwrites the previous position.
// See sfText_move to apply an offset based on the previous position instead.
// The default position of a text Text object is (0, 0).
// \param text     Text object
// \param position New position
// void sfText_setPosition(sfText* text, sfVector2f position);
func (self Text) SetPosition(x, y float32) {
	v := C.sfVector2f{C.float(x), C.float(y)}
	C.sfText_setPosition(self.Cref, v)
}

// Set the orientation of a text
// This function completely overwrites the previous rotation.
// See sfText_rotate to add an angle based on the previous rotation instead.
// The default rotation of a text Text object is 0.
// \param text  Text object
// \param angle New rotation, in degrees
// void sfText_setRotation(sfText* text, float angle);
func (self Text) SetRotation(angle float32) {
	C.sfText_setRotation(self.Cref, C.float(angle))
}

// Set the scale factors of a text
// This function completely overwrites the previous scale.
// See sfText_scale to add a factor based on the previous scale instead.
// The default scale of a text Text object is (1, 1).
// \param text  Text object
// \param scale New scale factors
// void sfText_setScale(sfText* text, sfVector2f scale);
func (self Text) SetScale(x, y float32) {
	v := C.sfVector2f{C.float(x), C.float(y)}
	C.sfText_setScale(self.Cref, v)
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
func (self Text) SetOrigin(x, y float32) {
	v := C.sfVector2f{C.float(x), C.float(y)}
	C.sfText_setOrigin(self.Cref, v)
}

// Get the position of a text
// \param text Text object
// \return Current position
// sfVector2f sfText_getPosition(const sfText* text);
func (self Text) Position(x, y float32) (float32, float32) {
	v := C.sfText_getPosition(self.Cref)
	return float32(v.x), float32(v.y)
}

// Get the orientation of a text
// The rotation is always in the range [0, 360].
// \param text Text object
// \return Current rotation, in degrees
// float sfText_getRotation(const sfText* text);
func (self Text) Rotation() float32 {
	return float32(C.sfText_getRotation(self.Cref))
}

// Get the current scale of a text
// \param text Text object
// \return Current scale factors
// sfVector2f sfText_getScale(const sfText* text);
func (self Text) GetScale() (x, y float32) {
	v := C.sfText_getScale(self.Cref)
	return float32(v.x), float32(v.y)
}

// Get the local origin of a text
// \param text Text object
// \return Current origin
// sfVector2f sfText_getOrigin(const sfText* text);
func (self Text) Origin() (x, y float32) {
	v := C.sfText_getOrigin(self.Cref)
	return float32(v.x), float32(v.y)
}

// Move a text by a given offset
// This function adds to the current position of the object,
// unlike sfText_setPosition which overwrites it.
// \param text   Text object
// \param offset Offset
// void sfText_move(sfText* text, sfVector2f offset);
func (self Text) Move(x, y float32) {
	v := C.sfVector2f{C.float(x), C.float(y)}
	C.sfText_move(self.Cref, v)
}

// Rotate a text
// This function adds to the current rotation of the object,
// unlike sfText_setRotation which overwrites it.
// \param text  Text object
// \param angle Angle of rotation, in degrees
// void sfText_rotate(sfText* text, float angle);
func (self Text) Rotate(angle float32) {
	C.sfText_rotate(self.Cref, C.float(angle))
}

// Scale a text
// This function multiplies the current scale of the object,
// unlike sfText_setScale which overwrites it.
// \param text    Text object
// \param factors Scale factors
// void sfText_scale(sfText* text, sfVector2f factors);
func (self Text) Scale(x, y float32) {
	v := C.sfVector2f{C.float(x), C.float(y)}
	C.sfText_scale(self.Cref, v)
}

// Get the combined transform of a text
// \param text Text object
// \return Transform combining the position/rotation/scale/origin of the object
// const sfTransform* sfText_getTransform(const sfText* text);
func (self Text) Transform() Transform {
	t := C.sfText_getTransform(self.Cref)
	return Transform{&t}
}

// Get the inverse of the combined transform of a text
// \param text Text object
// \return Inverse of the combined transformations applied to the object
// const sfTransform* sfText_getInverseTransform(const sfText* text);
func (self Text) InverseTransform() Transform {
	t :=  C.sfText_getInverseTransform(self.Cref)
	return Transform{&t}
}

// Set the string of a text (from an ANSI string)
// A text's string is empty by default.
// \param text   Text object
// \param string New string
// void sfText_setString(sfText* text, const char* string);
func (self Text) SetString(s string) {
	C.sfText_setString(self.Cref, C.CString(s))
}

// Set the string of a text (from a unicode string)
// \param text   Text object
// \param string New string
// void sfText_setUnicodeString(sfText* text, const sfUint32* string);
func (self Text) SetUnicodeString(s string) {
	/* Go string are encoded in UTF-8, whereas SFML unicode
	strings are encoded in UTF-32, so we have to make the
	conversion */
	bs := []byte(s)
	n := utf8.RuneCount(bs)
	runes := make([]rune, n+1)
	var size int
	i := 0
	j := 0
	for ; i < n; i += 1 {
		runes[i], size = utf8.DecodeRune(bs[j:])
		j += size
	}
	runes[i] = 0 /* don't forget the null terminator */
	C.sfText_setUnicodeString(self.Cref, (*C.sfUint32)(unsafe.Pointer(&runes[0])))
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
func (self Text) SetFont(font Font) {
	C.sfText_setFont(self.Cref, font.Cref)
}

// Set the character size of a text
// The default size is 30.
// \param text Text object
// \param size New character size, in pixels
// void sfText_setCharacterSize(sfText* text, unsigned int size);
func (self Text) SetCharacterSize(size int) {
	C.sfText_setCharacterSize(self.Cref, C.uint(size))
}

// Set the style of a text
// You can pass a combination of one or more styles, for
// example sfTextBold | sfTextItalic.
// The default style is sfTextRegular.
// \param text  Text object
// \param style New style
// void sfText_setStyle(sfText* text, sfUint32 style);
func (self Text) SetStyle(style uint32) {
	C.sfText_setStyle(self.Cref, C.sfUint32(style))
}

// Set the global color of a text
// By default, the text's color is opaque white.
// \param text  Text object
// \param color New color of the text
// void sfText_setColor(sfText* text, sfColor color);
func (self Text) SetColor(color Color) {
	C.sfText_setColor(self.Cref, color.Cref)
}

// Get the text of a text (returns an ANSI string)
// \param text Text object
// \return String an a locale-dependant ANSI string
// const char* sfText_getString(const sfText* text);
func (self Text) String() string {
	return C.GoString(C.sfText_getString(self.Cref))
}

// Get the string of a text (returns a unicode string)
// \param text Text object
// \return String as UTF-32
// const sfUint32* sfText_getUnicodeString(const sfText* text);
func (self Text) UnicodeString() string {
	/* We have to make the inverse conversion from UTF-32 to UTF-8 */
	s := C.sfText_getUnicodeString(self.Cref)
	length := C._UnicodeStringSize(s)
	b := C.GoBytes(unsafe.Pointer(s), length*4)
	buf := make([]byte, length*4)
	for i := 0; i < len(b); i += 4 {
		n := uint(b[i]) + uint(b[i+1]) << 8 + uint(b[i+2]) << 16 + uint(b[i+3]) << 24
		utf8.EncodeRune(buf[i:], rune(n))
	}
		
	return string(buf)
}

// Get the font used by a text
// \param text Text object
// \return Pointer to the font
// const sfFont* sfText_getFont(const sfText* text);
func (self Text) Font() Font {
	return Font{C.sfText_getFont(self.Cref)}
}

// Get the size of the characters of a text
// \param text Text object
// \return Size of the characters
// unsigned int sfText_getCharacterSize(const sfText* text);
func (self Text) CharacterSize() uint {
	return uint(C.sfText_getCharacterSize(self.Cref))
}

// Get the style of a text
// \param text Text object
// \return Current string style (see sfTextStyle enum)
// sfUint32 sfText_getStyle(const sfText* text);
func (self Text) Style() uint {
	return uint(C.sfText_getStyle(self.Cref))
}

// Get the global color of a text
// \param text Text object
// \return Global color of the text
// sfColor sfText_getColor(const sfText* text);
func (self Text) Color() Color {
	return Color{C.sfText_getColor(self.Cref)}
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
func (self Text) FindCharacterPos(index uint) Vector2f {
	return Vector2f{C.sfText_findCharacterPos(self.Cref, C.size_t(index))}
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
func (self Text) LocalBounds() FloatRect {
	r := C.sfText_getLocalBounds(self.Cref)
	return FloatRect{&r}
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
func (self Text) GlobalBounds() FloatRect {
	r := C.sfText_getGlobalBounds(self.Cref)
	return FloatRect{&r}
}
