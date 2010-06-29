package main

import (	
	"sfml/gfx"
	//"sfml/win"
	"sfml/sys"
	"fmt"
	//"math"
)

func Debug(x interface{}){	fmt.Printf("%#+v\n", x) }

func main() {
	clock := sys.ClockCreate()
	Debug(clock)

	mode := gfx.CreateVideoMode(1024, 1024, 32)
	Debug(mode)
	
	settings := gfx.CreateWindowSettings(24, 8, 1)
	Debug(settings)
	app := gfx.CreateRenderWindow(mode, "Go Wrapper for SFML", 1, settings)
	Debug(app)

	fnt := gfx.FontCreateFromFile("./Inconsolata.otf", 96)
	Debug(fnt)

	txt := gfx.StringCreate()
	txt.SetText("Hello Go")
	txt.SetFont(fnt)
	Debug(txt)

	img := gfx.ImageCreateFromFile("../test/gopher.png")
	gopher := gfx.CreateSprite()
	gopher.SetImage(img)
	gopher.SetX(200)
	gopher.SetY(200)
	
	//seagreen := gfx.Color_FromRGB_P(244,23,34)

	//seagreen := gfx.ColorFromRGB(0,34,23)
	//Debug(seagreen)

	//frame := 0 
	//evt := gfx.Event{}

	app.SetFramerateLimit(10)

	//for app.IsOpened() {			
	for {
		//Debug(app)
		//Debug(evt)
		
		//tick := clock.GetTime()
		txt.SetText(fmt.Sprintf("%v fps", app.GetFrameTime()))
		//gopher.SetRotation(tick*100)
		//scale := float(1 + (math.Sin(float64(tick)/10)))
		//gopher.SetScaleX(scale)
		//gopher.SetScaleY(scale)

		app.DrawSprite(gopher)
		//app.DrawString(txt)
		
		//seagreen.R += 1
		app.Display()
		//app.Clear(seagreen)
	}
}

