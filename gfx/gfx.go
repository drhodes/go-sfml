package gfx

// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Font.h>
// #include <SFML/Graphics/Image.h>
import "C"

import "runtime"


type Color struct { 
	Cref C.sfColor 
}

type Font struct {
	Cref *C.sfFont
}

type Image struct {	
	Cref *C.sfImage 
}

type IntRect struct{
	Cref *C.sfIntRect
}

type RenderWindow struct{
	Cref *C.sfRenderWindow
}


// _Color_
// -------------------------------------------------------------------------------
func ColorFromRGB(r, g, b int) Color {
	return Color{ C.sfColor_FromRGB(C.sfUint8(r), C.sfUint8(g), C.sfUint8(b) ) }
}

func ColorFromRGBA(r, g, b, a uint8) Color {
	return Color{ C.sfColor_FromRGBA(C.sfUint8(r), C.sfUint8(g), C.sfUint8(b), C.sfUint8(a) ) }
}

// Add two colors
// param other : Color
// return Component-wise saturated addition of the two colors
func (self *Color) Add(other Color) Color {
	return Color{ C.sfColor_Add(self.Cref, other.Cref) }
}

func (self *Color) AddMutate(other Color) {
	self.Cref = C.sfColor_Add(self.Cref, other.Cref) 
}

 
// Modulate two colors
// param other : Color
// return Component-wise multiplication of the two colors
func (self *Color) Modulate(other Color) Color {
	return Color { C.sfColor_Modulate(self.Cref, other.Cref) }
}

func (self *Color) ModulateMutate(other Color) {
	self.Cref = C.sfColor_Modulate(self.Cref, other.Cref) 
}



// _Font_
// -------------------------------------------------------------------------------
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




// _Image_
// -------------------------------------------------------------------------------
// Create a new empty image
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_Create();
func ImageCreate() Image {
    img := Image{ C.sfImage_Create() }
	runtime.SetFinalizer(&img, (*Image).Destroy)   
	return img
}

// Create a new image filled with a color
// param Width : Image width
// param Height : Image height
// param Col : Image color
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromColor(unsigned int Width, unsigned int Height, sfColor Color);
func ImageCreateFromColor(width, height uint, clr Color) Image {
    img := Image{ C.sfImage_CreateFromColor(C.uint(width), C.uint(height), clr.Cref) }
	runtime.SetFinalizer(&img, (*Image).Destroy)   
	return img
}

// Create a new image from an array of pixels in memory
// param Width : Image width
// param Height : Image height
// param Data : Pointer to the pixels in memory (assumed format is RGBA)
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromPixels(unsigned int Width, unsigned int Height, const sfUint8* Data);
func ImageCreateFromPixels(width uint, height uint, data []uint8) Image {
	cdata := C.sfUint8(data[0])
    img := Image{ C.sfImage_CreateFromPixels(C.uint(width), C.uint(height), &cdata) }
	runtime.SetFinalizer(&img, (*Image).Destroy)   
	return img
}

// Create a new image from a file
// param Filename : Path of the image file to load
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromFile(const char* Filename);
func ImageCreateFromFile(filename string) Image {
    img := Image{ C.sfImage_CreateFromFile(C.CString(filename)) }
	runtime.SetFinalizer(&img, (*Image).Destroy)   
	return img
}

// Create a new image from a file in memory
// param Data : Pointer to the file data in memory
// param SizeInBytes : Size of the data to load, in bytes
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromMemory(const char* Data, size_t SizeInBytes);
func ImageCreateFromMemory(data []uint8, sizeInByes int) Image {
	cdata := C.char(data[0])
    img := Image{ C.sfImage_CreateFromMemory(&cdata, C.size_t(sizeInByes)) }
	runtime.SetFinalizer(&img, (*Image).Destroy)   
	return img

}

// Destroy an existing image
// param Image : Image to delete
// sfImage_Destroy(sfImage* Image);
func (self *Image) Destroy() {
    C.sfImage_Destroy(self.Cref)
	self.Cref = nil
}

// this should go in a utility package.
func SfBool2GoBool(cVal C.sfBool) bool {
	return int(cVal) == 1
}

func GoBool2SfBool(goVal bool) C.sfBool {
	if goVal {
		return C.sfBool(1)
	}
	return C.sfBool(0)
}

// Save the content of an image to a file
// param Image : Image to save
// param Filename : Path of the file to save (overwritten if already exist)
// return sfTrue if saving was successful
// sfBool sfImage_SaveToFile(sfImage* Image, const char* Filename);
func (self *Image) SaveToFile(filename string) bool {
	return SfBool2GoBool(C.sfImage_SaveToFile(self.Cref, C.CString(filename)))
}

