package sfml

// #cgo LDFLAGS:-lcsfml-graphics
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/BlendMode.h>
// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/Graphics/Sprite.h>
// #include <SFML/System/Vector2.h>
import "C"

import "errors"

type Sprite struct {
	Cref *C.sfSprite
}

// Create a new sprite
// \return A new sfSprite object, or NULL if it failed
// sfSprite* sfSprite_create(void);
func NewSprite() (Sprite, error) {
	spr := C.sfSprite_create()
	if spr == nil {
		return Sprite{nil}, errors.New("Couldn't make a sprite")
	}
	return Sprite{spr}, nil
}

// Copy an existing sprite
// \param sprite Sprite to copy
// \return Copied object
// sfSprite* sfSprite_copy(sfSprite* sprite);
func (self Sprite) Copy() Sprite {
	return Sprite{C.sfSprite_copy(self.Cref)}
}

// Destroy an existing sprite
// \param sprite Sprite to delete
// void sfSprite_destroy(sfSprite* sprite);
func (self Sprite) Destroy() {
	C.sfSprite_destroy(self.Cref)
}

// Set the position of a sprite
// This function completely overwrites the previous position.
// See sfSprite_move to apply an offset based on the previous position instead.
// The default position of a sprite Sprite object is (0, 0).
// \param sprite   Sprite object
// \param position New position
// void sfSprite_setPosition(sfSprite* sprite, sfVector2f position);
func (self Sprite) SetPosition(x, y float32) {
	v := C.sfVector2f{C.float(x), C.float(y)}
	C.sfSprite_setPosition(self.Cref, v)
}

// Set the orientation of a sprite
// This function completely overwrites the previous rotation.
// See sfSprite_rotate to add an angle based on the previous rotation instead.
// The default rotation of a sprite Sprite object is 0.
// \param sprite Sprite object
// \param angle  New rotation, in degrees
// void sfSprite_setRotation(sfSprite* sprite, float angle);
func (self Sprite) SetRotation(angle float32) {
	C.sfSprite_setRotation(self.Cref, C.float(angle))
}

// Set the scale factors of a sprite
// This function completely overwrites the previous scale.
// See sfSprite_scale to add a factor based on the previous scale instead.
// The default scale of a sprite Sprite object is (1, 1).
// \param sprite Sprite object
// \param scale  New scale factors
// void sfSprite_setScale(sfSprite* sprite, sfVector2f scale);
func (self Sprite) SetScale(x, y float32) {
	v := C.sfVector2f{C.float(x), C.float(y)}
	C.sfSprite_setScale(self.Cref, v)
}

// Set the local origin of a sprite
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a sprite Sprite object is (0, 0).
// \param sprite Sprite object
// \param origin New origin
// void sfSprite_setOrigin(sfSprite* sprite, sfVector2f origin);
func (self Sprite) SetOrigin(x, y float32) {
	v := C.sfVector2f{C.float(x), C.float(y)}
	C.sfSprite_setOrigin(self.Cref, v)
}

// Get the position of a sprite
// \param sprite Sprite object
// \return Current position
// sfVector2f sfSprite_getPosition(const sfSprite* sprite);
func (self Sprite) Position() (float32, float32) {
	v := C.sfSprite_getPosition(self.Cref)
	return float32(v.x), float32(v.y)
}

// Get the orientation of a sprite
// The rotation is always in the range [0, 360].
// \param sprite Sprite object
// \return Current rotation, in degrees
// float sfSprite_getRotation(const sfSprite* sprite);
func (self Sprite) Rotation() float32 {
	return float32(C.sfSprite_getRotation(self.Cref))
}

// Get the current scale of a sprite
// \param sprite Sprite object
// \return Current scale factors
// sfVector2f sfSprite_getScale(const sfSprite* sprite);
func (self Sprite) GetScale() (float32, float32) {
	v := C.sfSprite_getScale(self.Cref)
	return float32(v.x), float32(v.y)
}

// Get the local origin of a sprite
// \param sprite Sprite object
// \return Current origin
// sfVector2f sfSprite_getOrigin(const sfSprite* sprite);
func (self Sprite) Origin() (float32, float32) {
	v := C.sfSprite_getOrigin(self.Cref)
	return float32(v.x), float32(v.y)
}

// Move a sprite by a given offset
// This function adds to the current position of the object,
// unlike sfSprite_setPosition which overwrites it.
// \param sprite Sprite object
// \param offset Offset
// void sfSprite_move(sfSprite* sprite, sfVector2f offset);
func (self Sprite) Move(x, y float32) {
	offset := C.sfVector2f{C.float(x), C.float(y)}
	C.sfSprite_move(self.Cref, offset)
}

// Rotate a sprite
// This function adds to the current rotation of the object,
// unlike sfSprite_setRotation which overwrites it.
// \param sprite Sprite object
// \param angle  Angle of rotation, in degrees
// void sfSprite_rotate(sfSprite* sprite, float angle);
func (self Sprite) Rotate(angle float32) {
	C.sfSprite_rotate(self.Cref, C.float(angle))
}

