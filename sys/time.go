package sys

// #cgo LDFLAGS:-lcsfml-system
// #include <SFML/System.h>
import "C"

type Time struct{
	Cref C.sfTime
}


func (self Time) Microseconds(amount int64) Time {
	return Time{C.sfMicroseconds(C.sfInt64(amount))}
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

func Seconds(amount float32) Time {
	return Time{C.sfSeconds(C.float(amount))}
}

func Milliseconds(amount int32) Time {
	return Time{C.sfMilliseconds(C.sfInt32(amount))}
}

//sfTime sfTime_Zero;

