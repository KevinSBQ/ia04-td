package ex3sprites

type Rectangle struct {
	hg, bd Point2D
}

func (r* Rectangle) Move(dx, dy float64) {
	r.hg.Move(dx, dy)
	r.bd.Move(dx, dy)
}