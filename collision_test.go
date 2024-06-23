package main

import (
	"slices"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

var emptyImage = ebiten.NewImage(200, 200)
var squareImage100x100 = ebiten.NewImage(100, 100)

func TestBoundingBoxVertices(t *testing.T) {

	givenSprite := Sprite{
		X:     100,
		Y:     50,
		Angle: 0,
		Scale: 2,
		Img:   emptyImage,
	}

	bb := NewBoundingBoxFromSprite(&givenSprite)

	expectedVertices := []vec2{{100, 50}, {500, 50}, {500, 450}, {100, 450}}
	actualVertices := bb.Vertices()
	equal := slices.Equal(expectedVertices, actualVertices)
	if !equal {
		t.Errorf("expected '%v' to be equal to '%v', got '%v'\n", actualVertices, expectedVertices, equal)
	}
}

func TestCollisionDetectionAABB(t *testing.T) {

	spriteA := Sprite{
		X:     0,
		Y:     0,
		Angle: 0,
		Scale: 1,
		Img:   squareImage100x100,
	}
	spriteB := Sprite{
		X:     50,
		Y:     50,
		Angle: 0,
		Scale: 1,
		Img:   squareImage100x100,
	}
	spriteC := Sprite{
		X:     500,
		Y:     500,
		Angle: 0,
		Scale: 1,
		Img:   squareImage100x100,
	}

	bbA := NewBoundingBoxFromSprite(&spriteA)
	bbB := NewBoundingBoxFromSprite(&spriteB)
	bbC := NewBoundingBoxFromSprite(&spriteC)

	collision := collisionDetectionAABB(bbA, bbB)
	if !collision {
		t.Errorf("expected '%v' to be aabb colliding with '%v', got %v\n", bbA, bbB, collision)
	}

	noCollision := collisionDetectionAABB(bbA, bbC)
	if noCollision {
		t.Errorf("expected '%v' not to be aabb colliding with '%v', got %v\n", bbA, bbC, noCollision)
	}
}
