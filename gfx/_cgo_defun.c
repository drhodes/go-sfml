
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

#pragma dynimport _cgo_sfColor_FromRGB _cgo_sfColor_FromRGB "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_FromRGB)(void*);

void
·_C_sfColor_FromRGB(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfColor_FromRGB, &p);
}

#pragma dynimport _cgo_sfColor_Modulate _cgo_sfColor_Modulate "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_Modulate)(void*);

void
·_C_sfColor_Modulate(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfColor_Modulate, &p);
}

#pragma dynimport _cgo_sfColor_FromRGBA _cgo_sfColor_FromRGBA "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_FromRGBA)(void*);

void
·_C_sfColor_FromRGBA(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfColor_FromRGBA, &p);
}

#pragma dynimport _cgo_sfColor_Add _cgo_sfColor_Add "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_Add)(void*);

void
·_C_sfColor_Add(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfColor_Add, &p);
}

