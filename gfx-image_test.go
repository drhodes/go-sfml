// Copyright 2012.  All rights reserved. 
// Use of this source code is governed by a BSD-style  
// license that can be found in the LICENSE file.
package sfml


import (
	"testing"
)

func TestImage(t *testing.T) {
	img, err := ImageFromFile("test/gopher.png")
	if err != nil {
		t.Fatal(err)
	}
	w, h := img.Size()
	if w != 153 || h != 55 {
		t.Fatal("Size don't return the good size (expected 153x55, got %dx%d)", w, h)
	}
}

func TestNonExistentImage(t *testing.T) {
	_, err := ImageFromFile("dont-exist.png")
	if err == nil {
		t.Fatal("Got no error when loading an image from a non existent file")
	}
}