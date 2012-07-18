package gfx

import (
	"testing"
	"log"
)

//const verb = true

func TestTranform(t *testing.T) {		
	it := IdentityTransform()
	log.Println(it.GetMatrix())
}



