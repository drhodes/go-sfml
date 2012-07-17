package sys

// #cgo LDFLAGS:-lcsfml-system
// #include <SFML/System.h>
// #include <SFML/System/Vector2.h>
//
// sfVector2u _NewVector2u(unsigned int x, unsigned int y) {
//   sfVector2u v;
//   v.x = x; v.y=y;
//   return v;
// }
// sfVector2i _NewVector2i(int x, int y) {
//   sfVector2i v;
//   v.x = x; v.y=y;
//   return v;
// }
// sfVector2f _NewVector2f(float x, float y) {
//   sfVector2f v;
//   v.x = x; v.y=y;
//   return v;
// }
import "C"

type (
	Vector2u struct {
		Cref C.sfVector2u
	}
	Vector2i struct{
		Cref C.sfVector2i
	}
	Vector2f struct{
		Cref C.sfVector2f
	}	
)

func NewVector2u(x, y uint) Vector2u {
	return Vector2u{C._NewVector2u(C.uint(x), C.uint(y))}
}

func NewVector2i(x, y int) Vector2i {
	return Vector2i{C._NewVector2i(C.int(x), C.int(y))}
}

func NewVector2f(x, y float32) Vector2f {
	return Vector2f{C._NewVector2f(C.float(x), C.float(y))}
}

