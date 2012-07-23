package main

import (
	"sfml/gfx"
	//"sfml/win"
	//"log"
)

type Ball struct {
	speed float32
	lastx, lasty float32
	x, y float32
	sprite *gfx.Sprite
}

func NewBall() (*Ball, error) {
	spr, err := gfx.NewSprite()
	if err != nil {
		return nil, E(err, "Couldn't make a ball")
	}

	b := Ball{5, 1, 1, 2, 2, spr}

	fname := "./assets/ball.png"
	tex, err := gfx.TextureFromFile(fname, gfx.NewIntRect(0,0,16,16))
	if err != nil {
		return nil, E(err, "Couldn't open image: " + fname)
	}

	b.sprite.SetTexture(*tex, false)	
	b.sprite.Scale(.5,.5)
	return &b, nil
}

func (self *Ball) xnudge(n float32) {
	self.lastx = self.x
	self.x += n
}
func (self *Ball) ynudge(n float32) {
	self.lasty = self.y
	self.y += n
}

// return false if ball hits floor
func (self *Ball) update(paddle *Paddle) bool {
	if paddle.collides(self) {
		self.speed += .1
		//log.Println("Bounce!")
		self.ynudge(-self.speed)					
	}

	switch {
	case self.x <= 0:
		// hit the left wall, start moving right
		//self.sprite.Scale(.99,.99)
		self.xnudge(self.speed)
		break
	case self.x >= WIDTH - 16:
		// hit the right wall, start moving left
		//self.sprite.Scale(.99,.99)

		self.xnudge(-self.speed)
		break
	case self.x > 0 && self.x < WIDTH && self.x < self.lastx:
		// in the middle moving left 
		self.xnudge(-self.speed)
		break
	case self.x > 0 && self.x < WIDTH && self.x > self.lastx:
		// in the middle moving left 
		self.xnudge(self.speed)
	}

	switch {
	case self.y <= 0:
		// hit the top wall, start moving down
		self.ynudge(+self.speed)
		break
	case self.y >= HEIGHT - 16:
		// hit the floor
		self.speed += .1
		//log.Println("BOOM")
		self.ynudge(-self.speed)			
		return true
	case self.y > 0 && self.y < HEIGHT && self.y > self.lasty:
		// in the middle moving down
		self.ynudge(+self.speed)
		break
	case self.y > 0 && self.y < HEIGHT && self.y < self.lasty:
		// in the middle moving up 
		self.ynudge(-self.speed)
	}
	self.sprite.SetPosition(self.x, self.y)
	return false
}