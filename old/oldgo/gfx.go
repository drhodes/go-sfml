package sfml

/*
#include "color_patch.c"
#include <SFML/Graphics/Font.h>
#include <SFML/Graphics/Image.h>
#include <SFML/Graphics/String.h>
#include <SFML/Graphics/View.h>
#include <SFML/Graphics/Shape.h>
#include <SFML/Graphics/Sprite.h>
#include <SFML/Graphics/PostFX.h>
#include <SFML/Graphics/RenderWindow.h>
*/
import "C"

import(
	"runtime"
	"unsafe"
)

func NewColor(Val C.sfColor) Color {
	return Color{ Val }
}

type Color struct { 
	Cref C.sfColor 
}


func Color_FromRGB_P(R, G, B uint8) C.sfColor_P {
	return C.sfColor_FromRGB_P(C.sfUint8(R), C.sfUint8(G), C.sfUint8(B))
}

/*
var	Black Color = NewColor( C.sfBlack )
var White Color = NewColor( C.sfWhite )
var Red Color = NewColor( C.sfRed )
var Green Color = NewColor( C.sfGreen )
var Blue Color = NewColor( C.sfBlue )
var Yellow Color = NewColor( C.sfYellow )
var Magenta Color = NewColor( C.sfMagenta )
var Cyan Color = NewColor( C.sfCyan )
*/

type WindowSettings struct { 
	Cref C.sfWindowSettings 
}

//-------------------------------------
func NewFont(val *C.sfFont) Font {
    tmp := Font{ val }
    runtime.SetFinalizer(&tmp, (*Font).Destroy)
    return tmp
}

type Font struct {
	Cref *C.sfFont
}
//-------------------------------------
func NewImage(val *C.sfImage) Image {
	tmp := Image{ val }
	runtime.SetFinalizer(&tmp, (*Image).Destroy)
	return tmp
}
type Image struct {	
	Cref *C.sfImage 
}
//-------------------------------------
func NewFloatRect(val C.sfFloatRect) FloatRect {
	var tmp FloatRect;
	tmp.Cref = val
	return tmp
}

type FloatRect struct{
	Cref C.sfFloatRect
}
//-------------------------------------
func NewIntRect(val C.sfIntRect) IntRect {
	return IntRect{ val }
}

type IntRect struct{
	Cref C.sfIntRect
}
//-------------------------------------
func NewRenderWindow(val *C.sfRenderWindow) RenderWindow {
	tmp := RenderWindow{ val }
	runtime.SetFinalizer(&tmp, (* RenderWindow).Destroy)
	return tmp
}
type RenderWindow struct{
	Cref *C.sfRenderWindow
}
//-------------------------------------
type String struct{
	Cref *C.sfString
}
func NewString(val *C.sfString) String {
	tmp := String{ val }
	runtime.SetFinalizer(&tmp, (*String).Destroy)
	return tmp
}
//-------------------------------------
func NewBlendMode(val C.sfBlendMode) BlendMode {
	tmp := BlendMode{ val }
	return tmp
}
type BlendMode struct{
	Cref C.sfBlendMode
}
//-------------------------------------
func NewView(val *C.sfView) View {
	tmp := View{ val }
	runtime.SetFinalizer(&tmp, (*View).Destroy)
	return tmp
}
type View struct {	
	Cref *C.sfView 
}
//-------------------------------------
func NewShape(val *C.sfShape) Shape {
	tmp := Shape{ val }
	runtime.SetFinalizer(&tmp, (*Shape).Destroy)
	return tmp
}
type Shape struct {	
	Cref *C.sfShape 
}
//-------------------------------------
func NewSprite(val *C.sfSprite) Sprite {
	tmp := Sprite{ val }
	runtime.SetFinalizer(&tmp, (*Sprite).Destroy)
	return tmp
}
type Sprite struct {	
	Cref *C.sfSprite 
}
//-------------------------------------
func NewPostFX(val *C.sfPostFX) PostFX {
	tmp := PostFX{ val }
	runtime.SetFinalizer(&tmp, (*PostFX).Destroy)
	return tmp
}
type PostFX struct {	
	Cref *C.sfPostFX 
}


// _Color_
// -------------------------------------------------------------------------------
func ColorFromRGB(R, G, B uint8) Color {
	//return NewColor( C.sfColor_FromRGB( C.sfUint8(R), C.sfUint8(G), C.sfUint8(B) ))
	hack := NewColor( C.sfColor_FromRGB(0,0,0))
	return hack
}

func ColorFromRGBA(R, G, B, A uint8) Color {
	return NewColor( C.sfColor_FromRGBA( C.sfUint8(R), C.sfUint8(G), C.sfUint8(B), C.sfUint8(A)))
}

/*
type SwigcptrSfColor uintptr

func (p SwigcptrSfColor) Swigcptr() uintptr {
	return (uintptr)(p)
}

func _swig_wrap_sfColor_r_set(SwigcptrSfColor, byte)

func (arg1 SwigcptrSfColor) SetR(arg2 byte) {
	C._wrap_sfColor_r_set(arg1, arg2)
}
 */

// Add two colors
// param other : Color
// return Component-wise saturated addition of the two colors
func (self *Color) Add(other Color) Color {
	return NewColor( C.sfColor_Add(self.Cref, other.Cref) )
}


// Modulate two colors
// param other : Color
// return Component-wise multiplication of the two colors
func (self *Color) Modulate(other Color) Color {
	return NewColor( C.sfColor_Modulate(self.Cref, other.Cref) )
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
    fnt := Font{ C.sfFont_CreateFromFile( C.CString(path), C.uint(charSize), nil) }
	runtime.SetFinalizer(&fnt, (*Font).Destroy)   
	return fnt	
}

// Create a new image font a file in memory
// param Data : Pointer to the file data in memory
// param SizeInBytes : Size of the data to load, in bytes
// param CharSize : Size of characters in bitmap - the bigger, the higher quality
// param Charset : Characters set to generate (just pass NULL to get the default charset)
// return A new sfFont object, or NULL if it failed

// -??
// func FontCreateFromMemory(data uint8, SizeInBytes uint8, CharSize uint) Font {
// 	cdata := C.char(data)
// 	//C.free(unsafe.Pointer(cfile))
	
//     return Font{ C.sfFont_CreateFromMemory( &cdata, 
// 			C.size_t(SizeInBytes), 
// 			C.uint(CharSize), nil )}
// }


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
    return NewImage( C.sfImage_Create() )
}

// Create a new image filled with a color
// param Width : Image width
// param Height : Image height
// param Col : Image color
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromColor(unsigned int Width, unsigned int Height, sfColor Color);
func ImageCreateFromColor(width, height uint, clr Color) Image {
    return NewImage( C.sfImage_CreateFromColor( C.uint(width), C.uint(height), clr.Cref) )
}

// Create a new image from an array of pixels in memory
// param Width : Image width
// param Height : Image height
// param Data : Pointer to the pixels in memory (assumed format is RGBA)
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromPixels(unsigned int Width, unsigned int Height, const sfUint8* Data);
func ImageCreateFromPixels(width uint, height uint, data []uint8) Image {
	cdata := C.sfUint8(data[0])
	return NewImage( C.sfImage_CreateFromPixels( C.uint(width), C.uint(height), &cdata) ) 
}

// Create a new image from a file
// param Filename : Path of the image file to load
// return A new sfImage object, or NULL if it failed
// sfImage* sfImage_CreateFromFile(const char* Filename);
func ImageCreateFromFile(filename string) Image {
    return NewImage( C.sfImage_CreateFromFile( C.CString(filename)) )
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
	return SfBool2GoBool( C.sfImage_SaveToFile(self.Cref, C.CString(filename)))
}

// Create a transparency mask for an image from a specified colorkey
// param Image : Image to modify
// param ColorKey : Color to become transparent
// param Alpha : Alpha value to use for transparent pixels
//  sfImage_CreateMaskFromColor(sfImage* Image, sfColor ColorKey, sfUint8 Alpha);
func (self *Image) CreateMaskFromColor(colorKey Color, alpha uint8) {
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
    C.sfImage_Copy(self.Cref, source.Cref, C.uint(destX), C.uint(destY), sourceRect.Cref)
}

// Create the image from the current contents of the given window
// param Image : Destination image
// param Window : Window to capture
// param SourceRect : Sub-rectangle of the screen to copy (empty by default - entire image)
// return True if creation was successful
// sfBool sfImage_CopyScreen(sfImage* Image, sfRenderWindow* Window, sfIntRect SourceRect);
func (self *Image) CopyScreen(window RenderWindow, sourceRect IntRect) bool {
    return SfBool2GoBool( C.sfImage_CopyScreen(self.Cref, window.Cref, sourceRect.Cref) )
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

//  ARRGgghh
// //func NewArray(typ interface{}, n int) Pointer

// func (self *Image) GetPixelsPtr() []byte {	
// 	__COLOR__ := ColorFromRGBA(0,0,0,0)
// 	cPtr := C.sfImage_GetPixelsPtr(self.Cref)
// 	arr := unsafe.NewArray(__COLOR__, self.Height * self.Width)
// 	return arr
// }

	
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
	return uint( C.sfImage_GetWidth(self.Cref))
}

