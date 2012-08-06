package sfml

//#include <SFML/Graphics/Font.h>
import "C"
import "runtime"

type Font struct {
	Cref *C.sfFont
}

// Create a new empty font
// return A new sfFont object, or NULL if it failed
func FontCreate() Font {
    fnt := Font{ C.sfFont_Create() }
	runtime.SetFinalizer(&fnt, (*Font).Destroy)   
	return fnt
}

// Create a new font from a file
// param Filename : Path of the font file to load
// param CharSize : Size of characters in bitmap - the bigger, the higher quality
// param Charset : Characters set to generate (just pass NULL to get the default charset)
// return A new sfFont object, or NULL if it failed
func FontCreateFromFile(path string, charSize int) Font {
    fnt := Font{ C.sfFont_CreateFromFile(C.CString(path), C.uint(charSize), nil) }
	runtime.SetFinalizer(&fnt, (*Font).Destroy)   
	return fnt	
}

// Create a new image font a file in memory
// param Data : Pointer to the file data in memory
// param SizeInBytes : Size of the data to load, in bytes
// param CharSize : Size of characters in bitmap - the bigger, the higher quality
// param Charset : Characters set to generate (just pass NULL to get the default charset)
// return A new sfFont object, or NULL if it failed
/* -untested
func FontCreateFromMemory(data uint8, SizeInBytes uint8, CharSize uint) Font {
	cdata := C.char(data)
	//C.free(unsafe.Pointer(cfile))
	
    return Font{ C.sfFont_CreateFromMemory( &cdata, 
			C.size_t(SizeInBytes), 
			C.uint(CharSize), nil )}
}
*/

// Destroy an existing font
// param Font : Font to delete
func (self *Font) Destroy() {	
	C.sfFont_Destroy(self.Cref);
	self.Cref = nil;
}

// Get the base size of characters in a font;
// All glyphs dimensions are based on this value
// param Font : Font object
// return Base size of characters
// Get the built-in default font (Arial)
// return Pointer to the default font
func GetDefaultFont() Font {
	return Font{ C.sfFont_GetDefaultFont() };
}