// Create a transparency mask for an image from a specified colorkey
// param Image : Image to modify
// param ColorKey : Color to become transparent
// param Alpha : Alpha value to use for transparent pixels
//  sfImage_CreateMaskFromColor(sfImage* Image, sfColor ColorKey, sfUint8 Alpha);
func (self *Image) CreateMaskFromColor(colorKey Color, alpha uint8)  {
    C.sfImage_CreateMaskFromColor(self.Cref, colorKey.Cref, C.sfUint8(alpha))
}

// Copy pixels from another image onto this one.
// This function does a slow pixel copy and should only
// be used at initialization time
// param Image : Destination image
// param Source : Source image to copy
// param DestX : X coordinate of the destination position
// param DestY : Y coordinate of the destination position
// param SourceRect : Sub-rectangle of the source image to copy
//  sfImage_Copy(sfImage* Image, sfImage* Source, unsigned int DestX, unsigned int DestY, sfIntRect SourceRect);
func (self *Image) Copy(source Image, destX, destY uint, sourceRect IntRect) {
    C.sfImage_Copy(self.Cref, source.Cref, C.uint(destX), C.uint(destY), *sourceRect.Cref)
}

// Create the image from the current contents of the given window
// param Image : Destination image
// param Window : Window to capture
// param SourceRect : Sub-rectangle of the screen to copy (empty by default - entire image)
// return True if creation was successful
// sfBool sfImage_CopyScreen(sfImage* Image, sfRenderWindow* Window, sfIntRect SourceRect);
func (self *Image) CopyScreen(window RenderWindow, sourceRect IntRect) bool {
    return SfBool2GoBool( C.sfImage_CopyScreen(self.Cref, window.Cref, *sourceRect.Cref) )
}

// Change the color of a pixel of an image
// Don't forget to call Update when you end modifying pixels
// param Image : Image to modify
// param X : X coordinate of pixel in the image
// param Y : Y coordinate of pixel in the image
// param Col : New color for pixel (X, Y)
// sfImage_SetPixel(sfImage* Image, unsigned int X, unsigned int Y, sfColor Color);
func (self *Image) SetPixel(x, y uint, color Color) {
    C.sfImage_SetPixel(self.Cref, C.uint(x), C.uint(y), color.Cref)
}

// Get a pixel from an image
// param Image : Image to read
// param X : X coordinate of pixel in the image
// param Y : Y coordinate of pixel in the image
// return Color of pixel (x, y)
// sfColor sfImage_GetPixel(sfImage* Image, unsigned int X, unsigned int Y);
func (self *Image) GetPixel(x, y uint) Color {
    return Color{ C.sfImage_GetPixel(self.Cref, C.uint(x), C.uint(y)) }
}

// Get a read-only pointer to the array of pixels of an image (8 bit integers RGBA)
// Array size is sfImage_GetWidth() x sfImage_GetHeight() x 4
// This pointer becomes invalid if you reload or resize the image
// param Image : Image to read
// return Pointer to the array of pixels
// const sfUint8* sfImage_GetPixelsPtr(sfImage* Image);

/* ARRGgghh
//func NewArray(typ interface{}, n int) Pointer

func (self *Image) GetPixelsPtr() []byte {	
	__COLOR__ := ColorFromRGBA(0,0,0,0)
	cPtr := C.sfImage_GetPixelsPtr(self.Cref)
	arr := unsafe.NewArray(__COLOR__, self.Height * self.Width)
	return arr
}
*/
	
// Bind the image for rendering
// param Image : Image to bind
// sfImage_Bind(sfImage* Image);
func (self *Image) Bind() {
    C.sfImage_Bind(self.Cref)
}

// Enable or disable image smooth filter
// param Image : Image to modify
// param Smooth : sfTrue to enable smoothing filter, false to disable it
// sfImage_SetSmooth(sfImage* Image, sfBool Smooth);
func (self *Image) SetSmooth(smooth bool) {
    C.sfImage_SetSmooth(self.Cref, GoBool2SfBool(smooth))
}

// Return the width of the image
// param Image : Image to read
// return Width in pixels
// unsigned int sfImage_GetWidth(sfImage* Image);
func (self *Image) GetWidth() uint {
	return uint(C.sfImage_GetWidth(self.Cref))
}

// Return the height of the image
// param Image : Image to read
// return Height in pixels
// unsigned int sfImage_GetHeight(sfImage* Image);
func (self *Image) GetHeight() uint {
	return uint(C.sfImage_GetHeight(self.Cref))
}

// Tells whether the smoothing filter is enabled or not on an image
// param Image : Image to read
// return sfTrue if the smoothing filter is enabled
// sfBool sfImage_IsSmooth(sfImage* Image);
func (self *Image) IsSmooth() bool {
    return SfBool2GoBool( C.sfImage_IsSmooth(self.Cref) )
}
