// Created by cgo - DO NOT EDIT
//line color.go:1
package gfx

// #include <SFML/Graphics.h>
// #include <SFML/Config.h>

//import "unsafe"

type Color struct {
	Cref _C_sfColor
}

func ColorFromRGB(r, g, b int) Color {
	return Color{_C_sfColor_FromRGB(_C_sfUint8(r), _C_sfUint8(g), _C_sfUint8(b))}
}


func ColorFromRGBA(r, g, b, a uint8) Color {
	return Color{_C_sfColor_FromRGBA(_C_sfUint8(r), _C_sfUint8(g), _C_sfUint8(b), _C_sfUint8(a))}
}


/// Add two colors
/// \param other : Color
/// \return Component-wise saturated addition of the two colors
func (self *Color) Add(other Color) Color {
	return Color{_C_sfColor_Add(self.Cref, other.Cref)}
}

/// Modulate two colors
/// \param other : Color
/// \return Component-wise multiplication of the two colors
func (self *Color) Modulate(other Color) Color {
	return Color{_C_sfColor_Modulate(self.Cref, other.Cref)}
}
