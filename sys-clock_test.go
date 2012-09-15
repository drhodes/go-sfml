// Copyright 2012.  All rights reserved. 
// Use of this source code is governed by a BSD-style  
// license that can be found in the LICENSE file.

package sfml


import (
	"testing"
)

func TestClockConstructors(t *testing.T) {
	clock := NewClock()
	clock.Restart()
}

func TestAsMilliseconds(t *testing.T) {
	sec := Seconds(123)
	if sec.AsMilliseconds() != 123000 {
		t.Fatalf("AsMilliseconds don't return the expected result (expected %d, got %d)", 123000, sec.AsMilliseconds())
	}
}

func TestAsMicroseconds(t *testing.T) {
	sec := Seconds(123)
	if sec.AsMicroseconds() != 123000000 {
		t.Fatalf("AsMicroSeconds don't return the expected result (expected %d, got %d)", 123000000, sec.AsMicroseconds())
	}
}

func TestAsSeconds(t *testing.T) {
	sec := Seconds(123)
	if sec.AsSeconds() != 123 {
		t.Fatalf("AsSeconds don't return the expected result (expected %d, got %d)", 123, sec.AsSeconds())
	}
}
