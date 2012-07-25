package gfx

// #cgo LDFLAGS:-lcsfml-graphics
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/Graphics/PrimitiveType.h>
// #include <SFML/Graphics/RenderStates.h>
// #include <SFML/Graphics/Vertex.h>
// #include <SFML/Graphics/RenderWindow.h>
// #include <SFML/Window/Event.h>
// #include <SFML/Window/VideoMode.h>
// #include <SFML/Window/WindowHandle.h>
// #include <SFML/Window/Window.h>
// #include <SFML/Window/Context.h>
// #include <SFML/System/Vector2.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
	. "sfml/win"
	"sfml/win"
	//"sfml/sys"
	//"log"
)

type RenderWindow struct {
    Cref *C.sfRenderWindow
}

// \brief Construct a new render window
//
// \param mode     Video mode to use
// \param title    Title of the window
// \param style    Window style
// \param settings Creation settings (pass NULL to use default values)
//
// sfRenderWindow* sfRenderWindow_create(sfVideoMode mode, const char* title, sfUint32 style, const sfContextSettings* settings);
func NewRenderWindow(mode VideoMode, title string, style uint32, settings ContextSettings) RenderWindow {
	mptr := unsafe.Pointer(mode.Cref)
	mp := (*C.sfVideoMode)(mptr)

	sptr := unsafe.Pointer(settings.Cref)
	sp := (*C.sfContextSettings)(sptr)

	ctitle := C.CString(title)
	//defer C.free(unsafe.Pointer(&ctitle))

	ref := C.sfRenderWindow_create(*mp, ctitle, C.sfUint32(style), sp)
	return RenderWindow{ref}
}

func NewRenderWindowDefault(mode VideoMode, title string) RenderWindow {
	style := win.StyleDefaultStyle
	mptr := unsafe.Pointer(mode.Cref)
	mp := (*C.sfVideoMode)(mptr)

	ctitle := C.CString(title)
	//defer C.free(unsafe.Pointer(&ctitle))

	ref := C.sfRenderWindow_create(*mp, ctitle, C.sfUint32(style), nil)
	return RenderWindow{ref}
}

// Construct a render window from an existing control
// \param handle   Platform-specific handle of the control
// \param settings Creation settings (pass NULL to use default values)
// sfRenderWindow* sfRenderWindow_createFromHandle(sfWindowHandle handle, const sfContextSettings* settings);
// func RenderWindowFromhandle(handle WindowHandle, settings ContextSettings) RenderWindow { 
// 	mptr := unsafe.Pointer(&mode)
// 	mp := (*C.sfVideoMode)(unsafe.Pointer(&handle))
//     return RenderWindow { 
// 		C.sfRenderWindow_createFromHandle(C.sfWindowHandle(handle.Cref), (*C.sfContextSettings)(settings.Cref)),
// 	}
// }
            
// Destroy an existing render window
// \param renderWindow Render window to destroy
// void sfRenderWindow_destroy(sfRenderWindow* renderWindow);
func (self RenderWindow) Destroy() { 
    C.sfRenderWindow_destroy(self.Cref)
}
            
// Close a render window (but doesn't destroy the internal data)
// \param renderWindow Render window to close
// void sfRenderWindow_close(sfRenderWindow* renderWindow);
func (self RenderWindow) Close() { 
    C.sfRenderWindow_close(self.Cref)
}
            
// Tell whether or not a render window is opened
// \param renderWindow Render window object
// sfBool sfRenderWindow_isOpen(const sfRenderWindow* renderWindow);
func (self RenderWindow) IsOpen() bool { 
    return C.sfRenderWindow_isOpen(self.Cref) == 1
}
            
