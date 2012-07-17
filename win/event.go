package win

// #cgo LDFLAGS:-lcsfml-window
// #include <SFML/Window/Export.h>
// #include <SFML/Window/Event.h>
// #include <SFML/Window/VideoMode.h>
// #include <SFML/Window/Window.h>
// #include <SFML/Window/WindowHandle.h>
// #include <SFML/Window/Types.h>
// #include <SFML/Window/Export.h>
// #include <SFML/Window/Joystick.h>
// #include <SFML/Window/Keyboard.h>
// #include <SFML/Window/Mouse.h>
import "C"

import (
	//"unsafe"
	//"fmt"
	//. "sfml/sys"
	//"errors"
)

type Event struct {
 	Cref *C.sfEvent
}

func NewEvent() Event {
	e := C.sfEvent{}
	return Event{ &e }
}

type EventType uint
const (
	EvtClosed EventType = iota
	EvtResized
	EvtLostFocus
	EvtGainedFocus
	EvtTextEntered
	EvtKeyPressed
	EvtKeyReleased
	EvtMouseWheelMoved
	EvtMouseButtonPressed
	EvtMouseButtonReleased
	EvtMouseMoved
	EvtMouseEntered
	EvtMouseLeft
	EvtJoystickButtonPressed
	EvtJoystickButtonReleased
	EvtJoystickMoved
	EvtJoystickConnected
	EvtJoystickDisconnected
)

////////////////////////////////////////////////////////////
/// \brief Keyboard event parameters
///
////////////////////////////////////////////////////////////
// struct sfKeyEvent
// {
//     sfEventType type;
//     sfKeyCode   code;
//     sfBool      alt;
//     sfBool      control;
//     sfBool      shift;
//     sfBool      system;
// };

func (self KeyEvent) Code() bool {
		return C.sfKeyEvent(self.Cref) == 1
}
// func (self KeyEvent) Alt() bool {
// }
// func (self KeyEvent) Control() bool {
// }
// func (self KeyEvent) Shift() bool {
// }
// func (self KeyEvent) System() bool {
// }





////////////////////////////////////////////////////////////
/// \brief Text event parameters
///
////////////////////////////////////////////////////////////
// struct sfTextEvent
// {
//     sfEventType type;
//     sfUint32    unicode;
// };

////////////////////////////////////////////////////////////
/// \brief Mouse move event parameters
///
////////////////////////////////////////////////////////////
// struct sfMouseMoveEvent
// {
//     sfEventType type;
//     int         x;
//     int         y;
// };

////////////////////////////////////////////////////////////
/// \brief Mouse buttons events parameters
///
////////////////////////////////////////////////////////////
// struct sfMouseButtonEvent
// {
//     sfEventType   type;
//     sfMouseButton button;
//     int           x;
//     int           y;
// };

////////////////////////////////////////////////////////////
/// \brief Mouse wheel events parameters
///
////////////////////////////////////////////////////////////
// struct sfMouseWheelEvent
// {
//     sfEventType type;
//     int         delta;
//     int         x;
//     int         y;
// };

////////////////////////////////////////////////////////////
/// \brief Joystick axis move event parameters
///
////////////////////////////////////////////////////////////
// struct sfJoystickMoveEvent
// {
//     sfEventType    type;
//     unsigned int   joystickId;
//     sfJoystickAxis axis;
//     float          position;
// };

////////////////////////////////////////////////////////////
/// \brief Joystick buttons events parameters
///
////////////////////////////////////////////////////////////
// struct sfJoystickButtonEvent
// {
//     sfEventType  type;
//     unsigned int joystickId;
//     unsigned int button;
// };

////////////////////////////////////////////////////////////
/// \brief Joystick connection/disconnection event parameters
///
////////////////////////////////////////////////////////////
// struct sfJoystickConnectEvent
// {
//     sfEventType  type;
//     unsigned int joystickId;
// };

////////////////////////////////////////////////////////////
/// \brief Size events parameters
///
////////////////////////////////////////////////////////////
// struct sfSizeEvent
// {
//     sfEventType  type;
//     unsigned int width;
//     unsigned int height;
// };

////////////////////////////////////////////////////////////
/// \brief sfEvent defines a system event and its parameters
///
////////////////////////////////////////////////////////////
//typedef union
// {
//     ////////////////////////////////////////////////////////////
//     // Member data
//     ////////////////////////////////////////////////////////////
//     sfEventType                   type; ///< Type of the event
//     struct sfSizeEvent            size;
//     struct sfKeyEvent             key;
//     struct sfTextEvent            text;
//     struct sfMouseMoveEvent       mouseMove;
//     struct sfMouseButtonEvent     mouseButton;
//     struct sfMouseWheelEvent      mouseWheel;
//     struct sfJoystickMoveEvent    joystickMove;
//     struct sfJoystickButtonEvent  joystickButton;
//     struct sfJoystickConnectEvent joystickConnect;
// } sfEvent;
