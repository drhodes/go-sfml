package main

import (
	"github.com/drhodes/go-sfml/gfx"
	"github.com/drhodes/go-sfml/win"
	"fmt"
	"log"
)

func Debug(x interface{}){	fmt.Printf("%#+v\n", x) }

func AwesomeKeyHandler(ke win.KeyEvent) {
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
	vm := win.NewVideoMode(512,512,24)

	w, err := win.NewWindow(
		vm,
		"PONG",
		win.StyleDefaultStyle,
		win.ContextSettings{},
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
		//t = c.GetElapsedTime()

		e := w.PollEvent()
		switch e.(type) {
		case win.KeyEvent:
			AwesomeKeyHandler(e.(win.KeyEvent))
		case win.MouseMoveEvent:
			me := e.(win.MouseMoveEvent)
			log.Printf("MouseMove <%d, %d>\n", me.X(), me.Y())
		case nil:
			log.Println("LOOOOOOOOOL")
		}
	}
	//log.Println(t.AsSeconds())
}



