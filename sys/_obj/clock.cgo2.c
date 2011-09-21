#line 3 "clock.go"
#include <SFML/System.h>



// Usual nonsense: if x and y are not equal, the type will be invalid
// (have a negative array count) and an inscrutable error will come
// out of the compiler and hopefully mention "name".
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

// Check at compile time that the sizes we use match our expectations.
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)

#include <errno.h>
#include <string.h>

void
_cgo_468de02d4734_Cfunc_sfClock_GetTime(void *v)
{
	struct {
		sfClock* p0;
		float r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfClock_GetTime(a->p0);
}

void
_cgo_468de02d4734_Cfunc_sfClock_Reset(void *v)
{
	struct {
		sfClock* p0;
	} __attribute__((__packed__)) *a = v;
	sfClock_Reset(a->p0);
}

void
_cgo_468de02d4734_Cfunc_sfClock_Create(void *v)
{
	struct {
		const sfClock* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (const sfClock*) sfClock_Create();
}

void
_cgo_468de02d4734_Cfunc_sfSleep(void *v)
{
	struct {
		float p0;
	} __attribute__((__packed__)) *a = v;
	sfSleep(a->p0);
}

void
_cgo_468de02d4734_Cfunc_sfClock_Destroy(void *v)
{
	struct {
		sfClock* p0;
	} __attribute__((__packed__)) *a = v;
	sfClock_Destroy(a->p0);
}

