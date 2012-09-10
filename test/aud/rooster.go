package main

import (
	"github.com/drhodes/go-sfml"
	"log"
)

const WIDTH = 400
const HEIGHT = 400

func main() {
	vm := sfml.NewVideoMode(WIDTH, HEIGHT, 24)
	w := sfml.NewRenderWindowDefault(vm, "Rooster")
	w.SetMouseCursorVisible(false)
	w.SetFramerateLimit(60)
	w.SetVerticalSyncEnabled(true)

	snd := sfml.

	for w.IsOpen() {
		e, _ := w.PollEvent()
		switch e.(type) {
		case sfml.KeyEvent:
			break		
		w.Drain()
	}
}
