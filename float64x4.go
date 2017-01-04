package float64x4

type T [4]float64

func (receiver T) X() float64 {
	return receiver[x]
}

func (receiver T) Y() float64 {
	return receiver[y]
}

func (receiver T) Z() float64 {
	return receiver[z]
}

func (receiver T) W() float64 {
	return receiver[w]
}


func New(x,y,z,w float64) T {
	return T{x, y, z, w}
}



func Zero() T {
	return T{0.0, 0.0, 0.0, 0.0}
}



func Splat(a float64) T {
	return T{a, a, a, a}
}
