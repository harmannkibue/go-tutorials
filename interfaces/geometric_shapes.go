package interfaces

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, length float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.length) * 2
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return math.pi * (c.radius * 2)
}

func main() {

}
