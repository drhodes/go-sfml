package sfml

import (
	"log"
	"testing"
)

func TestMode(t *testing.T) {
	mode := DesktopMode()
	log.Println("Width", mode.Width())
	log.Println("Height", mode.Height())
	log.Println("BitsPerPixel", mode.BitsPerPixel())
	log.Println("Mode Valid", mode.IsValid())
}

func TestModes(t *testing.T) {
	modes := FullscreenModes()
	for _, m := range modes {
		if !m.IsValid() {
			t.Fatalf("FullscreenModes returned an invalid mode: %s", m)
		}
	}
}
