package sys

// #include <SFML/System.h>
import "C"

type Clock struct{
	Cref *C.sfClock
}

func ClockCreate() Clock {
	return Clock{ C.sfClock_Create() }
}

func (self *Clock) Destroy() {
	C.sfClock_Destroy(self.Cref)
	self.Cref = nil
}

func (self *Clock) GetTime() float32 {
	return float32(C.sfClock_GetTime(self.Cref))
}

func (self *Clock) Reset() {
	C.sfClock_Reset(self.Cref);
}


//------------------------------------------------------
func Sleep(duration float32){
	C.sfSleep(C.float(duration))
}
