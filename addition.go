package float64x4



func (a T) Add(b T) T {

	return T{
		a[x] + b[x],
		a[y] + b[y],
		a[z] + b[z],
		a[w] + b[w],
	}
}



func (a *T) AddTo(b T) {

	a[x] += b[x]
	a[y] += b[y]
	a[z] += b[z]
	a[w] += b[w]
}
