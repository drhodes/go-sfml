package sfml

import (
	"log"
	"testing"
)

func TestConstructors(t *testing.T) {
	clock := NewClock()
	clock.Restart()
}

func BenchmarkGetTime(b *testing.B) {
}

func BenchmarkClockDestroy(b *testing.B) {	
}

func TestMicroseconds(t *testing.T) {
}

func TestAsSeconds(t *testing.T) {
}

func TestAsMilliseconds(t *testing.T) {
	sec := Seconds(123)
	if sec.AsMilliseconds() != 123000 {
		log.Println(sec)
		t.Fail()
	}
}

func TestAsMicroseconds(t *testing.T) {
	sec := Seconds(123)
	if sec.AsMicroseconds() != 123000000 {
		t.Fail()
	}
}

func TestSeconds(t *testing.T) {
	sec := Seconds(123)
	if sec.AsSeconds() != 123 {
		t.Fail()
	}
}

func TestMilliseconds(t *testing.T) {
}

/*
func BenchmarkCollidePoly(b *testing.B) {
	pg1 := new(Polygon)
	pg2 := new(Polygon)
	for i := 0; i < b.N; i++ {
		pg1 = pg1.RandomEl()
		pg2 = pg2.RandomEl()
		//Debug(pg1)
		//Debug(pg2)
		pg1.Collide(pg2)
		//Debug("-------------------------------------------")
	}
}
*/