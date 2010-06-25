package sys


import "testing"
import "fmt"
import _ "http/pprof"
import "sfml/sys"


func Debug(x interface{}){	fmt.Printf("%#v\n", x); }
func DebugT(x interface{}){ fmt.Printf("%#T\n", x); }
func DebugN(x interface{}){ fmt.Printf("%#v", x); }

func loop(n int, f func() ){
	for i:=0; i<n; i++ {
		DebugN(i)
		DebugN(": ")
		f()
	}	
}

func TestConstructors(t *testing.T) {		
	clock := sys.ClockCreate()
	Debug(clock)
	
	for clock.GetTime() < .1 {
		Debug(clock.GetTime())
	}

	clock.Reset()
	Debug(clock.GetTime())

	clock.Destroy()
	Debug(clock)

	/*
    for system.ClockGetTime(clock) < 5 {
		fmt.Printf("%v\n", system.ClockGetTime(clock))
		system.Sleep(0.5);
	}	
	
	fnt := GetDefaultFont()
	loop( 2000000, func(){
		fnt = FontCreateFromFile("../test/arial.ttf", 34)			
		//fnt.Destroy()
		Debug(fnt)
	})
	 */
}


/*
func BenchmarkPoly(b *testing.B) {
	pg := new(Polygon)
	for i := 0; i < b.N; i++ {
		pg.RandomEl();
	}
}

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