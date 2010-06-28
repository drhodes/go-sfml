// Created by cgo - DO NOT EDIT
//line gfx.go:1
package gfx

// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Font.h>
// #include <SFML/Graphics/Image.h>
// #include <SFML/Graphics/String.h>


import (
	"runtime"
	"unsafe"
)

type Color struct {
	Cref _C_sfColor
}

func NewFont(val *[0]byte) Font {
	tmp := Font{val}
	runtime.SetFinalizer(&tmp, (*Font).Destroy)
	return tmp
}

type Font struct {
	Cref *[0]byte
}

func NewImage(val *[0]byte) Image {
	tmp := Image{val}
	runtime.SetFinalizer(&tmp, (*Image).Destroy)
	return tmp
}

type Image struct {
	Cref *[0]byte
}

func NewFloatRect(val *_C_sfFloatRect) FloatRect {
	tmp := FloatRect{val}
	//runtime.SetFinalizer(&tmp, (* FloatRect).Destroy)
	return tmp
}

type FloatRect struct {
	Cref *_C_sfFloatRect
}

func NewIntRect(val *_C_sfIntRect) IntRect {
	tmp := IntRect{val}
	//runtime.SetFinalizer(&tmp, (* IntRect).Destroy)
	return tmp
}

type IntRect struct {
	Cref *_C_sfIntRect
}

func NewRenderWindow(val *[0]byte) RenderWindow {
	tmp := RenderWindow{val}
	//runtime.SetFinalizer(&tmp, (* RenderWindow).Destroy)
	return tmp
}

type RenderWindow struct {
	Cref *[0]byte
}

func NewString(val *[0]byte) String {
	tmp := String{val}
	runtime.SetFinalizer(&tmp, (*String).Destroy)
	return tmp
}

type String struct {
	Cref *[0]byte
}

func NewBlendMode(val *_C_sfBlendMode) BlendMode {
	tmp := BlendMode{val}
	//runtime.SetFinalizer(&tmp, (*BlendMode).Destroy)
	return tmp
}

type BlendMode struct {
	Cref *_C_sfBlendMode
}

// _Color_
// -------------------------------------------------------------------------------
func ColorFromRGB(r, g, b int) Color {
	return Color{_C_sfColor_FromRGB(_C_sfUint8(r), _C_sfUint8(g), _C_sfUint8(b))}
}

func ColorFromRGBA(r, g, b, a uint8) Color {
	return Color{_C_sfColor_FromRGBA(_C_sfUint8(r), _C_sfUint8(g), _C_sfUint8(b), _C_sfUint8(a))}
}

// Add two colors
// param other : Color
// return Component-wise saturated addition of the two colors
func (self *Color) Add(other Color) Color {
	return Color{_C_sfColor_Add(self.Cref, other.Cref)}
}

func (self *Color) AddMutate(other Color) {
	self.Cref = _C_sfColor_Add(self.Cref, other.Cref)
}


// Modulate two colors
// param other : Color
// return Component-wise multiplication of the two colors
func (self *Color) Modulate(other Color) Color {
	return Color{_C_sfColor_Modulate(self.Cref, other.Cref)}
}

func (self *Color) ModulateMutate(other Color) {
	self.Cref = _C_sfColor_Modulate(self.Cref, other.Cref)
}


// _Font_
// -------------------------------------------------------------------------------
// Create a new empty font
// return A new sfFont object, or NULL if it failed
func FontCreate() Font {
	fnt := Font{_C_sfFont_Create()}
	runtime.SetFinalizer(&fnt, (*Font).Destroy)
	return fnt
}

// Create a new font from a file
// param Filename : Path of the font file to load
// param CharSize : Size of characters in bitmap - the bigger, the higher quality
// param Charset : Characters set to generate (just pass NULL to get the default charset)
// return A new sfFont object, or NULL if it failed
func FontCreateFromFile(path string, charSize int) Font {
	fnt := Font{_C_sfFont_CreateFromFile(_C_CString(path), _C_uint(charSize), nil)}
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
	_C_sfFont_Destroy(self.Cref)
	self.Cref = nil
}

