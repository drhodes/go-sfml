
#include "runtime.h"
#include "cgocall.h"

#pragma dynimport initcgo initcgo "/usr/custom/goroot/pkg/linux_386/libcgo.so"
#pragma dynimport libcgo_thread_start libcgo_thread_start "/usr/custom/goroot/pkg/linux_386/libcgo.so"
#pragma dynimport libcgo_set_scheduler libcgo_set_scheduler "/usr/custom/goroot/pkg/linux_386/libcgo.so"
#pragma dynimport _cgo_malloc _cgo_malloc "/usr/custom/goroot/pkg/linux_386/libcgo.so"
#pragma dynimport _cgo_free _cgo_free "/usr/custom/goroot/pkg/linux_386/libcgo.so"

void
·_C_GoString(int8 *p, String s)
{
	s = gostring((byte*)p);
	FLUSH(&s);
}

void
·_C_CString(String s, int8 *p)
{
	p = cmalloc(s.len+1);
	mcpy((byte*)p, s.str, s.len);
	p[s.len] = 0;
	FLUSH(&p);
}

