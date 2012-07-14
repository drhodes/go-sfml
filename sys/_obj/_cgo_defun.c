
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

void _cgo_782fb323d796_Cfunc_sfClock_copy(void*);

void
·_Cfunc_sfClock_copy(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_782fb323d796_Cfunc_sfClock_copy, &p);
}

void _cgo_782fb323d796_Cfunc_sfClock_create(void*);

void
·_Cfunc_sfClock_create(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_782fb323d796_Cfunc_sfClock_create, &p);
}

void _cgo_782fb323d796_Cfunc_sfClock_destroy(void*);

void
·_Cfunc_sfClock_destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_782fb323d796_Cfunc_sfClock_destroy, &p);
}

