package gfx

import "testing"

func TestImageConstructors(t *testing.T) {			
	Space()
	Debug("Image ---------------------------------------------")

	img1 := ImageCreate()	
	Debug(img1)
	
	//func (self *Image) CreateFromColor(width, height uint, clr Color) Image 
	img2 := ImageCreateFromColor(100, 100, ColorFromRGB(1,2,3) )
	Debug(img2)

	img3 := ImageCreateFromFile("../test/gopher.png")
	Debug(img3)

	Debug(int(img3.GetHeight()))
	Debug(int(img3.GetWidth()))
	Debug(img3.IsSmooth())
	img3.SetSmooth(false)
	Debug(img3.IsSmooth())
	Debug(img3.SaveToFile("/tmp/gopher_menamena.png"))

}	
