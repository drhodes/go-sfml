package win

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

import (
	//"unsafe"
	//"fmt"
	//. "sfml/sys"
	//"errors"
)

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
	return Event{ &e }
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

type KeyCode int
const (
	KeyA = iota     ///< The A key
	KeyB            ///< The B key
	KeyC            ///< The C key
	KeyD            ///< The D key
	KeyE            ///< The E key
	KeyF            ///< The F key
	KeyG            ///< The G key
	KeyH            ///< The H key
	KeyI            ///< The I key
	KeyJ            ///< The J key
	KeyK            ///< The K key
	KeyL            ///< The L key
	KeyM            ///< The M key
	KeyN            ///< The N key
	KeyO            ///< The O key
	KeyP            ///< The P key
	KeyQ            ///< The Q key
	KeyR            ///< The R key
	KeyS            ///< The S key
	KeyT            ///< The T key
	KeyU            ///< The U key
	KeyV            ///< The V key
	KeyW            ///< The W key
	KeyX            ///< The X key
	KeyY            ///< The Y key
	KeyZ            ///< The Z key
	KeyNum0         ///< The 0 key
	KeyNum1         ///< The 1 key
	KeyNum2         ///< The 2 key
	KeyNum3         ///< The 3 key
	KeyNum4         ///< The 4 key
	KeyNum5         ///< The 5 key
	KeyNum6         ///< The 6 key
	KeyNum7         ///< The 7 key
	KeyNum8         ///< The 8 key
	KeyNum9         ///< The 9 key
	KeyEscape       ///< The Escape key
	KeyLControl     ///< The left Control key
	KeyLShift       ///< The left Shift key
	KeyLAlt         ///< The left Alt key
	KeyLSystem      ///< The left OS specific key: window (Windows and Linux), apple (MacOS X), ...
	KeyRControl     ///< The right Control key
	KeyRShift       ///< The right Shift key
	KeyRAlt         ///< The right Alt key
	KeyRSystem      ///< The right OS specific key: window (Windows and Linux), apple (MacOS X), ...
	KeyMenu         ///< The Menu key
	KeyLBracket     ///< The [ key
	KeyRBracket     ///< The ] key
	KeySemiColon    ///< The ; key
	KeyComma        ///< The , key
	KeyPeriod       ///< The . key
	KeyQuote        ///< The ' key
	KeySlash        ///< The / key
	KeyBackSlash    ///< The \ key
	KeyTilde        ///< The ~ key
	KeyEqual        ///< The = key
	KeyDash         ///< The - key
	KeySpace        ///< The Space key
	KeyReturn       ///< The Return key
	KeyBack         ///< The Backspace key
	KeyTab          ///< The Tabulation key
	KeyPageUp       ///< The Page up key
	KeyPageDown     ///< The Page down key
	KeyEnd          ///< The End key
	KeyHome         ///< The Home key
	KeyInsert       ///< The Insert key
	KeyDelete       ///< The Delete key
	KeyAdd          ///< +
	KeySubtract     ///< -
	KeyMultiply     ///< *
	KeyDivide       ///< /
	KeyLeft         ///< Left arrow
	KeyRight        ///< Right arrow
	KeyUp           ///< Up arrow
	KeyDown         ///< Down arrow
	KeyNumpad0      ///< The numpad 0 key
	KeyNumpad1      ///< The numpad 1 key
	KeyNumpad2      ///< The numpad 2 key
	KeyNumpad3      ///< The numpad 3 key
	KeyNumpad4      ///< The numpad 4 key
	KeyNumpad5      ///< The numpad 5 key
	KeyNumpad6      ///< The numpad 6 key
	KeyNumpad7      ///< The numpad 7 key
	KeyNumpad8      ///< The numpad 8 key
	KeyNumpad9      ///< The numpad 9 key
	KeyF1           ///< The F1 key
	KeyF2           ///< The F2 key
	KeyF3           ///< The F3 key
	KeyF4           ///< The F4 key
	KeyF5           ///< The F5 key
	KeyF6           ///< The F6 key
	KeyF7           ///< The F7 key
	KeyF8           ///< The F8 key
	KeyF9           ///< The F8 key
	KeyF10          ///< The F10 key
	KeyF11          ///< The F11 key
	KeyF12          ///< The F12 key
	KeyF13          ///< The F13 key
	KeyF14          ///< The F14 key
	KeyF15          ///< The F15 key
	KeyPause        ///< The Pause key
)

type KeyEvent struct {
	Type EventType	
	Cref *C.struct_sfKeyEvent
}
func (self KeyEvent) Code() KeyCode {
	return (KeyCode) (self.Cref.code)
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