// Scale a sprite
// This function multiplies the current scale of the object,
// unlike sfSprite_setScale which overwrites it.
// \param sprite  Sprite object
// \param factors Scale factors
// void sfSprite_scale(sfSprite* sprite, sfVector2f factors);
func (self Sprite) Scale(x, y float32) {
	v := C.sfVector2f{C.float(x), C.float(y)}
	C.sfSprite_scale(self.Cref, v)
}

// Get the combined transform of a sprite
// \param sprite Sprite object
// \return Transform combining the position/rotation/scale/origin of the object
// const sfTransform* sfSprite_getTransform(const sfSprite* sprite);
func (self Sprite) Transform() Transform {
	return Transform{C.sfSprite_getTransform(self.Cref)}
}

// Get the inverse of the combined transform of a sprite
// \param sprite Sprite object
// \return Inverse of the combined transformations applied to the object
// const sfTransform* sfSprite_getInverseTransform(const sfSprite* sprite);
func (self Sprite) InverseTransform() Transform {
	return Transform{C.sfSprite_getInverseTransform(self.Cref)}
}

// Change the source texture of a sprite
// The \a texture argument refers to a texture that must
// exist as long as the sprite uses it. Indeed, the sprite
// doesn't store its own copy of the texture, but rather keeps
// a pointer to the one that you passed to this function.
// If the source texture is destroyed and the sprite tries to
// use it, the behaviour is undefined.
// If \a resetRect is true, the TextureRect property of
// the sprite is automatically adjusted to the size of the new
// texture. If it is false, the texture rect is left unchanged.
// \param sprite    Sprite object
// \param texture   New texture
// \param resetRect Should the texture rect be reset to the size of the new texture?
// void sfSprite_setTexture(sfSprite* sprite, const sfTexture* texture, sfBool resetRect);
func (self Sprite) SetTexture(tex Texture, resetRect bool) {
	C.sfSprite_setTexture(self.Cref, tex.Cref, Bool(resetRect))
}

// Set the sub-rectangle of the texture that a sprite will display
// The texture rect is useful when you don't want to display
// the whole texture, but rather a part of it.
// By default, the texture rect covers the entire texture.
// \param sprite    Sprite object
// \param rectangle Rectangle defining the region of the texture to display
// void sfSprite_setTextureRect(sfSprite* sprite, sfIntRect rectangle);
func (self Sprite) SetTextureRect(rect IntRect) {
	C.sfSprite_setTextureRect(self.Cref, *rect.Cref)
}

// Set the global color of a sprite
// This color is modulated (multiplied) with the sprite's
// texture. It can be used to colorize the sprite, or change
// its global opacity.
// By default, the sprite's color is opaque white.
// \param sprite Sprite object
// \param color  New color of the sprite
// void sfSprite_setColor(sfSprite* sprite, sfColor color);
func (self Sprite) SetColor(color Color) {
	C.sfSprite_setColor(self.Cref, color.Cref)
}

// Get the source texture of a sprite
// If the sprite has no source texture, a NULL pointer is returned.
// The returned pointer is const, which means that you can't
// modify the texture when you retrieve it with this function.
// \param sprite Sprite object
// \return Pointer to the sprite's texture
// const sfTexture* sfSprite_getTexture(const sfSprite* sprite);
func (self *Texture) Texture() Texture {
	return Texture{C.sfSprite_getTexture(self.Cref)}
}

// Get the sub-rectangle of the texture displayed by a sprite
// \param sprite Sprite object
// \return Texture rectangle of the sprite
// sfIntRect sfSprite_getTextureRect(const sfSprite* sprite);
func (self Sprite) TextureRect() IntRect {
	ref := C.sfSprite_getTextureRect(self.Cref)
	return IntRect{&ref}
}

// Get the global color of a sprite
// \param sprite Sprite object
// \return Global color of the sprite
// sfColor sfSprite_getColor(const sfSprite* sprite);
func (self Sprite) Color() Color {
	return Color{C.sfSprite_getColor(self.Cref)}
}

// Get the local bounding rectangle of a sprite
// The returned rectangle is in local coordinates, which means
// that it ignores the transformations (translation, rotation,
// scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// entity in the entity's coordinate system.
// \param sprite Sprite object
// \return Local bounding rectangle of the entity
// sfFloatRect sfSprite_getLocalBounds(const sfSprite* sprite);
func (self Sprite) LocalBounds() FloatRect {
	ref := C.sfSprite_getLocalBounds(self.Cref)
	return FloatRect{&ref}
}

// Get the global bounding rectangle of a sprite
// The returned rectangle is in global coordinates, which means
// that it takes in account the transformations (translation,
// rotation, scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// sprite in the global 2D world's coordinate system.
// \param sprite Sprite object
// \return Global bounding rectangle of the entity
// sfFloatRect sfSprite_getGlobalBounds(const sfSprite* sprite);
func (self Sprite) GlobalBounds() FloatRect {
	ref := C.sfSprite_getGlobalBounds(self.Cref)
	return FloatRect{&ref}
}

func Bool(b bool) C.sfBool {
	if b {
		return C.sfBool(1)
	}
	return C.sfBool(0)
}
