/*
package gfx

import . "sfml/graphics"
import "runtime"

// ------------------------------------------------------------------------
type Font struct {
	Cval *SfFont
}

// func SfFont_Create() SfFont {
func FontCreate() Font { 
	defer runtime.GC()	
	tmp := SfFont_Create() 
	return Font{ &tmp }
}

// func SfFont_GetDefaultFont() SfFont {
func GetDefaultFont() Font { // no leak
	defer runtime.GC()	
	tmp :=SfFont_GetDefaultFont() 
	return Font{ &tmp }
}

// func SfFont_CreateFromFile(arg1 string, arg2 uint, arg3 SfUint32) SfFont {
func FontCreateFromFile(filename string, charSize int) Font { 
	defer runtime.GC()
	// if anyone knows what charset is ...
	tmp := SfFont_CreateFromFile2(filename, charSize)
	fnt := Font{ &tmp }
	// or how to pass a nil as a NULL into C ...
	// had to inline this function in graphics.i	
	return fnt
}

// func SfFont_CreateFromMemory(arg1 string, arg2 int, arg3 uint, arg4 SfUint32) SfFont {
func FontCreateFromMemory(data string, sizeInBytes int, charSize uint) Font {
	defer runtime.GC()
	tmp := SfFont_CreateFromMemory2(data, sizeInBytes, charSize)
	fnt := Font{ &tmp }
	return fnt
}

// func SfFont_Destroy(arg1 SfFont) {
func (self *Font) Destroy(){
	SfFont_Destroy(*self.Cval)
}

// func SfFont_GetCharacterSize(arg1 SfFont) uint {
func (self *Font) GetCharacterSize() uint {
	return SfFont_GetCharacterSize(*self.Cval)
}


*/