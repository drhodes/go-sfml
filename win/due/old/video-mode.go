//
// SFML - Simple and Fast Multimedia Library
// Copyright (C) 2007-2009 Laurent Gomila (laurent.gom@gmail.com)
//
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it freely,
// subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented;
// you must not claim that you wrote the original software.
// If you use this software in a product, an acknowledgment
// in the product documentation would be appreciated but is not required.
//
// 2. Altered source versions must be plainly marked as such,
// and must not be misrepresented as being the original software.
//
// 3. This notice may not be removed or altered from any source distribution.
//
#ifndef SFML_VIDEOMODE_H
#define SFML_VIDEOMODE_H
// Headers
#include <SFML/Config.h>
/// sfVideoMode defines a video mode (width, height, bpp, frequency)
/// and provides functions for getting modes supported
/// by the display device
typedef struct
{
unsigned int Width; ///< Video mode width, in pixels
unsigned int Height; ///< Video mode height, in pixels
unsigned int BitsPerPixel; ///< Video mode pixel depth, in bits per pixels
} sfVideoMode;
/// Get the current desktop video mode
/// return Current desktop video mode
// sfVideoMode sfVideoMode_GetDesktopMode();
func (self *VideoMode) GetDesktopMode() sfVideoMode {
    return C.sfVideoMode_GetDesktopMode(self.Cref)
}
/// Get a valid video mode
/// Index must be in range [0, GetModesCount()[
/// Modes are sorted from best to worst
/// param Index : Index of video mode to get
/// return Corresponding video mode (invalid mode if index is out of range)
// sfVideoMode sfVideoMode_GetMode(size_t Index);
func (self *VideoMode) GetMode() sfVideoMode {
    return C.sfVideoMode_GetMode(self.Cref)
}
/// Get valid video modes count
/// return Number of valid video modes available
// size_t sfVideoMode_GetModesCount();
func (self *VideoMode) GetModesCount() size_t {
    return C.sfVideoMode_GetModesCount(self.Cref)
}
/// Tell whether or not a video mode is supported
/// param Mode : Video mode to check
/// return True if video mode is supported, false otherwise
// sfBool sfVideoMode_IsValid(sfVideoMode Mode);
func (self *VideoMode) IsValid() sfBool {
    return C.sfVideoMode_IsValid(self.Cref)
}
#endif // SFML_VIDEOMODE_H