// Get the creation settings of a render window
// \param renderWindow Render window object
// \return Settings used to create the window
// sfContextSettings sfRenderWindow_getSettings(const sfRenderWindow* renderWindow);
// func (self RenderWindow) GetSettings() ContextSettings { 
// 	ref := C.sfRenderWindow_getSettings(self.Cref)
// 	rptr := unsafe.Pointer(&ref)
// 	rp := (*C.sfContextSettings)(rptr)
//     return ContextSettings{rp}
// }

            
// Get the event on top of events stack of a render window, if any, and pop it
// \param renderWindow Render window object
// \param event        Event to fill, if any
// \return sfTrue if an event was returned, sfFalse if events stack was empty
// sfBool sfRenderWindow_pollEvent(sfRenderWindow* renderWindow, sfEvent* event);
func (self RenderWindow) PollEvent() (interface{}, bool) { 
	// ok if got event.	
	e := NewEvent()
	mptr := unsafe.Pointer(e.Cref)
	mp := (*C.sfEvent)(mptr)

	ok := C.sfRenderWindow_pollEvent(self.Cref, mp) == 1
	if ok {
		// look at the first byte, it's the event type
		et := EventType((*e.Cref)[0])
		switch et {
		case EvtClosed:			
		case EvtResized:
		case EvtLostFocus:
		case EvtGainedFocus:			
		case EvtTextEntered:
			return e.ToTextEvent(), true
		case EvtKeyPressed, EvtKeyReleased:
			return e.ToKeyEvent(), true		
		case EvtMouseWheelMoved:
			return e.ToMouseWheelEvent(), true
		case EvtMouseButtonPressed, EvtMouseButtonReleased:
			return e.ToMouseButtonEvent(), true
		case EvtMouseMoved, EvtMouseEntered, EvtMouseLeft:
			return e.ToMouseMoveEvent(), true
		case EvtJoystickButtonPressed, EvtJoystickButtonReleased, EvtJoystickMoved:
			return e.ToJoystickMoveEvent(), true
		case EvtJoystickConnected:
		case EvtJoystickDisconnected:
			return e.ToJoystickConnectEvent(), true
		}
	}
	return EvtNone, false
	
	//return e
	// e := unsafe.Pointer(event.Cref)
	// ep := (*C.sfEvent)(e)
    // return C.sfRenderWindow_pollEvent(self.Cref, ep) == 1
}

func (self RenderWindow) Drain() {
    for {
        _, ok := self.PollEvent()
        if !ok { break }
    }
}


// Wait for an event and return it
// \param renderWindow Render window object
// \param event        Event to fill
// \return sfFalse if an error occured
// sfBool sfRenderWindow_waitEvent(sfRenderWindow* renderWindow, sfEvent* event);
func (self RenderWindow) WaitEvent(event Event) bool { 
	ptr := unsafe.Pointer(event.Cref)
	p := (*C.sfEvent)(ptr)
    return C.sfRenderWindow_waitEvent(self.Cref, p) == 1
}

/*                                  
// Get the position of a render window
// \param renderWindow Render window object
// \return Position in pixels
// sfVector2i sfRenderWindow_getPosition(const sfRenderWindow* renderWindow);

func (self RenderWindow) Getposition() Vector2i { 
    return C.sfRenderWindow_getPosition(self.Cref);
}
            
// Change the position of a render window on screen
// Only works for top-level windows
// \param renderWindow Render window object
// \param position     New position, in pixels
// void sfRenderWindow_setPosition(sfRenderWindow* renderWindow, sfVector2i position);

func (self RenderWindow) Setposition(position Vector2i) void { 
    return C.sfRenderWindow_setPosition(self.Cref, sfVector2i(position));
}
            
// Get the size of the rendering region of a render window
// \param renderWindow Render window object
// \return Size in pixels
// sfVector2u sfRenderWindow_getSize(const sfRenderWindow* renderWindow);

func (self RenderWindow) Getsize() Vector2u { 
    return C.sfRenderWindow_getSize(self.Cref);
}
  */         
