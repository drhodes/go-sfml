package sfml

// #include <SFML/Graphics/Shape.h>
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/System/Vector2.h>


type Shape struct {
	Cref *C.sfShape
}

// Create a new shape
// \param getPointCount Callback that provides the point count of the shape
// \param getPoint      Callback that provides the points of the shape
// \param userData      Data to pass to the callback functions
// \return A new sfShape object
// sfShape* sfShape_create(sfShapeGetPointCountCallback getPointCount, sfShapeGetPointCallback getPoint, void* userData);
func NewShape(getPoint ShapeGetPointCallback, userData *void) *Shape { 
    return C.sfShape_create(self.Cref, sfShapegetpointcallback(getPoint), sf*void(userData));
}
            
// Destroy an existing shape
// \param Shape Shape to delete
// void sfShape_destroy(sfShape* shape);
func (self Shape) Destroy() void { 
    return C.sfShape_destroy(self.Cref);
}
            
// Set the position of a shape
// This function completely overwrites the previous position.
// See sfShape_move to apply an offset based on the previous position instead.
// The default position of a circle Shape object is (0, 0).
// \param shape    Shape object
// \param position New position
// void sfShape_setPosition(sfShape* shape, sfVector2f position);
func (self Shape) Setposition(position Vector2f) void { 
    return C.sfShape_setPosition(self.Cref, sfVector2f(position));
}
            
// Set the orientation of a shape
// This function completely overwrites the previous rotation.
// See sfShape_rotate to add an angle based on the previous rotation instead.
// The default rotation of a circle Shape object is 0.
// \param shape Shape object
// \param angle New rotation, in degrees
// void sfShape_setRotation(sfShape* shape, float angle);
func (self Shape) Setrotation(angle float) void { 
    return C.sfShape_setRotation(self.Cref, sfFloat(angle));
}
            
// Set the scale factors of a shape
// This function completely overwrites the previous scale.
// See sfShape_scale to add a factor based on the previous scale instead.
// The default scale of a circle Shape object is (1, 1).
// \param shape Shape object
// \param scale New scale factors
// void sfShape_setScale(sfShape* shape, sfVector2f scale);
func (self Shape) Setscale(scale Vector2f) void { 
    return C.sfShape_setScale(self.Cref, sfVector2f(scale));
}
            
// Set the local origin of a shape
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a circle Shape object is (0, 0).
// \param shape  Shape object
// \param origin New origin
// void sfShape_setOrigin(sfShape* shape, sfVector2f origin);
func (self Shape) Setorigin(origin Vector2f) void { 
    return C.sfShape_setOrigin(self.Cref, sfVector2f(origin));
}
            
// Get the position of a shape
// \param shape Shape object
// \return Current position
// sfVector2f sfShape_getPosition(const sfShape* shape);
func (self Shape) Getposition() Vector2f { 
    return C.sfShape_getPosition(self.Cref);
}
            
// Get the orientation of a shape
// The rotation is always in the range [0, 360].
// \param shape Shape object
// \return Current rotation, in degrees
// float sfShape_getRotation(const sfShape* shape);
func (self Shape) Getrotation() float { 
    return C.sfShape_getRotation(self.Cref);
}
            
// Get the current scale of a shape
// \param shape Shape object
// \return Current scale factors
// sfVector2f sfShape_getScale(const sfShape* shape);
func (self Shape) Getscale() Vector2f { 
    return C.sfShape_getScale(self.Cref);
}
            
// Get the local origin of a shape
// \param shape Shape object
// \return Current origin
// sfVector2f sfShape_getOrigin(const sfShape* shape);
func (self Shape) Getorigin() Vector2f { 
    return C.sfShape_getOrigin(self.Cref);
}
            
// Move a shape by a given offset
// This function adds to the current position of the object,
// unlike sfShape_setPosition which overwrites it.
// \param shape  Shape object
// \param offset Offset
// void sfShape_move(sfShape* shape, sfVector2f offset);
func (self Shape) Move(offset Vector2f) void { 
    return C.sfShape_move(self.Cref, sfVector2f(offset));
}
            
// Rotate a shape
// This function adds to the current rotation of the object,
// unlike sfShape_setRotation which overwrites it.
// \param shape Shape object
// \param angle Angle of rotation, in degrees
// void sfShape_rotate(sfShape* shape, float angle);
func (self Shape) Rotate(angle float) void { 
    return C.sfShape_rotate(self.Cref, sfFloat(angle));
}
            
// Scale a shape
// This function multiplies the current scale of the object,
// unlike sfShape_setScale which overwrites it.
// \param shape   Shape object
// \param factors Scale factors
// void sfShape_scale(sfShape* shape, sfVector2f factors);
func (self Shape) Scale(factors Vector2f) void { 
    return C.sfShape_scale(self.Cref, sfVector2f(factors));
}
            
// Get the combined transform of a shape
// \param shape shape object
// \return Transform combining the position/rotation/scale/origin of the object
// const sfTransform* sfShape_getTransform(const sfShape* shape);
func (self *Transform) *Transform(Shape_getTransform)  { 
    return C.sf*Transform(self.Cref, sfVector2f(factors));
}
            
// Get the inverse of the combined transform of a shape
// \param shape shape object
// \return Inverse of the combined transformations applied to the object
// const sfTransform* sfShape_getInverseTransform(const sfShape* shape);
func (self *Transform) *Transform(Shape_getInverseTransform)  { 
    return C.sf*Transform(self.Cref, sfVector2f(factors));
}
            
