package sfml

import (
	"testing"
)

func TestImage(t *testing.T) {
	img, err := ImageFromFile("test/gopher.png")
	if err != nil {
		t.Fatal(err)
	}
	w, h := img.GetSize()
	if w != 153 || h != 55 {
		t.Fatal("Size reported by GetSize is wrong")
	}
}

func TestNonExistentImage(t *testing.T) {
	_, err := ImageFromFile("dont-exist.png")
	if err == nil {
		t.Fatal("Got no error when loading an image from a non existent file")
	}
}