package gfx

// #include <SFML/Graphics.h>
// #include <SFML/Config.h>
import "C"
//import "unsafe"

type Color struct{
	Cref C.sfColor
}

func ColorFromRGB(r, g, b int) Color {
	return Color{ C.sfColor_FromRGB(C.sfUint8(r), C.sfUint8(g), C.sfUint8(b) ) }
}


func ColorFromRGBA(r, g, b, a uint8) Color {
	return Color{ C.sfColor_FromRGBA(C.sfUint8(r), C.sfUint8(g), C.sfUint8(b), C.sfUint8(a) ) }
}


/// Add two colors
/// \param other : Color
/// \return Component-wise saturated addition of the two colors
func (self *Color) Add(other Color) Color {
	return Color{ C.sfColor_Add(self.Cref, other.Cref) }
}
 
/// Modulate two colors
/// \param other : Color
/// \return Component-wise multiplication of the two colors
func (self *Color) Modulate(other Color) Color {
	return Color { C.sfColor_Modulate(self.Cref, other.Cref) }
}


