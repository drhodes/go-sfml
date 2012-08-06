package sfml


import (
	"testing"
	"log"
)

//const verb = true

func TestImage(t *testing.T) {		
	img, err := ImageFromFile("./gopher.png")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	log.Println(img.Getsize())	
}
