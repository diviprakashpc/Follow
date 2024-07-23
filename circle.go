package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Circle struct {
	img        *ebiten.Image
	length     float32
	height     float32
	x          float32
	y          float32
	r          int
	angle      float64
	LerpFactor float32
	tail       []Circle
}

func (c *Circle) UpdateTail() {
	for i := 0; i < len(c.tail); i++ {
		parent := c
		if i != 0 {
			parent = &c.tail[i-1]
		}
		c.tail[i].Rot(int(parent.x), int(parent.y))
		destX := parent.x - parent.length
		destY := parent.y
		xx := (parent.x-destX)*float32(math.Cos(parent.angle)) - (parent.y-destY)*float32(math.Sin(parent.angle)) + parent.x
		yy := (parent.x-destX)*float32(math.Sin(parent.angle)) + (parent.y-destY)*float32(math.Cos(parent.angle)) + parent.y
		c.tail[i].LerpTowards(xx, yy, parent.LerpFactor)
	}
}

func (c *Circle) DrawTail(screen *ebiten.Image) {
	for i := 0; i < len(c.tail); i++ {
		c.tail[i].Draw(screen)
	}
}

func (c *Circle) Rot(x, y int) {
	xx := -(float32(x) - c.x)
	yy := -(float32(y) - c.y)
	targetAngle := math.Atan2(float64(yy), float64(xx))
	c.angle = targetAngle
}

func (c *Circle) LerpTowards(x, y, f float32) {
	c.x = c.Lerp(c.x, x, f)
	c.y = c.Lerp(c.y, y, f)
}

func (c *Circle) Lerp(a float32, b float32, f float32) float32 {
	return a + (b-a)*f
}

func (c *Circle) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(c.img.Bounds().Dx())/2, -float64(c.img.Bounds().Dy())/2)
	opt.GeoM.Rotate(c.angle)
	opt.GeoM.Translate(float64(c.img.Bounds().Dx())/2, float64(c.img.Bounds().Dy())/2)
	opt.GeoM.Translate(float64(c.x), float64(c.y))
	screen.DrawImage(c.img, opt)
}
