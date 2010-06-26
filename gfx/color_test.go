package gfx

import "testing"

func TestConstructors(t *testing.T) {		
	clr1 := ColorFromRGB(1,2,3)
	Debug(clr1)
	
	clr2 := ColorFromRGBA(1,2,3,254)
	Debug(clr2)

	Debug(clr1.Add(clr2))
	Debug(clr1.Modulate(clr2))
}	