// Get the base size of characters in a font;
// All glyphs dimensions are based on this value
// param Font : Font object
// return Base size of characters
// Get the built-in default font (Arial)
// return Pointer to the default font
func GetDefaultFont() Font {
	return Font{_C_sfFont_GetDefaultFont()}
}


// _Image_
// -------------------------------------------------------------------------------
// Create a new empty image
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_Create();
func ImageCreate() Image {
	return NewImage(_C_sfImage_Create())
}

// Create a new image filled with a color
// param Width : Image width
// param Height : Image height
// param Col : Image color
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromColor(unsigned int Width, unsigned int Height, sfColor Color);
func ImageCreateFromColor(width, height uint, clr Color) Image {
	return NewImage(_C_sfImage_CreateFromColor(_C_uint(width), _C_uint(height), clr.Cref))
}

// Create a new image from an array of pixels in memory
// param Width : Image width
// param Height : Image height
// param Data : Pointer to the pixels in memory (assumed format is RGBA)
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromPixels(unsigned int Width, unsigned int Height, const sfUint8* Data);
func ImageCreateFromPixels(width uint, height uint, data []uint8) Image {
	cdata := _C_sfUint8(data[0])
	return NewImage(_C_sfImage_CreateFromPixels(_C_uint(width), _C_uint(height), &cdata))
}

// Create a new image from a file
// param Filename : Path of the image file to load
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromFile(const char* Filename);
func ImageCreateFromFile(filename string) Image {
	return NewImage(_C_sfImage_CreateFromFile(_C_CString(filename)))
}

// Create a new image from a file in memory
// param Data : Pointer to the file data in memory
// param SizeInBytes : Size of the data to load, in bytes
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromMemory(const char* Data, size_t SizeInBytes);
func ImageCreateFromMemory(data []uint8, sizeInByes int) Image {
	cdata := _C_char(data[0])
	img := Image{_C_sfImage_CreateFromMemory(&cdata, _C_size_t(sizeInByes))}
	runtime.SetFinalizer(&img, (*Image).Destroy)
	return img

}

// Destroy an existing image
// param Image : Image to delete
// sfImage_Destroy(sfImage* Image);
func (self *Image) Destroy() {
	_C_sfImage_Destroy(self.Cref)
	self.Cref = nil
}

// this should go in a utility package.
func SfBool2GoBool(cVal _C_sfBool) bool {
	return int(cVal) == 1
}

func GoBool2SfBool(goVal bool) _C_sfBool {
	if goVal {
		return _C_sfBool(1)
	}
	return _C_sfBool(0)
}

// Save the content of an image to a file
// param Image : Image to save
// param Filename : Path of the file to save (overwritten if already exist)
// return sfTrue if saving was successful
// sfBool sfImage_SaveToFile(sfImage* Image, const char* Filename);
func (self *Image) SaveToFile(filename string) bool {
	return SfBool2GoBool(_C_sfImage_SaveToFile(self.Cref, _C_CString(filename)))
}

// Create a transparency mask for an image from a specified colorkey
// param Image : Image to modify
// param ColorKey : Color to become transparent
// param Alpha : Alpha value to use for transparent pixels
//  sfImage_CreateMaskFromColor(sfImage* Image, sfColor ColorKey, sfUint8 Alpha);
func (self *Image) CreateMaskFromColor(colorKey Color, alpha uint8) {
	_C_sfImage_CreateMaskFromColor(self.Cref, colorKey.Cref, _C_sfUint8(alpha))
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
	_C_sfImage_Copy(self.Cref, source.Cref, _C_uint(destX), _C_uint(destY), *sourceRect.Cref)
}

// Create the image from the current contents of the given window
// param Image : Destination image
// param Window : Window to capture
// param SourceRect : Sub-rectangle of the screen to copy (empty by default - entire image)
// return True if creation was successful
// sfBool sfImage_CopyScreen(sfImage* Image, sfRenderWindow* Window, sfIntRect SourceRect);
func (self *Image) CopyScreen(window RenderWindow, sourceRect IntRect) bool {
	return SfBool2GoBool(_C_sfImage_CopyScreen(self.Cref, window.Cref, *sourceRect.Cref))
}

