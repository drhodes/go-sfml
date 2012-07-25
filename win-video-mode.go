package sfml

// #cgo LDFLAGS:-lcsfml-window
// #include <SFML/Window/VideoMode.h>
// #include <string.h>
// 
//
import "C"
import (
	"unsafe"
	"fmt"
)

// sfVideoMode defines a video mode (width, height, bpp, frequency)
// and provides functions for getting modes supported by the display device
type VideoMode struct {
	Cref *C.sfVideoMode
}

func NewVideoMode(w,h,bpp uint) VideoMode {
	v := C.sfVideoMode{C.uint(w),C.uint(h),C.uint(bpp)}
	return VideoMode{&v}
}

func (self VideoMode) Nil() bool {
	return self.Cref == nil
}

func (self VideoMode) Width() uint {
	return uint(self.Cref.width)
}

func (self VideoMode) Height() uint {
	return uint(self.Cref.height)
}

func (self VideoMode) BitsPerPixel() uint {
	return uint(self.Cref.bitsPerPixel)
}

func (self VideoMode) Show() string {
	s := `
VideoMode:
   Width:        %d
   Height:       %d
   BitsPerPixel: %d
` 
	return fmt.Sprintf(s, self.Width(), self.Height(), self.BitsPerPixel())
}

// Get the current desktop video mode
// sfVideoMode sfVideoMode_getDesktopMode(void)
func GetDesktopMode() VideoMode { 
	temp := C.sfVideoMode_getDesktopMode()
    return VideoMode{&temp}
}


// Retrieve all the video modes supported in fullscreen mode
//
// When creating a fullscreen window, the video mode is restricted
// to be compatible with what the graphics driver and monitor
// support. This function returns the complete list of all video
// modes that can be used in fullscreen mode.
// The returned slice is sorted from best to worst, so that
// the first element will always give the best mode (higher
// width, height and bits-per-pixel).
// const sfVideoMode* sfVideoMode_getFullscreenModes(size_t* Count)
func GetFullscreenModes() []VideoMode {
	// get the size_of a sfVideoMode
	size := unsafe.Sizeof(*VideoMode{}.Cref)

	var nmodes C.size_t;
	cmodes := C.sfVideoMode_getFullscreenModes(&nmodes)

	modes := []VideoMode{}
	p := unsafe.Pointer(cmodes)
	ptr := uintptr(p)

	for nmodes > 0 {
		m := VideoMode{(*C.sfVideoMode)(p)}
		if m.IsValid() {	
			modes = append(modes, m)
		}
		ptr += size
		p = unsafe.Pointer(ptr)		
		nmodes--
	}
	return modes
}

            
// Tell whether or not a video mode is valid
// The validity of video modes is only relevant when using
// fullscreen windows otherwise any video mode can be used
// with no restriction.
// *return true if the video mode is valid for fullscreen mode
// sfBool sfVideoMode_isValid(sfVideoMode mode)
func (self VideoMode) IsValid() bool { 
    return int(C.sfVideoMode_isValid(*self.Cref)) == 1
}
            
