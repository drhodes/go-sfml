package sfml

import (
	"testing"
)

func TestClockConstructors(t *testing.T) {
	clock := NewClock()
	clock.Restart()
}

func TestAsMilliseconds(t *testing.T) {
	sec := Seconds(123)
	if sec.AsMilliseconds() != 123000 {
		t.Fail()
	}
}

func TestAsMicroseconds(t *testing.T) {
	sec := Seconds(123)
	if sec.AsMicroseconds() != 123000000 {
		t.Fail()
	}
}

func TestAsSeconds(t *testing.T) {
	sec := Seconds(123)
	if sec.AsSeconds() != 123 {
		t.Fail()
	}
}
