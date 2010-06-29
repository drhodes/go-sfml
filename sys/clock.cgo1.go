// Created by cgo - DO NOT EDIT
//line clock.go:1
package sys

// #include <SFML/System.h>


type Clock struct {
	Cref *[0]byte
}

func ClockCreate() Clock {
	return Clock{_C_sfClock_Create()}
}

func (self *Clock) Destroy() {
	_C_sfClock_Destroy(self.Cref)
	self.Cref = nil
}

func (self *Clock) GetTime() float {
	return float(_C_sfClock_GetTime(self.Cref))
}

func (self *Clock) Reset() {
	_C_sfClock_Reset(self.Cref)
}


//------------------------------------------------------
func Sleep(duration float) {
	_C_sfSleep(_C_float(duration))
}
