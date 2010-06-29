
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

#pragma dynimport _cgo_sfClock_GetTime _cgo_sfClock_GetTime "/usr/custom/goroot/pkg/linux_386/sfml/sys.so"
void (*_cgo_sfClock_GetTime)(void*);

void
·_C_sfClock_GetTime(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfClock_GetTime, &p);
}

#pragma dynimport _cgo_sfClock_Reset _cgo_sfClock_Reset "/usr/custom/goroot/pkg/linux_386/sfml/sys.so"
void (*_cgo_sfClock_Reset)(void*);

void
·_C_sfClock_Reset(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfClock_Reset, &p);
}

#pragma dynimport _cgo_sfClock_Create _cgo_sfClock_Create "/usr/custom/goroot/pkg/linux_386/sfml/sys.so"
void (*_cgo_sfClock_Create)(void*);

void
·_C_sfClock_Create(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfClock_Create, &p);
}

#pragma dynimport _cgo_sfSleep _cgo_sfSleep "/usr/custom/goroot/pkg/linux_386/sfml/sys.so"
void (*_cgo_sfSleep)(void*);

void
·_C_sfSleep(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfSleep, &p);
}

#pragma dynimport _cgo_sfClock_Destroy _cgo_sfClock_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/sys.so"
void (*_cgo_sfClock_Destroy)(void*);

void
·_C_sfClock_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfClock_Destroy, &p);
}

