package win

import (
	"testing"
	//"log"
)

func TestContextConstructor(t *testing.T) {		
	ctx := NewContext()
	ctx.SetActive(true)
	ctx.Destroy()
}

// func BenchmarkNewContext(b *testing.B) {
//     for i := 0; i < b.N; i++ {
// 		NewContext().Destroy()
// 	}
// }


