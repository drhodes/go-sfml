package sys

import "testing"

func TestConstructors(t *testing.T) {		
	clock := ClockCreate()
	Debug(clock)
	
	for clock.GetTime() < .5 {
		Debug(clock.GetTime())
		Sleep(clock.GetTime())
	}

	Debug(`Resetting:`)
	clock.Reset()

	Debug(`Sleeping .5 seconds:`)
	Sleep(.5)

	DebugN(`Getting time: `)
	Debug(clock.GetTime())

	Debug(clock)
	Debug(`Destroying Clock`)

	clock.Destroy()
	Debug(clock)
}

func BenchmarkGetTime(b *testing.B) {
	clock := ClockCreate()
	for i := 0; i < b.N; i++ {
		clock.GetTime()
	}
}


func BenchmarkClockDestroy(b *testing.B) {
	clock := ClockCreate()	
	for i := 0; i < b.N; i++ {
		clock.Destroy()
	}
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