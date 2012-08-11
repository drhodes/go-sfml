package sfml

// #cgo LDFLAGS:-lcsfml-graphics
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/Graphics/Texture.h>
// #include <SFML/Window/Types.h>
// #include <SFML/System/InputStream.h>
// #include <SFML/System/Vector2.h>
// #include <stddef.h>
import "C"

import (
	"errors"
	"unsafe"
)

type Texture struct {
	Cref *C.sfTexture
}

// \brief Create a new texture
// \param width  Texture width
// \param height Texture height
// \return A new sfTexture object, or NULL if it failed
// sfTexture* sfTexture_create(unsigned int width, unsigned int height);
func NewTexture(w, h int) Texture {
	return Texture{C.sfTexture_create(C.uint(w), C.uint(h))}
}

// \brief Create a new texture from a file
// \param filename Path of the image file to load
// \param area     Area of the source image to load (NULL to load the entire image)
// \return A new sfTexture object, or NULL if it failed
// sfTexture* sfTexture_createFromFile(const char* filename, const sfIntRect* area);
func TextureFromFile(fname string, area IntRect) (Texture, error) {
	tex := C.sfTexture_createFromFile(C.CString(fname), area.Cref)
	if tex == nil {
		return Texture{nil}, errors.New("Couldn't create texture from file: " + fname)
	}
	return Texture{tex}, nil
}

// TODO:  Is this possible to do in go?
// \brief Create a new texture from a file in memory
// \param data        Pointer to the file data in memory
// \param sizeInBytes Size of the data to load, in bytes
// \param area        Area of the source image to load (NULL to load the entire image)
// \return A new sfTexture object, or NULL if it failed
// sfTexture* sfTexture_createFromMemory(const void* data, size_t sizeInBytes, const sfIntRect* area);
// func (self Texture) Createfrommemory(sizeInBytes size_t, area *IntRect ) *Texture { 
//     return C.sfTexture_createFromMemory(self.Cref, sfSize_t(sizeInBytes), sf(*IntRect));
// }

// TODO: go doesn't have streams.    
// \brief Create a new texture from a custom stream
// \param stream Source stream to read from
// \param area   Area of the source image to load (NULL to load the entire image)
// \return A new sfTexture object, or NULL if it failed
// sfTexture* sfTexture_createFromStream(sfInputStream* stream, const sfIntRect* area);
// func (self Texture) Createfromstream(area *IntRect ) *Texture { 
//     return C.sfTexture_createFromStream(self.Cref, sf(*IntRect));
// }

// \brief Create a new texture from an image
// \param image Image to upload to the texture
// \param area  Area of the source image to load (NULL to load the entire image)
// \return A new sfTexture object, or NULL if it failed
// sfTexture* sfTexture_createFromImage(const sfImage* image, const sfIntRect* area);
func TextureFromImageRect(img Image, area IntRect) (Texture, error) {
	tex := C.sfTexture_createFromImage(img.Cref, area.Cref)
	if tex == nil {
		return Texture{nil}, errors.New("Couldn't create texture from image")
	}
	return Texture{tex}, nil
}

func TextureFromImageWhole(img Image) (Texture, error) {
	tex := C.sfTexture_createFromImage(img.Cref, nil)
	if tex == nil {
		return Texture{nil}, errors.New("Couldn't create texture from image")
	}
	return Texture{tex}, nil
}

// \brief Copy an existing texture
// \param texture Texture to copy
// \return Copied object
// sfTexture* sfTexture_copy(sfTexture* texture);
func (self Texture) Copy() Texture {
	return Texture{C.sfTexture_copy(self.Cref)}
}

// \brief Destroy an existing texture
// \param texture Texture to delete
// void sfTexture_destroy(sfTexture* texture);
func (self Texture) Destroy() {
	C.sfTexture_destroy(self.Cref)
}

// \brief Return the size of the texture
// \param texture Texture to read
// \return Size in pixels
// sfVector2u sfTexture_getSize(const sfTexture* texture);
func (self Texture) Size() (x, y uint) {
	v := C.sfTexture_getSize(self.Cref)
	return uint(v.x), uint(v.y)
}

// \brief Copy a texture's pixels to an image
// \param texture Texture to copy
// \return Image containing the texture's pixels
// sfImage* sfTexture_copyToImage(const sfTexture* texture);
func (self Texture) CopyToImage() Image {
	return Image{C.sfTexture_copyToImage(self.Cref)}
}

