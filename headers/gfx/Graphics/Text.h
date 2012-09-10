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

#ifndef SFML_TEXT_H
#define SFML_TEXT_H

////////////////////////////////////////////////////////////
// Headers
////////////////////////////////////////////////////////////
#include <SFML/Graphics/Export.h>
#include <SFML/Graphics/Color.h>
#include <SFML/Graphics/Rect.h>
#include <SFML/Graphics/Types.h>
#include <SFML/System/Vector2.h>
#include <stddef.h>


////////////////////////////////////////////////////////////
/// sfText styles
////////////////////////////////////////////////////////////
typedef enum
{
    sfTextRegular    = 0,      ///< Regular characters, no style
    sfTextBold       = 1 << 0, ///< Characters are bold
    sfTextItalic     = 1 << 1, ///< Characters are in italic
    sfTextUnderlined = 1 << 2  ///< Characters are underlined
} sfTextStyle;


////////////////////////////////////////////////////////////
/// \brief Create a new text
///
/// \return A new sfText object, or NULL if it failed
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfText* sfText_create(void);

////////////////////////////////////////////////////////////
/// \brief Copy an existing text
///
/// \param text Text to copy
///
/// \return Copied object
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfText* sfText_copy(sfText* text);

////////////////////////////////////////////////////////////
/// \brief Destroy an existing text
///
/// \param text Text to delete
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_destroy(sfText* text);

////////////////////////////////////////////////////////////
/// \brief Set the position of a text
///
/// This function completely overwrites the previous position.
/// See sfText_move to apply an offset based on the previous position instead.
/// The default position of a text Text object is (0, 0).
///
/// \param text     Text object
/// \param position New position
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setPosition(sfText* text, sfVector2f position);

////////////////////////////////////////////////////////////
/// \brief Set the orientation of a text
///
/// This function completely overwrites the previous rotation.
/// See sfText_rotate to add an angle based on the previous rotation instead.
/// The default rotation of a text Text object is 0.
///
/// \param text  Text object
/// \param angle New rotation, in degrees
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setRotation(sfText* text, float angle);

////////////////////////////////////////////////////////////
/// \brief Set the scale factors of a text
///
/// This function completely overwrites the previous scale.
/// See sfText_scale to add a factor based on the previous scale instead.
/// The default scale of a text Text object is (1, 1).
///
/// \param text  Text object
/// \param scale New scale factors
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setScale(sfText* text, sfVector2f scale);

////////////////////////////////////////////////////////////
/// \brief Set the local origin of a text
///
/// The origin of an object defines the center point for
/// all transformations (position, scale, rotation).
/// The coordinates of this point must be relative to the
/// top-left corner of the object, and ignore all
/// transformations (position, scale, rotation).
/// The default origin of a text object is (0, 0).
///
/// \param text   Text object
/// \param origin New origin
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setOrigin(sfText* text, sfVector2f origin);