// Change the color of a pixel of an image
// Don't forget to call Update when you end modifying pixels
// param Image : Image to modify
// param X : X coordinate of pixel in the image
// param Y : Y coordinate of pixel in the image
// param Col : New color for pixel (X, Y)
// sfImage_SetPixel(sfImage* Image, unsigned int X, unsigned int Y, sfColor Color);
func (self *Image) SetPixel(x, y uint, color Color) {
	_C_sfImage_SetPixel(self.Cref, _C_uint(x), _C_uint(y), color.Cref)
}

// Get a pixel from an image
// param Image : Image to read
// param X : X coordinate of pixel in the image
// param Y : Y coordinate of pixel in the image
// return Color of pixel (x, y)
// sfColor sfImage_GetPixel(sfImage* Image, unsigned int X, unsigned int Y);
func (self *Image) GetPixel(x, y uint) Color {
	return Color{_C_sfImage_GetPixel(self.Cref, _C_uint(x), _C_uint(y))}
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
	_C_sfImage_Bind(self.Cref)
}

// Enable or disable image smooth filter
// param Image : Image to modify
// param Smooth : sfTrue to enable smoothing filter, false to disable it
// sfImage_SetSmooth(sfImage* Image, sfBool Smooth);
func (self *Image) SetSmooth(smooth bool) {
	_C_sfImage_SetSmooth(self.Cref, GoBool2SfBool(smooth))
}

// Return the width of the image
// param Image : Image to read
// return Width in pixels
// unsigned int sfImage_GetWidth(sfImage* Image);
func (self *Image) GetWidth() uint {
	return uint(_C_sfImage_GetWidth(self.Cref))
}

// Return the height of the image
// param Image : Image to read
// return Height in pixels
// unsigned int sfImage_GetHeight(sfImage* Image);
func (self *Image) GetHeight() uint {
	return uint(_C_sfImage_GetHeight(self.Cref))
}

// Tells whether the smoothing filter is enabled or not on an image
// param Image : Image to read
// return sfTrue if the smoothing filter is enabled
// sfBool sfImage_IsSmooth(sfImage* Image);
func (self *Image) IsSmooth() bool {
	return SfBool2GoBool(_C_sfImage_IsSmooth(self.Cref))
}


// _Rect_
// -------------------------------------------------------------------------------
// Move a rectangle by the given offset
// param Rect : Rectangle to move
// param OffsetX : Horizontal offset
// param OffsetY : Vertical offset
// void sfFloatRect_Offset(sfFloatRect* Rect, float OffsetX, float OffsetY);
func (self *FloatRect) Offset(offsetX float, offsetY float) {
	_C_sfFloatRect_Offset(self.Cref, _C_float(offsetX), _C_float(offsetY))
}

// void sfIntRect_Offset(sfIntRect* Rect, int OffsetX, int OffsetY);
func (self *IntRect) Offset(offsetX int, offsetY int) {
	_C_sfIntRect_Offset(self.Cref, _C_int(offsetX), _C_int(offsetY))
}

// Check if a point is inside a rectangle's area
// param Rect : Rectangle to test
// param X : X coordinate of the point to test
// param Y : Y coordinate of the point to test
// return sfTrue if the point is inside
// sfBool sfFloatRect_Contains(sfFloatRect* Rect, float X, float Y);
func (self *FloatRect) Contains(x float, y float) bool {
	return SfBool2GoBool(_C_sfFloatRect_Contains(self.Cref, _C_float(x), _C_float(y)))
}

// sfBool sfIntRect_Contains(sfIntRect* Rect, int X, int Y);
func (self *IntRect) Contains(x int, y int) bool {
	return SfBool2GoBool(_C_sfIntRect_Contains(self.Cref, _C_int(x), _C_int(y)))
}

