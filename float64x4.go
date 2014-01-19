package float64x4



type T struct {
	X,Y,Z,W float64
}



func New(x,y,z,w float64) T {
	return T{X:x, Y:y, Z:z, W:w}
}



func Zero() T {
	return T{X:0.0, Y:0.0, Z:0.0, W:0.0}
}



func Splat(a float64) T {
	return T{X:a, Y:a, Z:a, W:a}
}
