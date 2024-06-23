package main

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	X, Y              float64
	Angle             float64
	Scale             float64
	Img               *ebiten.Image
	CollisionDetected bool
}

func (s Sprite) Width() float64 {
	return s.Scale * float64(s.Img.Bounds().Dx())
}

func (s Sprite) Height() float64 {
	return s.Scale * float64(s.Img.Bounds().Dy())
}