// Return the height of the image
// param Image : Image to read
// return Height in pixels
// unsigned int sfImage_GetHeight(sfImage* Image);
func (self *Image) GetHeight() uint {
	return uint( C.sfImage_GetHeight(self.Cref))
}

// Tells whether the smoothing filter is enabled or not on an image
// param Image : Image to read
// return sfTrue if the smoothing filter is enabled
// sfBool sfImage_IsSmooth(sfImage* Image);
func (self *Image) IsSmooth() bool {
    return SfBool2GoBool( C.sfImage_IsSmooth(self.Cref) )
}




// _Rect_
// -------------------------------------------------------------------------------
// Move a rectangle by the given offset
// param Rect : Rectangle to move
// param OffsetX : Horizontal offset
// param OffsetY : Vertical offset
// void sfFloatRect_Offset(sfFloatRect* Rect, float32 OffsetX, float32 OffsetY);
func (self *FloatRect) Offset(offsetX float32, offsetY float32) {
    C.sfFloatRect_Offset(&(self.Cref), C.float(offsetX), C.float(offsetY)) 
}

// void sfIntRect_Offset(sfIntRect* Rect, int OffsetX, int OffsetY);
func (self *IntRect) Offset(offsetX int, offsetY int) {
    C.sfIntRect_Offset(&(self.Cref), C.int(offsetX), C.int(offsetY)) 
}

// Check if a point is inside a rectangle's area
// param Rect : Rectangle to test
// param X : X coordinate of the point to test
// param Y : Y coordinate of the point to test
// return sfTrue if the point is inside
// sfBool sfFloatRect_Contains(sfFloatRect* Rect, float32 X, float32 Y);
func (self *FloatRect) Contains(x float32, y float32) bool {
    return SfBool2GoBool( C.sfFloatRect_Contains(&self.Cref, C.float(x), C.float(y)) )
}

// sfBool sfIntRect_Contains(sfIntRect* Rect, int X, int Y);
func (self *IntRect) Contains(x int, y int) bool {
    return SfBool2GoBool( C.sfIntRect_Contains(&self.Cref, C.int(x), C.int(y)) )
}

// Check intersection between two rectangles
// param Rect1 : First rectangle to test
// param Rect2 : Second rectangle to test
// param OverlappingRect : Rectangle to be filled with overlapping rect (can be NULL)
// return sfTrue if rectangles overlap
// sfBool sfFloatRect_Intersects(sfFloatRect* Rect1, sfFloatRect* Rect2, sfFloatRect* OverlappingRect);
func (self *FloatRect) Intersects(rect2 FloatRect, overlappingRect FloatRect) bool {
	return SfBool2GoBool( C.sfFloatRect_Intersects(&(self.Cref), 
		&(rect2.Cref), &(overlappingRect.Cref)))
}

// sfBool sfIntRect_Intersects(sfIntRect* Rect1, sfIntRect* Rect2, sfIntRect* OverlappingRect);
func (self *IntRect) Intersects(rect2 IntRect, overlappingRect IntRect) bool {
    return SfBool2GoBool( C.sfIntRect_Intersects(&(self.Cref), 
		&(rect2.Cref), 
		&(overlappingRect.Cref)))
}





// _String_
// -------------------------------------------------------------------------------
// Create a new string
// return A new sfString object, or NULL if it failed
// sfString* sfString_Create();
func StringCreate() String {
    return NewString( C.sfString_Create() )
}


// Destroy an existing string
// param String : String to delete
// sfString_Destroy(sfString* String);
func (self *String) Destroy() {
    C.sfString_Destroy(self.Cref)
	self.Cref = nil
}

// Set the X position of a string
// param String : String to modify
// param X : New X coordinate
// sfString_SetX(sfString* String, float32 X);
func (self *String) SetX(x float32) {
    C.sfString_SetX(self.Cref, C.float(x))
}

// Set the Y position of a string
// param String : String to modify
// param Y : New Y coordinate
// sfString_SetY(sfString* String, float32 Y);
func (self *String) SetY(y float32) {
    C.sfString_SetY(self.Cref, C.float(y))
}

// Set the position of a string
// param String : String to modify
// param Left : New left coordinate
// param Top : New top coordinate
// sfString_SetPosition(sfString* String, float32 Left, float32 Top);
func (self *String) SetPosition(left float32, top float32) {
    C.sfString_SetPosition(self.Cref, C.float(left), C.float(top))
}

// Set the horizontal scale of a string
// param String : String to modify
// param Scale : New scale (must be strictly positive)
// sfString_SetScaleX(sfString* String, float32 Scale);
func (self *String) SetScaleX(scale float32) {
    C.sfString_SetScaleX(self.Cref, C.float(scale))
}

// Set the vertical scale of a string
// param String : String to modify
// param Scale : New scale (must be strictly positive)
// sfString_SetScaleY(sfString* String, float32 Scale);
func (self *String) SetScaleY(scale float32) {
    C.sfString_SetScaleY(self.Cref, C.float(scale))
}

// Set the scale of a string
// param String : String to modify
// param ScaleX : New horizontal scale (must be strictly positive)
// param ScaleY : New vertical scale (must be strictly positive)
// sfString_SetScale(sfString* String, float32 ScaleX, float32 ScaleY);
func (self *String) SetScale(scaleX float32, scaleY float32) {
    C.sfString_SetScale(self.Cref, C.float(scaleX), C.float(scaleY))
}

// Set the orientation of a string
// param String : String to modify
// param Rotation : Angle of rotation, in degrees
// sfString_SetRotation(sfString* String, float32 Rotation);
func (self *String) SetRotation(rotation float32) {
    C.sfString_SetRotation(self.Cref, C.float(rotation))
}

// Set the center of a string, in coordinates
// relative to its left-top corner
// param String : String to modify
// param X : X coordinate of the center
// param Y : Y coordinate of the center
// sfString_SetCenter(sfString* String, float32 X, float32 Y);
func (self *String) SetCenter(x float32, y float32) {
    C.sfString_SetCenter(self.Cref, C.float(x), C.float(y))
}

// Set the color of a string
// param String : String to modify
// param Color : New color
// sfString_SetColor(sfString* String, sfColor Color);
func (self *String) SetColor(color Color) {
    C.sfString_SetColor(self.Cref, color.Cref)
}

// Set the blending mode for a string
// param String : String to modify
// param Mode : New blending mode
// sfString_SetBlendMode(sfString* String, sfBlendMode Mode);
func (self *String) SetBlendMode(mode BlendMode) {
    C.sfString_SetBlendMode(self.Cref, mode.Cref)
}

// Get the X position of a string
// param String : String to read
// return Current X position
// float32 sfString_GetX(sfString* String);
func (self *String) GetX() float32 {
    return float32( C.sfString_GetX(self.Cref))
}

// Get the top Y of a string
// param String : String to read
// return Current Y position
// float32 sfString_GetY(sfString* String);
func (self *String) GetY() float32 {
    return float32( C.sfString_GetY(self.Cref))
}

// Get the horizontal scale of a string
// param String : String to read
// return Current X scale factor (always positive)
// float32 sfString_GetScaleX(sfString* String);
func (self *String) GetScaleX() float32 {
    return float32( C.sfString_GetScaleX(self.Cref))
}

// Get the vertical scale of a string
// param String : String to read
// return Current Y scale factor (always positive)
// float32 sfString_GetScaleY(sfString* String);
func (self *String) GetScaleY() float32 {
    return float32( C.sfString_GetScaleY(self.Cref))
}

// Get the orientation of a string
// param String : String to read
// return Current rotation, in degrees
// float32 sfString_GetRotation(sfString* String);
func (self *String) GetRotation() float32 {
    return float32( C.sfString_GetRotation(self.Cref))
}

// Get the X position of the center a string
// param String : String to read
// return Current X center position
// float32 sfString_GetCenterX(sfString* String);
func (self *String) GetCenterX() float32 {
    return float32( C.sfString_GetCenterX(self.Cref))
}

// Get the top Y of the center of a string
// param String : String to read
// return Current Y center position
// float32 sfString_GetCenterY(sfString* String);
func (self *String) GetCenterY() float32 {
    return float32( C.sfString_GetCenterY(self.Cref))
}

// Get the color of a string
// param String : String to read
// return Current color
// sfColor sfString_GetColor(sfString* String);
func (self *String) GetColor() Color {
	//free?
    return Color{ C.sfString_GetColor(self.Cref) }
}

// Get the current blending mode of a string
// param String : String to read
// return Current blending mode
// sfBlendMode sfString_GetBlendMode(sfString* String);
func (self *String) GetBlendMode() BlendMode {
	tmp := C.sfString_GetBlendMode(self.Cref) 
	return BlendMode{ tmp }
}

// Move a string
// param String : String to modify
// param OffsetX : Offset on the X axis
// param OffsetY : Offset on the Y axis
// sfString_Move(sfString* String, float32 OffsetX, float32 OffsetY);
func (self *String) Move(offsetX float32, offsetY float32) {
    C.sfString_Move(self.Cref, C.float(offsetX), C.float(offsetY))
}

// Scale a string
// param String : String to modify
// param FactorX : Horizontal scaling factor (must be strictly positive)
// param FactorY : Vertical scaling factor (must be strictly positive)
// sfString_Scale(sfString* String, float32 FactorX, float32 FactorY);
func (self *String) Scale(factorX float32, factorY float32) {
    C.sfString_Scale(self.Cref, C.float(factorX), C.float(factorY))
}

