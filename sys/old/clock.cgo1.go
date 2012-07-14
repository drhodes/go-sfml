// Created by cgo - DO NOT EDIT
//line clock.go:1
package sys


type Clock struct {
	Cref *[0]byte
}

func ClockCreate() Clock {
	return Clock{_Cfunc_sfClock_Create()}
}

func (self *Clock) Destroy() {
	_Cfunc_sfClock_Destroy(self.Cref)
	self.Cref = nil
}

func (self *Clock) GetTime() float32 {
	return float32(_Cfunc_sfClock_GetTime(self.Cref))
}

func (self *Clock) Reset() {
	_Cfunc_sfClock_Reset(self.Cref)
}


func Sleep(duration float32) {
	_Cfunc_sfSleep(_Ctype_float(duration))
}
