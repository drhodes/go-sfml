package main

import (
	"github.com/drhodes/go-sfml"
)

type Background struct {
	x, y   float32
	sprite *sfml.Sprite
}

func NewBackground() (*Background, error) {
	spr, err := sfml.NewSprite()
	if err != nil {
		return nil, E(err, "Couldn't make a background")
	}

	b := Background{0, 0, spr}

	fname := "./assets/background.png"
	tex, err := sfml.TextureFromFile(fname, sfml.NewIntRect(0, 0, 800, 600))
	if err != nil {
		return nil, E(err, "Couldn't open image: "+fname)
	}

	b.sprite.SetTexture(*tex, false)
	return &b, nil
}
