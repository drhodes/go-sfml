package gfx

// #cgo LDFLAGS:-lcsfml-graphics
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Image.h>
// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/System/InputStream.h>
// #include <SFML/System/Vector2.h>
// #include <stddef.h>
import "C"

import (
	"unsafe"
	"errors"
	)

type Image struct {
	Cref *C.sfImage
}

// Create an image
// This image is filled with black pixels.
// width  Width of the image
// height Height of the image
// return A new Image object
// sfImage* sfImage_create(unsigned int width, unsigned int height);
func NewImage(width, height uint) Image { 
    return Image{C.sfImage_create(C.uint(width), C.uint(height))}
}
          
// Create an image and fill it with a unique color
// width  Width of the image
// height Height of the image
// color  Fill color
// return A new sfImage object
// sfImage* sfImage_createFromColor(unsigned int width, unsigned int height, sfColor color);
func ImageFromColor(width, height uint, color Color) Image {
    return Image{C.sfImage_createFromColor(C.uint(width), C.uint(height), color.Cref)};
}

            
// Create an image from an array of pixels
// The \a pixel array is assumed to contain 32-bits RGBA pixels,
// and have the given \a width and \a height. If not, this is
// an undefined behaviour.
// If \a pixels is null, an empty image is created.
// \param width  Width of the image
// \param height Height of the image
// \param pixels Array of pixels to copy to the image
// \return A new sfImage object
// sfImage* sfImage_createFromPixels(unsigned int width, unsigned int height, const sfUint8* pixels);
func (self Image) ImageFromPixels(width, height uint, pixels []uint8) Image { 
	ptr := unsafe.Pointer(&pixels[0])
	p := (*C.sfUint8)(ptr)
    cimg := C.sfImage_createFromPixels(C.uint(width), C.uint(height), p)
	return Image{cimg}
}
            
// Create an image from a file on disk
// The supported image formats are bmp, png, tga, jpg, gif,
// psd, hdr and pic. Some format options are not supported,
// like progressive jpeg.
// If this function fails, the image is left unchanged.
// \param filename Path of the image file to load
// \return A new sfImage object, or NULL if it failed
// sfImage* sfImage_createFromFile(const char* filename);
func ImageFromFile(fname string) (Image, error) { 
    cimg := C.sfImage_createFromFile(C.CString(fname))
	if cimg == nil {
		return Image{nil}, errors.New("Error loading image")
	}
	return Image{cimg}, nil
}
            
            
//  Copy an existing image
// \param image Image to copy
// \return Copied object
// sfImage* sfImage_copy(sfImage* image);

// func (self Image) Copy() *Image { 
//     return C.sfImage_copy();
// }
            
//  Destroy an existing image
// \param image Image to delete
// void sfImage_destroy(sfImage* image);

// func (self Image) Destroy() void { 
//     return C.sfImage_destroy();
// }
            
//  Save an image to a file on disk
// The format of the image is automatically deduced from
// the extension. The supported image formats are bmp, png,
// tga and jpg. The destination file is overwritten
// if it already exists. This function fails if the image is empty.
// \param image    Image object
// \param filename Path of the file to save
// \return sfTrue if saving was successful
// sfBool sfImage_saveToFile(const sfImage* image, const char* filename);

// func (self Image) Savetofile(filename *char ) Bool { 
//     return C.sfImage_saveToFile();
// }

// Return the size of an image
// \param image Image object
// \return Size in pixels
// sfVector2u sfImage_getSize(const sfImage* image);
func (self Image) Getsize() C.sfVector2u { 
    return C.sfImage_getSize(self.Cref);
}

// TODO: Not working
//  Create a transparency mask from a specified color-key
// This function sets the alpha value of every pixel matching
// the given color to \a alpha (0 by default), so that they
// become transparent.
// \param image Image object
// \param color Color to make transparent
// \param alpha Alpha value to assign to transparent pixels
// void sfImage_createMaskFromColor(sfColor color, sfUint8 alpha);
// func (self Image) Createmaskfromcolor(col Color, alpha uint8) { 
//     return C.sfImage_createMaskFromColor(col.Cref, C.sfUint8(alpha));
// }
            
