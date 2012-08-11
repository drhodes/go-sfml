package sfml

import (
	"testing"
)

func TestContextConstructor(t *testing.T) {
	ctx, err := NewContext()
	if err != nil {
		t.Fatal(err)
	}
	ctx.SetActive(true)
	ctx.Destroy()
}
