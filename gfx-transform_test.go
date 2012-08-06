package sfml

import (
	"testing"
)

func TestTranform(t *testing.T) {
	it := IdentityTransform()
	expected := [16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	m := it.GetMatrix()
	for i := 0; i < 16; i += 1 {
		if m[i] != expected[i] {
			t.Fatalf("IdentityTransform gives an incorrect matrix (%dth element: expected %f, got %f)", i+1, expected[i], m[i])
		}
	}
}
