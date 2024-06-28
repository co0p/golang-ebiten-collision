package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	obstacles     []*Sprite
	player        Sprite
	collisionType collisionType
}

type collisionType string

const (
	COLLISION_AABB collisionType = "aabb"
	COLLISION_LINE collisionType = "line"
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
	sprites := []*Sprite{
		&Sprite{
			X: 100, Y: 100,
			// Angle: 100,
			Scale: 1,
			Img:   crate,
		}, &Sprite{
			X: 400, Y: 400,
			// Angle: 40,
			Scale: 1,
			Img:   crate,
		}, &Sprite{
			X: 500, Y: 200,
			// Angle: 0,
			Scale: 1,
			Img:   crate,
		},
	}

	player := Sprite{
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

		bb := NewBoundingBoxFromSprite(v)

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

		ebitenutil.DebugPrint(screen, fmt.Sprintf("collision detected: %v", v.CollisionDetected))
	}

	// draw rect around the obstacles
	gray := color.RGBA{100, 100, 100, 100}
	for _, o := range g.obstacles {
		bb := NewBoundingBoxFromSprite(o)

		// lt, rt, rb, lb
		vertices := bb.Vertices()
		lt, rt, rb, lb := vertices[0], vertices[1], vertices[2], vertices[3]
		vector.StrokeLine(screen, float32(lt.X), float32(lt.Y), float32(rt.X), float32(rt.Y), 2, gray, false)
		vector.StrokeLine(screen, float32(rt.X), float32(rt.Y), float32(rb.X), float32(rb.Y), 2, gray, false)
		vector.StrokeLine(screen, float32(rb.X), float32(rb.Y), float32(lb.X), float32(lb.Y), 2, gray, false)
		vector.StrokeLine(screen, float32(lb.X), float32(lb.Y), float32(lt.X), float32(lt.Y), 2, gray, false)

	}

	// draw mouse following sprite
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.player.Scale, g.player.Scale)
	op.GeoM.Rotate(g.player.Angle)

	op.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(g.player.Img, op)

	red := color.RGBA{255, 50, 50, 100}
	bb := NewBoundingBoxFromSprite(&g.player)
	vertices := bb.Vertices()
	lt, rt, rb, lb := vertices[0], vertices[1], vertices[2], vertices[3]
	vector.StrokeLine(screen, float32(lt.X), float32(lt.Y), float32(rt.X), float32(rt.Y), 2, red, false)
	vector.StrokeLine(screen, float32(rt.X), float32(rt.Y), float32(rb.X), float32(rb.Y), 2, red, false)
	vector.StrokeLine(screen, float32(rb.X), float32(rb.Y), float32(lb.X), float32(lb.Y), 2, red, false)
	vector.StrokeLine(screen, float32(lb.X), float32(lb.Y), float32(lt.X), float32(lt.Y), 2, red, false)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
