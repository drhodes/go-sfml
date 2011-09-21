# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=win-example
GOFILES=win-example.go

include $(GOROOT)/src/Make.cmd


cleaner:
	rm -f *_wrap* *_gc* .~* runme
	rm -f core 
	rm -f *.o *.$(GOOBJEXT) *.a *.so 
	rm -f *.$(O)

cleanest:
	make clean
	make cleaner
	rm -f *~

nuke:
	@echo "nop"
