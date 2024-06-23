package main

type vec2 struct {
	X float64
	Y float64
}

type boundingBox struct {
	position      vec2
	width, height float64
	vertices      []vec2
}

func NewBoundingBoxFromSprite(s *Sprite) boundingBox {
	return boundingBox{
		position: vec2{s.X, s.Y},
		width:    s.Width(),
		height:   s.Height(),
		vertices: vertices(s.X, s.Y, s.Width(), s.Height()),
	}
}

func (bb *boundingBox) Vertices() []vec2 {
	return bb.vertices
}

func vertices(X, Y, Width, Height float64) []vec2 {
	lt := vec2{X, Y}
	rt := vec2{X + Width, Y}
	rb := vec2{X + Height, Y + Height}
	lb := vec2{X, Y + Height}

	return []vec2{lt, rt, rb, lb}
}

func collisionDetectionAABB(a boundingBox, b boundingBox) bool {
	return a.position.X < b.position.X+b.width &&
		a.position.X+a.width > b.position.X &&
		a.position.Y < b.position.Y+b.height &&
		a.position.Y+a.height > b.position.Y
}

func collisionDetectionLine(a boundingBox, b boundingBox) bool {
	return false
}
