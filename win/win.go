package gfx

// #include <SFML/Graphics/Color.h>
import "C"

import(
	"runtime"
	"unsafe"
)

func NewColor(val C.sfColor) Color {
	return Color{ val }
}

type Color struct { 
	Cref C.sfColor 
}
//-------------------------------------
func NewFont(val *C.sfFont) Font {
    tmp := Font{ val }
    runtime.SetFinalizer(&tmp, (*Font).Destroy)
    return tmp
}

type Font struct {
	Cref *C.sfFont
}

// _Color_
// -------------------------------------------------------------------------------
