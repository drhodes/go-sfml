package main

import (	
	"sfml/gfx"
	"sfml/sys"
	. "sfml/win"
	"sfml/win"
	"fmt"
	"log"
)

func Debug(x interface{}){	fmt.Printf("%#+v\n", x) }

func AwesomeKeyHandler(ke KeyEvent) {
	switch ke.Type {
	case win.EvtKeyPressed:
		log.Println("Key Pressed: ", ke)
	case win.EvtKeyReleased:
		log.Println("Key Released: ", ke)
	}
}

// type Ball struct {
// 	curpos 
// 	spr 
// }


func main() {
	vm := NewVideoMode(512,512,24)

	w, err := NewWindow(
		vm,		
		"PONG",
		StyleDefaultStyle,
		ContextSettings{},
	)
	if err != nil {
		log.Fatal(err)
	}

	img, err := gfx.ImageFromFile("./gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(img.Getsize())
	w.SetFramerateLimit(60);

    for w.IsOpen() {
		t = c.GetElapsedTime()			
		
		e := w.PollEvent()
		switch e.(type) {
		case KeyEvent:
			AwesomeKeyHandler(e.(KeyEvent))
		case MouseMoveEvent:		
			me := e.(MouseMoveEvent)
			log.Printf("MouseMove <%d, %d>\n", me.X(), me.Y())
		case nil:
			log.Println("LOOOOOOOOOL")
		}
	}
	log.Println(t.AsSeconds())
}



