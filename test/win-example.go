package main

import (
	"github.com/drhodes/go-sfml/gfx"
	"github.com/drhodes/go-sfml/win"
	"log"
)

func main() {
	vm := win.NewVideoMode(800, 600, 32)
	window := gfx.NewRenderWindowDefault(vm, "SFML window", win.StyleDefaultStyle)
	black := gfx.FromRGB(0, 0, 0)

	texture, err := gfx.TextureFromFile("gopher.png", gfx.IntRect{})
	if err != nil {
		log.Fatal(err)
	}

	sprite, err := gfx.NewSprite()
	if err != nil {
		log.Fatal(err)
	}
	sprite.SetTexture(texture, false)

	for window.IsOpen() {
		e := window.PollEvent()
		switch e.(type) {
		case win.KeyEvent:
			ev := e.(win.KeyEvent)
			if ev.Code() == win.KeyEscape {
				window.Close()
			}
		}
		window.Clear(black)
		window.DrawSprite(sprite)
		window.Display()
	}
}


/*
func Debug(x interface{}){	fmt.Printf("%#+v\n", x) }

func AwesomeKeyHandler(ke win.KeyEvent) {
	switch ke.Type {
	case win.EvtKeyPressed:
		log.Println("Key Pressed: ", ke)
	case win.EvtKeyReleased:
		log.Println("Key Released: ", ke)
	}
}

func main() {
	c := sys.NewClock()
	vm := win.NewVideoMode(512,512,24)

	w := gfx.NewRenderWindowDefault(
		vm,
		"HelloWorld",
		win.StyleDefaultStyle)
	black := gfx.FromRGB(128, 128, 128)

	tex, err := gfx.TextureFromFile("./gopher.png", gfx.NewIntRect(0, 0, 153, 55))
	if err != nil {
		log.Fatal(err)
	}
	gopher, _ := gfx.NewSprite()
	gopher.SetTexture(tex, false)
	gopher.SetPosition(200, 200)

	w.SetFramerateLimit(60);

	t := c.GetElapsedTime()
	w.SetActive(true)

	for w.IsOpen() {
		t = c.GetElapsedTime()
		fmt.Println(t.AsSeconds())
		
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

		w.Clear(black)
		w.DrawSprite(gopher)
	}
	log.Println(t.AsSeconds())
}



// 	// fnt := gfx.FontCreateFromFile("./Inconsolata.otf", 96)
// 	// Debug(fnt)

// 	// txt := gfx.StringCreate()
// 	// txt.SetText("Hello Go")
// 	// txt.SetFont(fnt)
// 	// Debug(txt)

// 	img := gfx.ImageCreateFromFile("../test/gopher.png")
// 	gopher := gfx.CreateSprite()
// 	gopher.SetImage(img)
// 	gopher.SetX(200)
// 	gopher.SetY(200)
	
// 	seagreen := gfx.FromRGB(244,23,34)
// 	// //seagreen := gfx.ColorFromRGBA(244,23,34,34)
// 	// //seagreen = gfx.ColorFromRGB(0,34,23)
// 	// Debug(seagreen)

// 	//frame := 0 
// 	//evt := gfx.Event{}

// 	// app.SetFramerateLimit(60)

// 	// var tick float32 = 0
// 	// for app.IsOpened() {
// 	// 	//tick := clock.GetTime()
// 	// 	//txt.SetText(fmt.Sprintf("%v fps", app.GetFrameTime()))
// 	// 	tick += .001;
// 	// 	gopher.SetRotation(tick*100)
// 	// 	scale := float32(1 + (math.Sin(float64(tick)/10)))
// 	// 	gopher.SetScaleX(scale)
// 	// 	gopher.SetScaleY(scale)

// 	// 	app.DrawSprite(gopher)
// 	// 	app.DrawString(txt)

// 	// 	seagreen.R += 1
// 	// 	seagreen.G -= 1
// 	// 	app.Display()
// 	// 	app.Clear(seagreen)
// 	// }
// }

*/
