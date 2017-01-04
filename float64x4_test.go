package float64x4


import "testing"
import "math/rand"


func TestNew(t *testing.T) {

	tests := []struct{
		x,y,z,w float64
	}{
		{
			x: 0.0,
			y: 0.0,
			z: 0.0,
			w: 0.0,
		},
		{
			x: 1.0,
			y: 1.0,
			z: 1.0,
			w: 1.0,
		},
		{
			x: 0.5,
			y: 0.5,
			z: 0.5,
			w: 0.5,
		},
		{
			x: 5.0,
			y: 5.0,
			z: 5.0,
			w: 5.0,
		},


		{
			x: 1.0,
			y: 0.0,
			z: 0.0,
			w: 0.0,
		},
		{
			x: 0.0,
			y: 1.0,
			z: 0.0,
			w: 0.0,
		},
		{
			x: 0.0,
			y: 0.0,
			z: 1.0,
			w: 0.0,
		},
		{
			x: 0.0,
			y: 0.0,
			z: 0.0,
			w: 1.0,
		},
	}



	for _,datum := range tests {

		vector := New(datum.x, datum.y, datum.z, datum.w)

		if datum.x != vector.X() {
			t.Errorf("Bad value for X. Expected %v, but instead got %v", datum.x, vector.X)
		}
		if datum.y != vector.Y() {
			t.Errorf("Bad value for Y. Expected %v, but instead got %v", datum.y, vector.Y)
		}
		if datum.z != vector.Z() {
			t.Errorf("Bad value for Z. Expected %v, but instead got %v", datum.z, vector.Z)
		}
		if datum.w != vector.W() {
			t.Errorf("Bad value for W. Expected %v, but instead got %v", datum.w, vector.W)
		}
	}
}

func TestNewRandomly(t *testing.T) {

	for i:=0; i < 200; i++ {

		x := rand.Float64()
		y := rand.Float64()
		z := rand.Float64()
		w := rand.Float64()


		vector := New(x, y, z, w)

		if x != vector.X() {
			t.Errorf("Bad value for X. Expected %v, but instead got %v", x, vector.X)
		}
		if y != vector.Y() {
			t.Errorf("Bad value for Y. Expected %v, but instead got %v", y, vector.Y)
		}
		if z != vector.Z() {
			t.Errorf("Bad value for Z. Expected %v, but instead got %v", z, vector.Z)
		}
		if w != vector.W() {
			t.Errorf("Bad value for W. Expected %v, but instead got %v", w, vector.W)
		}
	}
}


func TestZero(t *testing.T) {

	vector := Zero()

	if 0.0 != vector.X() {
		t.Errorf("Bad value for X. Expected %v, but instead got %v", 0.0, vector.X)
	}
	if 0.0 != vector.Y() {
		t.Errorf("Bad value for Y. Expected %v, but instead got %v", 0.0, vector.Y)
	}
	if 0.0 != vector.Z() {
		t.Errorf("Bad value for Z. Expected %v, but instead got %v", 0.0, vector.Z)
	}
	if 0.0 != vector.W() {
		t.Errorf("Bad value for W. Expected %v, but instead got %v", 0.0, vector.W)
	}

}


func TestSplat(t *testing.T) {

	tests := []struct{
		a float64
	}{
		{
			a: 0.0,
		},
		{
			a: 1.0,
		},
		{
			a: 0.5,
		},
		{
			a: 5.0,
		},
	}



	for _,datum := range tests {

		vector := Splat(datum.a)

		if datum.a != vector.X() {
			t.Errorf("Bad value for X. Expected %v, but instead got %v", datum.a, vector.X)
		}
		if datum.a != vector.Y() {
			t.Errorf("Bad value for Y. Expected %v, but instead got %v", datum.a, vector.Y)
		}
		if datum.a != vector.Z() {
			t.Errorf("Bad value for Z. Expected %v, but instead got %v", datum.a, vector.Z)
		}
		if datum.a != vector.W() {
			t.Errorf("Bad value for W. Expected %v, but instead got %v", datum.a, vector.W)
		}
	}
}

func TestSplatRandomly(t *testing.T) {

	for i:=0; i < 200; i++ {

		number := rand.Float64()


		vector := Splat(number)

		if number != vector.X() {
			t.Errorf("Bad value for X. Expected %v, but instead got %v", number, vector.X)
		}
		if number != vector.Y() {
			t.Errorf("Bad value for Y. Expected %v, but instead got %v", number, vector.Y)
		}
		if number != vector.Z() {
			t.Errorf("Bad value for Z. Expected %v, but instead got %v", number, vector.Z)
		}
		if number != vector.W() {
			t.Errorf("Bad value for W. Expected %v, but instead got %v", number, vector.W)
		}
	}
}
