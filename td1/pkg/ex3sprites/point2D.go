package ex3sprites

import (
	"math"
)

type Point2D struct {
	x,y float64
}

func (p* Point2D) X() float64 {
	return p.x
}
func (p* Point2D) Y() float64 {
	return p.y
}

func (p* Point2D) SetX(x float64) {
	p.x = x
}
func (p* Point2D) SetY(y float64) {
	p.y = y
}

func NewPoint2D(x, y float64) *Point2D {
	return &Point2D{x,y}
}

func (p* Point2D) Clone() *Point2D {
	return &Point2D{p.X(), p.Y()}
}

func (p* Point2D) Module() float64 {
	mod := math.Sqrt(p.X()*p.X() + p.Y()*p.Y())
	return mod
}

func (p* Point2D) Move(dx, dy float64) {
	p.x += dx
	p.y += dy
}
