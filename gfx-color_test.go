package sfml

import (
	"testing"
)

func TestColor(t *testing.T) {
	var (
		r uint8
		g uint8
		b uint8
		a uint8
	)
	c1 := FromRGBA(12, 12, 12, 12)
	c2 := FromRGB(12, 12, 12)
	r, g, b, a = c1.Components()
	if r != 12 || g != 12 || b != 12 || a != 12 {
		t.Fatalf("Components don't return the expected values (expected 12,12,12,12, got %d,%d,%d,%d)", r, g, b, a)
	}
	c1 = c1.Modulate(c2)
	r, g, b, a = c1.Components()
	if r != 0 || g != 0 || b != 0 || a != 12 {
		t.Fatalf("Modulate didn't produce the expected result (color should be 0,0,0,12, but is %d,%d,%d,%d)", r, g, b, a)
	}
	c1 = c1.Add(c2)
	r, g, b, a = c1.Components()
	if r != 12 || g != 12 || b != 12 || a != 255 {
		t.Fatalf("Add didn't produce the expected result (color should be 12,12,12,12, but is %d,%d,%d,%d)", r, g, b, a)
	}
}
