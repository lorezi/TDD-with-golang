package structmethodsinterfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func (r Rectangle) Area(width, height float64) float64 {

	return width * height
}
