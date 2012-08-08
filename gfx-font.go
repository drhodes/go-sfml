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

import "unsafe"

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
func (self Font) FontFromFile(fname string) Font {
	s := C.CString(fname)
	defer C.free(unsafe.Pointer(s))
	return Font{C.sfFont_createFromFile(s)}
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
func (self Font) GetGlyph(codePoint, characterSize uint32, bold bool) Glyph {
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
func (self Font) GetKerning(first, second uint32, characterSize uint) int {
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
func (self Font) GetLineSpacing(characterSize int) int {
	return int(C.sfFont_getLineSpacing(self.Cref, C.uint(characterSize)))
}

// \brief Get the texture containing the glyphs of a given size in a font
// \param font          Source font
// \param characterSize Character size, in pixels
// \return Read-only pointer to the texture
// const sfTexture* sfFont_getTexture(sfFont* font, unsigned int characterSize);
func (self Font) GetTexture(characterSize uint) Texture {
	return Texture{C.sfFont_getTexture(self.Cref, C.uint(characterSize))}
}

// \brief Get the built-in default font (Arial)
// \return Pointer to the default font
// const sfFont* sfFont_getDefaultFont(void);
func (self Font) GetDefaultFont() Font {
	return Font{C.sfFont_getDefaultFont()}
}

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
