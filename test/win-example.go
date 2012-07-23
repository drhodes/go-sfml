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
	window.SetFramerateLimit(60)

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

