
typedef struct { char *p; int n; } _GoString_;
_GoString_ GoString(char *p);
char *CString(_GoString_);
#include <SFML/Graphics.h>
#include <SFML/Config.h>



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

void
_cgo_sfColor_FromRGB(void *v)
{
	struct {
		sfUint8 p0;
		sfUint8 p1;
		sfUint8 p2;
		char __pad0[1];
		sfColor r;
	} *a = v;
	a->r = sfColor_FromRGB(a->p0, a->p1, a->p2);
}

void
_cgo_sfColor_Modulate(void *v)
{
	struct {
		sfColor p0;
		sfColor p1;
		sfColor r;
	} *a = v;
	a->r = sfColor_Modulate(a->p0, a->p1);
}

void
_cgo_sfColor_FromRGBA(void *v)
{
	struct {
		sfUint8 p0;
		sfUint8 p1;
		sfUint8 p2;
		sfUint8 p3;
		sfColor r;
	} *a = v;
	a->r = sfColor_FromRGBA(a->p0, a->p1, a->p2, a->p3);
}

void
_cgo_sfColor_Add(void *v)
{
	struct {
		sfColor p0;
		sfColor p1;
		sfColor r;
	} *a = v;
	a->r = sfColor_Add(a->p0, a->p1);
}

