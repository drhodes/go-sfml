package win

import "testing"
import . "sfml/debug"

func TestContextConstructors(t *testing.T) {			
	Space()
	Debug("Contect ---------------------------------------------")

	ctx := CreateContext()
	Debug(ctx)	
	ctx.SetActive(true)
	Debug(ctx)	
	ctx.SetActive(false)
	Debug(ctx)	
	ctx.Destroy()
	Debug(ctx)	
}	
