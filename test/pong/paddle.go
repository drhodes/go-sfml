package main

import (
	"sfml/gfx"
	//"sfml/win"
	//"log"
)

type Paddle struct {
	x, y float32
	sprite *gfx.Sprite
}

func NewPaddle() (*Paddle, error) {
	spr, err := gfx.NewSprite()
	if err != nil {
		return nil, E(err, "Couldn't allocate a paddle")
	}
	
	b := Paddle{HEIGHT-16, WIDTH/2, spr}

	fname := "./assets/paddle.png"
	tex, err := gfx.TextureFromFile(fname, gfx.NewIntRect(0,0,96,16))
	if err != nil {
		return nil, E(err, "Couldn't open image: " + fname)
	}

	b.sprite.SetTexture(*tex, false)	
	return &b, nil
}

func (self *Paddle) update(x float32) {
	if x*2 < WIDTH {
		self.sprite.SetPosition(x*2, HEIGHT-32)
	}
}

func (self *Paddle) collides(b *Ball) bool {
	prect := self.sprite.GetGlobalBounds()
	brect := b.sprite.GetGlobalBounds()
	_, hit := prect.Intersects(brect)
	return hit
}