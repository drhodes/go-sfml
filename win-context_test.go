package sfml

import (
	"testing"
)

func TestContextConstructor(t *testing.T) {
	ctx := NewContext()
	ctx.SetActive(true)
	ctx.Destroy()
}
