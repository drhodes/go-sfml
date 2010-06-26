package main

import (	
	"sfml/gfx"
	"sfml/win"
	"sfml/sys"
	"fmt"
	//"math"
)

func Debug(x interface{}){	fmt.Printf("%#+v\n", x) }

func main() {
	clock := sys.ClockCreate()
	Debug(clock)

	mode := win.VideoModeCreate(1024, 1024, 32)
	settings := win.WindowSettingsCreate(24, 8, 1)
	app := gfx.RenderWindowCreate(mode, "Go Wrapper for SFML", 1, settings)
	Debug(app)
		/*
	fnt := gfx.FontCreateFromFile("./Inconsolata.otf", 96)

	txt := gfx.StringCreate()
	txt.SetText("Hello Go")
	txt.SetFont(fnt)

	img := gfx.ImageCreateFromFile("/tmp/path6376.png")
	gopher := gfx.SpriteCreate()
	gopher.SetImage(img)
	gopher.SetX(200)
	gopher.SetY(200)

	seagreen := gfx.ColorFromRGBA(0x40, 0x62, 0x64, 255)
	
	frame := 0 
	//evt := gfx.Event{}

	app.SetFramerateLimit(11000)
	for app.IsOpened() {			

		//Debug(app.GetEvent(evt))
		//Debug(evt)
		
		tick := clock.GetTime()
		txt.SetText(fmt.Sprintf("%v fps", int(frame)/int(tick+1)))
		gopher.SetRotation(float32(tick)*100)
		gopher.SetScaleX(1 + float32(math.Sin(float64(tick)/10)))
		gopher.SetScaleY(1 + float32(math.Sin(float64(tick)/10)))

		app.DrawSprite(gopher)
		app.DrawString(txt)
	
		app.Display()
		app.Clear(seagreen)		
		frame += 1
	}
	 */
}