// Change the size of the rendering region of a render window
// \param renderWindow Render window object
// \param size         New size, in pixels
// void sfRenderWindow_setSize(sfRenderWindow* renderWindow, sfVector2u size);
// func (self RenderWindow) Setsize(x, y uint) { 		
//     C.sfRenderWindow_setSize(self.Cref, .Vector2u(x,y).Cref)
// }
/*            
// Change the title of a render window
// \param renderWindow Render window object
// \param title        New title
// void sfRenderWindow_setTitle(sfRenderWindow* renderWindow, const char* title);

func (self RenderWindow) Settitle(title *char ) void { 
    return C.sfRenderWindow_setTitle(self.Cref, sf(*char));
}
            
// Change a render window's icon
// \param renderWindow Renderw indow object
// \param width        Icon's width, in pixels
// \param height       Icon's height, in pixels
// \param pixels       Pointer to the pixels in memory, format must be RGBA 32 bits
// void sfRenderWindow_setIcon(sfRenderWindow* renderWindow, unsigned int width, unsigned int height, const sfUint8* pixels);

func (self RenderWindow) Seticon(width int , height int , pixels *Uint8 ) void { 
    return C.sfRenderWindow_setIcon(self.Cref, sf(int), sf(int), sf(*Uint8));
}
            
// Show or hide a render window
// \param renderWindow Render window object
// \param visible      sfTrue to show the window, sfFalse to hide it
// void sfRenderWindow_setVisible(sfRenderWindow* renderWindow, sfBool visible);

func (self RenderWindow) Setvisible(visible Bool) void { 
    return C.sfRenderWindow_setVisible(self.Cref, sfBool(visible));
}
            
// Show or hide the mouse cursor on a render window
// \param renderWindow Render window object
// \param show         sfTrue to show, sfFalse to hide
// void sfRenderWindow_setMouseCursorVisible(sfRenderWindow* renderWindow, sfBool show);
*/
func (self RenderWindow) SetMouseCursorVisible(show bool) { 
    C.sfRenderWindow_setMouseCursorVisible(self.Cref, Bool(show));
}
  /*          
// Enable / disable vertical synchronization on a render window
// \param renderWindow Render window object
// \param enabled      sfTrue to enable v-sync, sfFalse to deactivate
// void sfRenderWindow_setVerticalSyncEnabled(sfRenderWindow* renderWindow, sfBool enabled);
*/
func (self RenderWindow) SetVerticalSyncEnabled(enabled bool) { 
    C.sfRenderWindow_setVerticalSyncEnabled(self.Cref, Bool(enabled))
}

// Enable or disable automatic key-repeat for keydown events
// Automatic key-repeat is enabled by default
// \param renderWindow Render window object
// \param enabled      sfTrue to enable, sfFalse to disable
// void sfRenderWindow_setKeyRepeatEnabled(sfRenderWindow* renderWindow, sfBool enabled);
func (self RenderWindow) SetKeyRepeatEnabled(enabled bool) { 
    C.sfRenderWindow_setKeyRepeatEnabled(self.Cref, Bool(enabled));
}

 /*          
            
// Activate or deactivate a render window as the current target for rendering
// \param renderWindow Render window object
// \param active       sfTrue to activate, sfFalse to deactivate
// \return True if operation was successful, false otherwise
// sfBool sfRenderWindow_setActive(sfRenderWindow* renderWindow, sfBool active);

func (self RenderWindow) Setactive(active Bool) Bool { 
    return C.sfRenderWindow_setActive(self.Cref, sfBool(active));
}
*/            
// Display a render window on screen
// \param renderWindow Render window object
// void sfRenderWindow_display(sfRenderWindow* renderWindow);
func (self RenderWindow) Display() { 
    C.sfRenderWindow_display(self.Cref);
}
// Limit the framerate to a maximum fixed frequency for a render window
// \param renderWindow Render window object
// \param limit        Framerate limit, in frames per seconds (use 0 to disable limit)
//  void sfRenderWindow_setFramerateLimit(sfRenderWindow* renderWindow, unsigned int limit);
func (self RenderWindow) SetFrameRateLimit(limit uint) { 
    C.sfRenderWindow_setFramerateLimit(self.Cref, C.uint(limit))
}
/*            
// Change the joystick threshold, ie. the value below which no move event will be generated
// \param renderWindow Render window object
// \param threshold    New threshold, in range [0, 100]
// void sfRenderWindow_setJoystickThreshold(sfRenderWindow* renderWindow, float threshold);

func (self RenderWindow) Setjoystickthreshold(threshold float) void { 
    return C.sfRenderWindow_setJoystickThreshold(self.Cref, sfFloat(threshold));
}
            
// Retrieve the OS-specific handle of a render window
// \param renderWindow Render window object
// \return Window handle
// sfWindowHandle sfRenderWindow_getSystemHandle(const sfRenderWindow* renderWindow);

func (self RenderWindow) Getsystemhandle() WindowHandle { 
    return C.sfRenderWindow_getSystemHandle(self.Cref);
}
*/            

