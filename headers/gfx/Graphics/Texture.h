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

#ifndef SFML_TEXTURE_H
#define SFML_TEXTURE_H

////////////////////////////////////////////////////////////
// Headers
////////////////////////////////////////////////////////////
#include <SFML/Graphics/Export.h>
#include <SFML/Graphics/Rect.h>
#include <SFML/Graphics/Types.h>
#include <SFML/Window/Types.h>
#include <SFML/System/InputStream.h>
#include <SFML/System/Vector2.h>
#include <stddef.h>


////////////////////////////////////////////////////////////
/// \brief Create a new texture
///
/// \param width  Texture width
/// \param height Texture height
///
/// \return A new sfTexture object, or NULL if it failed
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfTexture* sfTexture_create(unsigned int width, unsigned int height);

////////////////////////////////////////////////////////////
/// \brief Create a new texture from a file
///
/// \param filename Path of the image file to load
/// \param area     Area of the source image to load (NULL to load the entire image)
///
/// \return A new sfTexture object, or NULL if it failed
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfTexture* sfTexture_createFromFile(const char* filename, const sfIntRect* area);

////////////////////////////////////////////////////////////
/// \brief Create a new texture from a file in memory
///
/// \param data        Pointer to the file data in memory
/// \param sizeInBytes Size of the data to load, in bytes
/// \param area        Area of the source image to load (NULL to load the entire image)
///
/// \return A new sfTexture object, or NULL if it failed
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfTexture* sfTexture_createFromMemory(const void* data, size_t sizeInBytes, const sfIntRect* area);

////////////////////////////////////////////////////////////
/// \brief Create a new texture from a custom stream
///
/// \param stream Source stream to read from
/// \param area   Area of the source image to load (NULL to load the entire image)
///
/// \return A new sfTexture object, or NULL if it failed
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfTexture* sfTexture_createFromStream(sfInputStream* stream, const sfIntRect* area);

////////////////////////////////////////////////////////////
/// \brief Create a new texture from an image
///
/// \param image Image to upload to the texture
/// \param area  Area of the source image to load (NULL to load the entire image)
///
/// \return A new sfTexture object, or NULL if it failed
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfTexture* sfTexture_createFromImage(const sfImage* image, const sfIntRect* area);

////////////////////////////////////////////////////////////
/// \brief Copy an existing texture
///
/// \param texture Texture to copy
///
/// \return Copied object
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfTexture* sfTexture_copy(sfTexture* texture);

////////////////////////////////////////////////////////////
/// \brief Destroy an existing texture
///
/// \param texture Texture to delete
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfTexture_destroy(sfTexture* texture);

////////////////////////////////////////////////////////////
/// \brief Return the size of the texture
///
/// \param texture Texture to read
///
/// \return Size in pixels
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfVector2u sfTexture_getSize(const sfTexture* texture);

////////////////////////////////////////////////////////////
/// \brief Copy a texture's pixels to an image
///
/// \param texture Texture to copy
///
/// \return Image containing the texture's pixels
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfImage* sfTexture_copyToImage(const sfTexture* texture);

////////////////////////////////////////////////////////////
/// \brief Update a texture from an array of pixels
///
/// \param texture Texture to update
/// \param pixels  Array of pixels to copy to the texture
/// \param width   Width of the pixel region contained in \a pixels
/// \param height  Height of the pixel region contained in \a pixels
/// \param x       X offset in the texture where to copy the source pixels
/// \param y       Y offset in the texture where to copy the source pixels
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfTexture_updateFromPixels(sfTexture* texture, const sfUint8* pixels, unsigned int width, unsigned int height, unsigned int x, unsigned int y);

////////////////////////////////////////////////////////////
/// \brief Update a texture from an image
///
/// \param texture Texture to update
/// \param image   Image to copy to the texture
/// \param x       X offset in the texture where to copy the source pixels
/// \param y       Y offset in the texture where to copy the source pixels
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfTexture_updateFromImage(sfTexture* texture, const sfImage* image, unsigned int x, unsigned int y);

////////////////////////////////////////////////////////////
/// \brief Update a texture from the contents of a window
///
/// \param texture Texture to update
/// \param window  Window to copy to the texture
/// \param x       X offset in the texture where to copy the source pixels
/// \param y       Y offset in the texture where to copy the source pixels
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfTexture_updateFromWindow(sfTexture* texture, const sfWindow* window, unsigned int x, unsigned int y);

////////////////////////////////////////////////////////////
/// \brief Update a texture from the contents of a render-window
///
/// \param texture      Texture to update
/// \param renderWindow Render-window to copy to the texture
/// \param x            X offset in the texture where to copy the source pixels
/// \param y            Y offset in the texture where to copy the source pixels
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfTexture_updateFromRenderWindow(sfTexture* texture, const sfRenderWindow* renderWindow, unsigned int x, unsigned int y);

////////////////////////////////////////////////////////////
/// \brief Activate a texture for rendering
///
/// \param texture Texture to bind
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfTexture_bind(const sfTexture* texture);

////////////////////////////////////////////////////////////
/// \brief Enable or disable the smooth filter on a texture
///
/// \param texture The texture object
/// \param smooth  sfTrue to enable smoothing, sfFalse to disable it
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfTexture_setSmooth(sfTexture* texture, sfBool smooth);

////////////////////////////////////////////////////////////
/// \brief Tell whether the smooth filter is enabled or not for a texture
///
/// \param texture The texture object
///
/// \return sfTrue if smoothing is enabled, sfFalse if it is disabled
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfBool sfTexture_isSmooth(const sfTexture* texture);

////////////////////////////////////////////////////////////
/// \brief Enable or disable repeating for a texture
///
/// Repeating is involved when using texture coordinates
/// outside the texture rectangle [0, 0, width, height].
/// In this case, if repeat mode is enabled, the whole texture
/// will be repeated as many times as needed to reach the
/// coordinate (for example, if the X texture coordinate is
/// 3 * width, the texture will be repeated 3 times).
/// If repeat mode is disabled, the "extra space" will instead
/// be filled with border pixels.
/// Warning: on very old graphics cards, white pixels may appear
/// when the texture is repeated. With such cards, repeat mode
/// can be used reliably only if the texture has power-of-two
/// dimensions (such as 256x128).
/// Repeating is disabled by default.
///
/// \param texture  The texture object
/// \param repeated True to repeat the texture, false to disable repeating
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API void sfTexture_setRepeated(sfTexture* texture, sfBool repeated);

////////////////////////////////////////////////////////////
/// \brief Tell whether a texture is repeated or not
///
/// \param texture The texture object
///
/// \return sfTrue if repeat mode is enabled, sfFalse if it is disabled
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API sfBool sfTexture_isRepeated(const sfTexture* texture);

////////////////////////////////////////////////////////////
/// \brief Get the maximum texture size allowed
///
/// \return Maximum size allowed for textures, in pixels
///
////////////////////////////////////////////////////////////
CSFML_GRAPHICS_API unsigned int sfTexture_getMaximumSize();


#endif // SFML_TEXTURE_H