//  Copy pixels from an image onto another
// This function does a slow pixel copy and should not be
// used intensively. It can be used to prepare a complex
// static image from several others, but if you need this
// kind of feature in real-time you'd better use sfRenderTexture.
// If \a sourceRect is empty, the whole image is copied.
// If \a applyAlpha is set to true, the transparency of
// source pixels is applied. If it is false, the pixels are
// copied unchanged with their alpha value.
// \param image      Image object
// \param source     Source image to copy
// \param destX      X coordinate of the destination position
// \param destY      Y coordinate of the destination position
// \param sourceRect Sub-rectangle of the source image to copy
// \param applyAlpha Should the copy take in account the source transparency?
// void sfImage_copyImage(sfImage* image, const sfImage* source, unsigned int destX, unsigned int destY, sfIntRect sourceRect, sfBool applyAlpha);

// func (self Image) Copyimage(source *Image, destX, destY int, 
// 	sourceRect IntRect, applyAlpha bool) { 
//     return C.sfImage_copyImage();
// }

/*            
            
//  Change the color of a pixel in an image
// This function doesn't check the validity of the pixel
// coordinates, using out-of-range values will result in
// an undefined behaviour.
// \param image Image object
// \param x     X coordinate of pixel to change
// \param y     Y coordinate of pixel to change
// \param color New color of the pixel
// void sfImage_setPixel(sfImage* image, unsigned int x, unsigned int y, sfColor color);

func (self Image) Setpixel(x int , y int , color Color) void { 
    return C.sfImage_setPixel();
}
            
//  Get the color of a pixel in an image
// This function doesn't check the validity of the pixel
// coordinates, using out-of-range values will result in
// an undefined behaviour.
// \param image Image object
// \param x     X coordinate of pixel to get
// \param y     Y coordinate of pixel to get
// \return Color of the pixel at coordinates (x, y)
// sfColor sfImage_getPixel(const sfImage* image, unsigned int x, unsigned int y);

func (self Image) Getpixel(x int , y int ) Color { 
    return C.sfImage_getPixel();
}
            
//  Get a read-only pointer to the array of pixels of an image
// The returned value points to an array of RGBA pixels made of
// 8 bits integers components. The size of the array is
// getWidth() * getHeight() * 4.
// Warning: the returned pointer may become invalid if you
// modify the image, so you should never store it for too long.
// If the image is empty, a null pointer is returned.
// \param image Image object
// \return Read-only pointer to the array of pixels
// const sfUint8* sfImage_getPixelsPtr(const sfImage* image);

func (self *Uint8) *Uint8(Image_getPixelsPtr)  { 
    return C.sf*Uint8();
}
            
*/
            
//  Flip an image horizontally (left <-> right)
// \param image Image object
// void sfImage_flipHorizontally(sfImage* image);
func (self Image) FlipHorizontally() { 
    C.sfImage_flipHorizontally(self.Cref)
}
            
//  Flip an image vertically (top <-> bottom)
// \param image Image object
// void sfImage_flipVertically(sfImage* image);

func (self Image) FlipVertically() { 
    C.sfImage_flipVertically(self.Cref)
}


// may or may not implement these

//  Create an image from a file in memory
// The supported image formats are bmp, png, tga, jpg, gif,
// psd, hdr and pic. Some format options are not supported,
// like progressive jpeg.
// If this function fails, the image is left unchanged.
// \param data Pointer to the file data in memory
// \param size Size of the data to load, in bytes
// \return A new sfImage object, or NULL if it failed
// sfImage* sfImage_createFromMemory(const void* data, size_t size);
// func (self Image) Createfrommemory(size size_t) *Image { 
//     return C.sfImage_createFromMemory();
// }
            
//  Create an image from a custom stream
// The supported image formats are bmp, png, tga, jpg, gif,
// psd, hdr and pic. Some format options are not supported,
// like progressive jpeg.
// If this function fails, the image is left unchanged.
// \param stream Source stream to read from
// \return A new sfImage object, or NULL if it failed
// sfImage* sfImage_createFromStream(sfInputStream* stream);
// func (self Image) Createfromstream() *Image { 
//     return C.sfImage_createFromStream();
// }


