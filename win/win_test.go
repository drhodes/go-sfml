package win

import (
	"testing"
	"log"
	//"sfml/sys"
)

const verbose = false;

func TestConstructors(t *testing.T) {		
	//clock := sys.NewClock()
}

func TestMode(t *testing.T) {		
	mode := GetDesktopMode()
	if verbose {
		log.Println("Width", mode.Width())
		log.Println("Height", mode.Height())
		log.Println("BitsPerPixel", mode.BitsPerPixel())
	}
}





func BenchmarkGetTime(b *testing.B) {
}

func BenchmarkClockDestroy(b *testing.B) {	
}

