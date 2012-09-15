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
// #include <SFML/Graphics/Font.h>
// #include <SFML/Graphics/Glyph.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/System/InputStream.h>
// #include <stddef.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"unsafe"
)

type Font struct {
	Cref *C.sfFont
}

type Glyph struct {
	Cref *C.sfGlyph
	// int       advance;     ///< Offset to move horizontically to the next character
	// sfIntRect bounds;      ///< Bounding rectangle of the glyph, in coordinates relative to the baseline
	// sfIntRect textureRect; ///< Texture coordinates of the glyph inside the font's image
}

// \brief Create a new font from a file
// \param filename Path of the font file to load
// \return A new sfFont object, or NULL if it failed
// sfFont* sfFont_createFromFile(const char* filename);
func FontFromFile(fname string) (Font, error) {
	s := C.CString(fname)
	defer C.free(unsafe.Pointer(s))
	font := C.sfFont_createFromFile(s)
	if font == nil {
		return Font{nil}, errors.New("Couldn't create font from file: " + fname)
	}
	return Font{font}, nil
}

// \brief Copy an existing font
// \param font Font to copy
// \return Copied object
// sfFont* sfFont_copy(sfFont* font);
func (self Font) Copy() Font {
	return Font{C.sfFont_copy(self.Cref)}
}

// \brief Destroy an existing font
// \param font Font to delete
// void sfFont_destroy(sfFont* font);
func (self Font) Destroy() {
	C.sfFont_destroy(self.Cref)
}

// \brief Get a glyph in a font
// \param font          Source font
// \param codePoint     Unicode code point of the character to get
// \param characterSize Character size, in pixels
// \param bold          Retrieve the bold version or the regular one?
// \return The corresponding glyph
// sfGlyph sfFont_getGlyph(sfFont* font, sfUint32 codePoint, unsigned int characterSize, sfBool bold);
func (self Font) Glyph(codePoint, characterSize uint32, bold bool) Glyph {
	g := C.sfFont_getGlyph(
		self.Cref,
		C.sfUint32(codePoint),
		C.uint(characterSize),
		Bool(bold),
	)
	return Glyph{&g}
}

// \brief Get the kerning value corresponding to a given pair of characters in a font
// \param font          Source font
// \param first         Unicode code point of the first character
// \param second        Unicode code point of the second character
// \param characterSize Character size, in pixels
// \return Kerning offset, in pixels
// int sfFont_getKerning(sfFont* font, sfUint32 first, sfUint32 second, unsigned int characterSize);
func (self Font) Kerning(first, second uint32, characterSize uint) int {
	k := C.sfFont_getKerning(self.Cref, C.sfUint32(first),
		C.sfUint32(second), C.uint(characterSize))
	return int(k)
}

// \brief Get the line spacing value
// \param font          Source font
// \param codePoint     Unicode code point of the character to get
// \param characterSize Character size, in pixels
// \return Line spacing, in pixels
//int sfFont_getLineSpacing(sfFont* font, unsigned int characterSize);
func (self Font) LineSpacing(characterSize int) int {
	return int(C.sfFont_getLineSpacing(self.Cref, C.uint(characterSize)))
}

// \brief Get the texture containing the glyphs of a given size in a font
// \param font          Source font
// \param characterSize Character size, in pixels
// \return Read-only pointer to the texture
// const sfTexture* sfFont_getTexture(sfFont* font, unsigned int characterSize);
func (self Font) Texture(characterSize uint) Texture {
	return Texture{C.sfFont_getTexture(self.Cref, C.uint(characterSize))}
}

// \brief Get the built-in default font (Arial)
// \return Pointer to the default font
// const sfFont* sfFont_getDefaultFont(void);
// func (self Font) DefaultFont() Font {
// 	return Font{C.sfFont_getDefaultFont()}
// }

// TODO:LATER            
// \brief Create a new image font a file in memory
// \param data        Pointer to the file data in memory
// \param sizeInBytes Size of the data to load, in bytes
// \return A new sfFont object, or NULL if it failed
// sfFont* sfFont_createFromMemory(const void* data, size_t sizeInBytes);
// func (self Font) CreateFromMemory(sizeInBytes size_t) Font { 
//     return C.sfFont_createFromMemory(self.Cref, sfSize_t(sizeInBytes));
// }

// \brief Create a new image font a custom stream
// \param stream Source stream to read from
// \return A new sfFont object, or NULL if it failed
// sfFont* sfFont_createFromStream(sfInputStream* stream);
// func (self Font) Createfromstream() *Font { 
//     return C.sfFont_createFromStream(self.Cref);
// }
