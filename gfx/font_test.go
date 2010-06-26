package gfx

import "testing"

func TestFontConstructors(t *testing.T) {			
	Space()
	Debug("Font --------------------------------------------")
	fnt := FontCreate()	
	Debug(fnt)

	fnt = FontCreateFromFile("../test/Inconsolata.otf", 34)					
	Debug(fnt)	 
}	