// Check intersection between two rectangles
// param Rect1 : First rectangle to test
// param Rect2 : Second rectangle to test
// param OverlappingRect : Rectangle to be filled with overlapping rect (can be NULL)
// return sfTrue if rectangles overlap
// sfBool sfFloatRect_Intersects(sfFloatRect* Rect1, sfFloatRect* Rect2, sfFloatRect* OverlappingRect);
func (self *FloatRect) Intersects(rect2 FloatRect, overlappingRect FloatRect) bool {
	return SfBool2GoBool(_C_sfFloatRect_Intersects(self.Cref, rect2.Cref, overlappingRect.Cref))
}

// sfBool sfIntRect_Intersects(sfIntRect* Rect1, sfIntRect* Rect2, sfIntRect* OverlappingRect);
func (self *IntRect) Intersects(rect2 IntRect, overlappingRect IntRect) bool {
	return SfBool2GoBool(_C_sfIntRect_Intersects(self.Cref, rect2.Cref, overlappingRect.Cref))
}


// _String_
// -------------------------------------------------------------------------------
// Create a new string
// return A new sfString object, or NULL if it failed
// sfString* sfString_Create();
func StringCreate() String {
	return NewString(_C_sfString_Create())
}


// Destroy an existing string
// param String : String to delete
// sfString_Destroy(sfString* String);
func (self *String) Destroy() {
	_C_sfString_Destroy(self.Cref)
	self.Cref = nil
}

// Set the X position of a string
// param String : String to modify
// param X : New X coordinate
// sfString_SetX(sfString* String, float X);
func (self *String) SetX(x float) {
	_C_sfString_SetX(self.Cref, _C_float(x))
}

// Set the Y position of a string
// param String : String to modify
// param Y : New Y coordinate
// sfString_SetY(sfString* String, float Y);
func (self *String) SetY(y float) {
	_C_sfString_SetY(self.Cref, _C_float(y))
}

// Set the position of a string
// param String : String to modify
// param Left : New left coordinate
// param Top : New top coordinate
// sfString_SetPosition(sfString* String, float Left, float Top);
func (self *String) SetPosition(left float, top float) {
	_C_sfString_SetPosition(self.Cref, _C_float(left), _C_float(top))
}

// Set the horizontal scale of a string
// param String : String to modify
// param Scale : New scale (must be strictly positive)
// sfString_SetScaleX(sfString* String, float Scale);
func (self *String) SetScaleX(scale float) {
	_C_sfString_SetScaleX(self.Cref, _C_float(scale))
}

// Set the vertical scale of a string
// param String : String to modify
// param Scale : New scale (must be strictly positive)
// sfString_SetScaleY(sfString* String, float Scale);
func (self *String) SetScaleY(scale float) {
	_C_sfString_SetScaleY(self.Cref, _C_float(scale))
}

// Set the scale of a string
// param String : String to modify
// param ScaleX : New horizontal scale (must be strictly positive)
// param ScaleY : New vertical scale (must be strictly positive)
// sfString_SetScale(sfString* String, float ScaleX, float ScaleY);
func (self *String) SetScale(scaleX float, scaleY float) {
	_C_sfString_SetScale(self.Cref, _C_float(scaleX), _C_float(scaleY))
}

// Set the orientation of a string
// param String : String to modify
// param Rotation : Angle of rotation, in degrees
// sfString_SetRotation(sfString* String, float Rotation);
func (self *String) SetRotation(rotation float) {
	_C_sfString_SetRotation(self.Cref, _C_float(rotation))
}

// Set the center of a string, in coordinates
// relative to its left-top corner
// param String : String to modify
// param X : X coordinate of the center
// param Y : Y coordinate of the center
// sfString_SetCenter(sfString* String, float X, float Y);
func (self *String) SetCenter(x float, y float) {
	_C_sfString_SetCenter(self.Cref, _C_float(x), _C_float(y))
}

// Set the color of a string
// param String : String to modify
// param Color : New color
// sfString_SetColor(sfString* String, sfColor Color);
func (self *String) SetColor(color Color) {
	_C_sfString_SetColor(self.Cref, color.Cref)
}

// Set the blending mode for a string
// param String : String to modify
// param Mode : New blending mode
// sfString_SetBlendMode(sfString* String, sfBlendMode Mode);
func (self *String) SetBlendMode(mode BlendMode) {
	_C_sfString_SetBlendMode(self.Cref, *mode.Cref)
}

