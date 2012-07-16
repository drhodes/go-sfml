package gfx


import (
	"testing"
	//"log"
)

//const verb = true

func TestConstructors(t *testing.T) {		
	//clock := sys.NewClock()
}

func TestColor(t *testing.T) {		
	c1 := FromRGBA(12,12,12,12)
	c2 := FromRGB(12,12,12)
	c1.Modulate(c2)
	c1.Add(c2)
}


func BenchmarkGetTime(b *testing.B) {
}

func BenchmarkClockDestroy(b *testing.B) {	
}

