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
import "C"

import (
	"unsafe"
	"github.com/drhodes/go-sfml/win"
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
func NewRenderWindow(mode win.VideoMode, title string, style uint32, settings win.ContextSettings) RenderWindow {
	mptr := unsafe.Pointer(&mode.Cref)
	mp := (*C.sfVideoMode)(mptr)

	sptr := unsafe.Pointer(&settings.Cref)
	sp := (*C.sfContextSettings)(sptr)

	ref := C.sfRenderWindow_create(*mp, C.CString(title), C.sfUint32(style), sp)
	return RenderWindow{ref}
}

func NewRenderWindowDefault(mode win.VideoMode, title string, style uint32) RenderWindow {
	mptr := unsafe.Pointer(&mode.Cref)
	mp := (*C.sfVideoMode)(mptr)

	ref := C.sfRenderWindow_create(*mp, C.CString(title), C.sfUint32(style), nil)
	return RenderWindow{ref}
}



// \brief Construct a render window from an existing control
//
// \param handle   Platform-specific handle of the control
// \param settings Creation settings (pass NULL to use default values)
//
// sfRenderWindow* sfRenderWindow_createFromHandle(sfWindowHandle handle, const sfContextSettings* settings);
/*
func RenderWindowFromhandle(handle win.WindowHandle, settings win.ContextSettings) RenderWindow {
	mp := (*C.sfVideoMode)(unsafe.Pointer(&handle))

	return RenderWindow {
		C.sfRenderWindow_createFromHandle(handle.Cref, settings.Cref),
	}
}
*/


// \brief Destroy an existing render window
//
// \param renderWindow Render window to destroy
//
// void sfRenderWindow_destroy(sfRenderWindow* renderWindow);

func (self RenderWindow) Destroy() {
	C.sfRenderWindow_destroy(self.Cref)
}

// \brief Close a render window (but doesn't destroy the internal data)
//
// \param renderWindow Render window to close
//
// void sfRenderWindow_close(sfRenderWindow* renderWindow);

func (self RenderWindow) Close() {
	C.sfRenderWindow_close(self.Cref)
}

// \brief Tell whether or not a render window is opened
//
// \param renderWindow Render window object
//
// sfBool sfRenderWindow_isOpen(const sfRenderWindow* renderWindow);

func (self RenderWindow) IsOpen() bool {
	b := C.sfRenderWindow_isOpen(self.Cref)
	return b == 1
}

// \brief Get the creation settings of a render window
//
// \param renderWindow Render window object
//
// \return Settings used to create the window
//
// sfContextSettings sfRenderWindow_getSettings(const sfRenderWindow* renderWindow);

/*
func (self RenderWindow) Getsettings() win.ContextSettings {
	ref := C.sfRenderWindow_getSettings(self.Cref)
	return win.ContextSettings{&ref}
}
*/

// \brief Get the event on top of events stack of a render window, if any, and pop it
//
// \param renderWindow Render window object
// \param event        Event to fill, if any
//
// \return sfTrue if an event was returned, sfFalse if events stack was empty
//
// sfBool sfRenderWindow_pollEvent(sfRenderWindow* renderWindow, sfEvent* event);

func (self RenderWindow) Pollevent() interface{} {
	// ok if got event.
	e := win.NewEvent()
	ok := C.sfRenderWindow_pollEvent(self.Cref, e.Cref) == 1
	if ok {
		// look at the first byte, it's the event type
		et := win.EventType((*e.Cref)[0])
		switch et {
		case win.EvtClosed:
		case win.EvtResized:
		case win.EvtLostFocus:
		case win.EvtGainedFocus:
		case win.EvtTextEntered:
			return e.ToTextEvent()
		case win.EvtKeyPressed, win.EvtKeyReleased:
			return e.ToKeyEvent()
		case win.EvtMouseWheelMoved:
			return e.ToMouseWheelEvent()
		case win.EvtMouseButtonPressed, win.EvtMouseButtonReleased:
			return e.ToMouseButtonEvent()
		case win.EvtMouseMoved, win.EvtMouseEntered, win.EvtMouseLeft:
			return e.ToMouseMoveEvent()
		case win.EvtJoystickButtonPressed, win.EvtJoystickButtonReleased, win.EvtJoystickMoved:
			return e.ToJoystickMoveEvent()
		case win.EvtJoystickConnected:
		case win.EvtJoystickDisconnected:
			return e.ToJoystickConnectEvent()
		case win.EvtNone:
		}
	}
	return win.EvtNone
}

/*
// \brief Wait for an event and return it
//
// \param renderWindow Render window object
// \param event        Event to fill
//
// \return sfFalse if an error occured
//
// sfBool sfRenderWindow_waitEvent(sfRenderWindow* renderWindow, sfEvent* event);

func (self RenderWindow) Waitevent(event *Event) Bool { 
    return C.sfRenderWindow_waitEvent(self.Cref, sf*event(event));
}
            
// \brief Get the position of a render window
//
// \param renderWindow Render window object
//
// \return Position in pixels
//
// sfVector2i sfRenderWindow_getPosition(const sfRenderWindow* renderWindow);

func (self RenderWindow) Getposition() Vector2i { 
    return C.sfRenderWindow_getPosition(self.Cref);
}
            
// \brief Change the position of a render window on screen
//
// Only works for top-level windows
//
// \param renderWindow Render window object
// \param position     New position, in pixels
//
// void sfRenderWindow_setPosition(sfRenderWindow* renderWindow, sfVector2i position);

func (self RenderWindow) Setposition(position Vector2i) void { 
    return C.sfRenderWindow_setPosition(self.Cref, sfVector2i(position));
}
            
// \brief Get the size of the rendering region of a render window
//
// \param renderWindow Render window object
//
// \return Size in pixels
//
// sfVector2u sfRenderWindow_getSize(const sfRenderWindow* renderWindow);

func (self RenderWindow) Getsize() Vector2u { 
    return C.sfRenderWindow_getSize(self.Cref);
}
            
// \brief Change the size of the rendering region of a render window
//
// \param renderWindow Render window object
// \param size         New size, in pixels
//
// void sfRenderWindow_setSize(sfRenderWindow* renderWindow, sfVector2u size);

func (self RenderWindow) Setsize(size Vector2u) void { 
    return C.sfRenderWindow_setSize(self.Cref, sfVector2u(size));
}
            
// \brief Change the title of a render window
//
// \param renderWindow Render window object
// \param title        New title
//
// void sfRenderWindow_setTitle(sfRenderWindow* renderWindow, const char* title);

func (self RenderWindow) Settitle(title *char ) void { 
    return C.sfRenderWindow_setTitle(self.Cref, sf(*char));
}
            
// \brief Change a render window's icon
//
// \param renderWindow Renderw indow object
// \param width        Icon's width, in pixels
// \param height       Icon's height, in pixels
// \param pixels       Pointer to the pixels in memory, format must be RGBA 32 bits
//
// void sfRenderWindow_setIcon(sfRenderWindow* renderWindow, unsigned int width, unsigned int height, const sfUint8* pixels);

func (self RenderWindow) Seticon(width int , height int , pixels *Uint8 ) void { 
    return C.sfRenderWindow_setIcon(self.Cref, sf(int), sf(int), sf(*Uint8));
}
            
// \brief Show or hide a render window
//
// \param renderWindow Render window object
// \param visible      sfTrue to show the window, sfFalse to hide it
//
// void sfRenderWindow_setVisible(sfRenderWindow* renderWindow, sfBool visible);

func (self RenderWindow) Setvisible(visible Bool) void { 
    return C.sfRenderWindow_setVisible(self.Cref, sfBool(visible));
}
            
// \brief Show or hide the mouse cursor on a render window
//
// \param renderWindow Render window object
// \param show         sfTrue to show, sfFalse to hide
//
// void sfRenderWindow_setMouseCursorVisible(sfRenderWindow* renderWindow, sfBool show);

func (self RenderWindow) Setmousecursorvisible(show Bool) void { 
    return C.sfRenderWindow_setMouseCursorVisible(self.Cref, sfBool(show));
}
            
// \brief Enable / disable vertical synchronization on a render window
//
// \param renderWindow Render window object
// \param enabled      sfTrue to enable v-sync, sfFalse to deactivate
//
// void sfRenderWindow_setVerticalSyncEnabled(sfRenderWindow* renderWindow, sfBool enabled);

func (self RenderWindow) Setverticalsyncenabled(enabled Bool) void { 
    return C.sfRenderWindow_setVerticalSyncEnabled(self.Cref, sfBool(enabled));
}
            
// \brief Enable or disable automatic key-repeat for keydown events
//
// Automatic key-repeat is enabled by default
//
// \param renderWindow Render window object
// \param enabled      sfTrue to enable, sfFalse to disable
//
// void sfRenderWindow_setKeyRepeatEnabled(sfRenderWindow* renderWindow, sfBool enabled);

func (self RenderWindow) Setkeyrepeatenabled(enabled Bool) void { 
    return C.sfRenderWindow_setKeyRepeatEnabled(self.Cref, sfBool(enabled));
}
            
// \brief Activate or deactivate a render window as the current target for rendering
//
// \param renderWindow Render window object
// \param active       sfTrue to activate, sfFalse to deactivate
//
// \return True if operation was successful, false otherwise
//
// sfBool sfRenderWindow_setActive(sfRenderWindow* renderWindow, sfBool active);

func (self RenderWindow) Setactive(active Bool) Bool { 
    return C.sfRenderWindow_setActive(self.Cref, sfBool(active));
}
            
// \brief Display a render window on screen
//
// \param renderWindow Render window object
//
// void sfRenderWindow_display(sfRenderWindow* renderWindow);

func (self RenderWindow) Display() void { 
    return C.sfRenderWindow_display(self.Cref);
}
*/

// \brief Limit the framerate to a maximum fixed frequency for a render window
//
// \param renderWindow Render window object
// \param limit        Framerate limit, in frames per seconds (use 0 to disable limit)
//
// void sfRenderWindow_setFramerateLimit(sfRenderWindow* renderWindow, unsigned int limit);

func (self RenderWindow) SetFramerateLimit(limit uint) {
	C.sfRenderWindow_setFramerateLimit(self.Cref, C.uint(limit))
}

/*

// \brief Change the joystick threshold, ie. the value below which no move event will be generated
//
// \param renderWindow Render window object
// \param threshold    New threshold, in range [0, 100]
//
// void sfRenderWindow_setJoystickThreshold(sfRenderWindow* renderWindow, float threshold);

func (self RenderWindow) Setjoystickthreshold(threshold float) void { 
    return C.sfRenderWindow_setJoystickThreshold(self.Cref, sfFloat(threshold));
}
            
// \brief Retrieve the OS-specific handle of a render window
//
// \param renderWindow Render window object
//
// \return Window handle
//
// sfWindowHandle sfRenderWindow_getSystemHandle(const sfRenderWindow* renderWindow);

func (self RenderWindow) Getsystemhandle() WindowHandle { 
    return C.sfRenderWindow_getSystemHandle(self.Cref);
}
            
// \brief Clear a render window with the given color
//
// \param renderWindow Render window object
// \param color        Fill color
//
// void sfRenderWindow_clear(sfRenderWindow* renderWindow, sfColor color);

func (self RenderWindow) Clear(color Color) void { 
    return C.sfRenderWindow_clear(self.Cref, sfColor(color));
}
            
// \brief Change the current active view of a render window
//
// \param renderWindow Render window object
// \param view         Pointer to the new view
//
// void sfRenderWindow_setView(sfRenderWindow* renderWindow, const sfView* view);

func (self RenderWindow) Setview(view *View ) void { 
    return C.sfRenderWindow_setView(self.Cref, sf(*View));
}
            
// \brief Get the current active view of a render window
//
// \param renderWindow Render window object
//
// \return Current active view
//
// const sfView* sfRenderWindow_getView(const sfRenderWindow* renderWindow);

func (self *View) *View(RenderWindow_getView)  { 
    return C.sf*View(self.Cref, sf(*View));
}
            
// \brief Get the default view of a render window
//
// \param renderWindow Render window object
//
// \return Default view of the render window
//
// const sfView* sfRenderWindow_getDefaultView(const sfRenderWindow* renderWindow);

func (self *View) *View(RenderWindow_getDefaultView)  { 
    return C.sf*View(self.Cref, sf(*View));
}
            
// \brief Get the viewport of a view applied to this target
//
// \param renderWindow Render window object
// \param view         Target view
//
// \return Viewport rectangle, expressed in pixels in the current target
//
// sfIntRect sfRenderWindow_getViewport(const sfRenderWindow* renderWindow, const sfView* view);

func (self RenderWindow) Getviewport(view *View ) IntRect { 
    return C.sfRenderWindow_getViewport(self.Cref, sf(*View));
}
            
// \brief Convert a point in window coordinates into view coordinates
//
// \param renderWindow Render window object
// \param point        Point to convert, relative to the window
// \param targetView   Target view to convert the point to (pass NULL to use the current view)
//
// \return The converted point, in "world" units
//
// sfVector2f sfRenderWindow_convertCoords(const sfRenderWindow* renderWindow, sfVector2i point, const sfView* targetView);

func (self RenderWindow) Convertcoords(point Vector2i, targetView *View ) Vector2f { 
    return C.sfRenderWindow_convertCoords(self.Cref, sfVector2i(point), sf(*View));
}
            
// \brief Draw a drawable object to the render-target
//
// \param renderWindow render window object
// \param object       Object to draw
// \param states       Render states to use for drawing (NULL to use the default states)
//
// void sfRenderWindow_drawSprite(sfRenderWindow* renderWindow, const sfSprite* object, const sfRenderStates* states);
*/

func (self RenderWindow) DrawSprite(object *Sprite) {
	C.sfRenderWindow_drawSprite(self.Cref, object.Cref, nil)
}

/*
// void sfRenderWindow_drawText(sfRenderWindow* renderWindow, const sfText* object, const sfRenderStates* states);

func (self RenderWindow) Drawtext(object *Text , states *RenderStates ) void { 
    return C.sfRenderWindow_drawText(self.Cref, sf(*Text), sf(*RenderStates));
}
            
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
            
// \brief Draw primitives defined by an array of vertices to a render window
//
// \param renderWindow render window object
// \param vertices     Pointer to the vertices
// \param vertexCount  Number of vertices in the array
// \param type         Type of primitives to draw
// \param states       Render states to use for drawing (NULL to use the default states)
//
// void sfRenderWindow_drawPrimitives(sfRenderWindow* renderWindow, const sfVertex* vertices, unsigned int vertexCount, sfPrimitiveType type, const sfRenderStates* states);

func (self RenderWindow) Drawprimitives(vertices *Vertex , vertexCount int , type PrimitiveType, states *RenderStates ) void { 
    return C.sfRenderWindow_drawPrimitives(self.Cref, sf(*Vertex), sf(int), sfPrimitivetype(type), sf(*RenderStates));
}
            
// \brief Save the current OpenGL render states and matrices
//
// This function can be used when you mix SFML drawing
// and direct OpenGL rendering. Combined with popGLStates,
// it ensures that:
// \li SFML's internal states are not messed up by your OpenGL code
// \li your OpenGL states are not modified by a call to a SFML function
//
// Note that this function is quite expensive: it saves all the
// possible OpenGL states and matrices, even the ones you
// don't care about. Therefore it should be used wisely.
// It is provided for convenience, but the best results will
// be achieved if you handle OpenGL states yourself (because
// you know which states have really changed, and need to be
// saved and restored). Take a look at the resetGLStates
// function if you do so.
//
// \param renderWindow render window object
//
// void sfRenderWindow_pushGLStates(sfRenderWindow* renderWindow);

func (self RenderWindow) Pushglstates() void { 
    return C.sfRenderWindow_pushGLStates(self.Cref);
}
            
// \brief Restore the previously saved OpenGL render states and matrices
//
// See the description of pushGLStates to get a detailed
// description of these functions.
//
// \param renderWindow render window object
//
// void sfRenderWindow_popGLStates(sfRenderWindow* renderWindow);

func (self RenderWindow) Popglstates() void { 
    return C.sfRenderWindow_popGLStates(self.Cref);
}
            
// \brief Reset the internal OpenGL states so that the target is ready for drawing
//
// This function can be used when you mix SFML drawing
// and direct OpenGL rendering, if you choose not to use
// pushGLStates/popGLStates. It makes sure that all OpenGL
// states needed by SFML are set, so that subsequent sfRenderWindow_draw*()
// calls will work as expected.
//
// \param renderWindow render window object
//
// void sfRenderWindow_resetGLStates(sfRenderWindow* renderWindow);

func (self RenderWindow) Resetglstates() void { 
    return C.sfRenderWindow_resetGLStates(self.Cref);
}
            
// \brief Copy the current contents of a render window to an image
//
// This is a slow operation, whose main purpose is to make
// screenshots of the application. If you want to update an
// image with the contents of the window and then use it for
// drawing, you should rather use a sfTexture and its
// update(sfWindow*) function.
// You can also draw things directly to a texture with the
// sfRenderWindow class.
//
// \param renderWindow Render window object
//
// \return New image containing the captured contents
//
// sfImage* sfRenderWindow_capture(const sfRenderWindow* renderWindow);

func (self RenderWindow) Capture() *Image { 
    return C.sfRenderWindow_capture(self.Cref);
}
            
*/
