package main

import (
	//"sfml/gfx"
	//"sfml/win"
	//"log"
)

type Background struct {
	x, y float32
	sprite *gfx.Sprite
}

func NewBackground() (*Background, error) {
	spr, err := gfx.NewSprite()
	if err != nil {
		return nil, E(err, "Couldn't make a background")
	}

	b := Background{0, 0, spr}

	fname := "./assets/background.png"
	tex, err := gfx.TextureFromFile(fname, gfx.NewIntRect(0,0,800,600))
	if err != nil {
		return nil, E(err, "Couldn't open image: " + fname)
	}

	b.sprite.SetTexture(*tex, false)	
	return &b, nil
}