// \brief Update a texture from an array of pixels
// \param texture Texture to update
// \param pixels  Array of pixels to copy to the texture
// \param width   Width of the pixel region contained in \a pixels
// \param height  Height of the pixel region contained in \a pixels
// \param x       X offset in the texture where to copy the source pixels
// \param y       Y offset in the texture where to copy the source pixels
// void sfTexture_updateFromPixels(sfTexture* texture, const sfUint8* pixels, unsigned int width, unsigned int height, unsigned int x, unsigned int y);
func (self Texture) UpdateFromPixels(pixels []uint8, w, h, x, y int) {
	ptr := unsafe.Pointer(&pixels[0])
	p := (*C.sfUint8)(ptr)
	C.sfTexture_updateFromPixels(self.Cref, p,
		C.uint(w), C.uint(h), C.uint(x), C.uint(y))
}

// \brief Update a texture from an image
// \param texture Texture to update
// \param image   Image to copy to the texture
// \param x       X offset in the texture where to copy the source pixels
// \param y       Y offset in the texture where to copy the source pixels
// void sfTexture_updateFromImage(sfTexture* texture, const sfImage* image, unsigned int x, unsigned int y);
func (self Texture) UpdateFromImage(img Image, x, y uint) {
	C.sfTexture_updateFromImage(self.Cref, img.Cref, C.uint(x), C.uint(y))
}

// \brief Update a texture from the contents of a window
// \param texture Texture to update
// \param window  Window to copy to the texture
// \param x       X offset in the texture where to copy the source pixels
// \param y       Y offset in the texture where to copy the source pixels
// void sfTexture_updateFromWindow(sfTexture* texture, const sfWindow* window, unsigned int x, unsigned int y);
func (self Texture) UpdateFromWindow(w Window, x, y uint) {
	C.sfTexture_updateFromWindow(self.Cref, w.Cref, C.uint(x), C.uint(y))
}

// \brief Update a texture from the contents of a render-window
// \param texture      Texture to update
// \param renderWindow Render-window to copy to the texture
// \param x            X offset in the texture where to copy the source pixels
// \param y            Y offset in the texture where to copy the source pixels
// void sfTexture_updateFromRenderWindow(sfTexture* texture, const sfRenderWindow* renderWindow, unsigned int x, unsigned int y);
func (self Texture) UpdateFromRenderWindow(win RenderWindow, x, y uint) {
	C.sfTexture_updateFromRenderWindow(self.Cref, win.Cref, C.uint(x), C.uint(y))
}

// \brief Activate a texture for rendering
// \param texture Texture to bind
// void sfTexture_bind(const sfTexture* texture);
func (self Texture) Bind() {
	C.sfTexture_bind(self.Cref)
}

// \brief Enable or disable the smooth filter on a texture
// \param texture The texture object
// \param smooth  sfTrue to enable smoothing, sfFalse to disable it
// void sfTexture_setSmooth(sfTexture* texture, sfBool smooth);
func (self Texture) SetSmooth(smooth bool) {
	C.sfTexture_setSmooth(self.Cref, Bool(smooth))
}

// \brief Tell whether the smooth filter is enabled or not for a texture
// \param texture The texture object
// \return sfTrue if smoothing is enabled, sfFalse if it is disabled
// sfBool sfTexture_isSmooth(const sfTexture* texture);
func (self Texture) IsSmooth() bool {
	return C.sfTexture_isSmooth(self.Cref) == 1
}

// \brief Enable or disable repeating for a texture
// Repeating is involved when using texture coordinates
// outside the texture rectangle [0, 0, width, height].
// In this case, if repeat mode is enabled, the whole texture
// will be repeated as many times as needed to reach the
// coordinate (for example, if the X texture coordinate is
// 3 * width, the texture will be repeated 3 times).
// If repeat mode is disabled, the "extra space" will instead
// be filled with border pixels.
// Warning: on very old graphics cards, white pixels may appear
// when the texture is repeated. With such cards, repeat mode
// can be used reliably only if the texture has power-of-two
// dimensions (such as 256x128).
// Repeating is disabled by default.
// \param texture  The texture object
// \param repeated True to repeat the texture, false to disable repeating
// void sfTexture_setRepeated(sfTexture* texture, sfBool repeated);
func (self Texture) SetRepeated(repeated bool) {
	C.sfTexture_setRepeated(self.Cref, Bool(repeated))
}

// \brief Tell whether a texture is repeated or not
// \param texture The texture object
// \return sfTrue if repeat mode is enabled, sfFalse if it is disabled
// sfBool sfTexture_isRepeated(const sfTexture* texture);
func (self Texture) IsRepeated() bool {
	return C.sfTexture_isRepeated(self.Cref) == 1
}

// \brief Get the maximum texture size allowed
// \return Maximum size allowed for textures, in pixels
// unsigned int sfTexture_getMaximumSize();
func TextureMaximumSize() uint {
	return uint(C.sfTexture_getMaximumSize())
}
