
#include "runtime.h"
#include "cgocall.h"

void ·_Cerrno(void*, int32);

void
·_Cfunc_GoString(int8 *p, String s)
{
	s = runtime·gostring((byte*)p);
	FLUSH(&s);
}

void
·_Cfunc_GoStringN(int8 *p, int32 l, String s)
{
	s = runtime·gostringn((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_GoBytes(int8 *p, int32 l, Slice s)
{
	s = runtime·gobytes((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_CString(String s, int8 *p)
{
	p = runtime·cmalloc(s.len+1);
	runtime·memmove((byte*)p, s.str, s.len);
	p[s.len] = 0;
	FLUSH(&p);
}

void _cgo_468de02d4734_Cfunc_sfClock_GetTime(void*);

void
·_Cfunc_sfClock_GetTime(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_468de02d4734_Cfunc_sfClock_GetTime, &p);
}

void _cgo_468de02d4734_Cfunc_sfClock_Reset(void*);

void
·_Cfunc_sfClock_Reset(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_468de02d4734_Cfunc_sfClock_Reset, &p);
}

void _cgo_468de02d4734_Cfunc_sfClock_Create(void*);

void
·_Cfunc_sfClock_Create(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_468de02d4734_Cfunc_sfClock_Create, &p);
}

void _cgo_468de02d4734_Cfunc_sfSleep(void*);

void
·_Cfunc_sfSleep(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_468de02d4734_Cfunc_sfSleep, &p);
}

void _cgo_468de02d4734_Cfunc_sfClock_Destroy(void*);

void
·_Cfunc_sfClock_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_468de02d4734_Cfunc_sfClock_Destroy, &p);
}

