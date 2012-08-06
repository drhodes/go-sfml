package sfml

/*
 #cgo LDFLAGS:-lcsfml-window
 #include <SFML/Window/Window.h>
*/
import "C"

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

	ref := C.sfContextSettings{
		C.uint(depthbits),
		C.uint(stencilbits),
		C.uint(antialiasinglevel),
		C.uint(majVersion),
		C.uint(minVersion),
	}

	return ContextSettings{&ref}
	// todo: How is this memory collected?
}

func (self ContextSettings) Nil() bool {
	return self.Cref == nil
}
func (self ContextSettings) GetDepthBits() uint {
	return uint(self.Cref.depthBits)
}
func (self ContextSettings) GetStencilBits() uint {
	return uint(self.Cref.stencilBits)
}
func (self ContextSettings) GetAntialiasingLevel() uint {
	return uint(self.Cref.antialiasingLevel)
}
func (self ContextSettings) GetMajorVersion() uint {
	return uint(self.Cref.majorVersion)
}
func (self ContextSettings) GetMinorVersion() uint {
	return uint(self.Cref.minorVersion)
}
