package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	c Circle
}

const (
	constLerp  = 0.05
	TailLength = 10
)

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	g.c.LerpTowards(float32(x), float32(y), constLerp)
	g.c.Rot(x, y)
	g.c.UpdateTail()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.c.Draw(screen)
	g.c.DrawTail(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	im := ebiten.NewImage(10, 10)
	im.Fill(color.White)
	im2 := ebiten.NewImage(10, 10)
	im2.Fill(color.RGBA{100, 0, 0, 255})
	im3 := ebiten.NewImage(10, 10)
	im3.Fill(color.RGBA{100, 100, 0, 255})
	g := &Game{
		Circle{
			img:        im,
			LerpFactor: 0.3,
			length:     10,
			height:     10,
		},
	}
	for i := 0; i <= TailLength; i++ {
		img := ebiten.NewImage(10, 10)
		r := rand.Intn(256)
		gr := rand.Intn(256)
		b := rand.Intn(256)
		a := rand.Intn(256)
		img.Fill(color.RGBA{uint8(r), uint8(gr), uint8(b), uint8(a)})
		t := Circle{
			img:        img,
			length:     10,
			height:     10,
			LerpFactor: 0.25,
		}
		g.c.tail = append(g.c.tail, t)
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
