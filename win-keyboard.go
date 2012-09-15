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


// #cgo LDFLAGS:-lcsfml-window
// #include <SFML/Window/Keyboard.h>
import "C"

type KeyCode uint

const (
	KeyA         KeyCode = iota // The A key
	KeyB                        // The B key
	KeyC                        // The C key
	KeyD                        // The D key
	KeyE                        // The E key
	KeyF                        // The F key
	KeyG                        // The G key
	KeyH                        // The H key
	KeyI                        // The I key
	KeyJ                        // The J key
	KeyK                        // The K key
	KeyL                        // The L key
	KeyM                        // The M key
	KeyN                        // The N key
	KeyO                        // The O key
	KeyP                        // The P key
	KeyQ                        // The Q key
	KeyR                        // The R key
	KeyS                        // The S key
	KeyT                        // The T key
	KeyU                        // The U key
	KeyV                        // The V key
	KeyW                        // The W key
	KeyX                        // The X key
	KeyY                        // The Y key
	KeyZ                        // The Z key
	KeyNum0                     // The 0 key
	KeyNum1                     // The 1 key
	KeyNum2                     // The 2 key
	KeyNum3                     // The 3 key
	KeyNum4                     // The 4 key
	KeyNum5                     // The 5 key
	KeyNum6                     // The 6 key
	KeyNum7                     // The 7 key
	KeyNum8                     // The 8 key
	KeyNum9                     // The 9 key
	KeyEscape                   // The Escape key
	KeyLControl                 // The left Control key
	KeyLShift                   // The left Shift key
	KeyLAlt                     // The left Alt key
	KeyLSystem                  // The left OS specific key [Win | apple key]
	KeyRControl                 // The right Control key
	KeyRShift                   // The right Shift key
	KeyRAlt                     // The right Alt key
	KeyRSystem                  // The right OS specific key [Win | apple key]
	KeyMenu                     // The Menu key
	KeyLBracket                 // The [ key
	KeyRBracket                 // The ] key
	KeySemiColon                // The ; key
	KeyComma                    // The  key
	KeyPeriod                   // The . key
	KeyQuote                    // The ' key
	KeySlash                    // The / key
	KeyBackSlash                // The \ key
	KeyTilde                    // The ~ key
	KeyEqual                    // The = key
	KeyDash                     // The - key
	KeySpace                    // The Space key
	KeyReturn                   // The Return key
	KeyBack                     // The Backspace key
	KeyTab                      // The Tabulation key
	KeyPageUp                   // The Page up key
	KeyPageDown                 // The Page down key
	KeyEnd                      // The End key
	KeyHome                     // The Home key
	KeyInsert                   // The Insert key
	KeyDelete                   // The Delete key
	KeyAdd                      // +
	KeySubtract                 // -
	KeyMultiply                 // *
	KeyDivide                   // /
	KeyLeft                     // Left arrow
	KeyRight                    // Right arrow
	KeyUp                       // Up arrow
	KeyDown                     // Down arrow
	KeyNumpad0                  // The numpad 0 key
	KeyNumpad1                  // The numpad 1 key
	KeyNumpad2                  // The numpad 2 key
	KeyNumpad3                  // The numpad 3 key
	KeyNumpad4                  // The numpad 4 key
	KeyNumpad5                  // The numpad 5 key
	KeyNumpad6                  // The numpad 6 key
	KeyNumpad7                  // The numpad 7 key
	KeyNumpad8                  // The numpad 8 key
	KeyNumpad9                  // The numpad 9 key
	KeyF1                       // The F1 key
	KeyF2                       // The F2 key
	KeyF3                       // The F3 key
	KeyF4                       // The F4 key
	KeyF5                       // The F5 key
	KeyF6                       // The F6 key
	KeyF7                       // The F7 key
	KeyF8                       // The F8 key
	KeyF9                       // The F8 key
	KeyF10                      // The F10 key
	KeyF11                      // The F11 key
	KeyF12                      // The F12 key
	KeyF13                      // The F13 key
	KeyF14                      // The F14 key
	KeyF15                      // The F15 key
	KeyPause                    // The Pause key

	KeyCount // Keep last -- the total number of keyboard keys
)

// Check if a key is pressed
// key Key to check
// return true if the key is pressed, false otherwise
// sfBool sfKeyboard_isKeyPressed(sfKeyCode key);
func IsKeyPressed(key KeyCode) bool {
	return C.sfKeyboard_isKeyPressed(C.sfKeyCode(key)) == 1
}
