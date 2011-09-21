#line 3 "win.go"
#include <SFML/Window/Context.h>
#include <SFML/Window/Input.h>
#include <SFML/Window/VideoMode.h>
#include <SFML/Window/Window.h>



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
_cgo_e901e7711f2a_Cfunc_sfWindow_GetSettings(void *v)
{
	struct {
		sfWindow* p0;
		sfWindowSettings r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfWindow_GetSettings(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfVideoMode_GetDesktopMode(void *v)
{
	struct {
		sfVideoMode r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfVideoMode_GetDesktopMode();
}

void
_cgo_e901e7711f2a_Cfunc_sfContext_Create(void *v)
{
	struct {
		const sfContext* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (const sfContext*) sfContext_Create();
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_GetHeight(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfWindow_GetHeight(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfContext_SetActive(void *v)
{
	struct {
		sfContext* p0;
		sfBool p1;
	} __attribute__((__packed__)) *a = v;
	sfContext_SetActive(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_UseVerticalSync(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
	} __attribute__((__packed__)) *a = v;
	sfWindow_UseVerticalSync(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfInput_IsJoystickButtonDown(void *v)
{
	struct {
		sfInput* p0;
		unsigned int p1;
		unsigned int p2;
		sfBool r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfInput_IsJoystickButtonDown(a->p0, a->p1, a->p2);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_SetPosition(void *v)
{
	struct {
		sfWindow* p0;
		int p1;
		int p2;
	} __attribute__((__packed__)) *a = v;
	sfWindow_SetPosition(a->p0, a->p1, a->p2);
}

void
_cgo_e901e7711f2a_Cfunc_sfVideoMode_IsValid(void *v)
{
	struct {
		sfVideoMode p0;
		sfBool r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfVideoMode_IsValid(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_GetInput(void *v)
{
	struct {
		sfWindow* p0;
		const sfInput* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (const sfInput*) sfWindow_GetInput(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_Show(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
	} __attribute__((__packed__)) *a = v;
	sfWindow_Show(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_GetFrameTime(void *v)
{
	struct {
		sfWindow* p0;
		float r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfWindow_GetFrameTime(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_Close(void *v)
{
	struct {
		sfWindow* p0;
	} __attribute__((__packed__)) *a = v;
	sfWindow_Close(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_Destroy(void *v)
{
	struct {
		sfWindow* p0;
	} __attribute__((__packed__)) *a = v;
	sfWindow_Destroy(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfInput_GetMouseY(void *v)
{
	struct {
		sfInput* p0;
		int r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfInput_GetMouseY(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_CreateFromHandle(void *v)
{
	struct {
		sfWindowHandle p0;
		sfWindowSettings p1;
		const sfWindow* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (const sfWindow*) sfWindow_CreateFromHandle(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_SetCursorPosition(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int p1;
		unsigned int p2;
	} __attribute__((__packed__)) *a = v;
	sfWindow_SetCursorPosition(a->p0, a->p1, a->p2);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_Create(void *v)
{
	struct {
		sfVideoMode p0;
		char* p1;
		long unsigned int p2;
		sfWindowSettings p3;
		const sfWindow* r;
	} __attribute__((__packed__)) *a = v;
	a->r = (const sfWindow*) sfWindow_Create(a->p0, a->p1, a->p2, a->p3);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_SetIcon(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int p1;
		unsigned int p2;
		sfUint8* p3;
	} __attribute__((__packed__)) *a = v;
	sfWindow_SetIcon(a->p0, a->p1, a->p2, a->p3);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_SetJoystickThreshold(void *v)
{
	struct {
		sfWindow* p0;
		float p1;
	} __attribute__((__packed__)) *a = v;
	sfWindow_SetJoystickThreshold(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_SetActive(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
		sfBool r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfWindow_SetActive(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfInput_IsKeyDown(void *v)
{
	struct {
		sfInput* p0;
		sfKeyCode p1;
		sfBool r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfInput_IsKeyDown(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_GetEvent(void *v)
{
	struct {
		sfWindow* p0;
		sfEvent* p1;
		sfBool r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfWindow_GetEvent(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfInput_GetJoystickAxis(void *v)
{
	struct {
		sfInput* p0;
		unsigned int p1;
		sfJoyAxis p2;
		float r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfInput_GetJoystickAxis(a->p0, a->p1, a->p2);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_Display(void *v)
{
	struct {
		sfWindow* p0;
	} __attribute__((__packed__)) *a = v;
	sfWindow_Display(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_SetFramerateLimit(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int p1;
	} __attribute__((__packed__)) *a = v;
	sfWindow_SetFramerateLimit(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_IsOpened(void *v)
{
	struct {
		sfWindow* p0;
		sfBool r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfWindow_IsOpened(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_EnableKeyRepeat(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
	} __attribute__((__packed__)) *a = v;
	sfWindow_EnableKeyRepeat(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfVideoMode_GetMode(void *v)
{
	struct {
		size_t p0;
		sfVideoMode r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfVideoMode_GetMode(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfContext_Destroy(void *v)
{
	struct {
		sfContext* p0;
	} __attribute__((__packed__)) *a = v;
	sfContext_Destroy(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfInput_GetMouseX(void *v)
{
	struct {
		sfInput* p0;
		int r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfInput_GetMouseX(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_GetWidth(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfWindow_GetWidth(a->p0);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_ShowMouseCursor(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
	} __attribute__((__packed__)) *a = v;
	sfWindow_ShowMouseCursor(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfInput_IsMouseButtonDown(void *v)
{
	struct {
		sfInput* p0;
		sfMouseButton p1;
		sfBool r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfInput_IsMouseButtonDown(a->p0, a->p1);
}

void
_cgo_e901e7711f2a_Cfunc_sfWindow_SetSize(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int p1;
		unsigned int p2;
	} __attribute__((__packed__)) *a = v;
	sfWindow_SetSize(a->p0, a->p1, a->p2);
}

void
_cgo_e901e7711f2a_Cfunc_sfVideoMode_GetModesCount(void *v)
{
	struct {
		size_t r;
	} __attribute__((__packed__)) *a = v;
	a->r = sfVideoMode_GetModesCount();
}

