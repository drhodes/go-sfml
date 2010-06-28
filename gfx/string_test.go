package gfx

import "testing"

func TestStringConstructors(t *testing.T) {			
	Space()
	Debug("String ---------------------------------------------")

	str1 := StringCreate()	
	Debug(str1)

	str1.SetText(`Hello World!`)
	Debug(str1.GetUnicodeText())

	var x, y *float;
	str1.GetCharacterPos(2, x, y)
	DebugN(x)
	DebugN(" ")
	Debug(y)


}	
