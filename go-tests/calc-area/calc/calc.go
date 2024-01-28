package calc

import "math"

type Forma interface {
	area() float64
}

type Retang struct {
	H float64
	Larg float64
}

func (r Retang) Area() float64 {
	return r.H * r.Larg
}

type Circle struct {
	H float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.H, 2) 
}