// Change the source texture of a shape
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
// void sfShape_setTexture(sfShape* shape, const sfTexture* texture, sfBool resetRect);
func (self Shape) Settexture(texture *Texture , resetRect Bool) void { 
    return C.sfShape_setTexture(self.Cref, sf(*Texture), sfBool(resetRect));
}
            
// Set the sub-rectangle of the texture that a shape will display
// The texture rect is useful when you don't want to display
// the whole texture, but rather a part of it.
// By default, the texture rect covers the entire texture.
// \param shape Shape object
// \param rect  Rectangle defining the region of the texture to display
// void sfShape_setTextureRect(sfShape* shape, sfIntRect rect);
func (self Shape) Settexturerect(rect IntRect) void { 
    return C.sfShape_setTextureRect(self.Cref, sfIntrect(rect));
}
            
// Set the fill color of a shape
// This color is modulated (multiplied) with the shape's
// texture if any. It can be used to colorize the shape,
// or change its global opacity.
// You can use sfTransparent to make the inside of
// the shape transparent, and have the outline alone.
// By default, the shape's fill color is opaque white.
// \param shape Shape object
// \param color New color of the shape
// void sfShape_setFillColor(sfShape* shape, sfColor color);
func (self Shape) Setfillcolor(color Color) void { 
    return C.sfShape_setFillColor(self.Cref, sfColor(color));
}
            
// Set the outline color of a shape
// You can use sfTransparent to disable the outline.
// By default, the shape's outline color is opaque white.
// \param shape Shape object
// \param color New outline color of the shape
// void sfShape_setOutlineColor(sfShape* shape, sfColor color);
func (self Shape) Setoutlinecolor(color Color) void { 
    return C.sfShape_setOutlineColor(self.Cref, sfColor(color));
}
            
// Set the thickness of a shape's outline
// This number cannot be negative. Using zero disables
// the outline.
// By default, the outline thickness is 0.
// \param shape     Shape object
// \param thickness New outline thickness
// void sfShape_setOutlineThickness(sfShape* shape, float thickness);
func (self Shape) Setoutlinethickness(thickness float) void { 
    return C.sfShape_setOutlineThickness(self.Cref, sfFloat(thickness));
}
            
// Get the source texture of a shape
// If the shape has no source texture, a NULL pointer is returned.
// The returned pointer is const, which means that you can't
// modify the texture when you retrieve it with this function.
// \param shape Shape object
// \return Pointer to the shape's texture
// const sfTexture* sfShape_getTexture(const sfShape* shape);
func (self *Texture) *Texture(Shape_getTexture)  { 
    return C.sf*Texture(self.Cref, sfFloat(thickness));
}
            
// Get the sub-rectangle of the texture displayed by a shape
// \param shape Shape object
// \return Texture rectangle of the shape
// sfIntRect sfShape_getTextureRect(const sfShape* shape);
func (self Shape) Gettexturerect() IntRect { 
    return C.sfShape_getTextureRect(self.Cref);
}
            
// Get the fill color of a shape
// \param shape Shape object
// \return Fill color of the shape
// sfColor sfShape_getFillColor(const sfShape* shape);
func (self Shape) Getfillcolor() Color { 
    return C.sfShape_getFillColor(self.Cref);
}
            
// Get the outline color of a shape
// \param shape Shape object
// \return Outline color of the shape
// sfColor sfShape_getOutlineColor(const sfShape* shape);
func (self Shape) Getoutlinecolor() Color { 
    return C.sfShape_getOutlineColor(self.Cref);
}
            
// Get the outline thickness of a shape
// \param shape Shape object
// \return Outline thickness of the shape
// float sfShape_getOutlineThickness(const sfShape* shape);
func (self Shape) Getoutlinethickness() float { 
    return C.sfShape_getOutlineThickness(self.Cref);
}
            
// Get the total number of points of a shape
// \param shape Shape object
// \return Number of points of the shape
// unsigned int sfShape_getPointCount(const sfShape* shape);
func (self int) int(Shape_getPointCount)  { 
    return C.sfint(self.Cref);
}
            
// Get a point of a shape
// The result is undefined if \a index is out of the valid range.
// \param shape Shape object
// \param index Index of the point to get, in range [0 .. getPointCount() - 1]
// \return Index-th point of the shape
// sfVector2f sfShape_getPoint(const sfShape* shape, unsigned int index);
func (self Shape) Getpoint(index int ) Vector2f { 
    return C.sfShape_getPoint(self.Cref, sf(int));
}
            
// Get the local bounding rectangle of a shape
// The returned rectangle is in local coordinates, which means
// that it ignores the transformations (translation, rotation,
// scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// entity in the entity's coordinate system.
// \param shape Shape object
// \return Local bounding rectangle of the entity
// sfFloatRect sfShape_getLocalBounds(const sfShape* shape);
func (self Shape) Getlocalbounds() FloatRect { 
    return C.sfShape_getLocalBounds(self.Cref);
}
            
// Get the global bounding rectangle of a shape
// The returned rectangle is in global coordinates, which means
// that it takes in account the transformations (translation,
// rotation, scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// sprite in the global 2D world's coordinate system.
// \param shape Shape object
// \return Global bounding rectangle of the entity
// sfFloatRect sfShape_getGlobalBounds(const sfShape* shape);
func (self Shape) Getglobalbounds() FloatRect { 
    return C.sfShape_getGlobalBounds(self.Cref);
}
            
// Recompute the internal geometry of a shape
// This function must be called by specialized shape objects
// everytime their points change (ie. the result of either
// the getPointCount or getPoint callbacks is different).
// void sfShape_update(sfShape* shape);
func (self Shape) Update() void { 
    return C.sfShape_update(self.Cref);
}
            
