package float64x4



func (a T) Multiply(b T) T {

	return T{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
		W: a.W * b.W,
	}
}



func (a *T) MultiplyBy(b T) {

	a.X *= b.X
	a.Y *= b.Y
	a.Z *= b.Z
	a.W *= b.W
}
