package win

// #cgo LDFLAGS:-lcsfml-window
// #include <SFML/Window/VideoMode.h>
// 
// unsigned int VideoModeWidth(sfVideoMode* vm) {
// 	return vm->width;
// }
// unsigned int VideoModeHeight(sfVideoMode* vm) {
// 	return vm->height;
// }
// unsigned int VideoModeBitsPerPixel(sfVideoMode* vm) {
// 	return vm->bitsPerPixel;
// }
// 
import "C"

// sfVideoMode defines a video mode (width, height, bpp, frequency)
// and provides functions for getting modes supported by the display device
type VideoMode struct {
	Cref *C.sfVideoMode
}

func (self VideoMode) Width() uint {
	return uint(C.VideoModeWidth(self.Cref))
}

func (self VideoMode) Height() uint {
	return uint(C.VideoModeHeight(self.Cref))
}

func (self VideoMode) BitsPerPixel() uint {
	return uint(C.VideoModeBitsPerPixel(self.Cref))
}

//  Get the current desktop video mode
///
/// \return Current desktop video mode
///
// sfVideoMode sfVideoMode_getDesktopMode(void)
func GetDesktopMode() VideoMode { 
	temp := C.sfVideoMode_getDesktopMode()
    return VideoMode{&temp}
}

////////////////////////////////////////////////////////////
///  Retrieve all the video modes supported in fullscreen mode
///
/// When creating a fullscreen window, the video mode is restricted
/// to be compatible with what the graphics driver and monitor
/// support. This function returns the complete list of all video
/// modes that can be used in fullscreen mode.
/// The returned array is sorted from best to worst, so that
/// the first element will always give the best mode (higher
/// width, height and bits-per-pixel).
///
/// \param count Pointer to a variable that will be filled with the number of modes in the array
///
/// \return Pointer to an array containing all the supported fullscreen modes
///
////////////////////////////////////////////////////////////
// const sfVideoMode* sfVideoMode_getFullscreenModes(size_t* Count)
// func GetFullscreenModes(count int) VideoMode {
// 	t := C.size_t(count)
//     v := C.sfVideoMode_getFullscreenModes(&t)
// 	return VideoMode{&v}
// }
            
////////////////////////////////////////////////////////////
///  Tell whether or not a video mode is valid
///
/// The validity of video modes is only relevant when using
/// fullscreen windows otherwise any video mode can be used
/// with no restriction.
///
/// \param mode Video mode
///
/// \return sfTrue if the video mode is valid for fullscreen mode
///
////////////////////////////////////////////////////////////
// sfBool sfVideoMode_isValid(sfVideoMode mode)

// func (self VideoMode) Isvalid() Bool { 
//     return C.sfVideoMode_isValid()
// }
            
