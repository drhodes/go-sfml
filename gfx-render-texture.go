package sfml

// #cgo LDFLAGS:-lcsfml-graphics
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/Graphics/PrimitiveType.h>
// #include <SFML/Graphics/RenderStates.h>
// #include <SFML/Graphics/Vertex.h>
// #include <SFML/System/Vector2.h>
// #include <SFML/Graphics/RenderTexture.h>
import "C"

type RenderTexture struct {
	Cref *C.sfRenderTexture
}

// Construct a new render texture
// \param width       Width of the render texture
// \param height      Height of the render texture
// \param depthBuffer Do you want a depth-buffer attached? (useful only if you're doing 3D OpenGL on the rendertexture)
// \return A new sfRenderTexture object, or NULL if it failed
// sfRenderTexture* sfRenderTexture_create(unsigned int width, unsigned int height, sfBool depthBuffer);
func NewRenderTexture(width, height uint, depthBuffer bool) RenderTexture {
	ref := C.sfRenderTexture_create(C.uint(width), C.uint(height), Bool(depthBuffer))
	return RenderTexture{ref}
}

// Destroy an existing render texture
// \param renderTexture Render texture to destroy
// void sfRenderTexture_destroy(sfRenderTexture* renderTexture);
func (self RenderTexture) Destroy() {
	C.sfRenderTexture_destroy(self.Cref)
}

// Get the size of the rendering region of a render texture
// \param renderTexture Render texture object
// \return Size in pixels
// sfVector2u sfRenderTexture_getSize(const sfRenderTexture* renderTexture);
func (self RenderTexture) Size() (uint, uint) {
	v := C.sfRenderTexture_getSize(self.Cref)
	return uint(v.x), uint(v.y)
}

// Activate or deactivate a render texture as the current target for rendering
// \param renderTexture Render texture object
// \param active        sfTrue to activate, sfFalse to deactivate
// \return True if operation was successful, false otherwise
// sfBool sfRenderTexture_setActive(sfRenderTexture* renderTexture, sfBool active);
func (self RenderTexture) SetActive(active bool) bool {
	return C.sfRenderTexture_setActive(self.Cref, Bool(active)) == 1
}

// Update the contents of the target texture
// \param renderTexture Render texture object
// void sfRenderTexture_display(sfRenderTexture* renderTexture);
func (self RenderTexture) Display() {
	C.sfRenderTexture_display(self.Cref)
}

// Clear the rendertexture with the given color
// \param renderTexture Render texture object
// \param color         Fill color
// void sfRenderTexture_clear(sfRenderTexture* renderTexture, sfColor color);
func (self RenderTexture) Clear(color Color) {
	C.sfRenderTexture_clear(self.Cref, color.Cref)
}

/*
// Change the current active view of a render texture
// \param renderTexture Render texture object
// \param view          Pointer to the new view
// void sfRenderTexture_setView(sfRenderTexture* renderTexture, const sfView* view);

// Get the current active view of a render texture
// \param renderTexture Render texture object
// \return Current active view
// const sfView* sfRenderTexture_getView(const sfRenderTexture* renderTexture);

// Get the default view of a render texture
// \param renderTexture Render texture object
// \return Default view of the rendertexture
// const sfView* sfRenderTexture_getDefaultView(const sfRenderTexture* renderTexture);

// Get the viewport of a view applied to this target
// \param renderTexture Render texture object
// \param view          Target view
// \return Viewport rectangle, expressed in pixels in the current target
// sfIntRect sfRenderTexture_getViewport(const sfRenderTexture* renderTexture, const sfView* view);

// Convert a point in texture coordinates into view coordinates
// \param renderTexture Render texture object
// \param point         Point to convert, relative to the texture
// \param targetView    Target view to convert the point to (pass NULL to use the current view)
/// \return The converted point, in "world" units
// sfVector2f sfRenderTexture_convertCoords(const sfRenderTexture* renderTexture, sfVector2i point, const sfView* targetView);
*/

// Draw a drawable object to the render-target
// \param renderTexture Render texture object
// \param object        Object to draw
// \param states        Render states to use for drawing (NULL to use the default states)
// void sfRenderTexture_drawSprite(sfRenderTexture* renderTexture, const sfSprite* object, const sfRenderStates* states);
func (self RenderTexture) DrawSpriteDefault(obj Sprite) {
	C.sfRenderTexture_drawSprite(self.Cref, obj.Cref, nil)
}

