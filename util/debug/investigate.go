package main

import . "sfml/gfx"

//import "testing"
import "fmt"
//import "runtime"
import "runtime/pprof"
//import . "sfml/graphics"
import "os"
//import "os"

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

func TestConstructors() {		
	c1 := new(Color)
	Debug(c1)
	
	f, _ := os.Open("./test.log", os.O_CREATE|os.O_WRONLY, 0777) 
	loop( 20000, func(){
		fnt := FontCreateFromFile("../test/arial.ttf", 34)			
		pprof.WriteHeapProfile(f)
		//fnt.Destroy()
		Debug(fnt)
		//runtime.GC()
	})
}


func main(){
	TestConstructors()
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