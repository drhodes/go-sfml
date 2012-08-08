package main

import (
	"github.com/drhodes/go-sfml"
	"log"
)

func main() {
	vm := sfml.NewVideoMode(800, 600, 32)
	window := sfml.NewRenderWindowDefault(vm, "SFML window")
	black := sfml.FromRGB(0, 0, 0)
	window.SetFramerateLimit(60)

	texture, err := sfml.TextureFromFile("gopher.png", sfml.IntRect{})
	if err != nil {
		log.Fatal(err)
	}

	sprite, err := sfml.NewSprite()
	if err != nil {
		log.Fatal(err)
	}
	sprite.SetTexture(texture, false)

	font, err := sfml.FontFromFile("Inconsolata.otf")
	if err != nil {
		log.Fatal(err)
	}

	text, err := sfml.NewText()
	if err != nil {
		log.Fatal(err)
	}
	text.SetString("Hello Go-SFML")
	text.SetFont(font)
	text.SetCharacterSize(50)

	for window.IsOpen() {
		e, _ := window.PollEvent()
		switch e.(type) {
		case sfml.KeyEvent:
			ev := e.(sfml.KeyEvent)
			if ev.Code() == sfml.KeyEscape {
				window.Close()
			}
		}
		window.Clear(black)
		window.DrawSpriteDefault(sprite)
		window.DrawTextDefault(text)
		window.Display()
	}
}

