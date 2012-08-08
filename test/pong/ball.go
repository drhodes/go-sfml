package main

import (
	"github.com/drhodes/go-sfml"
)

type Ball struct {
	speedx float32
	speedy float32
	x, y   float32
	sprite sfml.Sprite
}

func NewBall() (*Ball, error) {
	spr, err := sfml.NewSprite()
	if err != nil {
		return nil, E(err, "Couldn't make a ball")
	}

	b := Ball{0, 5, 1, WIDTH / 2, spr}

	fname := "./assets/ball.png"
	tex, err := sfml.TextureFromFile(fname, sfml.NewIntRect(0, 0, 16, 16))
	if err != nil {
		return nil, E(err, "Couldn't open image: "+fname)
	}

	b.sprite.SetTexture(tex, false)
	b.sprite.Scale(.5, .5)
	return &b, nil
}
func (self *Ball) Width() float32 {
	return self.sprite.GetLocalBounds().Width()
}

// return false if ball hits floor
func (self *Ball) update(paddle *Paddle) bool {
	if paddle.collides(self) {
		factor := paddle.ImpulseX(self)
		self.speedx += factor
		self.speedy *= -1
	}

	switch {
	case self.x <= 0:
		// hit the left wall, start moving right
		self.speedx *= -1
		break
	case self.x >= WIDTH-16:
		// hit the right wall, start moving left
		self.speedx *= -1
		break
	}

	switch {
	case self.y <= 0:
		// hit the top wall, start moving down
		self.speedy *= -1
		break
	case self.y >= HEIGHT-16:
		// hit the floor
		self.speedy *= -1
		self.y = 10
		self.speedx = 0
		self.speedy = 5
		return true
	}
	self.x += self.speedx
	self.y += self.speedy
	self.sprite.SetPosition(self.x, self.y)
	return false
}