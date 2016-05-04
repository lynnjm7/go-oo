type Drawer interface {
	Draw()
}

type Magnituder interface {
	Magnitude() float
}

// This is somewhat contrived... but illustrates the point
type MagAndDraw interface {
	Drawer
	Magnituder
}
