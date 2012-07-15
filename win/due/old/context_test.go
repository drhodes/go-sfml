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

	mode := GetDesktopMode()
	Debug(mode)


	Debug(GetMode(1))
	Debug(GetModesCount())
	Debug(mode.IsValid())
	mode.SetWidth(45)
	Debug(mode.IsValid())
}