// Get the X position of a string
// param String : String to read
// return Current X position
// float sfString_GetX(sfString* String);
func (self *String) GetX() float {
	return float(_C_sfString_GetX(self.Cref))
}

// Get the top Y of a string
// param String : String to read
// return Current Y position
// float sfString_GetY(sfString* String);
func (self *String) GetY() float {
	return float(_C_sfString_GetY(self.Cref))
}

// Get the horizontal scale of a string
// param String : String to read
// return Current X scale factor (always positive)
// float sfString_GetScaleX(sfString* String);
func (self *String) GetScaleX() float {
	return float(_C_sfString_GetScaleX(self.Cref))
}

// Get the vertical scale of a string
// param String : String to read
// return Current Y scale factor (always positive)
// float sfString_GetScaleY(sfString* String);
func (self *String) GetScaleY() float {
	return float(_C_sfString_GetScaleY(self.Cref))
}

// Get the orientation of a string
// param String : String to read
// return Current rotation, in degrees
// float sfString_GetRotation(sfString* String);
func (self *String) GetRotation() float {
	return float(_C_sfString_GetRotation(self.Cref))
}

// Get the X position of the center a string
// param String : String to read
// return Current X center position
// float sfString_GetCenterX(sfString* String);
func (self *String) GetCenterX() float {
	return float(_C_sfString_GetCenterX(self.Cref))
}

// Get the top Y of the center of a string
// param String : String to read
// return Current Y center position
// float sfString_GetCenterY(sfString* String);
func (self *String) GetCenterY() float {
	return float(_C_sfString_GetCenterY(self.Cref))
}

// Get the color of a string
// param String : String to read
// return Current color
// sfColor sfString_GetColor(sfString* String);
func (self *String) GetColor() Color {
	//free?
	return Color{_C_sfString_GetColor(self.Cref)}
}

// Get the current blending mode of a string
// param String : String to read
// return Current blending mode
// sfBlendMode sfString_GetBlendMode(sfString* String);
func (self *String) GetBlendMode() BlendMode {
	tmp := _C_sfString_GetBlendMode(self.Cref)
	return BlendMode{&tmp}
}

// Move a string
// param String : String to modify
// param OffsetX : Offset on the X axis
// param OffsetY : Offset on the Y axis
// sfString_Move(sfString* String, float OffsetX, float OffsetY);
func (self *String) Move(offsetX float, offsetY float) {
	_C_sfString_Move(self.Cref, _C_float(offsetX), _C_float(offsetY))
}

// Scale a string
// param String : String to modify
// param FactorX : Horizontal scaling factor (must be strictly positive)
// param FactorY : Vertical scaling factor (must be strictly positive)
// sfString_Scale(sfString* String, float FactorX, float FactorY);
func (self *String) Scale(factorX float, factorY float) {
	_C_sfString_Scale(self.Cref, _C_float(factorX), _C_float(factorY))
}

// Rotate a string
// param String : String to modify
// param Angle : Angle of rotation, in degrees
// sfString_Rotate(sfString* String, float Angle);
func (self *String) Rotate(angle float) {
	_C_sfString_Rotate(self.Cref, _C_float(angle))
}

// Transform a point from global coordinates into the string's local coordinates
// (ie it applies the inverse of object's center, translation, rotation and scale to the point)
// param String : String object
// param PointX : X coordinate of the point to transform
// param PointY : Y coordinate of the point to transform
// param X : Value to fill with the X coordinate of the converted point
// param Y : Value to fill with the y coordinate of the converted point
// sfString_TransformToLocal(sfString* String, float PointX, float PointY, float* X, float* Y);
func (self *String) TransformToLocal(pointX float, pointY float, x *float, y *float) {
	xPtr := _C_float(*x)	// wow bad.
	yPtr := _C_float(*y)
	_C_sfString_TransformToLocal(self.Cref, _C_float(pointX), _C_float(pointY), &xPtr, &yPtr)
}

