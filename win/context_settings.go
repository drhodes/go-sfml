package win

/*
 #cgo LDFLAGS:-lcsfml-window
 #include <SFML/Window/Window.h>

 unsigned int _GetDepthBits(sfContextSettings cs) {
     return cs.depthBits;
 }
 unsigned int _GetStencilBits(sfContextSettings cs) {
     return cs.stencilBits;
 }
 unsigned int _GetAntialiasingLevel(sfContextSettings cs) {
     return cs.antialiasingLevel;
 }
 unsigned int _GetMajorVersion(sfContextSettings cs) {
     return cs.majorVersion;
 }
 unsigned int _GetMinorVersion(sfContextSettings cs) {
     return cs.minorVersion;
 }
 sfContextSettings _NewContextSettings(
 unsigned int depthbits,
 unsigned int stencilbits,
 unsigned int antialiasinglevel,
 unsigned int majorversion,
 unsigned int minorversion
 ){
 sfContextSettings cs;
 cs.depthBits = depthbits;
 cs.stencilBits = stencilbits;
 cs.antialiasingLevel = antialiasinglevel;
 cs.majorVersion = majorversion;
 cs.minorVersion = minorversion;
 return cs;
 }


*/
import "C"

import (
//"unsafe"
//"fmt"
)

//
// Structure defining the window's creation settings
type ContextSettings struct {
	Cref *C.sfContextSettings
	// DepthBits         uint //< Bits of the depth buffer
	// StencilBits       uint //< Bits of the stencil buffer
	// AntialiasingLevel uint //< Level of antialiasing
	// MajorVersion      uint //< Major number of the context version to create
	// MinorVersion      uint //< Minor number of the context version to create
}

func NewContextSettings(depthbits, stencilbits, antialiasinglevel, 
	majVersion, minVersion uint) ContextSettings {
	ref := C._NewContextSettings(
		C.uint(depthbits),
		C.uint(stencilbits),
		C.uint(antialiasinglevel),
		C.uint(majVersion),
		C.uint(minVersion),
		)
	return ContextSettings{&ref}
	// todo: How is this memory collected?
}

func (self ContextSettings) Nil() bool {
	return self.Cref == nil
}
func (self ContextSettings) GetDepthBits() uint {
	return uint(C._GetDepthBits(*self.Cref))
}
func (self ContextSettings) GetStencilBits() uint {
	return uint(C._GetStencilBits(*self.Cref))
}
func (self ContextSettings) GetAntialiasingLevel() uint {
	return uint(C._GetAntialiasingLevel(*self.Cref))
}
func (self ContextSettings) GetMajorVersion() uint {
	return uint(C._GetMajorVersion(*self.Cref))
}
func (self ContextSettings) GetMinorVersion() uint {
	return uint(C._GetMinorVersion(*self.Cref))
}
