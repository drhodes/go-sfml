# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=sfml/cgo

CGOFILES:= util.go

CGO_LDFLAGS:= -lcsfml-system -lcsfml-graphics

CLEANFILES+:= util

#cleaner:
#	@rm -f *~ cheat*

include $(GOROOT)/src/Make.pkg
