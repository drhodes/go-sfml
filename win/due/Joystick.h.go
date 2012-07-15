////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
#include <SFML/Window/Export.h>
////////////////////////////////////////////////////////////
/// \brief Global joysticks capabilities
///
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
/// \brief Axes supported by SFML joysticks
///
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
/// \brief Check if a joystick is connected
///
/// \param joystick Index of the joystick to check
///
/// \return sfTrue if the joystick is connected, sfFalse otherwise
///
////////////////////////////////////////////////////////////
// sfBool sfJoystick_isConnected(unsigned int joystick);

func (self Joystick) Isconnected() Bool { 
    return C.sfJoystick_isConnected();
}
            
////////////////////////////////////////////////////////////
/// \brief Return the number of buttons supported by a joystick
///
/// If the joystick is not connected, this function returns 0.
///
/// \param joystick Index of the joystick
///
/// \return Number of buttons supported by the joystick
///
////////////////////////////////////////////////////////////
// unsigned int sfJoystick_getButtonCount(unsigned int joystick);

func (self int) int(Joystick_getButtonCount)  { 
    return C.sfint();
}
            
////////////////////////////////////////////////////////////
/// \brief Check if a joystick supports a given axis
///
/// If the joystick is not connected, this function returns false.
///
/// \param joystick Index of the joystick
/// \param axis     Axis to check
///
/// \return sfTrue if the joystick supports the axis, sfFalse otherwise
///
////////////////////////////////////////////////////////////
// sfBool sfJoystick_hasAxis(unsigned int joystick, sfJoystickAxis axis);

func (self Joystick) Hasaxis(axis JoystickAxis) Bool { 
    return C.sfJoystick_hasAxis();
}
            
////////////////////////////////////////////////////////////
/// \brief Check if a joystick button is pressed
///
/// If the joystick is not connected, this function returns false.
///
/// \param joystick Index of the joystick
/// \param button   Button to check
///
/// \return sfTrue if the button is pressed, sfFalse otherwise
///
////////////////////////////////////////////////////////////
// sfBool sfJoystick_isButtonPressed(unsigned int joystick, unsigned int button);

func (self Joystick) Isbuttonpressed(button int ) Bool { 
    return C.sfJoystick_isButtonPressed();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the current position of a joystick axis
///
/// If the joystick is not connected, this function returns 0.
///
/// \param joystick Index of the joystick
/// \param axis     Axis to check
///
/// \return Current position of the axis, in range [-100 .. 100]
///
////////////////////////////////////////////////////////////
// float sfJoystick_getAxisPosition(unsigned int joystick, sfJoystickAxis axis);

func (self Joystick) Getaxisposition(axis JoystickAxis) float { 
    return C.sfJoystick_getAxisPosition();
}
            
////////////////////////////////////////////////////////////
/// \brief Update the states of all joysticks
///
/// This function is used internally by SFML, so you normally
/// don't have to call it explicitely. However, you may need to
/// call it if you have no window yet (or no window at all):
/// in this case the joysticks states are not updated automatically.
///
////////////////////////////////////////////////////////////
// void sfJoystick_update(void);

func (self Joystick) Update() void { 
    return C.sfJoystick_update();
}
            
