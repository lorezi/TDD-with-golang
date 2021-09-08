package structmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func Area(r Rectangle) float64 {

	return r.Width * r.Height
}
