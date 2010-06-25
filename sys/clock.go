package sys

// #include <SFML/System/Clock.h>
import "C"

type Clock struct{
	Cref *C.sfClock
}

func ClockCreate() Clock {
	return Clock{ C.sfClock_Create() }
}

func (self *Clock) Destroy() {
	C.sfClock_Destroy(self.Cref)
}

func (self *Clock) GetTime() float {
	return float(C.sfClock_GetTime(self.Cref))
}

func (self *Clock) Reset() {
	C.sfClock_Reset(self.Cref);
}


