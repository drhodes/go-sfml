package sfml

// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/ConvexShape.h>
// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/System/Vector2.h>
import "C"

type ConvexShape struct {
	Cref *C.sfConvexShape
}

// \brief Create a new convex shape
// \return A new sfConvexShape object, or NULL if it failed
// sfConvexShape* sfConvexShape_create(void);
func NewConvexShape() ConvexShape { 
    return ConvexShape{C.sfConvexShape_create()}
}

// \brief Copy an existing convex shape
// \param shape Shape to copy
// \return Copied object
// sfConvexShape* sfConvexShape_copy(sfConvexShape* shape);
func (self ConvexShape) Copy() ConvexShape { 
    return ConvexShape{C.sfConvexShape_copy(self.Cref)}
}
            
// \brief Destroy an existing convex Shape
// \param Shape Shape to delete
// void sfConvexShape_destroy(sfConvexShape* shape);
func (self ConvexShape) Destroy() { 
    C.sfConvexShape_destroy(self.Cref)
}
            
// \brief Set the position of a convex shape
// This function completely overwrites the previous position.
// See sfConvexShape_move to apply an offset based on the previous position instead.
// The default position of a circle Shape object is (0, 0).
// \param shape    Shape object
// \param position New position
// void sfConvexShape_setPosition(sfConvexShape* shape, sfVector2f position);
func (self ConvexShape) Setposition(position Vector2f) {
	C.sfConvexShape_setPosition(self.Cref, position.Cref)
}

// \brief Set the orientation of a convex shape
// This function completely overwrites the previous rotation.
// See sfConvexShape_rotate to add an angle based on the previous rotation instead.
// The default rotation of a circle Shape object is 0.
// \param shape Shape object
// \param angle New rotation, in degrees
// void sfConvexShape_setRotation(sfConvexShape* shape, float angle);
func (self ConvexShape) Setrotation(angle float32) { 
    C.sfConvexShape_setRotation(self.Cref, C.float(angle))
}
            
// \brief Set the scale factors of a convex shape
// This function completely overwrites the previous scale.
// See sfConvexShape_scale to add a factor based on the previous scale instead.
// The default scale of a circle Shape object is (1, 1).
// \param shape Shape object
// \param scale New scale factors
// void sfConvexShape_setScale(sfConvexShape* shape, sfVector2f scale);
func (self ConvexShape) SetScale(x, y float32) {
	C.sfConvexShape_setScale(self.Cref, C.sfVector2f{
		C.float(x), C.float(y),
	})
}

// \brief Set the local origin of a convex shape
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a circle Shape object is (0, 0).
// \param shape  Shape object
// \param origin New origin
// void sfConvexShape_setOrigin(sfConvexShape* shape, sfVector2f origin);
func (self ConvexShape) Setorigin(x, y float32) { 
    C.sfConvexShape_setOrigin(self.Cref, C.sfVector2f {
		C.float(x), C.float(y),
	})
}
            
// \brief Get the position of a convex shape
// \param shape Shape object
// \return Current position
// sfVector2f sfConvexShape_getPosition(const sfConvexShape* shape);
func (self ConvexShape) GetPosition() (x, y float32) { 
    p := C.sfConvexShape_getPosition(self.Cref);
	return float32(p.x), float32(p.y)
}
            
// \brief Get the orientation of a convex shape
// The rotation is always in the range [0, 360].
// \param shape Shape object
// \return Current rotation, in degrees
// float sfConvexShape_getRotation(const sfConvexShape* shape);
func (self ConvexShape) GetRotation() float32 { 
    return float32(C.sfConvexShape_getRotation(self.Cref))
}
          
// \brief Get the current scale of a convex shape
// \param shape Shape object
// \return Current scale factors
// sfVector2f sfConvexShape_getScale(const sfConvexShape* shape);
func (self ConvexShape) GetScale() (x, y float32) { 
    p := C.sfConvexShape_getScale(self.Cref)
	return float32(p.x), float32(p.y)
}
            
// \brief Get the local origin of a convex shape
// \param shape Shape object
// \return Current origin
// sfVector2f sfConvexShape_getOrigin(const sfConvexShape* shape);
func (self ConvexShape) GetOrigin() (x, y float32) { 
    p := C.sfConvexShape_getOrigin(self.Cref)
	return float32(p.x), float32(p.y)	
}
            
// \brief Move a convex shape by a given offset
// This function adds to the current position of the object,
// unlike sfConvexShape_setPosition which overwrites it.
// \param shape  Shape object
// \param offset Offset
// void sfConvexShape_move(sfConvexShape* shape, sfVector2f offset);
func (self ConvexShape) Move(x, y float32) { 
    C.sfConvexShape_move(self.Cref, C.sfVector2f {
		C.float(x), C.float(y),
	})
}
            