// void sfRenderTexture_drawText(sfRenderTexture* renderTexture, const sfText* object, const sfRenderStates* states);
func (self RenderTexture) DrawTextDefault(obj Text) {
	C.sfRenderTexture_drawText(self.Cref, obj.Cref, nil)
}

// void sfRenderTexture_drawShape(sfRenderTexture* renderTexture, const sfShape* object, const sfRenderStates* states);

// void sfRenderTexture_drawCircleShape(sfRenderTexture* renderTexture, const sfCircleShape* object, const sfRenderStates* states);
func (self RenderTexture) DrawCircleShapeDefault(obj CircleShape) {
	C.sfRenderTexture_drawCircleShape(self.Cref, obj.Cref, nil)
}

// void sfRenderTexture_drawConvexShape(sfRenderTexture* renderTexture, const sfConvexShape* object, const sfRenderStates* states);

// void sfRenderTexture_drawRectangleShape(sfRenderTexture* renderTexture, const sfRectangleShape* object, const sfRenderStates* states);
func (self RenderTexture) DrawRectangleShapeDefault(obj RectangleShape) {
	C.sfRenderTexture_drawRectangleShape(self.Cref, obj.Cref, nil)
}

// void sfRenderTexture_drawVertexArray(sfRenderTexture* renderTexture, const sfVertexArray* object, const sfRenderStates* states);

////////////////////////////////////////////////////////////
/// \brief Draw primitives defined by an array of vertices to a render texture
///
/// \param renderTexture Render texture object
/// \param vertices      Pointer to the vertices
/// \param vertexCount   Number of vertices in the array
/// \param type          Type of primitives to draw
/// \param states        Render states to use for drawing (NULL to use the default states)
///
////////////////////////////////////////////////////////////
// void sfRenderTexture_drawPrimitives(sfRenderTexture* renderTexture,
//                                     const sfVertex* vertices, unsigned int vertexCount,
//                                     sfPrimitiveType type, const sfRenderStates* states);

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
// \param renderTexture Render texture object
// void sfRenderTexture_pushGLStates(sfRenderTexture* renderTexture);
func (self RenderTexture) PushGLStates() {
	C.sfRenderTexture_pushGLStates(self.Cref)
}

// Restore the previously saved OpenGL render states and matrices
// See the description of pushGLStates to get a detailed
// description of these functions.
// \param renderTexture Render texture object
// void sfRenderTexture_popGLStates(sfRenderTexture* renderTexture);
func (self RenderTexture) PopGLStates() {
	C.sfRenderTexture_popGLStates(self.Cref)
}

// Reset the internal OpenGL states so that the target is ready for drawing
// This function can be used when you mix SFML drawing
// and direct OpenGL rendering, if you choose not to use
// pushGLStates/popGLStates. It makes sure that all OpenGL
// states needed by SFML are set, so that subsequent sfRenderTexture_draw*()
// calls will work as expected.
// \param renderTexture Render texture object
// void sfRenderTexture_resetGLStates(sfRenderTexture* renderTexture);
func (self RenderTexture) ResetGLStates() {
	C.sfRenderTexture_resetGLStates(self.Cref)
}

// Get the target texture of a render texture
// \param renderTexture Render texture object
// \return Pointer to the target texture
// const sfTexture* sfRenderTexture_getTexture(const sfRenderTexture* renderTexture);
func (self RenderTexture) Texture() Texture {
	return Texture{C.sfRenderTexture_getTexture(self.Cref)}
}

// Enable or disable the smooth filter on a render texture
// \param renderTexture Render texture object
// \param smooth        sfTrue to enable smoothing, sfFalse to disable it
// void sfRenderTexture_setSmooth(sfRenderTexture* renderTexture, sfBool smooth);
func (self RenderTexture) SetSmooth(smooth bool) {
	C.sfRenderTexture_setSmooth(self.Cref, Bool(smooth))
}

// Tell whether the smooth filter is enabled or not for a render texture
// \param renderTexture Render texture object
// \return sfTrue if smoothing is enabled, sfFalse if it is disabled
// sfBool sfRenderTexture_isSmooth(const sfRenderTexture* renderTexture);
func (self RenderTexture) IsSmooth() bool {
	return C.sfRenderTexture_isSmooth(self.Cref) == 1
}
