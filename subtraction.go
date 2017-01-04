package float64x4



func (a T) Subtract(b T) T {

	return T{
		a[x] - b[x],
		a[y] - b[y],
		a[z] - b[z],
		a[w] - b[w],
	}
}



func (a *T) SubtractFrom(b T) {

	a[x] -= b[x]
	a[y] -= b[y]
	a[z] -= b[z]
	a[w] -= b[w]
}
