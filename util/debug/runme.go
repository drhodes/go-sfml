package main
// This graphics illustrates how C++ classes can be used from Go using SWIG.

import gfx "sfml/graphics" 
import "fmt"

func Debug(x interface{}){
	fmt.Printf("%#v\n", x)
}

func main() {
	fmt.Println("Creating some objects:");	
	img := gfx.ImageCreateFromFile("cute_image.jpg");
	Debug(img);

	/*
	clr := gfx.NewColor(1,2,3,4);
	clr.Add(clr)

	fnt := gfx.NewFont("../test/arial.ttf", 10)
	Debug(fnt)
	 */

	//snd := aud.SoundCreate()
	//fmt.Printf("%#v\n", snd)
}
 