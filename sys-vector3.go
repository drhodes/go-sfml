package sfml

// #cgo LDFLAGS:-lcsfml-system
// #include <SFML/System.h>
// #include <SFML/System/Vector3.h>
//
// sfVector3f _NewVector3f(float x, float y, float z) {
//   sfVector3f v;
//   v.x = x; v.y=y; v.z=z;
//   return v;
// }
import "C"

type Vector3f struct {
	Cref C.sfVector3f
}

func NewVector3f(x, y, z float32) Vector3f {
	return Vector3f{C._NewVector3f(C.float(x), C.float(y), C.float(z))}
}

func (self Vector3f) X() float32 {
	return float32(self.Cref.x)
}

func (self Vector3f) Y() float32 {
	return float32(self.Cref.y)
}

func (self Vector3f) Z() float32 {
	return float32(self.Cref.z)
}

// internal
func CNewVector3f(x, y, z C.float) Vector3f {
	return Vector3f{C._NewVector3f(x, y, z)}
}
