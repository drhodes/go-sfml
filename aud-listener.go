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