// \brief Rotate a convex shape
// This function adds to the current rotation of the object,
// unlike sfConvexShape_setRotation which overwrites it.
// \param shape Shape object
// \param angle Angle of rotation, in degrees
// void sfConvexShape_rotate(sfConvexShape* shape, float angle);
func (self ConvexShape) Rotate(angle float32) { 
    C.sfConvexShape_rotate(self.Cref, C.float(angle))
}

// \brief Scale a convex shape
// This function multiplies the current scale of the object,
// unlike sfConvexShape_setScale which overwrites it.
// \param shape   Shape object
// \param factors Scale factors
// void sfConvexShape_scale(sfConvexShape* shape, sfVector2f factors);
func (self ConvexShape) Scale(x, y float32) {
	C.sfConvexShape_scale(self.Cref, C.sfVector2f{
		C.float(x), C.float(y),
	})
}

// \brief Get the combined transform of a convex shape
// \param shape shape object
// \return Transform combining the position/rotation/scale/origin of the object
// const sfTransform* sfConvexShape_getTransform(const sfConvexShape* shape);
func (self ConvexShape) GetTransform() Transform {
	t := C.sfConvexShape_getTransform(self.Cref)
	return Transform{&t}
}

// \brief Get the inverse of the combined transform of a convex shape
// \param shape shape object
// \return Inverse of the combined transformations applied to the object
// const sfTransform* sfConvexShape_getInverseTransform(const sfConvexShape* shape);
func (self ConvexShape) GetInverseTransform() Transform {
	t := C.sfConvexShape_getInverseTransform(self.Cref)
	return Transform{&t}
}

// \brief Change the source texture of a convex shape
// The \a texture argument refers to a texture that must
// exist as long as the shape uses it. Indeed, the shape
// doesn't store its own copy of the texture, but rather keeps
// a pointer to the one that you passed to this function.
// If the source texture is destroyed and the shape tries to
// use it, the behaviour is undefined.
// \a texture can be NULL to disable texturing.
// If \a resetRect is true, the TextureRect property of
// the shape is automatically adjusted to the size of the new
// texture. If it is false, the texture rect is left unchanged.
// \param shape     Shape object
// \param texture   New texture
// \param resetRect Should the texture rect be reset to the size of the new texture?
// void sfConvexShape_setTexture(sfConvexShape* shape, const sfTexture* texture, sfBool resetRect);
func (self ConvexShape) SetTexture(txt Texture, resetRect bool) { 
    C.sfConvexShape_setTexture(self.Cref, txt.Cref, Bool(resetRect))
}
            
// \brief Set the sub-rectangle of the texture that a convex shape will display
// The texture rect is useful when you don't want to display
// the whole texture, but rather a part of it.
// By default, the texture rect covers the entire texture.
// \param shape Shape object
// \param rect  Rectangle defining the region of the texture to display
// void sfConvexShape_setTextureRect(sfConvexShape* shape, sfIntRect rect);
func (self ConvexShape) Settexturerect(rect IntRect) {
	C.sfConvexShape_setTextureRect(self.Cref, *rect.Cref)
}

// \brief Set the fill color of a convex shape
// This color is modulated (multiplied) with the shape's
// texture if any. It can be used to colorize the shape,
// or change its global opacity.
// You can use sfTransparent to make the inside of
// the shape transparent, and have the outline alone.
// By default, the shape's fill color is opaque white.
// \param shape Shape object
// \param color New color of the shape
// void sfConvexShape_setFillColor(sfConvexShape* shape, sfColor color);
func (self ConvexShape) SetFillColor(color Color) { 
	C.sfConvexShape_setFillColor(self.Cref, color.Cref)
}
            
// \brief Set the outline color of a convex shape
// You can use sfTransparent to disable the outline.
// By default, the shape's outline color is opaque white.
// \param shape Shape object
// \param color New outline color of the shape
// void sfConvexShape_setOutlineColor(sfConvexShape* shape, sfColor color);
func (self ConvexShape) SetOutlineColor(color Color) { 
    C.sfConvexShape_setOutlineColor(self.Cref, color.Cref)
}
            
// \brief Set the thickness of a convex shape's outline
// This number cannot be negative. Using zero disables
// the outline.
// By default, the outline thickness is 0.
// \param shape     Shape object
// \param thickness New outline thickness
// void sfConvexShape_setOutlineThickness(sfConvexShape* shape, float thickness);
func (self ConvexShape) SetOutlineThickness(thickness float32) {
	C.sfConvexShape_setOutlineThickness(self.Cref, C.float(thickness))
}

