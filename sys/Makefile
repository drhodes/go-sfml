# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=sfml/sys

GOFILES:=debug.go

CGOFILES:=clock.go
CGO_LDFLAGS:=-lcsfml-system

CLEANFILES+=clock

include $(GOROOT)/src/Make.pkg
