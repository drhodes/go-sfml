// Copyright 2012.  All rights reserved. 
// Use of this source code is governed by a BSD-style  
// license that can be found in the LICENSE file.

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
