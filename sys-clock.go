package sfml

// #cgo LDFLAGS:-lcsfml-system
// #include <SFML/System.h>
import "C"

type Clock struct {
	Cref *C.sfClock
}

func NewClock() Clock {
	return Clock{C.sfClock_create()}
}

func (self Clock) Copy() Clock {
	return Clock{C.sfClock_copy(self.Cref)}
}

func (self Clock) Destroy() {
	C.sfClock_destroy(self.Cref)
}

func (self Clock) ElapsedTime() Time {
	return Time{C.sfClock_getElapsedTime(self.Cref)}
}

func (self Clock) Restart() Time {
	return Time{C.sfClock_restart(self.Cref)}
}
