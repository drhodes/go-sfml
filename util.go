package sfml

// #cgo LDFLAGS:-lcsfml-graphics
// #include <SFML/Graphics/BlendMode.h>
// #include <SFML/Graphics/Color.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/Graphics/Sprite.h>
// #include <SFML/System/Vector2.h>
import "C"

func Bool(b bool) C.sfBool {
	if b {
		return C.sfBool(1)
	}
	return C.sfBool(0)
}

