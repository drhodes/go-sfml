package sfml

// #cgo LDFLAGS:-lcsfml-graphics -ljpeg
// #include <SFML/Graphics/Color.h>
// #include <SFML/Window/Context.h>
// #include <SFML/Window/Export.h>
// #include <SFML/Window/Types.h>
// #include <SFML/Graphics/Export.h>
import "C"


// Utility class for manpulating RGBA colors
type Color struct {
	Cref C.sfColor
}

// Construct a color from its 3 RGB components
//
// \param red   Red component   (0 .. 255)
// \param green Green component (0 .. 255)
// \param blue  Blue component  (0 .. 255)
//
// \return sfColor constructed from the components
// sfColor sfColor_fromRGB(sfUint8 red, sfUint8 green, sfUint8 blue);
func FromRGB(red uint8, green uint8, blue uint8) Color {
	return Color{C.sfColor_fromRGB(
		C.sfUint8(red),
		C.sfUint8(green),
		C.sfUint8(blue),
		)};
}

// Construct a color from its 4 RGBA components
//
// \param red   Red component   (0 .. 255)
// \param green Green component (0 .. 255)
// \param blue  Blue component  (0 .. 255)
// \param alpha Alpha component (0 .. 255)
//
// \return sfColor constructed from the components
//
// sfColor sfColor_fromRGBA(sfUint8 red, sfUint8 green, sfUint8 blue, sfUint8 alpha);
func FromRGBA(red uint8, green uint8, blue uint8, alpha uint8) Color {
	return Color{C.sfColor_fromRGBA(
		C.sfUint8(red),
		C.sfUint8(green),
		C.sfUint8(blue),	
		C.sfUint8(alpha),		
		)};
}

// Add two colors
//
// \param color2 Second color
//
// \return Component-wise saturated addition of the two colors
//
// sfColor sfColor_add(sfColor color1, sfColor color2);
func (self Color) Add(color2 Color) Color {
	return Color{C.sfColor_add(self.Cref, color2.Cref)}
}

// Modulate two colors
//
// \param color1 First color
// \param color2 Second color
//
// \return Component-wise multiplication of the two colors
//

// sfColor sfColor_modulate(sfColor color1, sfColor color2);

func (self Color) Modulate(color2 Color) Color {
	return Color{C.sfColor_modulate(self.Cref, color2.Cref)}
}
