// Copyright 2012.  All rights reserved. 
// Use of this source code is governed by a BSD-style  
// license that can be found in the LICENSE file.

////////////////////////////////////////////////////////////
//
// SFML - Simple and Fast Multimedia Library
// Copyright (C) 2007-2012 Laurent Gomila (laurent.gom@gmail.com)
//
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it freely,
// subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented;
//    you must not claim that you wrote the original software.
//    If you use this software in a product, an acknowledgment
//    in the product documentation would be appreciated but is not required.
//
// 2. Altered source versions must be plainly marked as such,
//    and must not be misrepresented as being the original software.
//
// 3. This notice may not be removed or altered from any source distribution.
//
////////////////////////////////////////////////////////////
package sfml


//#include <SFML/Audio/Export.h>
//#include <SFML/Audio/Listener.h>
//#include <SFML/Audio/Types.h>
//#include <SFML/System/Vector3.h>
import "C"

// \brief Change the global volume of all the sounds and musics
// The volume is a number between 0 and 100; it is combined with
// the individual volume of each sound / music.
// The default value for the volume is 100 (maximum).
// \param volume New global volume, in the range [0, 100]
// sfListener_setGlobalVolume(float volume);
func SetGlobalVolume(volume float32) { 
    C.sfListener_setGlobalVolume(C.float(volume));
}
            
// \brief Get the current value of the global volume
// \return Current global volume, in the range [0, 100]
// float sfListener_getGlobalVolume(void);
func GetGlobalVolume() float32 { 
    return float32(C.sfListener_getGlobalVolume())
}

// \brief Set the position of the listener in the scene
// The default listener's position is (0, 0, 0).
// \param position New position of the listener
// sfListener_setPosition(sfVector3f position);
func SetPosition(pos Vector3f) { 
    C.sfListener_setPosition(pos.Cref);
}
            
// \brief Get the current position of the listener in the scene
// \return The listener's position
// sfVector3f sfListener_getPosition();
func GetPosition() Vector3f { 
    return Vector3f{C.sfListener_getPosition()}
}
            
// \brief Set the orientation of the listener in the scene
// The orientation defines the 3D axes of the listener
// (left, up, front) in the scene. The orientation vector
// doesn't have to be normalized.
// The default listener's orientation is (0, 0, -1).
// \param position New direction of the listener
// sfListener_setDirection(sfVector3f orientation);
func SetDirection(orientation Vector3f) { 
    C.sfListener_setDirection(orientation.Cref);
}
            
// \brief Get the current orientation of the listener in the scene
// \return The listener's direction
// sfVector3f sfListener_getDirection();
func GetDirection() Vector3f { 
    return Vector3f{C.sfListener_getDirection()}
}