// Clear a render window with the given color
// \param renderWindow Render window object
// \param color        Fill color
// void sfRenderWindow_clear(sfRenderWindow* renderWindow, sfColor color);
func (self RenderWindow) Clear(color Color) { 
    C.sfRenderWindow_clear(self.Cref, color.Cref)
}

/*            
// Change the current active view of a render window
// \param renderWindow Render window object
// \param view         Pointer to the new view
// void sfRenderWindow_setView(sfRenderWindow* renderWindow, const sfView* view);
func (self RenderWindow) Setview(view *View ) void { 
    return C.sfRenderWindow_setView(self.Cref, sf(*View));
}
            
// Get the current active view of a render window
// \param renderWindow Render window object
// \return Current active view
// const sfView* sfRenderWindow_getView(const sfRenderWindow* renderWindow);
func (self *View) *View(RenderWindow_getView)  { 
    return C.sf*View(self.Cref, sf(*View));
}
            
// Get the default view of a render window
// \param renderWindow Render window object
// \return Default view of the render window
// const sfView* sfRenderWindow_getDefaultView(const sfRenderWindow* renderWindow);

func (self *View) *View(RenderWindow_getDefaultView)  { 
    return C.sf*View(self.Cref, sf(*View));
}
            
// Get the viewport of a view applied to this target
// \param renderWindow Render window object
// \param view         Target view
// \return Viewport rectangle, expressed in pixels in the current target
// sfIntRect sfRenderWindow_getViewport(const sfRenderWindow* renderWindow, const sfView* view);
func (self RenderWindow) Getviewport(view *View ) IntRect { 
    return C.sfRenderWindow_getViewport(self.Cref, sf(*View));
}
            
// Convert a point in window coordinates into view coordinates
// \param renderWindow Render window object
// \param point        Point to convert, relative to the window
// \param targetView   Target view to convert the point to (pass NULL to use the current view)
// \return The converted point, in "world" units
// sfVector2f sfRenderWindow_convertCoords(const sfRenderWindow* renderWindow, sfVector2i point, const sfView* targetView);
func (self RenderWindow) Convertcoords(point Vector2i, targetView *View ) Vector2f { 
    return C.sfRenderWindow_convertCoords(self.Cref, sfVector2i(point), sf(*View));
}
            
*/
// Draw a drawable object to the render-target
// \param renderWindow render window object
// \param object       Object to draw
// \param states       Render states to use for drawing (NULL to use the default states)
// void sfRenderWindow_drawSprite(sfRenderWindow* renderWindow, const sfSprite* object, const sfRenderStates* states);
// func (self RenderWindow) DrawSprite(object *Sprite , states *RenderStates ) { 
//     return C.sfRenderWindow_drawSprite(self.Cref, sf(*Sprite), sf(*RenderStates));
// }
func (self RenderWindow) DrawSpriteDefault(obj *Sprite) { 
    C.sfRenderWindow_drawSprite(self.Cref, obj.Cref, nil)
}


// void sfRenderWindow_drawText(sfRenderWindow* renderWindow, const sfText* object, const sfRenderStates* states);
// func (self RenderWindow) Drawtext(object *Text , states *RenderStates ) void { 
//     return C.sfRenderWindow_drawText(self.Cref, sf(*Text), sf(*RenderStates));
// }
// func (self RenderWindow) DrawTextDefault(object *Text) { 
//     return C.sfRenderWindow_drawText(self.Cref, object.Cref, nil)
// }