// Rotate a string
// param String : String to modify
// param Angle : Angle of rotation, in degrees
// sfString_Rotate(sfString* String, float32 Angle);
func (self *String) Rotate(angle float32) {
    C.sfString_Rotate(self.Cref, C.float(angle))
}

// Transform a point from global coordinates into the string's local coordinates
// (ie it applies the inverse of object's center, translation, rotation and scale to the point)
// param String : String object
// param PointX : X coordinate of the point to transform
// param PointY : Y coordinate of the point to transform
// param X : Value to fill with the X coordinate of the converted point
// param Y : Value to fill with the y coordinate of the converted point
// sfString_TransformToLocal(sfString* String, float32 PointX, float32 PointY, float32* X, float32* Y);
func (self *String) TransformToLocal(pointX float32, pointY float32, x *float32, y *float32) {
	xPtr := C.float(*x) // wow bad. 
	yPtr := C.float(*y)
    C.sfString_TransformToLocal(self.Cref, C.float(pointX), C.float(pointY), &xPtr, &yPtr)
}

// Transform a point from the string's local coordinates into global coordinates
// (ie it applies the object's center, translation, rotation and scale to the point)
// param String : String object
// param PointX : X coordinate of the point to transform
// param PointY : Y coordinate of the point to transform
// param X : Value to fill with the X coordinate of the converted point
// param Y : Value to fill with the y coordinate of the converted point
// sfString_TransformToGlobal(sfString* String, float32 PointX, float32 PointY, float32* X, float32* Y);
func (self *String) TransformToGlobal(pointX float32, pointY float32, x *float32, y *float32) {
	xPtr := C.float(*x) // wow bad. 
	yPtr := C.float(*y)
    C.sfString_TransformToGlobal(self.Cref, C.float(pointX), C.float(pointY), &xPtr, &yPtr)
}

// Set the text of a string (from a multibyte string)
// param String : String to modify
// param Text : New text
// sfString_SetText(sfString* String, const char* Text);
func (self *String) SetText(text string) {
    C.sfString_SetText(self.Cref, C.CString(text))
}

// Set the text of a string (from a unicode string)
// param String : String to modify
// param Text : New text
// sfString_SetUnicodeText(sfString* String, const sfUint32* Text);
func (self *String) SetUnicodeText(text *uint32) {
	ctext := C.sfUint32(*text)
	//free?
    C.sfString_SetUnicodeText(self.Cref, &ctext)
}

// Set the font of a string
// param String : String to modify
// param Font : Font to use
// sfString_SetFont(sfString* String, sfFont* Font);
func (self *String) SetFont(font Font) {
    C.sfString_SetFont(self.Cref, font.Cref)
}

// Set the size of a string
// param String : String to modify
// param Size : New size, in pixels
// sfString_SetSize(sfString* String, float32 Size);
func (self *String) SetSize(size float32) {
    C.sfString_SetSize(self.Cref, C.float(size))
}

// Set the style of a string
// param String : String to modify
// param Size : New style (see sfStringStyle enum)
// sfString_SetStyle(sfString* String, unsigned long Style);
func (self *String) SetStyle(style uint64) {
    C.sfString_SetStyle(self.Cref, C.ulong(style))
}

// Get the text of a string (returns a unicode string)
// param String : String to read
// return Text as UTF-32
// const sfUint32* sfString_GetUnicodeText(sfString* String);

const MaxString int = 65536

func (self *String) GetUnicodeText() string {
	ptr := C.sfString_GetUnicodeText(self.Cref)
	var data *[MaxString]uint32
	data = (*[MaxString]uint32)(unsafe.Pointer(ptr)) 

	i := 0
	str := ""
	for data[i] != 0x0 {
		str += string(data[i])
		i++
	}
	return str	
}

//  !! Skipping this, because string are unicode in go !!	
// // Get the text of a string (returns an ANSI string)
// // param String : String to read
// // return Text an a locale-dependant ANSI string
// // const char* sfString_GetText(sfString* String);
// func (self *String) GetText() string {
// 	return string( C.sfString_GetText(self.Cref) )
// }



// Get the font used by a string
// param String : String to read
// return Pointer to the font
// sfFont* sfString_GetFont(sfString* String);
func (self *String) GetFont() Font {
    return NewFont( C.sfString_GetFont(self.Cref) )
}

