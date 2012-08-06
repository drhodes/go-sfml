package sfml

import (
	"log"
	"testing"
)

func TestMode(t *testing.T) {
	mode := GetDesktopMode()
	log.Println("Width", mode.Width())
	log.Println("Height", mode.Height())
	log.Println("BitsPerPixel", mode.BitsPerPixel())
	log.Println("Mode Valid", mode.IsValid())
}

func TestGetModes(t *testing.T) {
	modes := GetFullscreenModes()
	for _, m := range modes {
		if !m.IsValid() {
			t.Fatalf("GetFullscreenModes returned an invalid mode: %s", m)
		}
	}
}
