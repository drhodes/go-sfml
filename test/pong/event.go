package main

import (
	"github.com/drhodes/go-sfml"
	"log"
)

func KeyHandler(ke sfml.KeyEvent) {
	switch ke.Type {
	case sfml.EvtKeyPressed:
		if sfml.IsKeyPressed(sfml.KeyEscape) {
			log.Fatal("Quit.")
		}
		if sfml.IsKeyPressed(sfml.KeyQ) {
			log.Fatal("Quit.")
		}

	case sfml.EvtKeyReleased:
		//log.Println("Key Released: ", ke)
	}
}

func MouseHandle(ke sfml.MouseMoveEvent, p *Paddle) {
	p.update(float32(ke.X()))
}