// Get the size of the characters of a string
// param String : String to read
// return Size of the characters
// float32 sfString_GetSize(sfString* String);
func (self *String) GetSize() float32 {
    return float32( C.sfString_GetSize(self.Cref))
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
// sfString_GetCharacterPos(sfString* String, size_t Index, float32* X, float32* Y);
//  HELP!
// func (self *String) GetCharacterPos(index int, x *float, y *float) {
// 	//free?
// 	xPtr := (*C.float)(x)
// 	yPtr := (*C.float)(unsafe.Pointer(y))
//     C.sfString_GetCharacterPos(self.Cref, C.size_t(index), xPtr, yPtr)
// }


// Get the bounding rectangle of a string on screen
// param String : String to read
// return Rectangle contaning the string in screen coordinates
// sfFloatRect sfString_GetRect(sfString* String);
func (self *String) GetRect() FloatRect {
	//free?
	return NewFloatRect( C.sfString_GetRect(self.Cref) )
}




// _View_
// -------------------------------------------------------------------------------
// Construct a default view (1000x1000)
// sfView* sfView_Create();
func CreateView() View {
	return NewView( C.sfView_Create() )
}

// Construct a view from a rectangle
// param Rect : Rectangle defining the bounds of the view
// sfView* sfView_CreateFromRect(sfFloatRect Rect);
func CreateViewFromRect(rect FloatRect) View {
    return NewView( C.sfView_CreateFromRect(rect.Cref) )
}

// Destroy an existing view
// param View : View to destroy
// sfView_Destroy(sfView* View);
func (self *View) Destroy() {
    C.sfView_Destroy(self.Cref)
	self.Cref = nil
}

// Change the center of a view
// param View : View to modify
// param X : X coordinate of the new center
// param Y : Y coordinate of the new center
// sfView_SetCenter(sfView* View, float32 X, float32 Y);
func (self *View) SetCenter(x, y float32) {
    C.sfView_SetCenter(self.Cref, C.float(x), C.float(y))
}

// Change the half-size of a view
// param View : View to modify
// param HalfWidth : New half-width
// param HalfHeight : New half-height
// sfView_SetHalfSize(sfView* View, float32 HalfWidth, float32 HalfHeight);
func (self *View) SetHalfSize(halfWidth float32, halfHeight float32) {
    C.sfView_SetHalfSize(self.Cref, C.float(halfWidth), C.float(halfHeight))
}

// Rebuild a view from a rectangle
// param View : View to modify
// param ViewRect : Rectangle defining the position and size of the view
// sfView_SetFromRect(sfView* View, sfFloatRect ViewRect);
func (self *View) SetFromRect(viewRect FloatRect) {
    C.sfView_SetFromRect(self.Cref, viewRect.Cref)
}

// Get the X coordinate of the center of a view
// param View : View to read
// return X coordinate of the center of the view
// float32 sfView_GetCenterX(sfView* View);
func (self *View) GetCenterX() float32 {
    return float32( C.sfView_GetCenterX(self.Cref))
}

// Get the Y coordinate of the center of a view
// param View : View to read
// return Y coordinate of the center of the view
// float32 sfView_GetCenterY(sfView* View);
func (self *View) GetCenterY() float32 {
    return float32( C.sfView_GetCenterY(self.Cref))
}

// Get the half-width of the view
// param View : View to read
// return Half-width of the view
// float32 sfView_GetHalfSizeX(sfView* View);
func (self *View) GetHalfSizeX() float32 {
    return float32( C.sfView_GetHalfSizeX(self.Cref))
}

// Get the half-height of the view
// param View : View to read
// return Half-height of the view
// float32 sfView_GetHalfSizeY(sfView* View);
func (self *View) GetHalfSizeY() float32 {
    return float32( C.sfView_GetHalfSizeY(self.Cref))
}

// Get the bounding rectangle of a view
// param View : View to read
// return Bounding rectangle of the view
// sfFloatRect sfView_GetRect(sfView* View);
func (self *View) GetRect() FloatRect {
    return NewFloatRect( C.sfView_GetRect(self.Cref))
}

// Move a view
// param View : View to move
// param OffsetX : Offset to move the view, on X axis
// param OffsetY : Offset to move the view, on Y axis
// sfView_Move(sfView* View, float32 OffsetX, float32 OffsetY);
func (self *View) Move(offsetX, offsetY float32) {
    C.sfView_Move(self.Cref, C.float(offsetX), C.float(offsetY))
}

// Resize a view rectangle to simulate a zoom / unzoom effect
// param View : View to zoom
// param Factor : Zoom factor to apply, relative to the current zoom
// sfView_Zoom(sfView* View, float32 Factor);
func (self *View) Zoom(factor float32) {
    C.sfView_Zoom(self.Cref, C.float(factor))
}




// _Shape_
// -------------------------------------------------------------------------------
// Create a new shape
// return A new sfShape object, or NULL if it failed
// sfShape* sfShape_Create();
func CreateShape() Shape {
    return NewShape( C.sfShape_Create())
}


// Create a new shape made of a single line
// param P1X, P1Y : Position of the first point
// param P2X, P2Y : Position second point
// param Thickness : Line thickness
// param Col : Color used to draw the line
// param Outline : Outline width
// param OutlineCol : Color used to draw the outline
// sfShape* sfShape_CreateLine(float P1X, float32 P1Y, float32 P2X, float32 P2Y, float32 Thickness, sfColor Col, float32 Outline, sfColor OutlineCol);
func CreateLine(p1Y, p2X, p2Y, thickness float32, col Color, outline float32, outlineCol Color) Shape {
    //return NewShape( C.sfShape_CreateLine(self.Cref, p1Y, p2X, p2Y, thickness, col, outline, outlineCol) )
	return NewShape( C.sfShape_CreateLine( C.float(p2X), C.float(p1Y), C.float(p2X), C.float(p2Y), C.float(thickness), col.Cref, C.float(outline), outlineCol.Cref))
}

// Create a new shape made of a single rectangle
// param P1X, P1Y : Position of the first point
// param P2X, P2Y : Position second point
// param Col : Color used to fill the rectangle
// param Outline : Outline width
// param OutlineCol : Color used to draw the outline
// sfShape* sfShape_CreateRectangle(float P1X, float32 P1Y, float32 P2X, float32 P2Y, sfColor Col, float32 Outline, sfColor OutlineCol);
func CreateRectangle(p1X, p1Y, p2X, p2Y float32, col Color, outline float32, outlineCol Color) Shape {
    return NewShape( C.sfShape_CreateRectangle( C.float(p1X), C.float(p1Y), 
		C.float(p2X), C.float(p2Y), col.Cref, C.float(outline), outlineCol.Cref))
}

// Create a new shape made of a single circle
// param X, Y : Position of the center
// param Radius : Radius
// param Col : Color used to fill the circle
// param Outline : Outline width
// param OutlineCol : Color used to draw the outline
// sfShape* sfShape_CreateCircle(float X, float32 Y, float32 Radius, sfColor Col, float32 Outline, sfColor OutlineCol);
func CreateCircle(x, y, radius float32, col Color, outline float32, outlineCol Color) Shape {
    return NewShape( C.sfShape_CreateCircle( C.float(x), C.float(y), C.float(radius), col.Cref, C.float(outline), outlineCol.Cref))
}

// Destroy an existing Shape
// param Shape : Shape to delete
// sfShape_Destroy(sfShape* Shape);
func (self *Shape) Destroy() {
    C.sfShape_Destroy(self.Cref)
}



// Set the X position of a shape
// param Shape : Shape to modify
// param X : New X coordinate
// sfShape_SetX(sfShape* Shape, float32 X);
func (self *Shape) SetX(x float32) {
    C.sfShape_SetX(self.Cref, C.float(x))
}

// Set the Y position of a shape
// param Shape : Shape to modify
// param Y : New Y coordinate
// sfShape_SetY(sfShape* Shape, float32 Y);
func (self *Shape) SetY(y float32) {
    C.sfShape_SetY(self.Cref, C.float(y))
}

// Set the position of a shape
// param Shape : Shape to modify
// param X : New X coordinate
// param Y : New Y coordinate
// sfShape_SetPosition(sfShape* Shape, float32 X, float32 Y);
func (self *Shape) SetPosition(x, y float32) {
    C.sfShape_SetPosition(self.Cref, C.float(x), C.float(y))
}

// Set the horizontal scale of a shape
// param Shape : Shape to modify
// param Scale : New scale (must be strictly positive)
// sfShape_SetScaleX(sfShape* Shape, float32 Scale);
func (self *Shape) SetScaleX(scale float32) {
    C.sfShape_SetScaleX(self.Cref, C.float(scale))
}

// Set the vertical scale of a shape
// param Shape : Shape to modify
// param Scale : New scale (must be strictly positive)
// sfShape_SetScaleY(sfShape* Shape, float32 Scale);
func (self *Shape) SetScaleY(scale float32) {
    C.sfShape_SetScaleY(self.Cref, C.float(scale))
}

// Set the scale of a shape
// param Shape : Shape to modify
// param ScaleX : New horizontal scale (must be strictly positive)
// param ScaleY : New vertical scale (must be strictly positive)
// sfShape_SetScale(sfShape* Shape, float32 ScaleX, float32 ScaleY);
func (self *Shape) SetScale(scaleX, scaleY float32) {
    C.sfShape_SetScale(self.Cref, C.float(scaleX), C.float(scaleY))
}

// Set the orientation of a shape
// param Shape : Shape to modify
// param Rotation : Angle of rotation, in degrees
// sfShape_SetRotation(sfShape* Shape, float32 Rotation);
func (self *Shape) SetRotation(rotation float32) {
    C.sfShape_SetRotation(self.Cref, C.float(rotation))
}

// Set the center of a shape, in coordinates relative to
// its left-top corner
// param Shape : Shape to modify
// param X : X coordinate of the center
// param Y : Y coordinate of the center
// sfShape_SetCenter(sfShape* Shape, float32 X, float32 Y);
func (self *Shape) SetCenter(x float32, y float32) {
    C.sfShape_SetCenter(self.Cref, C.float(x), C.float(y))
}

// Set the color of a shape
// param Shape : Shape to modify
// param Color : New color
// sfShape_SetColor(sfShape* Shape, sfColor Color);
func (self *Shape) SetColor(color Color) {
    C.sfShape_SetColor(self.Cref, color.Cref)
}

// Set the blending mode for a shape
// param Shape : Shape to modify
// param Mode : New blending mode
// sfShape_SetBlendMode(sfShape* Shape, sfBlendMode Mode);
func (self *Shape) SetBlendMode(mode BlendMode) {
    C.sfShape_SetBlendMode(self.Cref, mode.Cref)
}

// Get the X position of a shape
// param Shape : Shape to read
// return Current X position
// float32 sfShape_GetX(sfShape* Shape);
func (self *Shape) GetX() float32 {
    return float32( C.sfShape_GetX(self.Cref))
}

// Get the Y position of a shape
// param Shape : Shape to read
// return Current Y position
// float32 sfShape_GetY(sfShape* Shape);
func (self *Shape) GetY() float32 {
    return float32( C.sfShape_GetY(self.Cref))
}

// Get the horizontal scale of a shape
// param Shape : Shape to read
// return Current X scale factor (always positive)
// float32 sfShape_GetScaleX(sfShape* Shape);
func (self *Shape) GetScaleX() float32 {
    return float32( C.sfShape_GetScaleX(self.Cref))
}

// Get the vertical scale of a shape
// param Shape : Shape to read
// return Current Y scale factor (always positive)
// float32 sfShape_GetScaleY(sfShape* Shape);
func (self *Shape) GetScaleY() float32 {
    return float32( C.sfShape_GetScaleY(self.Cref))
}

// Get the orientation of a shape
// param Shape : Shape to read
// return Current rotation, in degrees
// float32 sfShape_GetRotation(sfShape* Shape);
func (self *Shape) GetRotation() float32 {
    return float32( C.sfShape_GetRotation(self.Cref))
}

// Get the X position of the center a shape
// param Shape : Shape to read
// return Current X center
// float32 sfShape_GetCenterX(sfShape* Shape);
func (self *Shape) GetCenterX() float32 {
    return float32( C.sfShape_GetCenterX(self.Cref))
}

// Get the Y position of the center a shape
// param Shape : Shape to read
// return Current Y center
// float32 sfShape_GetCenterY(sfShape* Shape);
func (self *Shape) GetCenterY() float32 {
    return float32( C.sfShape_GetCenterY(self.Cref))
}

// Get the color of a shape
// param Shape : Shape to read
// return Current color
// sfColor sfShape_GetColor(sfShape* Shape);
func (self *Shape) GetColor() Color {
    return NewColor( C.sfShape_GetColor(self.Cref))
}

// Get the current blending mode of a shape
// param Shape : Shape to read
// return Current blending mode
// sfBlendMode sfShape_GetBlendMode(sfShape* Shape);
func (self *Shape) GetBlendMode() BlendMode {
    return NewBlendMode( C.sfShape_GetBlendMode(self.Cref))
}

// Move a shape
// param Shape : Shape to modify
// param OffsetX : Offset on the X axis
// param OffsetY : Offset on the Y axis
// sfShape_Move(sfShape* Shape, float32 OffsetX, float32 OffsetY);
func (self *Shape) Move(offsetX float32, offsetY float32) {
    C.sfShape_Move(self.Cref, C.float(offsetX), C.float(offsetY))
}

// Scale a shape
// param Shape : Shape to modify
// param FactorX : Horizontal scaling factor (must be strictly positive)
// param FactorY : Vertical scaling factor (must be strictly positive)
// sfShape_Scale(sfShape* Shape, float32 FactorX, float32 FactorY);
func (self *Shape) Scale(factorX float32, factorY float32) {
    C.sfShape_Scale(self.Cref, C.float(factorX), C.float(factorY))
}

// Rotate a shape
// param Shape : Shape to modify
// param Angle : Angle of rotation, in degrees
// sfShape_Rotate(sfShape* Shape, float32 Angle);
func (self *Shape) Rotate(angle float32) {
    C.sfShape_Rotate(self.Cref, C.float(angle))
}

// Transform a point from global coordinates into the shape's local coordinates
// (ie it applies the inverse of object's center, translation, rotation and scale to the point)
// param Shape : Shape object
// param PointX : X coordinate of the point to transform
// param PointY : Y coordinate of the point to transform
// param X : Value to fill with the X coordinate of the converted point
// param Y : Value to fill with the y coordinate of the converted point
// sfShape_TransformToLocal(sfShape* Shape, float32 PointX, float32 PointY, float32* X, float32* Y);

//  ??
// func (self *Shape) TransformToLocal(pointX float32, pointY float32, x float32*, y float32*) {
//     return C.sfShape_TransformToLocal(self.Cref, C.float(pointX), C.float(pointY), C.float(x), C.float(y))
// }

// // Transform a point from the shape's local coordinates into global coordinates
// // (ie it applies the object's center, translation, rotation and scale to the point)
// // param Shape : Shape object
// // param PointX : X coordinate of the point to transform
// // param PointY : Y coordinate of the point to transform
// // param X : Value to fill with the X coordinate of the converted point
// // param Y : Value to fill with the y coordinate of the converted point
// // sfShape_TransformToGlobal(sfShape* Shape, float32 PointX, float32 PointY, float32* X, float32* Y);
//  ??
// func (self *Shape) TransformToGlobal(pointX float32, pointY float32, x float32*, y float32*) {
//     return C.sfShape_TransformToGlobal(self.Cref, pointX, pointY, x, y)
// }


// Add a point to a shape
// param Shape : Shape to modify
// param X, Y : Position of the point
// param Col : Color of the point
// param OutlineCol : Outline color of the point
// sfShape_AddPoint(sfShape* Shape, float32 X, float32 Y, sfColor Col, sfColor OutlineCol);
func (self *Shape) AddPoint(x, y float32, col Color, outlineCol Color) {
    C.sfShape_AddPoint(self.Cref, C.float(x), C.float(y), col.Cref, outlineCol.Cref)
}

// Enable or disable filling a shape.
// Fill is enabled by default
// param Enable : True to enable, false to disable
// sfShape_EnableFill(sfShape* Shape, sfBool Enable);
func (self *Shape) EnableFill(enable bool) {
    C.sfShape_EnableFill(self.Cref, GoBool2SfBool(enable))
}

// Enable or disable drawing a shape outline.
// Outline is enabled by default
// param Shape : Shape to modify
// param Enable : True to enable, false to disable
// sfShape_EnableOutline(sfShape* Shape, sfBool Enable);
func (self *Shape) EnableOutline(enable bool) {
    C.sfShape_EnableOutline(self.Cref, GoBool2SfBool(enable))
}

// Change the width of a shape outline
// param Shape : Shape to modify
// param Width : New width
// sfShape_SetOutlineWidth(sfShape* Shape, float32 Width);
func (self *Shape) SetOutlineWidth(width float32) {
    C.sfShape_SetOutlineWidth(self.Cref, C.float(width))
}

// Get the width of a shape outline
// param Shape : Shape to read
// param return Current outline width
// float32 sfShape_GetOutlineWidth(sfShape* Shape);
func (self *Shape) GetOutlineWidth() float32 {
    return float32( C.sfShape_GetOutlineWidth(self.Cref))
}

// Get the number of points composing a shape
// param Shape : Shape to read
// return Total number of points
// unsigned int sfShape_GetNbPoints(sfShape* Shape);
func (self *Shape) GetNbPoints() uint {
	return uint( C.sfShape_GetNbPoints(self.Cref))
}

// Get a the position of a shape's point
// param Shape : Shape to read
// param Index : Index of the point to get
// param X : Variable to fill with the X coordinate of the point
// param Y : Variable to fill with the Y coordinate of the point
// sfShape_GetPointPosition(sfShape* Shape, unsigned int Index, float32* X, float32* Y);

// ??
//func (self *Shape) GetPointPosition(index uint, x float32*, y float32*) {
//    return C.sfShape_GetPointPosition(self.Cref, index, x, y)
//}


// Get a the color of a shape's point
// param Shape : Shape to read
// param Index : Index of the point to get
// return Color of the point
// sfColor sfShape_GetPointColor(sfShape* Shape, unsigned int Index);
func (self *Shape) GetPointColor(index uint) Color {
    return NewColor( C.sfShape_GetPointColor(self.Cref, C.uint(index)))
}

// Get a the outline color of a shape's point
// param Shape : Shape to read
// param Index : Index of the point to get
// return Outline color of the point
// sfColor sfShape_GetPointOutlineColor(sfShape* Shape, unsigned int Index);
func (self *Shape) GetPointOutlineColor(index uint) Color {
    return NewColor( C.sfShape_GetPointOutlineColor(self.Cref, C.uint(index)))
}

// Set a the position of a shape's point
// param Shape : Shape to read
// param Index : Index of the point to get
// param X : X coordinate of the point
// param Y : Y coordinate of the point
// sfShape_SetPointPosition(sfShape* Shape, unsigned int Index, float32 X, float32 Y);
func (self *Shape) SetPointPosition(index uint, x, y float32) {
    C.sfShape_SetPointPosition(self.Cref, C.uint(index), C.float(x), C.float(y))
}

// Set a the color of a shape's point
// param Shape : Shape to read
// param Index : Index of the point to get
// param Color : Color of the point
// sfShape_SetPointColor(sfShape* Shape, unsigned int Index, sfColor Color);
func (self *Shape) SetPointColor(index uint, color Color) {
    C.sfShape_SetPointColor(self.Cref, C.uint(index), color.Cref)
}

// Set a the outline color of a shape's point
// param Shape : Shape to read
// param Index : Index of the point to get
// param Color : Outline color of the point
// sfShape_SetPointOutlineColor(sfShape* Shape, unsigned int Index, sfColor Color);
func (self *Shape) SetPointOutlineColor(index uint, color Color) {
    C.sfShape_SetPointOutlineColor(self.Cref, C.uint(index), color.Cref)
}




// _Sprite_
// -------------------------------------------------------------------------------
// Create a new sprite
// return A new Sprite object, or NULL if it failed
// Sprite* Sprite_Create();
func CreateSprite() Sprite {
    return NewSprite( C.sfSprite_Create())
}

// Destroy an existing sprite
// param Sprite : Sprite to delete
// void Sprite_Destroy(sfSprite* Sprite);
func (self *Sprite) Destroy() {
    C.sfSprite_Destroy(self.Cref)
	self.Cref = nil
}

// Set the X position of a sprite
// param Sprite : Sprite to modify
// param X : New X coordinate
// void Sprite_SetX(sfSprite* Sprite, float32 X);
func (self *Sprite) SetX(x float32) {
    C.sfSprite_SetX(self.Cref, C.float(x))
}

// Set the T position of a sprite
// param Sprite : Sprite to modify
// param Y : New Y coordinate
// void Sprite_SetY(sfSprite* Sprite, float32 C.FLOAT(Y));
func (self *Sprite) SetY(y float32) {
    C.sfSprite_SetY(self.Cref, C.float(y))
}

// Set the position of a sprite
// param Sprite : Sprite to modify
// param X : New X coordinate
// param Y : New Y coordinate
// void Sprite_SetPosition(sfSprite* Sprite, float32 C.FLOAT(X), float32 C.FLOAT(Y));
func (self *Sprite) SetPosition(x float32, y float32) {
    C.sfSprite_SetPosition(self.Cref, C.float(x), C.float(y))
}

// Set the horizontal scale of a sprite
// param Sprite : Sprite to modify
// param Scale : New scale (must be strictly positive)
// void Sprite_SetScaleX(sfSprite* Sprite, float32 Scale);
func (self *Sprite) SetScaleX(scale float32) {
    C.sfSprite_SetScaleX(self.Cref, C.float(scale))
}

// Set the vertical scale of a sprite
// param Sprite : Sprite to modify
// param Scale : New scale (must be strictly positive)
// void Sprite_SetScaleY(sfSprite* Sprite, float32 Scale);
func (self *Sprite) SetScaleY(scale float32) {
    C.sfSprite_SetScaleY(self.Cref, C.float(scale))
}

// Set the scale of a sprite
// param Sprite : Sprite to modify
// param ScaleX : New horizontal scale (must be strictly positive)
// param ScaleY : New vertical scale (must be strictly positive)
// void Sprite_SetScale(sfSprite* Sprite, float32 ScaleC.FLOAT(X), float32 ScaleC.FLOAT(Y));
func (self *Sprite) SetScale(scaleX float32, scaleY float32) {
    C.sfSprite_SetScale(self.Cref, C.float(scaleX), C.float(scaleY))
}

// Set the orientation of a sprite
// param Sprite : Sprite to modify
// param Rotation : Angle of rotation, in degrees
// void Sprite_SetRotation(sfSprite* Sprite, float32 Rotation);
func (self *Sprite) SetRotation(rotation float32) {
    C.sfSprite_SetRotation(self.Cref, C.float(rotation))
}

// Set the center of a sprite, in coordinates relative to
// its left-top corner
// param Sprite : Sprite to modify
// param X : X coordinate of the center
// param Y : Y coordinate of the center
// void Sprite_SetCenter(sfSprite* Sprite, float32 C.FLOAT(X), float32 C.FLOAT(Y));
func (self *Sprite) SetCenter(x float32, y float32) {
    C.sfSprite_SetCenter(self.Cref, C.float(x), C.float(y))
}

// Set the color of a sprite
// param Sprite : Sprite to modify
// param Color : New color
// void Sprite_SetColor(sfSprite* Sprite, Color Color);
func (self *Sprite) SetColor(color Color) {
    C.sfSprite_SetColor(self.Cref, color.Cref)
}

// Set the blending mode for a sprite
// param Sprite : Sprite to modify
// param Mode : New blending mode
// void Sprite_SetBlendMode(sfSprite* Sprite, BlendMode Mode);
func (self *Sprite) SetBlendMode(mode BlendMode) {
    C.sfSprite_SetBlendMode(self.Cref, mode.Cref)
}

// Get the X position of a sprite
// param Sprite : Sprite to read
// return Current X position
// float32 Sprite_GetX(sfSprite* Sprite);
func (self *Sprite) GetX() float32 {
    return float32( C.sfSprite_GetX(self.Cref))
}

// Get the Y position of a sprite
// param Sprite : Sprite to read
// return Current Y position
// float32 Sprite_GetY(sfSprite* Sprite);
func (self *Sprite) GetY() float32 {
    return float32( C.sfSprite_GetY(self.Cref))
}

// Get the horizontal scale of a sprite
// param Sprite : Sprite to read
// return Current X scale factor (always positive)
// float32 Sprite_GetScaleX(sfSprite* Sprite);
func (self *Sprite) GetScaleX() float32 {
    return float32( C.sfSprite_GetScaleX(self.Cref))
}

// Get the vertical scale of a sprite
// param Sprite : Sprite to read
// return Current Y scale factor (always positive)
// float32 Sprite_GetScaleY(sfSprite* Sprite);
func (self *Sprite) GetScaleY() float32 {
    return float32( C.sfSprite_GetScaleY(self.Cref))
}

// Get the orientation of a sprite
// param Sprite : Sprite to read
// return Current rotation, in degrees
// float32 Sprite_GetRotation(sfSprite* Sprite);
func (self *Sprite) GetRotation() float32 {
    return float32( C.sfSprite_GetRotation(self.Cref))
}

// Get the X position of the center a sprite
// param Sprite : Sprite to read
// return Current X center
// float32 Sprite_GetCenterX(sfSprite* Sprite);
func (self *Sprite) GetCenterX() float32 {
    return float32( C.sfSprite_GetCenterX(self.Cref))
}

// Get the Y position of the center a sprite
// param Sprite : Sprite to read
// return Current Y center
// float32 Sprite_GetCenterY(sfSprite* Sprite);
func (self *Sprite) GetCenterY() float32 {
    return float32( C.sfSprite_GetCenterY(self.Cref))
}

// Get the color of a sprite
// param Sprite : Sprite to read
// return Current color
// Color Sprite_GetColor(sfSprite* Sprite);
func (self *Sprite) GetColor() Color {
    return NewColor( C.sfSprite_GetColor(self.Cref) )
}

// Get the current blending mode of a sprite
// param Sprite : Sprite to read
// return Current blending mode
// BlendMode Sprite_GetBlendMode(sfSprite* Sprite);
func (self *Sprite) GetBlendMode() BlendMode {
    return NewBlendMode( C.sfSprite_GetBlendMode(self.Cref) )
}

// Move a sprite
// param Sprite : Sprite to modify
// param OffsetX : Offset on the X axis
// param OffsetY : Offset on the Y axis
// void Sprite_Move(sfSprite* Sprite, float32 OffsetC.FLOAT(X), float32 OffsetC.FLOAT(Y));
func (self *Sprite) Move(offsetX float32, offsetY float32) {
    C.sfSprite_Move(self.Cref, C.float(offsetX), C.float(offsetY))
}

// Scale a sprite
// param Sprite : Sprite to modify
// param FactorX : Horizontal scaling factor (must be strictly positive)
// param FactorY : Vertical scaling factor (must be strictly positive)
// void Sprite_Scale(sfSprite* Sprite, float32 FactorC.FLOAT(X), float32 FactorC.FLOAT(Y));
func (self *Sprite) Scale(factorX float32, factorY float32) {
    C.sfSprite_Scale(self.Cref, C.float(factorX), C.float(factorY))
}

// Rotate a sprite
// param Sprite : Sprite to modify
// param Angle : Angle of rotation, in degrees
// void Sprite_Rotate(sfSprite* Sprite, float32 Angle);
func (self *Sprite) Rotate(angle float32) {
    C.sfSprite_Rotate(self.Cref, C.float(angle))
}

// Transform a point from global coordinates into the sprite's local coordinates
// (ie it applies the inverse of object's center, translation, rotation and scale to the point)
// param Sprite : Sprite object
// param PointX : X coordinate of the point to transform
// param PointY : Y coordinate of the point to transform
// param X : Value to fill with the X coordinate of the converted point
// param Y : Value to fill with the y coordinate of the converted point
// void Sprite_TransformToLocal(sfSprite* Sprite, float32 PointC.FLOAT(X), float32 PointY, float32* C.FLOAT(X), float32* C.FLOAT(Y));

//func (self *Sprite) TransformToLocal(pointX float32, pointY float32, x float32*, y float32*) {
//    return C.sfSprite_TransformToLocal(self.Cref, pointC.FLOAT(X), pointY, C.float(x), C.float(y))
//}


// Transform a point from the sprite's local coordinates into global coordinates
// (ie it applies the object's center, translation, rotation and scale to the point)
// param Sprite : Sprite object
// param PointX : X coordinate of the point to transform
// param PointY : Y coordinate of the point to transform
// param X : Value to fill with the X coordinate of the converted point
// param Y : Value to fill with the y coordinate of the converted point
// void Sprite_TransformToGlobal(sfSprite* Sprite, float32 PointC.FLOAT(X), float32 PointY, float32* C.FLOAT(X), float32* C.FLOAT(Y));

//func (self *Sprite) TransformToGlobal(pointX float32, pointY float32, x float32*, y float32*) {
//    return C.sfSprite_TransformToGlobal(self.Cref, pointC.FLOAT(X), pointY, C.float(x), C.float(y))
//}


// Change the image of a sprite
// param Sprite : Sprite to modify
// param Image : New image
// void Sprite_SetImage(sfSprite* Sprite, Image* Image);
func (self *Sprite) SetImage(image Image) {
    C.sfSprite_SetImage(self.Cref, image.Cref)
}

// Set the sub-rectangle of a sprite inside the source image
// param Sprite : Sprite to modify
// param SubRect : New sub-rectangle
// void Sprite_SetSubRect(sfSprite* Sprite, IntRect SubRect);
func (self *Sprite) SetSubRect(subRect IntRect) {
    C.sfSprite_SetSubRect(self.Cref, subRect.Cref)
}

// Resize a sprite (by changing its scale factors)
// param Sprite : Sprite to modify
// param Width : New width (must be strictly positive)
// param Height : New height (must be strictly positive)
// void Sprite_Resize(sfSprite* Sprite, float32 Width, float32 Height);
func (self *Sprite) Resize(width float32, height float32) {
    C.sfSprite_Resize(self.Cref, C.float(width), C.float(height))
}

// Flip a sprite horizontally
// param Sprite : Sprite to modify
// param Flipped : True to flip the sprite
// void Sprite_FlipX(sfSprite* Sprite, Bool Flipped);
func (self *Sprite) FlipX(flipped bool) {
    C.sfSprite_FlipX(self.Cref, GoBool2SfBool(flipped))
}

// Flip a sprite vertically
// param Sprite : Sprite to modify
// param Flipped : True to flip the sprite
// void Sprite_FlipY(sfSprite* Sprite, Bool Flipped);
func (self *Sprite) FlipY(flipped bool) {
    C.sfSprite_FlipY(self.Cref, GoBool2SfBool(flipped))
}

// Get the source image of a sprite
// param Sprite : Sprite to read
// return Pointer to the image (can be NULL)
// Image* Sprite_GetImage(sfSprite* Sprite);
func (self *Sprite) GetImage() Image {
    return NewImage( C.sfSprite_GetImage(self.Cref) )
}

// Get the sub-rectangle of a sprite inside the source image
// param Sprite : Sprite to read
// return Sub-rectangle
// IntRect Sprite_GetSubRect(sfSprite* Sprite);
func (self *Sprite) GetSubRect() IntRect {
    return NewIntRect( C.sfSprite_GetSubRect(self.Cref) )
}

// Get a sprite width
// param Sprite : Sprite to read
// return Width of the sprite
// float32 Sprite_GetWidth(sfSprite* Sprite);
func (self *Sprite) GetWidth() float32 {
    return float32( C.sfSprite_GetWidth(self.Cref))
}

// Get a sprite height
// param Sprite : Sprite to read
// return Height of the sprite
// float32 Sprite_GetHeight(sfSprite* Sprite);
func (self *Sprite) GetHeight() float32 {
    return float32( C.sfSprite_GetHeight(self.Cref))
}

// Get the color of a given pixel in a sprite
// param Sprite : Sprite to read
// param X : X coordinate of the pixel to get
// param Y : Y coordinate of the pixel to get
// return Color of pixel ( C.FLOAT(X), C.FLOAT(Y))
// Color Sprite_GetPixel(sfSprite* Sprite, unsigned int C.FLOAT(X), unsigned int C.FLOAT(Y));
func (self *Sprite) GetPixel(x uint, y uint) Color {
    return NewColor( C.sfSprite_GetPixel(self.Cref, C.uint(x), C.uint(y)) )
}




// _PostFX_
// -------------------------------------------------------------------------------
// Create a new post-fx from a file
// param Filename : File to load
// return A new PostFX object, or NULL if it failed
// PostFX* PostFX_CreateFromFile(const char* Filename);
func CreatePostFXFromFile(filename string) PostFX {
    return NewPostFX( C.sfPostFX_CreateFromFile( C.CString(filename)) )
}

// Create a new post-fx from an effect source code
// param Effect : Source code of the effect
// return A new PostFX object, or NULL if it failed
// PostFX* PostFX_CreateFromMemory(const char* Effect);
// ??
//func CreatePostFXFromMemory() PostFX {
//    return NewPostFX( C.sfPostFX_CreateFromMemory(self.Cref) )
//}


// Destroy an existing post-fx
// param PostFX : PostFX to delete
// void PostFX_Destroy(sfPostFX* PostFX);
func (self *PostFX) Destroy() {
    C.sfPostFX_Destroy(self.Cref)
}

// Change a parameter of a post-fx (1 float32)
// param PostFX : Post-effect to modify
// param Name : Parameter name in the effect
// param X : Value to assign
// void PostFX_SetParameter1(sfPostFX* PostFX, const char* Name, float32 X);
func (self *PostFX) SetParameter1(name string, x float32) {
    C.sfPostFX_SetParameter1(self.Cref, C.CString(name), C.float(x))
}

// Change a parameter of a post-fx (2 float32s)
// param PostFX : Post-effect to modify
// param Name : Parameter name in the effect
// param X, Y : Values to assign
// void PostFX_SetParameter2(sfPostFX* PostFX, const char* Name, float32 X, float32 Y);
func (self *PostFX) SetParameter2(name string, x, y float32) {
    C.sfPostFX_SetParameter2(self.Cref, C.CString(name), C.float(x), C.float(y))
}

// Change a parameter of a post-fx (3 float32s)
// param PostFX : Post-effect to modify
// param Name : Parameter name in the effect
// param X, Y, Z : Values to assign
// void PostFX_SetParameter3(sfPostFX* PostFX, const char* Name, float32 X, float32 Y, float32 Z);
func (self *PostFX) SetParameter3(name string, x, y, z float32) {
    C.sfPostFX_SetParameter3(self.Cref, C.CString(name), C.float(x), C.float(y), C.float(z))
}

// Change a parameter of a post-fx (4 float32s)
// param PostFX : Post-effect to modify
// param Name : Parameter name in the effect
// param X, Y, Z, W : Values to assign
// void PostFX_SetParameter4(sfPostFX* PostFX, const char* Name, float32 X, float32 Y, float32 Z, float32 W);
func (self *PostFX) SetParameter4(name string, x, y, z, w float32) {
    C.sfPostFX_SetParameter4(self.Cref, C.CString(name), C.float(x), C.float(y), C.float(z), C.float(w))
}

// Set a texture parameter in a post-fx
// param PostFX : Post-effect to modify
// param Name : Texture name in the effect
// param Texture : Image to set (pass NULL to use content of current framebuffer)
// void PostFX_SetTexture(sfPostFX* PostFX, const char* Name, Image* Texture);
func (self *PostFX) SetTexture(name string, texture Image) {
    C.sfPostFX_SetTexture(self.Cref, C.CString(name), texture.Cref)
}

// Tell whether or not the system supports post-effects
// return True if the system can use post-effects
// Bool PostFX_CanUsePostFX();
func CanUsePostFX() bool {
    return SfBool2GoBool( C.sfPostFX_CanUsePostFX() )
}





// _RenderWindow_
// -------------------------------------------------------------------------------

type VideoMode struct {
	Cref C.sfVideoMode
} 

func CreateVideoMode(height, width, bpp uint) VideoMode {
	return VideoMode{ C.sfVideoMode{ C.uint(height), C.uint(width), C.uint(bpp)}}
}

func NewWindowSettings( val C.sfWindowSettings ) WindowSettings{
	return WindowSettings{ val }
}



func CreateWindowSettings(depthBits, stencilBits, antialiasingLevel uint) WindowSettings {
	return WindowSettings{ C.sfWindowSettings{ C.uint(depthBits), C.uint(stencilBits), C.uint(antialiasingLevel)}}
}

// Construct a new renderwindow
// param Mode : Video mode to use
// param Title : Title of the window
// param Style : Window style
// param Params : Creation settings
// sfRenderWindow* sfRenderWindow_Create(sfVideoMode Mode, const string Title, unsigned long Style, sfWindowSettings Params);

func CreateRenderWindow(mode VideoMode, title string, style uint64, params WindowSettings) RenderWindow {
	return NewRenderWindow( C.sfRenderWindow_Create(mode.Cref, C.CString(title), C.ulong(style), params.Cref) )
}



 // Construct a renderwindow from an existing control
 // param Handle : Platform-specific handle of the control
 // param Params : Creation settings
 // sfRenderWindow* sfRenderWindow_CreateFromHandle(sfWindowHandle Handle, sfWindowSettings Params);
// func (self *RenderWindow) CreateFromHandle(params WindowSettings) RenderWindow {
// return C.sfRenderWindow_CreateFromHandle(self.Cref, params.Cref)
// }


// Destroy an existing renderwindow
// param RenderWindow : Renderwindow to destroy
// void sfRenderWindow_Destroy(sfRenderWindow* RenderWindow);
func (self *RenderWindow) Destroy() {
	C.sfRenderWindow_Destroy(self.Cref)
}

// Close a renderwindow (but doesn't destroy the internal data)
// param RenderWindow : Renderwindow to close
// void sfRenderWindow_Close(sfRenderWindow* RenderWindow);
func (self *RenderWindow) Close() {
	C.sfRenderWindow_Close(self.Cref)
}


// Tell whether or not a renderwindow is opened
// param RenderWindow : Renderwindow object
// sfBool sfRenderWindow_IsOpened(sfRenderWindow* RenderWindow);
func (self *RenderWindow) IsOpened() bool {
	return SfBool2GoBool( C.sfRenderWindow_IsOpened(self.Cref) )
}

// Get the width of the rendering region of a window
// param RenderWindow : Renderwindow object
// return Width in pixels
// unsigned int sfRenderWindow_GetWidth(sfRenderWindow* RenderWindow);
func (self *RenderWindow) GetWidth() uint {
	return uint( C.sfRenderWindow_GetWidth(self.Cref) )
}

// Get the height of the rendering region of a window
// param RenderWindow : Renderwindow object
// return Height in pixels
// unsigned int sfRenderWindow_GetHeight(sfRenderWindow* RenderWindow);
func (self *RenderWindow) GetHeight() uint {
	return uint( C.sfRenderWindow_GetHeight(self.Cref) )
}

// Get the creation settings of a window
// param RenderWindow : Renderwindow object
// return Settings used to create the window
// sfWindowSettings sfRenderWindow_GetSettings(sfRenderWindow* RenderWindow);
func (self *RenderWindow) GetSettings() WindowSettings {
	return NewWindowSettings( C.sfRenderWindow_GetSettings(self.Cref) )
}

// Get the event on top of events stack of a window, if any, and pop it
// param RenderWindow : Renderwindow object
// param Event : Event to fill, if any
// return sfTrue if an event was returned, sfFalse if events stack was empty
// sfBool sfRenderWindow_GetEvent(sfRenderWindow* RenderWindow, sfEvent* Event);

//func (self *RenderWindow) GetEvent(event Event) bool {
//    return SfBool2GoBool( C.sfRenderWindow_GetEvent(self.Cref, event.Cref) )
//}


// Enable / disable vertical synchronization on a window
// param RenderWindow : Renderwindow object
// param Enabled : sfTrue to enable v-sync, sfFalse to deactivate
// void sfRenderWindow_UseVerticalSync(sfRenderWindow* RenderWindow, sfBool Enabled);
func (self *RenderWindow) UseVerticalSync(enabled bool) {
    C.sfRenderWindow_UseVerticalSync(self.Cref, GoBool2SfBool(enabled))
}

// Show or hide the mouse cursor on a window
// param RenderWindow : RenderWindow object
// param Show : sfTrue to show, sfFalse to hide
// void sfRenderWindow_ShowMouseCursor(sfRenderWindow* RenderWindow, sfBool Show);
func (self *RenderWindow) ShowMouseCursor(show bool) {
    C.sfRenderWindow_ShowMouseCursor(self.Cref, GoBool2SfBool(show))
}

// Change the position of the mouse cursor on a window
// param RenderWindow : Renderwindow object
// param Left : Left coordinate of the cursor, relative to the window
// param Top : Top coordinate of the cursor, relative to the window
// void sfRenderWindow_SetCursorPosition(sfRenderWindow* RenderWindow, unsigned int Left, unsigned int Top);
func (self *RenderWindow) SetCursorPosition(left uint, top uint) {
    C.sfRenderWindow_SetCursorPosition(self.Cref, C.uint(left), C.uint(top))
}

// Change the position of a window on screen.
// Only works for top-level windows
// param RenderWindow : Renderwindow object
// param Left : Left position
// param Top : Top position
// void sfRenderWindow_SetPosition(sfRenderWindow* RenderWindow, int Left, int Top);
func (self *RenderWindow) SetPosition(left int, top int) {
    C.sfRenderWindow_SetPosition(self.Cref, C.int(left), C.int(top))
}

// Change the size of the rendering region of a window
// param RenderWindow : Renderwindow object
// param Width : New Width
// param Height : New Height
// void sfRenderWindow_SetSize(sfRenderWindow* RenderWindow, unsigned int Width, unsigned int Height);
func (self *RenderWindow) SetSize(width uint, height uint) {
    C.sfRenderWindow_SetSize(self.Cref, C.uint(width), C.uint(height))
}

// Show or hide a window
// param RenderWindow : Renderwindow object
// param State : sfTrue to show, sfFalse to hide
// void sfRenderWindow_Show(sfRenderWindow* RenderWindow, sfBool State);
func (self *RenderWindow) Show(state bool) {
    C.sfRenderWindow_Show(self.Cref, GoBool2SfBool(state))
}

// Enable or disable automatic key-repeat for keydown events.
// Automatic key-repeat is enabled by default
// param RenderWindow : Renderwindow object
// param Enabled : sfTrue to enable, sfFalse to disable
// void sfRenderWindow_EnableKeyRepeat(sfRenderWindow* RenderWindow, sfBool Enabled);
func (self *RenderWindow) EnableKeyRepeat(enabled bool) {
    C.sfRenderWindow_EnableKeyRepeat(self.Cref, GoBool2SfBool(enabled))
}

// Change the window's icon
// param RenderWindow : Renderwindow object
// param Width : Icon's width, in pixels
// param Height : Icon's height, in pixels
// param Pixels : Pointer to the pixels in memory, format must be RGBA 32 bits
// void sfRenderWindow_SetIcon(sfRenderWindow* RenderWindow, unsigned int Width, unsigned int Height, sfUint8* Pixels);
//func (self *RenderWindow) SetIcon(width uint, height uint, pixels Uint8) {
//    C.sfRenderWindow_SetIcon(self.Cref, width, height, pixels.Cref)
//}


// Activate or deactivate a window as the current target for rendering
// param RenderWindow : Renderwindow object
// param Active : sfTrue to activate, sfFalse to deactivate
// return True if operation was successful, false otherwise
// sfBool sfRenderWindow_SetActive(sfRenderWindow* RenderWindow, sfBool Active);
func (self *RenderWindow) SetActive(active bool) bool {
    return SfBool2GoBool(C.sfRenderWindow_SetActive(self.Cref, GoBool2SfBool(active)))
}

// Display a window on screen
// param RenderWindow : Renderwindow object
// void sfRenderWindow_Display(sfRenderWindow* RenderWindow);
func (self *RenderWindow) Display() {
    C.sfRenderWindow_Display(self.Cref)
}

// Get the input manager of a window
// param RenderWindow : Renderwindow object
// return Reference to the input
// sfInput* sfRenderWindow_GetInput(sfRenderWindow* RenderWindow);

//func (self *RenderWindow) GetInput() Input {
//    return C.sfRenderWindow_GetInput(self.Cref)
//}


// Limit the framerate to a maximum fixed frequency for a window
// param RenderWindow : Renderwindow object
// param Limit : Framerate limit, in frames per seconds (use 0 to disable limit)
// void sfRenderWindow_SetFramerateLimit(sfRenderWindow* RenderWindow, unsigned int Limit);
func (self *RenderWindow) SetFramerateLimit(limit uint) {
    C.sfRenderWindow_SetFramerateLimit(self.Cref, C.uint(limit))
}

// Get time elapsed since last frame of a window
// param RenderWindow : Renderwindow object
// return Time elapsed, in seconds
// float32 sfRenderWindow_GetFrameTime(sfRenderWindow* RenderWindow);
func (self *RenderWindow) GetFrameTime() float32 {
	return float32(12)
    //return float32(C.sfRenderWindow_GetFrameTime(self.Cref))
}

// Change the joystick threshold, ie. the value below which
// no move event will be generated
// param RenderWindow : Renderwindow object
// param Threshold : New threshold, in range [0, 100]
// void sfRenderWindow_SetJoystickThreshold(sfRenderWindow* RenderWindow, float32 Threshold);
func (self *RenderWindow) SetJoystickThreshold(threshold float32) {
    C.sfRenderWindow_SetJoystickThreshold(self.Cref, C.float(threshold))
}

// Draw something on a renderwindow
// param RenderWindow : Renderwindow to draw in
// param PostFX / Sprite / String / shape : Object to draw
// void sfRenderWindow_DrawPostFX(sfRenderWindow* RenderWindow, sfPostFX* PostFX);
func (self *RenderWindow) DrawPostFX(postFX PostFX) {
    C.sfRenderWindow_DrawPostFX(self.Cref, postFX.Cref)
}

// void sfRenderWindow_DrawSprite(sfRenderWindow* RenderWindow, sfSprite* Sprite);
func (self *RenderWindow) DrawSprite(sprite Sprite) {
    C.sfRenderWindow_DrawSprite(self.Cref, sprite.Cref)
}

// void sfRenderWindow_DrawShape (sfRenderWindow* RenderWindow, sfShape* Shape);
func (self *RenderWindow) DrawShape(shape Shape) {
    C.sfRenderWindow_DrawShape(self.Cref, shape.Cref)
}

// void sfRenderWindow_DrawString(sfRenderWindow* RenderWindow, sfString* String);
func (self *RenderWindow) DrawString(string String) {
    C.sfRenderWindow_DrawString(self.Cref, string.Cref)
}

// Save the content of a renderwindow to an image
// param RenderWindow : Renderwindow to capture
// return Image instance containing the contents of the screen
// sfImage* sfRenderWindow_Capture(sfRenderWindow* RenderWindow);
func (self *RenderWindow) Capture() Image {
    return NewImage( C.sfRenderWindow_Capture(self.Cref) )
}

// Clear the screen with the given color
// param RenderWindow : Renderwindow to modify
// param Color : Fill color
// void sfRenderWindow_Clear(sfRenderWindow* RenderWindow, sfColor Color);
func (self *RenderWindow) Clear(color C.sfColor_P) {
    C.sfRenderWindow_Clear(self.Cref, C.sfUnwraP(color))
}

// Change the current active view of a renderwindow
// param RenderWindow : Renderwindow to modify
// param NewView : Pointer to the new view
// void sfRenderWindow_SetView(sfRenderWindow* RenderWindow, sfView* View);
func (self *RenderWindow) SetView(view View) {
    C.sfRenderWindow_SetView(self.Cref, view.Cref)
}

// Get the current active view of a renderwindow
// param RenderWindow : Renderwindow
// return Current active view
// const sfView* sfRenderWindow_GetView(sfRenderWindow* RenderWindow);
// Get the default view of a renderwindow
// param RenderWindow : Renderwindow
// return Default view of the render window
// sfView* sfRenderWindow_GetDefaultView(sfRenderWindow* RenderWindow);
func (self *RenderWindow) GetDefaultView() View {
    return NewView( C.sfRenderWindow_GetDefaultView(self.Cref))
}

// Convert a point in window coordinates into view coordinates
// param RenderWindow : Target Renderwindow
// param WindowX : X coordinate of the point to convert, relative to the window
// param WindowY : Y coordinate of the point to convert, relative to the window
// param ViewX : Pointer to fill with the X coordinate of the converted point
// param ViewY : Pointer to fill with the Y coordinate of the converted point
// param TargetView : Target view to convert the point to (pass NULL to use the current view)
// void sfRenderWindow_ConvertCoords(sfRenderWindow* RenderWindow, unsigned int WindowX, unsigned int WindowY, float32* ViewX, float32* ViewY, sfView* TargetView);

// func (self *RenderWindow) ConvertCoords(windowX, windowY uint, viewX, viewY float32, targetView View) {
//     C.sfRenderWindow_ConvertCoords(self.Cref, C.uint(windowX), C.uint(windowY), C.float(viewX), C.float(viewY), targetView.Cref)
// }

// Tell SFML to preserve external OpenGL states, at the expense of
// more CPU charge. Use this function if you don't want SFML
// to mess up your own OpenGL states (if any).
// Don't enable state preservation if not needed, as it will allow
// SFML to do internal optimizations and improve performances.
// This parameter is false by default
// param RenderWindow : Target Renderwindow
// param Preserve : True to preserve OpenGL states, false to let SFML optimize
// void sfRenderWindow_PreserveOpenGLStates(sfRenderWindow* RenderWindow, sfBool Preserve);

func (self *RenderWindow) PreserveOpenGLStates(preserve bool) {
    C.sfRenderWindow_PreserveOpenGLStates(self.Cref, GoBool2SfBool(preserve))
}


