////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
#include <SFML/Window/Export.h>
#include <SFML/Window/Types.h>
#include <SFML/System/Vector2.h>
////////////////////////////////////////////////////////////
/// \brief Mouse buttons
///
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
/// \brief Check if a mouse button is pressed
///
/// \param button Button to check
///
/// \return sfTrue if the button is pressed, sfFalse otherwise
///
////////////////////////////////////////////////////////////
// sfBool sfMouse_isButtonPressed(sfMouseButton button);

func (self Mouse) Isbuttonpressed() Bool { 
    return C.sfMouse_isButtonPressed();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the current position of the mouse
///
/// This function returns the current position of the mouse
/// cursor relative to the given window, or desktop if NULL is passed.
///
/// \param relativeTo Reference window
///
/// \return Position of the mouse cursor, relative to the given window
///
////////////////////////////////////////////////////////////
// sfVector2i sfMouse_getPosition(const sfWindow* relativeTo);

func (self Mouse) Getposition() Vector2i { 
    return C.sfMouse_getPosition();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the current position of the mouse
///
/// This function sets the current position of the mouse
/// cursor relative to the given window, or desktop if NULL is passed.
///
/// \param position   New position of the mouse
/// \param relativeTo Reference window
///
////////////////////////////////////////////////////////////
// void sfMouse_setPosition(sfVector2i position, const sfWindow* relativeTo);

func (self Mouse) Setposition(relativeTo *Window ) void { 
    return C.sfMouse_setPosition();
}
            
