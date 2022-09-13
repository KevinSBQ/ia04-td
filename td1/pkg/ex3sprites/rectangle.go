package ex3sprites

type Rectangle struct {
	hg, bd Point2D
}

func NewRectangle (hg, bd Point2D) *Rectangle {
	return &Rectangle{hg, bd}
}

func (r* Rectangle) GetRectangle() *Rectangle {
	return r
}

func (r* Rectangle) SetRectangle(newhg, newbd Point2D) *Rectangle {
	if newhg.x < newbd.x && newhg.y < newbd.y {
		r.bd = newbd
		r.hg = newhg
	}
	return r
}

func (r* Rectangle) GetHG() *Point2D {
	return &r.hg
}

func (r* Rectangle) SetHG(p Point2D) {
	if p.x < r.bd.x && p.y < r.bd.y {
	r.hg = p
	}
}

func (r* Rectangle) GetBD() *Point2D {
	return &r.bd
}

func (r* Rectangle) SetBD(p Point2D) {
	if p.x > r.bd.x && p.y > r.bd.y {
	r.bd = p
	}
}

func (r* Rectangle) Move(dx, dy float64) {
	r.hg.Move(dx, dy)
	r.bd.Move(dx, dy)
}