// \brief Get the source texture of a convex shape
// If the shape has no source texture, a NULL pointer is returned.
// The returned pointer is const, which means that you can't
// modify the texture when you retrieve it with this function.
// \param shape Shape object
// \return Pointer to the shape's texture
// const sfTexture* sfConvexShape_getTexture(const sfConvexShape* shape);
func (self ConvexShape) GetTexture() Texture { 
    t := C.sfConvexShape_getTexture(self.Cref)
	return Texture{t}
}
            
// \brief Get the sub-rectangle of the texture displayed by a convex shape
// \param shape Shape object
// \return Texture rectangle of the shape
// sfIntRect sfConvexShape_getTextureRect(const sfConvexShape* shape);
func (self ConvexShape) GetTextureRect() IntRect { 
	r := C.sfConvexShape_getTextureRect(self.Cref) 
    return IntRect{&r}
}

// \brief Get the fill color of a convex shape
// \param shape Shape object
// \return Fill color of the shape
// sfColor sfConvexShape_getFillColor(const sfConvexShape* shape);
func (self ConvexShape) GetFillColor() Color { 
    return Color{ C.sfConvexShape_getFillColor(self.Cref) } 
}
            
// \brief Get the outline color of a convex shape
// \param shape Shape object
// \return Outline color of the shape
// sfColor sfConvexShape_getOutlineColor(const sfConvexShape* shape);
func (self ConvexShape) GetOutlineColor() Color { 
    return Color{ C.sfConvexShape_getOutlineColor(self.Cref) }
}
            
// \brief Get the outline thickness of a convex shape
// \param shape Shape object
// \return Outline thickness of the shape
// float sfConvexShape_getOutlineThickness(const sfConvexShape* shape);
func (self ConvexShape) GetOutlineThickness() float32 { 
    return float32(C.sfConvexShape_getOutlineThickness(self.Cref))
}
            
// \brief Get the total number of points of a convex shape
// \param shape Shape object
// \return Number of points of the shape
// unsigned int sfConvexShape_getPointCount(const sfConvexShape* shape);
func (self ConvexShape) GetPointCount() uint {
    return uint(C.sfConvexShape_getPointCount(self.Cref))
}
            
// \brief Get a point of a convex shape
// The result is undefined if \a index is out of the valid range.
// \param shape Shape object
// \param index Index of the point to get, in range [0 .. getPointCount() - 1]
// \return Index-th point of the shape
// sfVector2f sfConvexShape_getPoint(const sfConvexShape* shape, unsigned int index);
func (self ConvexShape) GetPoint(index uint) (x, y float32) { 
    p := C.sfConvexShape_getPoint(self.Cref, C.uint(index))
	return float32(p.x), float32(p.y)
}
            
// \brief Set the number of points of a convex shap
// \a count must be greater than 2 to define a valid shape.
// \param shape Shape object
// \param count New number of points of the shape
// void sfConvexShape_setPointCount(sfConvexShape* shape, unsigned int count);
func (self ConvexShape) SetPointCount(count uint) { 
    C.sfConvexShape_setPointCount(self.Cref, C.uint(count))	
}
           
// \brief Set the position of a point in a convex shape
// Don't forget that the polygon must remain convex, and
// the points need to stay ordered!
// setPointCount must be called first in order to set the total
// number of points. The result is undefined if \a index is out
// of the valid range.
// \param shape Shape object
// \param index Index of the point to change, in range [0 .. GetPointCount() - 1]
// \param point New point
// void sfConvexShape_setPoint(sfConvexShape* shape, unsigned int index, sfVector2f point);
func (self ConvexShape) SetPoint(index uint, x, y float32) {
	C.sfConvexShape_setPoint(
		self.Cref,
		C.uint(index),
		C.sfVector2f{C.float(x), C.float(y),
	})
}

// \brief Get the local bounding rectangle of a convex shape
// The returned rectangle is in local coordinates, which means
// that it ignores the transformations (translation, rotation,
// scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// entity in the entity's coordinate system.
// \param shape Shape object
// \return Local bounding rectangle of the entity
// sfFloatRect sfConvexShape_getLocalBounds(const sfConvexShape* shape);
func (self ConvexShape) GetLocalBounds() FloatRect { 
	r := C.sfConvexShape_getLocalBounds(self.Cref)
    return FloatRect{&r}
}
            
// \brief Get the global bounding rectangle of a convex shape
// The returned rectangle is in global coordinates, which means
// that it takes in account the transformations (translation,
// rotation, scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// sprite in the global 2D world's coordinate system.
// \param shape Shape object
// \return Global bounding rectangle of the entity
// sfFloatRect sfConvexShape_getGlobalBounds(const sfConvexShape* shape);
func (self ConvexShape) GetGlobalBounds() FloatRect { 
	r := C.sfConvexShape_getGlobalBounds(self.Cref);
    return FloatRect{&r}
}
