package sfml

import (
	"testing"
	"log"
)

const verb = false

func TestConstructors(t *testing.T) {		
	//clock := sys.NewClock()
}

func TestMode(t *testing.T) {		
	mode := GetDesktopMode()
	if verb {
		log.Println("Width", mode.Width())
		log.Println("Height", mode.Height())
		log.Println("BitsPerPixel", mode.BitsPerPixel())
		log.Println("Mode Valid", mode.IsValid())
	}
}

func TestGetModes(t *testing.T) {		
	modes := GetFullscreenModes()
	for _, m := range modes {
		if !m.IsValid() {
			t.Fail()
		}
		if verb {
			log.Println(m.Show())
		}
	}
}




func BenchmarkGetTime(b *testing.B) {
}

func BenchmarkClockDestroy(b *testing.B) {	
}

