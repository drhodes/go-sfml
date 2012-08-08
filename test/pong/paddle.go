package main

import (
	"github.com/drhodes/go-sfml"
)

type Paddle struct {
	x, y   float32
	sprite sfml.Sprite
}

func NewPaddle() (*Paddle, error) {
	spr, err := sfml.NewSprite()
	if err != nil {
		return nil, E(err, "Couldn't allocate a paddle")
	}

	b := Paddle{HEIGHT - 16, WIDTH / 2, spr}

	fname := "./assets/paddle.png"
	tex, err := sfml.TextureFromFile(fname, sfml.NewIntRect(0, 0, 96, 16))
	if err != nil {
		return nil, E(err, "Couldn't open image: "+fname)
	}

	b.sprite.SetTexture(tex, false)
	return &b, nil
}

func (self *Paddle) update(x float32) {
	self.x = x * 2
	if self.x < WIDTH {
		self.sprite.SetPosition(self.x, HEIGHT-32)
	}
}

func (self *Paddle) collides(b *Ball) bool {
	prect := self.sprite.GetGlobalBounds()
	brect := b.sprite.GetGlobalBounds()
	_, hit := prect.Intersects(brect)
	return hit
}

// find an impulse depending on where the ball hits the paddle 
func (self *Paddle) ImpulseX(b *Ball) float32 {
	// -2 <-----|-----> 2
	padleft := self.x
	padwidth := self.sprite.GetGlobalBounds().Width()
	padcenter := padleft + padwidth/2
	ballcenter := b.x

	factor := -((padcenter - ballcenter) / (padwidth/2))	
	return factor * 2
}