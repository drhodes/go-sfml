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


/*
 #cgo LDFLAGS:-lcsfml-window
 #include <SFML/Window/Event.h>
 #include <SFML/Window/Export.h>
 #include <SFML/Window/Joystick.h>
 #include <SFML/Window/Keyboard.h>
 #include <SFML/Window/Mouse.h>
 #include <SFML/Window/Types.h>
 #include <SFML/Window/VideoMode.h>
 #include <SFML/Window/Window.h>
 #include <SFML/Window/WindowHandle.h>

 struct sfKeyEvent*
 Event2KeyEvent(sfEvent* e) { return &e->key; }
 struct sfSizeEvent*
 Event2SizeEvent(sfEvent* e) { return &e->size; }
 struct sfTextEvent*
 Event2TextEvent(sfEvent* e) { return &e->text; }
 struct sfMouseMoveEvent*
 Event2MouseMoveEvent(sfEvent* e) { return &e->mouseMove; }
 struct sfMouseButtonEvent*
 Event2MouseButtonEvent(sfEvent* e) { return &e->mouseButton; }
 struct sfMouseWheelEvent*
 Event2MouseWheelEvent(sfEvent* e) { return &e->mouseWheel; }
 struct sfJoystickMoveEvent*
 Event2JoystickMoveEvent(sfEvent* e) { return &e->joystickMove; }
 struct sfJoystickButtonEvent*
 Event2JoystickButtonEvent(sfEvent* e) { return &e->joystickButton; }
 struct sfJoystickConnectEvent*
 Event2JoystickConnectEvent(sfEvent* e) { return &e->joystickConnect; }
 sfEventType*
 GetEventType(sfEvent* e) { return &e->type; }

*/
import "C"

type CEvent *C.sfEvent
type Event struct {
	Cref *C.sfEvent
}

func (self Event) ToSizeEvent() SizeEvent {
	return SizeEvent{
		EventType(*C.GetEventType(self.Cref)),
		C.Event2SizeEvent(self.Cref),
	}
}
func (self Event) ToKeyEvent() KeyEvent {
	return KeyEvent{
		EventType(*C.GetEventType(self.Cref)),
		C.Event2KeyEvent(self.Cref),
	}
}
func (self Event) ToTextEvent() TextEvent {
	return TextEvent{
		EventType(*C.GetEventType(self.Cref)),
		C.Event2TextEvent(self.Cref),
	}
}
func (self Event) ToMouseMoveEvent() MouseMoveEvent {
	return MouseMoveEvent{
		EventType(*C.GetEventType(self.Cref)),
		C.Event2MouseMoveEvent(self.Cref),
	}
}
func (self Event) ToMouseButtonEvent() MouseButtonEvent {
	return MouseButtonEvent{
		EventType(*C.GetEventType(self.Cref)),
		C.Event2MouseButtonEvent(self.Cref),
	}
}
func (self Event) ToMouseWheelEvent() MouseWheelEvent {
	return MouseWheelEvent{
		EventType(*C.GetEventType(self.Cref)),
		C.Event2MouseWheelEvent(self.Cref),
	}
}
func (self Event) ToJoystickMoveEvent() JoystickMoveEvent {
	return JoystickMoveEvent{
		EventType(*C.GetEventType(self.Cref)),
		C.Event2JoystickMoveEvent(self.Cref),
	}
}
func (self Event) ToJoystickButtonEvent() JoystickButtonEvent {
	return JoystickButtonEvent{
		EventType(*C.GetEventType(self.Cref)),
		C.Event2JoystickButtonEvent(self.Cref),
	}
}
func (self Event) ToJoystickConnectEvent() JoystickConnectEvent {
	return JoystickConnectEvent{
		EventType(*C.GetEventType(self.Cref)),
		C.Event2JoystickConnectEvent(self.Cref),
	}
}

func NewEvent() Event {
	e := C.sfEvent{}
	return Event{&e}
}

//var EvChan = make(chan, interface{})

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
	EvtNone // go introduction.
)

type NullEvent uint

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

type KeyEvent struct {
	Type EventType
	Cref *C.struct_sfKeyEvent
}

func (self KeyEvent) Code() KeyCode {
	return (KeyCode)(self.Cref.code)
}
func (self KeyEvent) Alt() bool {
	return self.Cref.alt == 1
}
func (self KeyEvent) Control() bool {
	return self.Cref.control == 1
}
func (self KeyEvent) Shift() bool {
	return self.Cref.shift == 1
}
func (self KeyEvent) System() bool {
	return self.Cref.system == 1
}

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
type SizeEvent struct {
	Type EventType
	Cref *C.struct_sfSizeEvent
}

func (self SizeEvent) Width() int {
	return int(self.Cref.width)
}

func (self SizeEvent) Height() int {
	return int(self.Cref.height)
}

////////////////////////////////////////////////////////////
/// \brief Text event parameters
///
////////////////////////////////////////////////////////////
// struct sfTextEvent
// {
//     sfEventType type;
//     sfUint32    unicode;
// };
type TextEvent struct {
	Type EventType
	Cref *C.struct_sfTextEvent
}

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
type MouseMoveEvent struct {
	Type EventType
	Cref *C.struct_sfMouseMoveEvent
}

func (self MouseMoveEvent) X() int {
	return int(self.Cref.x)
}

func (self MouseMoveEvent) Y() int {
	return int(self.Cref.y)
}

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
type MouseButtonEvent struct {
	Type EventType
	Cref *C.struct_sfMouseButtonEvent
}

func (self MouseButtonEvent) Button() MouseButton {
	return (MouseButton) (uint(self.Cref.button))
}

func (self MouseButtonEvent) X() int {
	return int(self.Cref.x)
}

func (self MouseButtonEvent) Y() int {
	return int(self.Cref.y)
}

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
type MouseWheelEvent struct {
	Type EventType
	Cref *C.struct_sfMouseWheelEvent
}

func (self MouseWheelEvent) Delta() int {
	return int(self.Cref.delta)
}

func (self MouseWheelEvent) X() int {
	return int(self.Cref.x)
}

func (self MouseWheelEvent) Y() int {
	return int(self.Cref.y)
}

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
type JoystickMoveEvent struct {
	Type EventType
	Cref *C.struct_sfJoystickMoveEvent
}

func (self JoystickMoveEvent) JoystickID() uint {
	return uint(self.Cref.joystickId)
}

func (self JoystickMoveEvent) Axis() JoystickAxis {
	axis := self.Cref.axis
	return JoystickAxis{&axis}
}

func (self JoystickMoveEvent) Position() float32 {
	return float32(self.Cref.position)
}

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
type JoystickButtonEvent struct {
	Type EventType
	Cref *C.struct_sfJoystickButtonEvent
}

func (self JoystickButtonEvent) JoystickID() uint {
	return uint(self.Cref.joystickId)
}

func (self JoystickButtonEvent) Button() uint {
	return uint(self.Cref.button)
}

////////////////////////////////////////////////////////////
/// \brief Joystick connection/disconnection event parameters
///
////////////////////////////////////////////////////////////
// struct sfJoystickConnectEvent
// {
//     sfEventType  type;
//     unsigned int joystickId;
// };
type JoystickConnectEvent struct {
	Type EventType
	Cref *C.struct_sfJoystickConnectEvent
}

func (self JoystickConnectEvent) JoystickID() uint {
	return uint(self.Cref.joystickId)
}

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
