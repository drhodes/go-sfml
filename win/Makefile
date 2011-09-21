# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=sfml/win

#GOFILES:= 

CGOFILES:= win.go

CGO_LDFLAGS:= -lcsfml-system -lcsfml-window

CLEANFILES+:= win

#cleaner:
#	@rm -f *~ cheat*

include $(GOROOT)/src/Make.pkg
