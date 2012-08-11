package main

import (
	"github.com/drhodes/go-sfml"
	"log"
)

const WIDTH = 800
const HEIGHT = 600

func main() {
	vm := sfml.NewVideoMode(WIDTH, HEIGHT, 24)
	w := sfml.NewRenderWindowDefault(vm, "Pong")

	w.SetMouseCursorVisible(false)

	background, err := NewBackground()
	if err != nil {
		log.Fatal(E(err, `Couldn't create a ball in main`))
	}

	ball, err := NewBall()
	if err != nil {
		log.Fatal(E(err, `Couldn't create a ball in main`))
	}

	paddle, err := NewPaddle()
	if err != nil {
		log.Fatal(E(err, `Couldn't create a paddle in main`))
	}

	w.SetFramerateLimit(60)
	w.SetVerticalSyncEnabled(true)

	for w.IsOpen() {
		e, _ := w.PollEvent()
		switch e.(type) {
		case sfml.KeyEvent:
			KeyHandler(e.(sfml.KeyEvent))
		case sfml.MouseMoveEvent:
			MouseHandle(e.(sfml.MouseMoveEvent), paddle)
		}

		w.Drain()
		w.Clear(sfml.FromRGB(uint8(ball.y/7), uint8(ball.y/7), 45))

		hitfloor := ball.update(paddle)
		if hitfloor {
			w.Clear(sfml.FromRGB(200, 200, 200))
		}

		w.DrawSpriteDefault(background.sprite)
		w.DrawSpriteDefault(paddle.sprite)
		w.DrawSpriteDefault(ball.sprite)
		w.Display()
	}
}
