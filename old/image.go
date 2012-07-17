package gfx

//#include <SFML/Graphics/Color.h>
//#include <SFML/Graphics/Rect.h>
//#include <SFML/Graphics/Types.h>
import "C"

type Image struct {
	Cref *C.sfImage
}

// Create a new empty image
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_Create();
func ImageCreate() Image {
    return C.sfImage_Create(self.Cref)
}

// Create a new image filled with a color
// param Width : Image width
// param Height : Image height
// param Col : Image color
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromColor(unsigned int Width, unsigned int Height, sfColor Color);
func (self *Image) CreateFromColor(uint Height, sfColor Color) Image {
    return C.sfImage_CreateFromColor(self.Cref, uint, sfColor)
}

// Create a new image from an array of pixels in memory
// param Width : Image width
// param Height : Image height
// param Data : Pointer to the pixels in memory (assumed format is RGBA)
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromPixels(unsigned int Width, unsigned int Height, const sfUint8* Data);
func (self *Image) CreateFromPixels(uint Height, sfUint8* Data) Image {
    return C.sfImage_CreateFromPixels(self.Cref, uint, sfUint8*)
}

// Create a new image from a file
// param Filename : Path of the image file to load
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromFile(const char* Filename);
func (self *Image) CreateFromFile() Image {
    return C.sfImage_CreateFromFile(self.Cref)
}

// Create a new image from a file in memory
// param Data : Pointer to the file data in memory
// param SizeInBytes : Size of the data to load, in bytes
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromMemory(const char* Data, size_t SizeInBytes);
func (self *Image) CreateFromMemory(size_t SizeInBytes) Image {
    return C.sfImage_CreateFromMemory(self.Cref, size_t)
}

// Destroy an existing image
// param Image : Image to delete
// void sfImage_Destroy(sfImage* Image);
func (self *Image) Destroy() void {
    return C.sfImage_Destroy(self.Cref)
}

// Save the content of an image to a file
// param Image : Image to save
// param Filename : Path of the file to save (overwritten if already exist)
// return sfTrue if saving was successful
// sfBool sfImage_SaveToFile(sfImage* Image, const char* Filename);
func (self *Image) SaveToFile(char* Filename) sfBool {
    return C.sfImage_SaveToFile(self.Cref, char*)
}

// Create a transparency mask for an image from a specified colorkey
// param Image : Image to modify
// param ColorKey : Color to become transparent
// param Alpha : Alpha value to use for transparent pixels
// void sfImage_CreateMaskFromColor(sfImage* Image, sfColor ColorKey, sfUint8 Alpha);
func (self *Image) CreateMaskFromColor(sfColor ColorKey, sfUint8 Alpha) void {
    return C.sfImage_CreateMaskFromColor(self.Cref, sfColor, sfUint8)
}

// Copy pixels from another image onto this one.
// This function does a slow pixel copy and should only
// be used at initialization time
// param Image : Destination image
// param Source : Source image to copy
// param DestX : X coordinate of the destination position
// param DestY : Y coordinate of the destination position
// param SourceRect : Sub-rectangle of the source image to copy
// void sfImage_Copy(sfImage* Image, sfImage* Source, unsigned int DestX, unsigned int DestY, sfIntRect SourceRect);
func (self *Image) Copy(sfImage* Source, uint DestX, uint DestY, sfIntRect SourceRect) void {
    return C.sfImage_Copy(self.Cref, sfImage*, uint, uint, sfIntRect)
}

// Create the image from the current contents of the
// given window
// param Image : Destination image
// param Window : Window to capture
// param SourceRect : Sub-rectangle of the screen to copy (empty by default - entire image)
// return True if creation was successful
// sfBool sfImage_CopyScreen(sfImage* Image, sfRenderWindow* Window, sfIntRect SourceRect);
func (self *Image) CopyScreen(sfRenderWindow* Window, sfIntRect SourceRect) sfBool {
    return C.sfImage_CopyScreen(self.Cref, sfRenderWindow*, sfIntRect)
}

// Change the color of a pixel of an image
// Don't forget to call Update when you end modifying pixels
// param Image : Image to modify
// param X : X coordinate of pixel in the image
// param Y : Y coordinate of pixel in the image
// param Col : New color for pixel (X, Y)
// void sfImage_SetPixel(sfImage* Image, unsigned int X, unsigned int Y, sfColor Color);
func (self *Image) SetPixel(uint X, uint Y, sfColor Color) void {
    return C.sfImage_SetPixel(self.Cref, uint, uint, sfColor)
}

// Get a pixel from an image
// param Image : Image to read
// param X : X coordinate of pixel in the image
// param Y : Y coordinate of pixel in the image
// return Color of pixel (x, y)
// sfColor sfImage_GetPixel(sfImage* Image, unsigned int X, unsigned int Y);
func (self *Image) GetPixel(uint X, uint Y) sfColor {
    return C.sfImage_GetPixel(self.Cref, uint, uint)
}

// Get a read-only pointer to the array of pixels of an image (8 bits integers RGBA)
// Array size is sfImage_GetWidth() x sfImage_GetHeight() x 4
// This pointer becomes invalid if you reload or resize the image
// param Image : Image to read
// return Pointer to the array of pixels
// const sfUint8* sfImage_GetPixelsPtr(sfImage* Image);
// Bind the image for rendering
// param Image : Image to bind
// void sfImage_Bind(sfImage* Image);
func (self *Image) Bind() void {
    return C.sfImage_Bind(self.Cref)
}

// Enable or disable image smooth filter
// param Image : Image to modify
// param Smooth : sfTrue to enable smoothing filter, false to disable it
// void sfImage_SetSmooth(sfImage* Image, sfBool Smooth);
func (self *Image) SetSmooth(sfBool Smooth) void {
    return C.sfImage_SetSmooth(self.Cref, sfBool)
}

// Return the width of the image
// param Image : Image to read
// return Width in pixels
// unsigned int sfImage_GetWidth(sfImage* Image);
// Return the height of the image
// param Image : Image to read
// return Height in pixels
// unsigned int sfImage_GetHeight(sfImage* Image);
// Tells whether the smoothing filter is enabled or not on an image
// param Image : Image to read
// return sfTrue if the smoothing filter is enabled
// sfBool sfImage_IsSmooth(sfImage* Image);
func (self *Image) IsSmooth() sfBool {
    return C.sfImage_IsSmooth(self.Cref)
}

#endif // SFML_IMAGE_H
