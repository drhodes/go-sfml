# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=sfml/aud

#GOFILES:= 

CGOFILES:= aud.go

CGO_LDFLAGS:= -lcsfml-system -lcsfml-audio

CLEANFILES+:= aud

#cleaner:
#	@rm -f *~ cheat*

include $(GOROOT)/src/Make.pkg
