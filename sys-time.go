package sfml

// #cgo LDFLAGS:-lcsfml-system
// #include <SFML/System.h>
import "C"

import "fmt"

type Time struct {
	Cref C.sfTime
}

func (self Time) AsSeconds() float32 {
	return float32(C.sfTime_asSeconds(self.Cref))
}

func (self Time) AsMilliseconds() int32 {
	return int32(C.sfTime_asMilliseconds(self.Cref))
}

func (self Time) AsMicroseconds() int64 {
	return int64(C.sfTime_asMicroseconds(self.Cref))
}

func (self Time) String() string {
	return fmt.Sprintf("%d.%ds", self.AsSeconds(), self.AsMilliseconds())
}

func Seconds(amount float32) Time {
	return Time{C.sfSeconds(C.float(amount))}
}

func Milliseconds(amount int32) Time {
	return Time{C.sfMilliseconds(C.sfInt32(amount))}
}

func Microseconds(amount int64) Time {
	return Time{C.sfMicroseconds(C.sfInt64(amount))}
}

var Zero Time = Seconds(0)
