package sys

import "testing"

func TestConstructors(t *testing.T) {		
	clock := ClockCreate()	
}

func BenchmarkGetTime(b *testing.B) {
}


func BenchmarkClockDestroy(b *testing.B) {
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