// Transform a point from the string's local coordinates into global coordinates
// (ie it applies the object's center, translation, rotation and scale to the point)
// param String : String object
// param PointX : X coordinate of the point to transform
// param PointY : Y coordinate of the point to transform
// param X : Value to fill with the X coordinate of the converted point
// param Y : Value to fill with the y coordinate of the converted point
// sfString_TransformToGlobal(sfString* String, float PointX, float PointY, float* X, float* Y);
func (self *String) TransformToGlobal(pointX float, pointY float, x *float, y *float) {
	xPtr := _C_float(*x)	// wow bad.
	yPtr := _C_float(*y)
	_C_sfString_TransformToGlobal(self.Cref, _C_float(pointX), _C_float(pointY), &xPtr, &yPtr)
}

// Set the text of a string (from a multibyte string)
// param String : String to modify
// param Text : New text
// sfString_SetText(sfString* String, const char* Text);
func (self *String) SetText(text string) {
	_C_sfString_SetText(self.Cref, _C_CString(text))
}

// Set the text of a string (from a unicode string)
// param String : String to modify
// param Text : New text
// sfString_SetUnicodeText(sfString* String, const sfUint32* Text);
func (self *String) SetUnicodeText(text *uint32) {
	ctext := _C_sfUint32(*text)
	//free?
	_C_sfString_SetUnicodeText(self.Cref, &ctext)
}

// Set the font of a string
// param String : String to modify
// param Font : Font to use
// sfString_SetFont(sfString* String, sfFont* Font);
func (self *String) SetFont(font Font) {
	_C_sfString_SetFont(self.Cref, font.Cref)
}

// Set the size of a string
// param String : String to modify
// param Size : New size, in pixels
// sfString_SetSize(sfString* String, float Size);
func (self *String) SetSize(size float) {
	_C_sfString_SetSize(self.Cref, _C_float(size))
}

// Set the style of a string
// param String : String to modify
// param Size : New style (see sfStringStyle enum)
// sfString_SetStyle(sfString* String, unsigned long Style);
func (self *String) SetStyle(style uint64) {
	_C_sfString_SetStyle(self.Cref, _C_ulong(style))
}

// Get the text of a string (returns a unicode string)
// param String : String to read
// return Text as UTF-32
// const sfUint32* sfString_GetUnicodeText(sfString* String);

const MaxString int = 65536

func (self *String) GetUnicodeText() string {
	ptr := _C_sfString_GetUnicodeText(self.Cref)
	var data []uint32
	data = (*[MaxString]uint32)(unsafe.Pointer(ptr))

	i := 0
	str := ""
	for data[i] != 0x0 {
		str += string(data[i])
		i++
	}
	return str
}

/* !! Skipping this, because string are unicode in go !!
// Get the text of a string (returns an ANSI string)
// param String : String to read
// return Text an a locale-dependant ANSI string
// const char* sfString_GetText(sfString* String);
func (self *String) GetText() string {
	return string( C.sfString_GetText(self.Cref) )
}
*/


// Get the font used by a string
// param String : String to read
// return Pointer to the font
// sfFont* sfString_GetFont(sfString* String);
func (self *String) GetFont() Font {
	return NewFont(_C_sfString_GetFont(self.Cref))
}

// Get the size of the characters of a string
// param String : String to read
// return Size of the characters
// float sfString_GetSize(sfString* String);
func (self *String) GetSize() float {
	return float(_C_sfString_GetSize(self.Cref))
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
func (self *String) GetCharacterPos(index int, x *float, y *float) {
	//free?
	xPtr := (*_C_float)(x)
	yPtr := (*_C_float)(unsafe.Pointer(y))
	_C_sfString_GetCharacterPos(self.Cref, _C_size_t(index), xPtr, yPtr)
}

// Get the bounding rectangle of a string on screen
// param String : String to read
// return Rectangle contaning the string in screen coordinates
// sfFloatRect sfString_GetRect(sfString* String);
func (self *String) GetRect() FloatRect {
	//free?
	tmp := _C_sfString_GetRect(self.Cref)
	return FloatRect{&tmp}
}
