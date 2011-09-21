// Created by cgo - DO NOT EDIT

//line gfx.go:1
package gfx

import (
	"runtime"
	"unsafe"
)

func NewColor(Val _Ctypedef_sfColor) Color {
	return Color{Val}
}

type Color struct {
	Cref _Ctypedef_sfColor
}

func Color_FromRGB_P(R, G, B uint8) _Ctypedef_sfColor_P {
	return _Cfunc_sfColor_FromRGB_P(_Ctypedef_sfUint8(R), _Ctypedef_sfUint8(G), _Ctypedef_sfUint8(B))
}

type WindowSettings struct {
	Cref _Ctypedef_sfWindowSettings
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

func NewFloatRect(val _Ctypedef_sfFloatRect) FloatRect {
	var tmp FloatRect
	tmp.Cref = val
	return tmp
}

type FloatRect struct {
	Cref _Ctypedef_sfFloatRect
}

func NewIntRect(val _Ctypedef_sfIntRect) IntRect {
	return IntRect{val}
}

type IntRect struct {
	Cref _Ctypedef_sfIntRect
}

func NewRenderWindow(val *[0]byte) RenderWindow {
	tmp := RenderWindow{val}
	runtime.SetFinalizer(&tmp, (*RenderWindow).Destroy)
	return tmp
}

type RenderWindow struct {
	Cref *[0]byte
}

type String struct {
	Cref *[0]byte
}

func NewString(val *[0]byte) String {
	tmp := String{val}
	runtime.SetFinalizer(&tmp, (*String).Destroy)
	return tmp
}

func NewBlendMode(val _Ctypedef_sfBlendMode) BlendMode {
	tmp := BlendMode{val}
	return tmp
}

type BlendMode struct {
	Cref _Ctypedef_sfBlendMode
}

func NewView(val *[0]byte) View {
	tmp := View{val}
	runtime.SetFinalizer(&tmp, (*View).Destroy)
	return tmp
}

type View struct {
	Cref *[0]byte
}

func NewShape(val *[0]byte) Shape {
	tmp := Shape{val}
	runtime.SetFinalizer(&tmp, (*Shape).Destroy)
	return tmp
}

type Shape struct {
	Cref *[0]byte
}

func NewSprite(val *[0]byte) Sprite {
	tmp := Sprite{val}
	runtime.SetFinalizer(&tmp, (*Sprite).Destroy)
	return tmp
}

type Sprite struct {
	Cref *[0]byte
}

func NewPostFX(val *[0]byte) PostFX {
	tmp := PostFX{val}
	runtime.SetFinalizer(&tmp, (*PostFX).Destroy)
	return tmp
}

type PostFX struct {
	Cref *[0]byte
}

func ColorFromRGB(R, G, B uint8) Color {

	hack := NewColor(_Cfunc_sfColor_FromRGB(0, 0, 0))
	return hack
}

func ColorFromRGBA(R, G, B, A uint8) Color {
	return NewColor(_Cfunc_sfColor_FromRGBA(_Ctypedef_sfUint8(R), _Ctypedef_sfUint8(G), _Ctypedef_sfUint8(B), _Ctypedef_sfUint8(A)))
}

func (self *Color) Add(other Color) Color {
	return NewColor(_Cfunc_sfColor_Add(self.Cref, other.Cref))
}

func (self *Color) Modulate(other Color) Color {
	return NewColor(_Cfunc_sfColor_Modulate(self.Cref, other.Cref))
}

func FontCreate() Font {
	fnt := Font{_Cfunc_sfFont_Create()}
	runtime.SetFinalizer(&fnt, (*Font).Destroy)
	return fnt
}

func FontCreateFromFile(path string, charSize int) Font {
	fnt := Font{_Cfunc_sfFont_CreateFromFile(_Cfunc_CString(path), _Ctype_uint(charSize), nil)}
	runtime.SetFinalizer(&fnt, (*Font).Destroy)
	return fnt
}

func (self *Font) Destroy() {
	_Cfunc_sfFont_Destroy(self.Cref)
	self.Cref = nil
}

func GetDefaultFont() Font {
	return Font{_Cfunc_sfFont_GetDefaultFont()}
}

func ImageCreate() Image {
	return NewImage(_Cfunc_sfImage_Create())
}

func ImageCreateFromColor(width, height uint, clr Color) Image {
	return NewImage(_Cfunc_sfImage_CreateFromColor(_Ctype_uint(width), _Ctype_uint(height), clr.Cref))
}

func ImageCreateFromPixels(width uint, height uint, data []uint8) Image {
	cdata := _Ctypedef_sfUint8(data[0])
	return NewImage(_Cfunc_sfImage_CreateFromPixels(_Ctype_uint(width), _Ctype_uint(height), &cdata))
}

func ImageCreateFromFile(filename string) Image {
	return NewImage(_Cfunc_sfImage_CreateFromFile(_Cfunc_CString(filename)))
}

func ImageCreateFromMemory(data []uint8, sizeInByes int) Image {
	cdata := _Ctype_char(data[0])
	img := Image{_Cfunc_sfImage_CreateFromMemory(&cdata, _Ctypedef_size_t(sizeInByes))}
	runtime.SetFinalizer(&img, (*Image).Destroy)
	return img

}

func (self *Image) Destroy() {
	_Cfunc_sfImage_Destroy(self.Cref)
	self.Cref = nil
}

func SfBool2GoBool(cVal _Ctypedef_sfBool) bool {
	return int(cVal) == 1
}

func GoBool2SfBool(goVal bool) _Ctypedef_sfBool {
	if goVal {
		return _Ctypedef_sfBool(1)
	}
	return _Ctypedef_sfBool(0)
}

func (self *Image) SaveToFile(filename string) bool {
	return SfBool2GoBool(_Cfunc_sfImage_SaveToFile(self.Cref, _Cfunc_CString(filename)))
}

func (self *Image) CreateMaskFromColor(colorKey Color, alpha uint8) {
	_Cfunc_sfImage_CreateMaskFromColor(self.Cref, colorKey.Cref, _Ctypedef_sfUint8(alpha))
}

func (self *Image) Copy(source Image, destX, destY uint, sourceRect IntRect) {
	_Cfunc_sfImage_Copy(self.Cref, source.Cref, _Ctype_uint(destX), _Ctype_uint(destY), sourceRect.Cref)
}

func (self *Image) CopyScreen(window RenderWindow, sourceRect IntRect) bool {
	return SfBool2GoBool(_Cfunc_sfImage_CopyScreen(self.Cref, window.Cref, sourceRect.Cref))
}

func (self *Image) SetPixel(x, y uint, color Color) {
	_Cfunc_sfImage_SetPixel(self.Cref, _Ctype_uint(x), _Ctype_uint(y), color.Cref)
}

func (self *Image) GetPixel(x, y uint) Color {
	return Color{_Cfunc_sfImage_GetPixel(self.Cref, _Ctype_uint(x), _Ctype_uint(y))}
}

func (self *Image) Bind() {
	_Cfunc_sfImage_Bind(self.Cref)
}

func (self *Image) SetSmooth(smooth bool) {
	_Cfunc_sfImage_SetSmooth(self.Cref, GoBool2SfBool(smooth))
}

func (self *Image) GetWidth() uint {
	return uint(_Cfunc_sfImage_GetWidth(self.Cref))
}

func (self *Image) GetHeight() uint {
	return uint(_Cfunc_sfImage_GetHeight(self.Cref))
}

func (self *Image) IsSmooth() bool {
	return SfBool2GoBool(_Cfunc_sfImage_IsSmooth(self.Cref))
}

func (self *FloatRect) Offset(offsetX float32, offsetY float32) {
	_Cfunc_sfFloatRect_Offset(&(self.Cref), _Ctype_float(offsetX), _Ctype_float(offsetY))
}

func (self *IntRect) Offset(offsetX int, offsetY int) {
	_Cfunc_sfIntRect_Offset(&(self.Cref), _Ctype_int(offsetX), _Ctype_int(offsetY))
}

func (self *FloatRect) Contains(x float32, y float32) bool {
	return SfBool2GoBool(_Cfunc_sfFloatRect_Contains(&self.Cref, _Ctype_float(x), _Ctype_float(y)))
}

func (self *IntRect) Contains(x int, y int) bool {
	return SfBool2GoBool(_Cfunc_sfIntRect_Contains(&self.Cref, _Ctype_int(x), _Ctype_int(y)))
}

func (self *FloatRect) Intersects(rect2 FloatRect, overlappingRect FloatRect) bool {
	return SfBool2GoBool(_Cfunc_sfFloatRect_Intersects(&(self.Cref),
		&(rect2.Cref), &(overlappingRect.Cref)))
}

func (self *IntRect) Intersects(rect2 IntRect, overlappingRect IntRect) bool {
	return SfBool2GoBool(_Cfunc_sfIntRect_Intersects(&(self.Cref),
		&(rect2.Cref),
		&(overlappingRect.Cref)))
}

func StringCreate() String {
	return NewString(_Cfunc_sfString_Create())
}

func (self *String) Destroy() {
	_Cfunc_sfString_Destroy(self.Cref)
	self.Cref = nil
}

func (self *String) SetX(x float32) {
	_Cfunc_sfString_SetX(self.Cref, _Ctype_float(x))
}

func (self *String) SetY(y float32) {
	_Cfunc_sfString_SetY(self.Cref, _Ctype_float(y))
}

func (self *String) SetPosition(left float32, top float32) {
	_Cfunc_sfString_SetPosition(self.Cref, _Ctype_float(left), _Ctype_float(top))
}

func (self *String) SetScaleX(scale float32) {
	_Cfunc_sfString_SetScaleX(self.Cref, _Ctype_float(scale))
}

func (self *String) SetScaleY(scale float32) {
	_Cfunc_sfString_SetScaleY(self.Cref, _Ctype_float(scale))
}

func (self *String) SetScale(scaleX float32, scaleY float32) {
	_Cfunc_sfString_SetScale(self.Cref, _Ctype_float(scaleX), _Ctype_float(scaleY))
}

func (self *String) SetRotation(rotation float32) {
	_Cfunc_sfString_SetRotation(self.Cref, _Ctype_float(rotation))
}

func (self *String) SetCenter(x float32, y float32) {
	_Cfunc_sfString_SetCenter(self.Cref, _Ctype_float(x), _Ctype_float(y))
}

func (self *String) SetColor(color Color) {
	_Cfunc_sfString_SetColor(self.Cref, color.Cref)
}

func (self *String) SetBlendMode(mode BlendMode) {
	_Cfunc_sfString_SetBlendMode(self.Cref, mode.Cref)
}

func (self *String) GetX() float32 {
	return float32(_Cfunc_sfString_GetX(self.Cref))
}

func (self *String) GetY() float32 {
	return float32(_Cfunc_sfString_GetY(self.Cref))
}

func (self *String) GetScaleX() float32 {
	return float32(_Cfunc_sfString_GetScaleX(self.Cref))
}

func (self *String) GetScaleY() float32 {
	return float32(_Cfunc_sfString_GetScaleY(self.Cref))
}

func (self *String) GetRotation() float32 {
	return float32(_Cfunc_sfString_GetRotation(self.Cref))
}

func (self *String) GetCenterX() float32 {
	return float32(_Cfunc_sfString_GetCenterX(self.Cref))
}

func (self *String) GetCenterY() float32 {
	return float32(_Cfunc_sfString_GetCenterY(self.Cref))
}

func (self *String) GetColor() Color {

	return Color{_Cfunc_sfString_GetColor(self.Cref)}
}

func (self *String) GetBlendMode() BlendMode {
	tmp := _Cfunc_sfString_GetBlendMode(self.Cref)
	return BlendMode{tmp}
}

func (self *String) Move(offsetX float32, offsetY float32) {
	_Cfunc_sfString_Move(self.Cref, _Ctype_float(offsetX), _Ctype_float(offsetY))
}

func (self *String) Scale(factorX float32, factorY float32) {
	_Cfunc_sfString_Scale(self.Cref, _Ctype_float(factorX), _Ctype_float(factorY))
}

func (self *String) Rotate(angle float32) {
	_Cfunc_sfString_Rotate(self.Cref, _Ctype_float(angle))
}

func (self *String) TransformToLocal(pointX float32, pointY float32, x *float32, y *float32) {
	xPtr := _Ctype_float(*x)
	yPtr := _Ctype_float(*y)
	_Cfunc_sfString_TransformToLocal(self.Cref, _Ctype_float(pointX), _Ctype_float(pointY), &xPtr, &yPtr)
}

func (self *String) TransformToGlobal(pointX float32, pointY float32, x *float32, y *float32) {
	xPtr := _Ctype_float(*x)
	yPtr := _Ctype_float(*y)
	_Cfunc_sfString_TransformToGlobal(self.Cref, _Ctype_float(pointX), _Ctype_float(pointY), &xPtr, &yPtr)
}

func (self *String) SetText(text string) {
	_Cfunc_sfString_SetText(self.Cref, _Cfunc_CString(text))
}

func (self *String) SetUnicodeText(text *uint32) {
	ctext := _Ctypedef_sfUint32(*text)
	_Cfunc_sfString_SetUnicodeText(self.Cref, &ctext)
}

func (self *String) SetFont(font Font) {
	_Cfunc_sfString_SetFont(self.Cref, font.Cref)
}

func (self *String) SetSize(size float32) {
	_Cfunc_sfString_SetSize(self.Cref, _Ctype_float(size))
}

func (self *String) SetStyle(style uint64) {
	_Cfunc_sfString_SetStyle(self.Cref, _Ctype_ulong(style))
}

const MaxString int = 65536

func (self *String) GetUnicodeText() string {
	ptr := _Cfunc_sfString_GetUnicodeText(self.Cref)
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

func (self *String) GetFont() Font {
	return NewFont(_Cfunc_sfString_GetFont(self.Cref))
}

func (self *String) GetSize() float32 {
	return float32(_Cfunc_sfString_GetSize(self.Cref))
}

func (self *String) GetRect() FloatRect {

	return NewFloatRect(_Cfunc_sfString_GetRect(self.Cref))
}

func CreateView() View {
	return NewView(_Cfunc_sfView_Create())
}

func CreateViewFromRect(rect FloatRect) View {
	return NewView(_Cfunc_sfView_CreateFromRect(rect.Cref))
}

func (self *View) Destroy() {
	_Cfunc_sfView_Destroy(self.Cref)
	self.Cref = nil
}

func (self *View) SetCenter(x, y float32) {
	_Cfunc_sfView_SetCenter(self.Cref, _Ctype_float(x), _Ctype_float(y))
}

func (self *View) SetHalfSize(halfWidth float32, halfHeight float32) {
	_Cfunc_sfView_SetHalfSize(self.Cref, _Ctype_float(halfWidth), _Ctype_float(halfHeight))
}

func (self *View) SetFromRect(viewRect FloatRect) {
	_Cfunc_sfView_SetFromRect(self.Cref, viewRect.Cref)
}

func (self *View) GetCenterX() float32 {
	return float32(_Cfunc_sfView_GetCenterX(self.Cref))
}

func (self *View) GetCenterY() float32 {
	return float32(_Cfunc_sfView_GetCenterY(self.Cref))
}

func (self *View) GetHalfSizeX() float32 {
	return float32(_Cfunc_sfView_GetHalfSizeX(self.Cref))
}

func (self *View) GetHalfSizeY() float32 {
	return float32(_Cfunc_sfView_GetHalfSizeY(self.Cref))
}

func (self *View) GetRect() FloatRect {
	return NewFloatRect(_Cfunc_sfView_GetRect(self.Cref))
}

func (self *View) Move(offsetX, offsetY float32) {
	_Cfunc_sfView_Move(self.Cref, _Ctype_float(offsetX), _Ctype_float(offsetY))
}

func (self *View) Zoom(factor float32) {
	_Cfunc_sfView_Zoom(self.Cref, _Ctype_float(factor))
}

func CreateShape() Shape {
	return NewShape(_Cfunc_sfShape_Create())
}

func CreateLine(p1Y, p2X, p2Y, thickness float32, col Color, outline float32, outlineCol Color) Shape {

	return NewShape(_Cfunc_sfShape_CreateLine(_Ctype_float(p2X), _Ctype_float(p1Y), _Ctype_float(p2X), _Ctype_float(p2Y), _Ctype_float(thickness), col.Cref, _Ctype_float(outline), outlineCol.Cref))
}

func CreateRectangle(p1X, p1Y, p2X, p2Y float32, col Color, outline float32, outlineCol Color) Shape {
	return NewShape(_Cfunc_sfShape_CreateRectangle(_Ctype_float(p1X), _Ctype_float(p1Y), _Ctype_float(p2X), _Ctype_float(p2Y), col.Cref, _Ctype_float(outline), outlineCol.Cref))
}

func CreateCircle(x, y, radius float32, col Color, outline float32, outlineCol Color) Shape {
	return NewShape(_Cfunc_sfShape_CreateCircle(_Ctype_float(x), _Ctype_float(y), _Ctype_float(radius), col.Cref, _Ctype_float(outline), outlineCol.Cref))
}

func (self *Shape) Destroy() {
	_Cfunc_sfShape_Destroy(self.Cref)
}

func (self *Shape) SetX(x float32) {
	_Cfunc_sfShape_SetX(self.Cref, _Ctype_float(x))
}

func (self *Shape) SetY(y float32) {
	_Cfunc_sfShape_SetY(self.Cref, _Ctype_float(y))
}

func (self *Shape) SetPosition(x, y float32) {
	_Cfunc_sfShape_SetPosition(self.Cref, _Ctype_float(x), _Ctype_float(y))
}

func (self *Shape) SetScaleX(scale float32) {
	_Cfunc_sfShape_SetScaleX(self.Cref, _Ctype_float(scale))
}

func (self *Shape) SetScaleY(scale float32) {
	_Cfunc_sfShape_SetScaleY(self.Cref, _Ctype_float(scale))
}

func (self *Shape) SetScale(scaleX, scaleY float32) {
	_Cfunc_sfShape_SetScale(self.Cref, _Ctype_float(scaleX), _Ctype_float(scaleY))
}

func (self *Shape) SetRotation(rotation float32) {
	_Cfunc_sfShape_SetRotation(self.Cref, _Ctype_float(rotation))
}

func (self *Shape) SetCenter(x float32, y float32) {
	_Cfunc_sfShape_SetCenter(self.Cref, _Ctype_float(x), _Ctype_float(y))
}

func (self *Shape) SetColor(color Color) {
	_Cfunc_sfShape_SetColor(self.Cref, color.Cref)
}

func (self *Shape) SetBlendMode(mode BlendMode) {
	_Cfunc_sfShape_SetBlendMode(self.Cref, mode.Cref)
}

func (self *Shape) GetX() float32 {
	return float32(_Cfunc_sfShape_GetX(self.Cref))
}

func (self *Shape) GetY() float32 {
	return float32(_Cfunc_sfShape_GetY(self.Cref))
}

func (self *Shape) GetScaleX() float32 {
	return float32(_Cfunc_sfShape_GetScaleX(self.Cref))
}

func (self *Shape) GetScaleY() float32 {
	return float32(_Cfunc_sfShape_GetScaleY(self.Cref))
}

func (self *Shape) GetRotation() float32 {
	return float32(_Cfunc_sfShape_GetRotation(self.Cref))
}

func (self *Shape) GetCenterX() float32 {
	return float32(_Cfunc_sfShape_GetCenterX(self.Cref))
}

func (self *Shape) GetCenterY() float32 {
	return float32(_Cfunc_sfShape_GetCenterY(self.Cref))
}

func (self *Shape) GetColor() Color {
	return NewColor(_Cfunc_sfShape_GetColor(self.Cref))
}

func (self *Shape) GetBlendMode() BlendMode {
	return NewBlendMode(_Cfunc_sfShape_GetBlendMode(self.Cref))
}

func (self *Shape) Move(offsetX float32, offsetY float32) {
	_Cfunc_sfShape_Move(self.Cref, _Ctype_float(offsetX), _Ctype_float(offsetY))
}

func (self *Shape) Scale(factorX float32, factorY float32) {
	_Cfunc_sfShape_Scale(self.Cref, _Ctype_float(factorX), _Ctype_float(factorY))
}

func (self *Shape) Rotate(angle float32) {
	_Cfunc_sfShape_Rotate(self.Cref, _Ctype_float(angle))
}

func (self *Shape) AddPoint(x, y float32, col Color, outlineCol Color) {
	_Cfunc_sfShape_AddPoint(self.Cref, _Ctype_float(x), _Ctype_float(y), col.Cref, outlineCol.Cref)
}

func (self *Shape) EnableFill(enable bool) {
	_Cfunc_sfShape_EnableFill(self.Cref, GoBool2SfBool(enable))
}

func (self *Shape) EnableOutline(enable bool) {
	_Cfunc_sfShape_EnableOutline(self.Cref, GoBool2SfBool(enable))
}

func (self *Shape) SetOutlineWidth(width float32) {
	_Cfunc_sfShape_SetOutlineWidth(self.Cref, _Ctype_float(width))
}

func (self *Shape) GetOutlineWidth() float32 {
	return float32(_Cfunc_sfShape_GetOutlineWidth(self.Cref))
}

func (self *Shape) GetNbPoints() uint {
	return uint(_Cfunc_sfShape_GetNbPoints(self.Cref))
}

func (self *Shape) GetPointColor(index uint) Color {
	return NewColor(_Cfunc_sfShape_GetPointColor(self.Cref, _Ctype_uint(index)))
}

func (self *Shape) GetPointOutlineColor(index uint) Color {
	return NewColor(_Cfunc_sfShape_GetPointOutlineColor(self.Cref, _Ctype_uint(index)))
}

func (self *Shape) SetPointPosition(index uint, x, y float32) {
	_Cfunc_sfShape_SetPointPosition(self.Cref, _Ctype_uint(index), _Ctype_float(x), _Ctype_float(y))
}

func (self *Shape) SetPointColor(index uint, color Color) {
	_Cfunc_sfShape_SetPointColor(self.Cref, _Ctype_uint(index), color.Cref)
}

func (self *Shape) SetPointOutlineColor(index uint, color Color) {
	_Cfunc_sfShape_SetPointOutlineColor(self.Cref, _Ctype_uint(index), color.Cref)
}

func CreateSprite() Sprite {
	return NewSprite(_Cfunc_sfSprite_Create())
}

func (self *Sprite) Destroy() {
	_Cfunc_sfSprite_Destroy(self.Cref)
	self.Cref = nil
}

func (self *Sprite) SetX(x float32) {
	_Cfunc_sfSprite_SetX(self.Cref, _Ctype_float(x))
}

func (self *Sprite) SetY(y float32) {
	_Cfunc_sfSprite_SetY(self.Cref, _Ctype_float(y))
}

func (self *Sprite) SetPosition(x float32, y float32) {
	_Cfunc_sfSprite_SetPosition(self.Cref, _Ctype_float(x), _Ctype_float(y))
}

func (self *Sprite) SetScaleX(scale float32) {
	_Cfunc_sfSprite_SetScaleX(self.Cref, _Ctype_float(scale))
}

func (self *Sprite) SetScaleY(scale float32) {
	_Cfunc_sfSprite_SetScaleY(self.Cref, _Ctype_float(scale))
}

func (self *Sprite) SetScale(scaleX float32, scaleY float32) {
	_Cfunc_sfSprite_SetScale(self.Cref, _Ctype_float(scaleX), _Ctype_float(scaleY))
}

func (self *Sprite) SetRotation(rotation float32) {
	_Cfunc_sfSprite_SetRotation(self.Cref, _Ctype_float(rotation))
}

func (self *Sprite) SetCenter(x float32, y float32) {
	_Cfunc_sfSprite_SetCenter(self.Cref, _Ctype_float(x), _Ctype_float(y))
}

func (self *Sprite) SetColor(color Color) {
	_Cfunc_sfSprite_SetColor(self.Cref, color.Cref)
}

func (self *Sprite) SetBlendMode(mode BlendMode) {
	_Cfunc_sfSprite_SetBlendMode(self.Cref, mode.Cref)
}

func (self *Sprite) GetX() float32 {
	return float32(_Cfunc_sfSprite_GetX(self.Cref))
}

func (self *Sprite) GetY() float32 {
	return float32(_Cfunc_sfSprite_GetY(self.Cref))
}

func (self *Sprite) GetScaleX() float32 {
	return float32(_Cfunc_sfSprite_GetScaleX(self.Cref))
}

func (self *Sprite) GetScaleY() float32 {
	return float32(_Cfunc_sfSprite_GetScaleY(self.Cref))
}

func (self *Sprite) GetRotation() float32 {
	return float32(_Cfunc_sfSprite_GetRotation(self.Cref))
}

func (self *Sprite) GetCenterX() float32 {
	return float32(_Cfunc_sfSprite_GetCenterX(self.Cref))
}

func (self *Sprite) GetCenterY() float32 {
	return float32(_Cfunc_sfSprite_GetCenterY(self.Cref))
}

func (self *Sprite) GetColor() Color {
	return NewColor(_Cfunc_sfSprite_GetColor(self.Cref))
}

func (self *Sprite) GetBlendMode() BlendMode {
	return NewBlendMode(_Cfunc_sfSprite_GetBlendMode(self.Cref))
}

func (self *Sprite) Move(offsetX float32, offsetY float32) {
	_Cfunc_sfSprite_Move(self.Cref, _Ctype_float(offsetX), _Ctype_float(offsetY))
}

func (self *Sprite) Scale(factorX float32, factorY float32) {
	_Cfunc_sfSprite_Scale(self.Cref, _Ctype_float(factorX), _Ctype_float(factorY))
}

func (self *Sprite) Rotate(angle float32) {
	_Cfunc_sfSprite_Rotate(self.Cref, _Ctype_float(angle))
}

func (self *Sprite) SetImage(image Image) {
	_Cfunc_sfSprite_SetImage(self.Cref, image.Cref)
}

func (self *Sprite) SetSubRect(subRect IntRect) {
	_Cfunc_sfSprite_SetSubRect(self.Cref, subRect.Cref)
}

func (self *Sprite) Resize(width float32, height float32) {
	_Cfunc_sfSprite_Resize(self.Cref, _Ctype_float(width), _Ctype_float(height))
}

func (self *Sprite) FlipX(flipped bool) {
	_Cfunc_sfSprite_FlipX(self.Cref, GoBool2SfBool(flipped))
}

func (self *Sprite) FlipY(flipped bool) {
	_Cfunc_sfSprite_FlipY(self.Cref, GoBool2SfBool(flipped))
}

func (self *Sprite) GetImage() Image {
	return NewImage(_Cfunc_sfSprite_GetImage(self.Cref))
}

func (self *Sprite) GetSubRect() IntRect {
	return NewIntRect(_Cfunc_sfSprite_GetSubRect(self.Cref))
}

func (self *Sprite) GetWidth() float32 {
	return float32(_Cfunc_sfSprite_GetWidth(self.Cref))
}

func (self *Sprite) GetHeight() float32 {
	return float32(_Cfunc_sfSprite_GetHeight(self.Cref))
}

func (self *Sprite) GetPixel(x uint, y uint) Color {
	return NewColor(_Cfunc_sfSprite_GetPixel(self.Cref, _Ctype_uint(x), _Ctype_uint(y)))
}

func CreatePostFXFromFile(filename string) PostFX {
	return NewPostFX(_Cfunc_sfPostFX_CreateFromFile(_Cfunc_CString(filename)))
}

func (self *PostFX) Destroy() {
	_Cfunc_sfPostFX_Destroy(self.Cref)
}

func (self *PostFX) SetParameter1(name string, x float32) {
	_Cfunc_sfPostFX_SetParameter1(self.Cref, _Cfunc_CString(name), _Ctype_float(x))
}

func (self *PostFX) SetParameter2(name string, x, y float32) {
	_Cfunc_sfPostFX_SetParameter2(self.Cref, _Cfunc_CString(name), _Ctype_float(x), _Ctype_float(y))
}

func (self *PostFX) SetParameter3(name string, x, y, z float32) {
	_Cfunc_sfPostFX_SetParameter3(self.Cref, _Cfunc_CString(name), _Ctype_float(x), _Ctype_float(y), _Ctype_float(z))
}

func (self *PostFX) SetParameter4(name string, x, y, z, w float32) {
	_Cfunc_sfPostFX_SetParameter4(self.Cref, _Cfunc_CString(name), _Ctype_float(x), _Ctype_float(y), _Ctype_float(z), _Ctype_float(w))
}

func (self *PostFX) SetTexture(name string, texture Image) {
	_Cfunc_sfPostFX_SetTexture(self.Cref, _Cfunc_CString(name), texture.Cref)
}

func CanUsePostFX() bool {
	return SfBool2GoBool(_Cfunc_sfPostFX_CanUsePostFX())
}

type VideoMode struct {
	Cref _Ctypedef_sfVideoMode
}

func CreateVideoMode(height, width, bpp uint) VideoMode {
	return VideoMode{_Ctypedef_sfVideoMode{_Ctype_uint(height), _Ctype_uint(width), _Ctype_uint(bpp)}}
}

func NewWindowSettings(val _Ctypedef_sfWindowSettings) WindowSettings {
	return WindowSettings{val}
}

func CreateWindowSettings(depthBits, stencilBits, antialiasingLevel uint) WindowSettings {
	return WindowSettings{_Ctypedef_sfWindowSettings{_Ctype_uint(depthBits), _Ctype_uint(stencilBits), _Ctype_uint(antialiasingLevel)}}
}

func CreateRenderWindow(mode VideoMode, title string, style uint64, params WindowSettings) RenderWindow {
	return NewRenderWindow(_Cfunc_sfRenderWindow_Create(mode.Cref, _Cfunc_CString(title), _Ctype_ulong(style), params.Cref))
}

func (self *RenderWindow) Destroy() {
	_Cfunc_sfRenderWindow_Destroy(self.Cref)
}

func (self *RenderWindow) Close() {
	_Cfunc_sfRenderWindow_Close(self.Cref)
}

func (self *RenderWindow) IsOpened() bool {
	return SfBool2GoBool(_Cfunc_sfRenderWindow_IsOpened(self.Cref))
}

func (self *RenderWindow) GetWidth() uint {
	return uint(_Cfunc_sfRenderWindow_GetWidth(self.Cref))
}

func (self *RenderWindow) GetHeight() uint {
	return uint(_Cfunc_sfRenderWindow_GetHeight(self.Cref))
}

func (self *RenderWindow) GetSettings() WindowSettings {
	return NewWindowSettings(_Cfunc_sfRenderWindow_GetSettings(self.Cref))
}

func (self *RenderWindow) UseVerticalSync(enabled bool) {
	_Cfunc_sfRenderWindow_UseVerticalSync(self.Cref, GoBool2SfBool(enabled))
}

func (self *RenderWindow) ShowMouseCursor(show bool) {
	_Cfunc_sfRenderWindow_ShowMouseCursor(self.Cref, GoBool2SfBool(show))
}

func (self *RenderWindow) SetCursorPosition(left uint, top uint) {
	_Cfunc_sfRenderWindow_SetCursorPosition(self.Cref, _Ctype_uint(left), _Ctype_uint(top))
}

func (self *RenderWindow) SetPosition(left int, top int) {
	_Cfunc_sfRenderWindow_SetPosition(self.Cref, _Ctype_int(left), _Ctype_int(top))
}

func (self *RenderWindow) SetSize(width uint, height uint) {
	_Cfunc_sfRenderWindow_SetSize(self.Cref, _Ctype_uint(width), _Ctype_uint(height))
}

func (self *RenderWindow) Show(state bool) {
	_Cfunc_sfRenderWindow_Show(self.Cref, GoBool2SfBool(state))
}

func (self *RenderWindow) EnableKeyRepeat(enabled bool) {
	_Cfunc_sfRenderWindow_EnableKeyRepeat(self.Cref, GoBool2SfBool(enabled))
}

func (self *RenderWindow) SetActive(active bool) bool {
	return SfBool2GoBool(_Cfunc_sfRenderWindow_SetActive(self.Cref, GoBool2SfBool(active)))
}

func (self *RenderWindow) Display() {
	_Cfunc_sfRenderWindow_Display(self.Cref)
}

func (self *RenderWindow) SetFramerateLimit(limit uint) {
	_Cfunc_sfRenderWindow_SetFramerateLimit(self.Cref, _Ctype_uint(limit))
}

func (self *RenderWindow) GetFrameTime() float32 {
	return float32(12)

}

func (self *RenderWindow) SetJoystickThreshold(threshold float32) {
	_Cfunc_sfRenderWindow_SetJoystickThreshold(self.Cref, _Ctype_float(threshold))
}

func (self *RenderWindow) DrawPostFX(postFX PostFX) {
	_Cfunc_sfRenderWindow_DrawPostFX(self.Cref, postFX.Cref)
}

func (self *RenderWindow) DrawSprite(sprite Sprite) {
	_Cfunc_sfRenderWindow_DrawSprite(self.Cref, sprite.Cref)
}

func (self *RenderWindow) DrawShape(shape Shape) {
	_Cfunc_sfRenderWindow_DrawShape(self.Cref, shape.Cref)
}

func (self *RenderWindow) DrawString(string String) {
	_Cfunc_sfRenderWindow_DrawString(self.Cref, string.Cref)
}

func (self *RenderWindow) Capture() Image {
	return NewImage(_Cfunc_sfRenderWindow_Capture(self.Cref))
}

func (self *RenderWindow) Clear(color _Ctypedef_sfColor_P) {
	_Cfunc_sfRenderWindow_Clear(self.Cref, _Cfunc_sfUnwraP(color))
}

func (self *RenderWindow) SetView(view View) {
	_Cfunc_sfRenderWindow_SetView(self.Cref, view.Cref)
}

func (self *RenderWindow) GetDefaultView() View {
	return NewView(_Cfunc_sfRenderWindow_GetDefaultView(self.Cref))
}

func (self *RenderWindow) PreserveOpenGLStates(preserve bool) {
	_Cfunc_sfRenderWindow_PreserveOpenGLStates(self.Cref, GoBool2SfBool(preserve))
}
