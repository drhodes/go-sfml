// Created by cgo - DO NOT EDIT

//line clock.go:1
package sys
//line clock.go:8

//line clock.go:7
type Clock struct {
	Cref *[0]byte
}
//line clock.go:13

//line clock.go:12
func Sleep(duration float32) {
//line clock.go:15

//line clock.go:14
}
//line clock.go:19

//line clock.go:18
func NewClock() Clock {
	return Clock{_Cfunc_sfClock_create()}
}
//line clock.go:23

//line clock.go:22
func (self Clock) Copy() Clock {
	return Clock{_Cfunc_sfClock_copy(self.Cref)}
}
func (self Clock) Destroy() {
//line clock.go:25
	_Cfunc_sfClock_destroy(sfClock * clock)
//line clock.go:27
}
