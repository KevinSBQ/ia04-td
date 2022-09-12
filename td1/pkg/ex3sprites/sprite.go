package ex3sprites

import (
	"math"
)

type Sprite struct {
	pos Point2D
	hitbox Rectangle
	zoom float64
	nom string
}

func (s* Sprite) Move(dx, dy float64){
	s.pos.Move(dx, dy)
	s.hitbox.Move(dx, dy)
}

func Collision(s1, s2 Sprite) *Rectangle{
	x1 := math.Max(s1.zoom * (s1.hitbox.hg.x - s1.pos.x) + s1.pos.x, s2.zoom * (s2.hitbox.hg.x - s2.pos.x) + s2.pos.x)
	y1 := math.Max(s1.zoom * (s1.hitbox.hg.y - s1.pos.y) + s1.pos.y, s2.zoom * (s2.hitbox.hg.y - s2.pos.y) + s2.pos.y)
	x2 := math.Min(s1.zoom * (s1.hitbox.hg.x - s1.pos.x) + s1.pos.x, s2.zoom * (s2.hitbox.hg.x - s2.pos.x) + s2.pos.x)
	y2 := math.Min(s1.zoom * (s1.hitbox.hg.y - s1.pos.y) + s1.pos.y, s2.zoom * (s2.hitbox.hg.y - s2.pos.y) + s2.pos.y)
	// x1 y1 = inter hg
	// x2 y2 = inter bd
	if x1>x2 || y1>y2 {
		return nil
	}
	return &Rectangle{*NewPoint2D(x1, y1), *NewPoint2D(x2, y2)} // NewRectangle()
}