package main

import (
	"fmt"
	"github.com/drhodes/go-sfml"
	"log"
)

const WIDTH = 800
const HEIGHT = 600

func main() {
	vm := win.NewVideoMode(WIDTH, HEIGHT, 24)
	w := gfx.NewRenderWindowDefault(vm, "Pong")

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

	w.SetFrameRateLimit(60)
	w.SetVerticalSyncEnabled(true)

	for w.IsOpen() {
		e, _ := w.PollEvent()
		switch e.(type) {
		case win.KeyEvent:
			KeyHandler(e.(win.KeyEvent))
		case win.MouseMoveEvent:
			MouseHandle(e.(win.MouseMoveEvent), paddle)
		}

		w.Drain()
		w.Clear(gfx.FromRGB(uint8(ball.y/7), uint8(ball.y/7), 45))

		hitfloor := ball.update(paddle)
		if hitfloor {
			w.Clear(gfx.FromRGB(200, 200, 200))
		}

		w.DrawSpriteDefault(background.sprite)
		w.DrawSpriteDefault(paddle.sprite)
		w.DrawSpriteDefault(ball.sprite)
		w.Display()
	}
}