////////////////////////////////////////////////////////////
/// \brief Get the position of a text
///
/// \param text Text object
///
/// \return Current position
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfVector2f sfText_getPosition(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the orientation of a text
///
/// The rotation is always in the range [0, 360].
///
/// \param text Text object
///
/// \return Current rotation, in degrees
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API float sfText_getRotation(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the current scale of a text
///
/// \param text Text object
///
/// \return Current scale factors
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfVector2f sfText_getScale(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the local origin of a text
///
/// \param text Text object
///
/// \return Current origin
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfVector2f sfText_getOrigin(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Move a text by a given offset
///
/// This function adds to the current position of the object,
/// unlike sfText_setPosition which overwrites it.
///
/// \param text   Text object
/// \param offset Offset
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_move(sfText* text, sfVector2f offset);

////////////////////////////////////////////////////////////
/// \brief Rotate a text
///
/// This function adds to the current rotation of the object,
/// unlike sfText_setRotation which overwrites it.
///
/// \param text  Text object
/// \param angle Angle of rotation, in degrees
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_rotate(sfText* text, float angle);

////////////////////////////////////////////////////////////
/// \brief Scale a text
///
/// This function multiplies the current scale of the object,
/// unlike sfText_setScale which overwrites it.
///
/// \param text    Text object
/// \param factors Scale factors
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_scale(sfText* text, sfVector2f factors);

////////////////////////////////////////////////////////////
/// \brief Get the combined transform of a text
///
/// \param text Text object
///
/// \return Transform combining the position/rotation/scale/origin of the object
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API const sfTransform* sfText_getTransform(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the inverse of the combined transform of a text
///
/// \param text Text object
///
/// \return Inverse of the combined transformations applied to the object
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API const sfTransform* sfText_getInverseTransform(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Set the string of a text (from an ANSI string)
///
/// A text's string is empty by default.
///
/// \param text   Text object
/// \param string New string
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setString(sfText* text, const char* string);

////////////////////////////////////////////////////////////
/// \brief Set the string of a text (from a unicode string)
///
/// \param text   Text object
/// \param string New string
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setUnicodeString(sfText* text, const sfUint32* string);

////////////////////////////////////////////////////////////
/// \brief Set the font of a text
///
/// The \a font argument refers to a texture that must
/// exist as long as the text uses it. Indeed, the text
/// doesn't store its own copy of the font, but rather keeps
/// a pointer to the one that you passed to this function.
/// If the font is destroyed and the text tries to
/// use it, the behaviour is undefined.
/// Texts have a valid font by default, which the built-in
/// sfFont_getDefaultFont().
///
/// \param text Text object
/// \param font New font
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setFont(sfText* text, const sfFont* font);

////////////////////////////////////////////////////////////
/// \brief Set the character size of a text
///
/// The default size is 30.
///
/// \param text Text object
/// \param size New character size, in pixels
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setCharacterSize(sfText* text, unsigned int size);

////////////////////////////////////////////////////////////
/// \brief Set the style of a text
///
/// You can pass a combination of one or more styles, for
/// example sfTextBold | sfTextItalic.
/// The default style is sfTextRegular.
///
/// \param text  Text object
/// \param style New style
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setStyle(sfText* text, sfUint32 style);

////////////////////////////////////////////////////////////
/// \brief Set the global color of a text
///
/// By default, the text's color is opaque white.
///
/// \param text  Text object
/// \param color New color of the text
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfText_setColor(sfText* text, sfColor color);

////////////////////////////////////////////////////////////
/// \brief Get the text of a text (returns an ANSI string)
///
/// \param text Text object
///
/// \return String an a locale-dependant ANSI string
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API const char* sfText_getString(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the string of a text (returns a unicode string)
///
/// \param text Text object
///
/// \return String as UTF-32
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API const sfUint32* sfText_getUnicodeString(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the font used by a text
///
/// \param text Text object
///
/// \return Pointer to the font
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API const sfFont* sfText_getFont(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the size of the characters of a text
///
/// \param text Text object
///
/// \return Size of the characters
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API unsigned int sfText_getCharacterSize(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the style of a text
///
/// \param text Text object
///
/// \return Current string style (see sfTextStyle enum)
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfUint32 sfText_getStyle(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the global color of a text
///
/// \param text Text object
///
/// \return Global color of the text
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfColor sfText_getColor(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Return the position of the \a index-th character in a text
///
/// This function computes the visual position of a character
/// from its index in the string. The returned position is
/// in global coordinates (translation, rotation, scale and
/// origin are applied).
/// If \a index is out of range, the position of the end of
/// the string is returned.
///
/// \param text  Text object
/// \param index Index of the character
///
/// \return Position of the character
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfVector2f sfText_findCharacterPos(const sfText* text, size_t index);

////////////////////////////////////////////////////////////
/// \brief Get the local bounding rectangle of a text
///
/// The returned rectangle is in local coordinates, which means
/// that it ignores the transformations (translation, rotation,
/// scale, ...) that are applied to the entity.
/// In other words, this function returns the bounds of the
/// entity in the entity's coordinate system.
///
/// \param text Text object
///
/// \return Local bounding rectangle of the entity
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfFloatRect sfText_getLocalBounds(const sfText* text);

////////////////////////////////////////////////////////////
/// \brief Get the global bounding rectangle of a text
///
/// The returned rectangle is in global coordinates, which means
/// that it takes in account the transformations (translation,
/// rotation, scale, ...) that are applied to the entity.
/// In other words, this function returns the bounds of the
/// text in the global 2D world's coordinate system.
///
/// \param text Text object
///
/// \return Global bounding rectangle of the entity
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfFloatRect sfText_getGlobalBounds(const sfText* text);


#endif // SFML_TEXT_H
