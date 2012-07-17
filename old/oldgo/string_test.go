package gfx

import "testing"

func TestStringConstructors(t *testing.T) {			
	Space()
	Debug("String ---------------------------------------------")

	str1 := StringCreate()	
	Debug(str1)

	str1.SetText(`Hello World!`)
	Debug(str1.GetUnicodeText())

	// str1.GetCharacterPos(2, x, y)  
	// Have to investigate this further.
}	
