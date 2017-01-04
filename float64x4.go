package float64x4

// T represents the floating point approximation of ℝ⁴.
// I.e., 4 float64.
//
// Note that these can be used these can be used for
// affine transformation in ℝ³.
type T [4]float64

// X returns the 1st dimension.
func (receiver T) X() float64 {
	return receiver[x]
}

// Y returns the 2nd dimension.
func (receiver T) Y() float64 {
	return receiver[y]
}

// Z returns the 3rd dimension.
func (receiver T) Z() float64 {
	return receiver[z]
}

// W returns the 4th dimension.
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
