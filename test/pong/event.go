package main

import (	
	"sfml/win"
	"log"
)

func KeyHandler(ke win.KeyEvent) {
	switch ke.Type {
	case win.EvtKeyPressed:
		if win.IsKeyPressed(win.KeyEscape) {
			log.Fatal("Quit.")
		}
		if win.IsKeyPressed(win.KeyQ) {
			log.Fatal("Quit.")
		}

	case win.EvtKeyReleased:
		//log.Println("Key Released: ", ke)
	}
}

func MouseHandle(ke win.MouseMoveEvent, p *Paddle) {
	p.update(float32(ke.X()))
}