/*            

            
// void sfRenderWindow_drawShape(sfRenderWindow* renderWindow, const sfShape* object, const sfRenderStates* states);

func (self RenderWindow) Drawshape(object *Shape , states *RenderStates ) void { 
    return C.sfRenderWindow_drawShape(self.Cref, sf(*Shape), sf(*RenderStates));
}
            
// void sfRenderWindow_drawCircleShape(sfRenderWindow* renderWindow, const sfCircleShape* object, const sfRenderStates* states);

func (self RenderWindow) Drawcircleshape(object *CircleShape , states *RenderStates ) void { 
    return C.sfRenderWindow_drawCircleShape(self.Cref, sf(*CircleShape), sf(*RenderStates));
}
            
// void sfRenderWindow_drawConvexShape(sfRenderWindow* renderWindow, const sfConvexShape* object, const sfRenderStates* states);

func (self RenderWindow) Drawconvexshape(object *ConvexShape , states *RenderStates ) void { 
    return C.sfRenderWindow_drawConvexShape(self.Cref, sf(*ConvexShape), sf(*RenderStates));
}
            
// void sfRenderWindow_drawRectangleShape(sfRenderWindow* renderWindow, const sfRectangleShape* object, const sfRenderStates* states);

func (self RenderWindow) Drawrectangleshape(object *RectangleShape , states *RenderStates ) void { 
    return C.sfRenderWindow_drawRectangleShape(self.Cref, sf(*RectangleShape), sf(*RenderStates));
}
            
// void sfRenderWindow_drawVertexArray(sfRenderWindow* renderWindow, const sfVertexArray* object, const sfRenderStates* states);

func (self RenderWindow) Drawvertexarray(object *VertexArray , states *RenderStates ) void { 
    return C.sfRenderWindow_drawVertexArray(self.Cref, sf(*VertexArray), sf(*RenderStates));
}
            
// Draw primitives defined by an array of vertices to a render window
// \param renderWindow render window object
// \param vertices     Pointer to the vertices
// \param vertexCount  Number of vertices in the array
// \param type         Type of primitives to draw
// \param states       Render states to use for drawing (NULL to use the default states)
// void sfRenderWindow_drawPrimitives(sfRenderWindow* renderWindow, const sfVertex* vertices, unsigned int vertexCount, sfPrimitiveType type, const sfRenderStates* states);

func (self RenderWindow) Drawprimitives(vertices *Vertex , vertexCount int , type PrimitiveType, states *RenderStates ) void { 
    return C.sfRenderWindow_drawPrimitives(self.Cref, sf(*Vertex), sf(int), sfPrimitivetype(type), sf(*RenderStates));
}
            
// Save the current OpenGL render states and matrices
// This function can be used when you mix SFML drawing
// and direct OpenGL rendering. Combined with popGLStates,
// it ensures that:
// \li SFML's internal states are not messed up by your OpenGL code
// \li your OpenGL states are not modified by a call to a SFML function
// Note that this function is quite expensive: it saves all the
// possible OpenGL states and matrices, even the ones you
// don't care about. Therefore it should be used wisely.
// It is provided for convenience, but the best results will
// be achieved if you handle OpenGL states yourself (because
// you know which states have really changed, and need to be
// saved and restored). Take a look at the resetGLStates
// function if you do so.
// \param renderWindow render window object
// void sfRenderWindow_pushGLStates(sfRenderWindow* renderWindow);

func (self RenderWindow) Pushglstates() void { 
    return C.sfRenderWindow_pushGLStates(self.Cref);
}
            
// Restore the previously saved OpenGL render states and matrices
// See the description of pushGLStates to get a detailed
// description of these functions.
// \param renderWindow render window object
// void sfRenderWindow_popGLStates(sfRenderWindow* renderWindow);

func (self RenderWindow) Popglstates() void { 
    return C.sfRenderWindow_popGLStates(self.Cref);
}
            
// Reset the internal OpenGL states so that the target is ready for drawing
// This function can be used when you mix SFML drawing
// and direct OpenGL rendering, if you choose not to use
// pushGLStates/popGLStates. It makes sure that all OpenGL
// states needed by SFML are set, so that subsequent sfRenderWindow_draw*()
// calls will work as expected.
// \param renderWindow render window object
// void sfRenderWindow_resetGLStates(sfRenderWindow* renderWindow);

func (self RenderWindow) Resetglstates() void { 
    return C.sfRenderWindow_resetGLStates(self.Cref);
}
            
// Copy the current contents of a render window to an image
// This is a slow operation, whose main purpose is to make
// screenshots of the application. If you want to update an
// image with the contents of the window and then use it for
// drawing, you should rather use a sfTexture and its
// update(sfWindow*) function.
// You can also draw things directly to a texture with the
// sfRenderWindow class.
// \param renderWindow Render window object
// \return New image containing the captured contents
// sfImage* sfRenderWindow_capture(const sfRenderWindow* renderWindow);

func (self RenderWindow) Capture() *Image { 
    return C.sfRenderWindow_capture(self.Cref);
}
            
*/