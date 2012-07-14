package sys
// #cgo LDFLAGS: -lcsfml-system
// #include <SFML/Config.h>
// #include <SFML/System.h>
import "C"

type Clock struct{
	Cref *C.sfClock
}

//------------------------------------------------------
func Sleep(duration float32){
	//C.sfSleep(C.float(duration))
}



func NewClock() Clock {
	return Clock{ C.sfClock_create() }
}

func (self Clock) Copy() Clock {
	return Clock{C.sfClock_copy(self.Cref)}
}
func (self Clock) Destroy() {	
	C.sfClock_destroy(sfClock* clock);
}
