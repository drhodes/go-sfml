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
#ifndef SFML_RENDERWINDOW_H
#define SFML_RENDERWINDOW_H
// Headers
#include <SFML/Config.h>
#include <SFML/Graphics/Color.h>
#include <SFML/Graphics/Rect.h>
#include <SFML/Graphics/Types.h>
#include <SFML/Window/Event.h>
#include <SFML/Window/VideoMode.h>
#include <SFML/Window/WindowHandle.h>
#include <SFML/Window/Window.h>
/// Construct a new renderwindow
/// param Mode : Video mode to use
/// param Title : Title of the window
/// param Style : Window style
/// param Params : Creation settings
// sfRenderWindow* sfRenderWindow_Create(sfVideoMode Mode, const char* Title, unsigned long Style, sfWindowSettings Params);
func (self *RenderWindow) Create(title char*, style ulong, params WindowSettings) sfRenderWindow* {
    return C.sfRenderWindow_Create(self.Cref, title, style, params.Cref)
}
/// Construct a renderwindow from an existing control
/// param Handle : Platform-specific handle of the control
/// param Params : Creation settings
// sfRenderWindow* sfRenderWindow_CreateFromHandle(sfWindowHandle Handle, sfWindowSettings Params);
func (self *RenderWindow) CreateFromHandle(params WindowSettings) sfRenderWindow* {
    return C.sfRenderWindow_CreateFromHandle(self.Cref, params.Cref)
}
/// Destroy an existing renderwindow
/// param RenderWindow : Renderwindow to destroy
// void sfRenderWindow_Destroy(sfRenderWindow* RenderWindow);
func (self *RenderWindow) Destroy() void {
    return C.sfRenderWindow_Destroy(self.Cref)
}
/// Close a renderwindow (but doesn't destroy the internal data)
/// param RenderWindow : Renderwindow to close
// void sfRenderWindow_Close(sfRenderWindow* RenderWindow);
func (self *RenderWindow) Close() void {
    return C.sfRenderWindow_Close(self.Cref)
}
/// Tell whether or not a renderwindow is opened
/// param RenderWindow : Renderwindow object
// sfBool sfRenderWindow_IsOpened(sfRenderWindow* RenderWindow);
func (self *RenderWindow) IsOpened() sfBool {
    return C.sfRenderWindow_IsOpened(self.Cref)
}
/// Get the width of the rendering region of a window
/// param RenderWindow : Renderwindow object
/// return Width in pixels
// unsigned int sfRenderWindow_GetWidth(sfRenderWindow* RenderWindow);
/// Get the height of the rendering region of a window
/// param RenderWindow : Renderwindow object
/// return Height in pixels
// unsigned int sfRenderWindow_GetHeight(sfRenderWindow* RenderWindow);
/// Get the creation settings of a window
/// param RenderWindow : Renderwindow object
/// return Settings used to create the window
// sfWindowSettings sfRenderWindow_GetSettings(sfRenderWindow* RenderWindow);
func (self *RenderWindow) GetSettings() sfWindowSettings {
    return C.sfRenderWindow_GetSettings(self.Cref)
}
/// Get the event on top of events stack of a window, if any, and pop it
/// param RenderWindow : Renderwindow object
/// param Event : Event to fill, if any
/// return sfTrue if an event was returned, sfFalse if events stack was empty
// sfBool sfRenderWindow_GetEvent(sfRenderWindow* RenderWindow, sfEvent* Event);
func (self *RenderWindow) GetEvent(event Event*) sfBool {
    return C.sfRenderWindow_GetEvent(self.Cref, event.Cref)
}
/// Enable / disable vertical synchronization on a window
/// param RenderWindow : Renderwindow object
/// param Enabled : sfTrue to enable v-sync, sfFalse to deactivate
// void sfRenderWindow_UseVerticalSync(sfRenderWindow* RenderWindow, sfBool Enabled);
func (self *RenderWindow) UseVerticalSync(enabled Bool) void {
    return C.sfRenderWindow_UseVerticalSync(self.Cref, enabled.Cref)
}
/// Show or hide the mouse cursor on a window
/// param RenderWindow : RenderWindow object
/// param Show : sfTrue to show, sfFalse to hide
// void sfRenderWindow_ShowMouseCursor(sfRenderWindow* RenderWindow, sfBool Show);
func (self *RenderWindow) ShowMouseCursor(show Bool) void {
    return C.sfRenderWindow_ShowMouseCursor(self.Cref, show.Cref)
}
/// Change the position of the mouse cursor on a window
/// param RenderWindow : Renderwindow object
/// param Left : Left coordinate of the cursor, relative to the window
/// param Top : Top coordinate of the cursor, relative to the window
// void sfRenderWindow_SetCursorPosition(sfRenderWindow* RenderWindow, unsigned int Left, unsigned int Top);
func (self *RenderWindow) SetCursorPosition(left uint, top uint) void {
    return C.sfRenderWindow_SetCursorPosition(self.Cref, left, top)
}
/// Change the position of a window on screen.
/// Only works for top-level windows
/// param RenderWindow : Renderwindow object
/// param Left : Left position
/// param Top : Top position
// void sfRenderWindow_SetPosition(sfRenderWindow* RenderWindow, int Left, int Top);
func (self *RenderWindow) SetPosition(left int, top int) void {
    return C.sfRenderWindow_SetPosition(self.Cref, left, top)
}
/// Change the size of the rendering region of a window
/// param RenderWindow : Renderwindow object
/// param Width : New Width
/// param Height : New Height
// void sfRenderWindow_SetSize(sfRenderWindow* RenderWindow, unsigned int Width, unsigned int Height);
func (self *RenderWindow) SetSize(width uint, height uint) void {
    return C.sfRenderWindow_SetSize(self.Cref, width, height)
}
/// Show or hide a window
/// param RenderWindow : Renderwindow object
/// param State : sfTrue to show, sfFalse to hide
// void sfRenderWindow_Show(sfRenderWindow* RenderWindow, sfBool State);
func (self *RenderWindow) Show(state Bool) void {
    return C.sfRenderWindow_Show(self.Cref, state.Cref)
}
/// Enable or disable automatic key-repeat for keydown events.
/// Automatic key-repeat is enabled by default
/// param RenderWindow : Renderwindow object
/// param Enabled : sfTrue to enable, sfFalse to disable
// void sfRenderWindow_EnableKeyRepeat(sfRenderWindow* RenderWindow, sfBool Enabled);
func (self *RenderWindow) EnableKeyRepeat(enabled Bool) void {
    return C.sfRenderWindow_EnableKeyRepeat(self.Cref, enabled.Cref)
}
/// Change the window's icon
/// param RenderWindow : Renderwindow object
/// param Width : Icon's width, in pixels
/// param Height : Icon's height, in pixels
/// param Pixels : Pointer to the pixels in memory, format must be RGBA 32 bits
// void sfRenderWindow_SetIcon(sfRenderWindow* RenderWindow, unsigned int Width, unsigned int Height, sfUint8* Pixels);
func (self *RenderWindow) SetIcon(width uint, height uint, pixels Uint8*) void {
    return C.sfRenderWindow_SetIcon(self.Cref, width, height, pixels.Cref)
}
/// Activate or deactivate a window as the current target for rendering
/// param RenderWindow : Renderwindow object
/// param Active : sfTrue to activate, sfFalse to deactivate
/// return True if operation was successful, false otherwise
// sfBool sfRenderWindow_SetActive(sfRenderWindow* RenderWindow, sfBool Active);
func (self *RenderWindow) SetActive(active Bool) sfBool {
    return C.sfRenderWindow_SetActive(self.Cref, active.Cref)
}
/// Display a window on screen
/// param RenderWindow : Renderwindow object
// void sfRenderWindow_Display(sfRenderWindow* RenderWindow);
func (self *RenderWindow) Display() void {
    return C.sfRenderWindow_Display(self.Cref)
}
/// Get the input manager of a window
/// param RenderWindow : Renderwindow object
/// return Reference to the input
// sfInput* sfRenderWindow_GetInput(sfRenderWindow* RenderWindow);
func (self *RenderWindow) GetInput() sfInput* {
    return C.sfRenderWindow_GetInput(self.Cref)
}
/// Limit the framerate to a maximum fixed frequency for a window
/// param RenderWindow : Renderwindow object
/// param Limit : Framerate limit, in frames per seconds (use 0 to disable limit)
// void sfRenderWindow_SetFramerateLimit(sfRenderWindow* RenderWindow, unsigned int Limit);
func (self *RenderWindow) SetFramerateLimit(limit uint) void {
    return C.sfRenderWindow_SetFramerateLimit(self.Cref, limit)
}
/// Get time elapsed since last frame of a window
/// param RenderWindow : Renderwindow object
/// return Time elapsed, in seconds
// float sfRenderWindow_GetFrameTime(sfRenderWindow* RenderWindow);
func (self *RenderWindow) GetFrameTime() float {
    return C.sfRenderWindow_GetFrameTime(self.Cref)
}
/// Change the joystick threshold, ie. the value below which
/// no move event will be generated
/// param RenderWindow : Renderwindow object
/// param Threshold : New threshold, in range [0, 100]
// void sfRenderWindow_SetJoystickThreshold(sfRenderWindow* RenderWindow, float Threshold);
func (self *RenderWindow) SetJoystickThreshold(threshold float) void {
    return C.sfRenderWindow_SetJoystickThreshold(self.Cref, threshold)
}
/// Draw something on a renderwindow
/// param RenderWindow : Renderwindow to draw in
/// param PostFX / Sprite / String / shape : Object to draw
// void sfRenderWindow_DrawPostFX(sfRenderWindow* RenderWindow, sfPostFX* PostFX);
func (self *RenderWindow) DrawPostFX(postFX PostFX*) void {
    return C.sfRenderWindow_DrawPostFX(self.Cref, postFX.Cref)
}
// void sfRenderWindow_DrawSprite(sfRenderWindow* RenderWindow, sfSprite* Sprite);
func (self *RenderWindow) DrawSprite(sprite Sprite*) void {
    return C.sfRenderWindow_DrawSprite(self.Cref, sprite.Cref)
}
// void sfRenderWindow_DrawShape (sfRenderWindow* RenderWindow, sfShape* Shape);
func (self *RenderWindow) DrawShape(shape Shape*) void {
    return C.sfRenderWindow_DrawShape(self.Cref, shape.Cref)
}
// void sfRenderWindow_DrawString(sfRenderWindow* RenderWindow, sfString* String);
func (self *RenderWindow) DrawString(string String*) void {
    return C.sfRenderWindow_DrawString(self.Cref, string.Cref)
}
/// Save the content of a renderwindow to an image
/// param RenderWindow : Renderwindow to capture
/// return Image instance containing the contents of the screen
// sfImage* sfRenderWindow_Capture(sfRenderWindow* RenderWindow);
func (self *RenderWindow) Capture() sfImage* {
    return C.sfRenderWindow_Capture(self.Cref)
}
/// Clear the screen with the given color
/// param RenderWindow : Renderwindow to modify
/// param Color : Fill color
// void sfRenderWindow_Clear(sfRenderWindow* RenderWindow, sfColor Color);
func (self *RenderWindow) Clear(color Color) void {
    return C.sfRenderWindow_Clear(self.Cref, color.Cref)
}
/// Change the current active view of a renderwindow
/// param RenderWindow : Renderwindow to modify
/// param NewView : Pointer to the new view
// void sfRenderWindow_SetView(sfRenderWindow* RenderWindow, sfView* View);
func (self *RenderWindow) SetView(view View*) void {
    return C.sfRenderWindow_SetView(self.Cref, view.Cref)
}
/// Get the current active view of a renderwindow
/// param RenderWindow : Renderwindow
/// return Current active view
// const sfView* sfRenderWindow_GetView(sfRenderWindow* RenderWindow);
/// Get the default view of a renderwindow
/// param RenderWindow : Renderwindow
/// return Default view of the render window
// sfView* sfRenderWindow_GetDefaultView(sfRenderWindow* RenderWindow);
func (self *RenderWindow) GetDefaultView() sfView* {
    return C.sfRenderWindow_GetDefaultView(self.Cref)
}
/// Convert a point in window coordinates into view coordinates
/// param RenderWindow : Target Renderwindow
/// param WindowX : X coordinate of the point to convert, relative to the window
/// param WindowY : Y coordinate of the point to convert, relative to the window
/// param ViewX : Pointer to fill with the X coordinate of the converted point
/// param ViewY : Pointer to fill with the Y coordinate of the converted point
/// param TargetView : Target view to convert the point to (pass NULL to use the current view)
// void sfRenderWindow_ConvertCoords(sfRenderWindow* RenderWindow, unsigned int WindowX, unsigned int WindowY, float* ViewX, float* ViewY, sfView* TargetView);
func (self *RenderWindow) ConvertCoords(windowX uint, windowY uint, viewX float*, viewY float*, targetView View*) void {
    return C.sfRenderWindow_ConvertCoords(self.Cref, windowX, windowY, viewX, viewY, targetView.Cref)
}
/// Tell SFML to preserve external OpenGL states, at the expense of
/// more CPU charge. Use this function if you don't want SFML
/// to mess up your own OpenGL states (if any).
/// Don't enable state preservation if not needed, as it will allow
/// SFML to do internal optimizations and improve performances.
/// This parameter is false by default
/// param RenderWindow : Target Renderwindow
/// param Preserve : True to preserve OpenGL states, false to let SFML optimize
// void sfRenderWindow_PreserveOpenGLStates(sfRenderWindow* RenderWindow, sfBool Preserve);
func (self *RenderWindow) PreserveOpenGLStates(preserve Bool) void {
    return C.sfRenderWindow_PreserveOpenGLStates(self.Cref, preserve.Cref)
}
#endif // SFML_RENDERWINDOW_H
