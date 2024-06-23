package main

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type Game struct {
	obstacles     []Sprite
	player        Sprite
	collisionType collisionType
}

type collisionType string

const (
	COLLISION_AABB collisionType = "aabb"
	COLLISION_LINE               = "line"
)

func NewGame(ct string) *Game {

	selectedCollisionDetection := collisionType(ct)

	// load images
	img, _, err := image.Decode(bytes.NewReader(crateImage))
	if err != nil {
		log.Fatal(err)
	}
	crate := ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(dummyImage))
	if err != nil {
		log.Fatal(err)
	}
	dummy := ebiten.NewImageFromImage(img)

	// create some sprites
	sprites := []Sprite{
		{
			X: 100, Y: 100,
			//Angle: 100,
			Scale: 2,
			Img:   crate,
		}, {
			X: 400, Y: 400,
			//Angle: 40,
			Scale: 1,
			Img:   crate,
		}, {
			X: 500, Y: 200,
			Angle: 0,
			Scale: 4,
			Img:   crate,
		},
	}

	player := Sprite{
		X: -100, Y: -100,
		Img:   dummy,
		Angle: 0,
		Scale: 1,
	}

	fmt.Printf("using collision detection: %v\n", selectedCollisionDetection)

	return &Game{
		obstacles:     sprites,
		player:        player,
		collisionType: selectedCollisionDetection,
	}
}

func (g *Game) Update() error {

	// update player based on mouse pos
	mx, my := ebiten.CursorPosition()
	g.player.X = float64(mx)
	g.player.Y = float64(my)

	p := NewBoundingBoxFromSprite(&g.player)

	// do collision
	for _, v := range g.obstacles {

		bb := NewBoundingBoxFromSprite(&v)

		switch g.collisionType {
		case COLLISION_AABB:
			v.CollisionDetected = collisionDetectionAABB(p, bb)
		case COLLISION_LINE:
			v.CollisionDetected = collisionDetectionLine(bb, p)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// draw rectangles
	for _, v := range g.obstacles {
		if !v.CollisionDetected {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(v.Scale, v.Scale)
			op.GeoM.Rotate(v.Angle)
			op.GeoM.Translate(v.X, v.Y)
			screen.DrawImage(v.Img, op)
		}

		if v.CollisionDetected {
			op := &colorm.DrawImageOptions{}
			op.GeoM.Scale(v.Scale, v.Scale)
			op.GeoM.Rotate(v.Angle)
			op.GeoM.Translate(v.X, v.Y)

			var c colorm.ColorM
			c.ChangeHSV(20, 11, 20)
			colorm.DrawImage(screen, v.Img, c, op)
		}
	}

	// draw mouse following sprite
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.player.Scale, g.player.Scale)
	op.GeoM.Rotate(g.player.Angle)

	px, py := g.player.X, g.player.Y
	px -= g.player.Width() / 2
	py -= g.player.Height() / 2

	op.GeoM.Translate(px, py)
	screen.DrawImage(g.player.Img, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
