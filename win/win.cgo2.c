
typedef struct { char *p; int n; } _GoString_;
_GoString_ GoString(char *p);
char *CString(_GoString_);
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

void
_cgo_sfWindow_GetSettings(void *v)
{
	struct {
		sfWindow* p0;
		sfWindowSettings r;
	} *a = v;
	a->r = sfWindow_GetSettings(a->p0);
}

void
_cgo_sfVideoMode_GetDesktopMode(void *v)
{
	struct {
		sfVideoMode r;
	} *a = v;
	a->r = sfVideoMode_GetDesktopMode();
}

void
_cgo_sfContext_Create(void *v)
{
	struct {
		sfContext* r;
	} *a = v;
	a->r = sfContext_Create();
}

void
_cgo_sfWindow_GetHeight(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int r;
	} *a = v;
	a->r = sfWindow_GetHeight(a->p0);
}

void
_cgo_sfContext_SetActive(void *v)
{
	struct {
		sfContext* p0;
		sfBool p1;
	} *a = v;
	sfContext_SetActive(a->p0, a->p1);
}

void
_cgo_sfWindow_UseVerticalSync(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
	} *a = v;
	sfWindow_UseVerticalSync(a->p0, a->p1);
}

void
_cgo_sfInput_IsJoystickButtonDown(void *v)
{
	struct {
		sfInput* p0;
		unsigned int p1;
		unsigned int p2;
		sfBool r;
	} *a = v;
	a->r = sfInput_IsJoystickButtonDown(a->p0, a->p1, a->p2);
}

void
_cgo_sfWindow_SetPosition(void *v)
{
	struct {
		sfWindow* p0;
		int p1;
		int p2;
	} *a = v;
	sfWindow_SetPosition(a->p0, a->p1, a->p2);
}

void
_cgo_sfVideoMode_IsValid(void *v)
{
	struct {
		sfVideoMode p0;
		sfBool r;
	} *a = v;
	a->r = sfVideoMode_IsValid(a->p0);
}

void
_cgo_sfWindow_GetInput(void *v)
{
	struct {
		sfWindow* p0;
		sfInput* r;
	} *a = v;
	a->r = sfWindow_GetInput(a->p0);
}

void
_cgo_sfWindow_Show(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
	} *a = v;
	sfWindow_Show(a->p0, a->p1);
}

void
_cgo_sfWindow_GetFrameTime(void *v)
{
	struct {
		sfWindow* p0;
		float r;
	} *a = v;
	a->r = sfWindow_GetFrameTime(a->p0);
}

void
_cgo_sfWindow_Close(void *v)
{
	struct {
		sfWindow* p0;
	} *a = v;
	sfWindow_Close(a->p0);
}

void
_cgo_sfWindow_Destroy(void *v)
{
	struct {
		sfWindow* p0;
	} *a = v;
	sfWindow_Destroy(a->p0);
}

void
_cgo_sfInput_GetMouseY(void *v)
{
	struct {
		sfInput* p0;
		int r;
	} *a = v;
	a->r = sfInput_GetMouseY(a->p0);
}

void
_cgo_sfWindow_CreateFromHandle(void *v)
{
	struct {
		sfWindowHandle p0;
		sfWindowSettings p1;
		sfWindow* r;
	} *a = v;
	a->r = sfWindow_CreateFromHandle(a->p0, a->p1);
}

void
_cgo_sfWindow_SetCursorPosition(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int p1;
		unsigned int p2;
	} *a = v;
	sfWindow_SetCursorPosition(a->p0, a->p1, a->p2);
}

void
_cgo_sfWindow_Create(void *v)
{
	struct {
		sfVideoMode p0;
		char* p1;
		long unsigned int p2;
		sfWindowSettings p3;
		sfWindow* r;
	} *a = v;
	a->r = sfWindow_Create(a->p0, a->p1, a->p2, a->p3);
}

void
_cgo_sfWindow_SetIcon(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int p1;
		unsigned int p2;
		sfUint8* p3;
	} *a = v;
	sfWindow_SetIcon(a->p0, a->p1, a->p2, a->p3);
}

void
_cgo_sfWindow_SetJoystickThreshold(void *v)
{
	struct {
		sfWindow* p0;
		float p1;
	} *a = v;
	sfWindow_SetJoystickThreshold(a->p0, a->p1);
}

void
_cgo_sfWindow_SetActive(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
		sfBool r;
	} *a = v;
	a->r = sfWindow_SetActive(a->p0, a->p1);
}

void
_cgo_sfInput_IsKeyDown(void *v)
{
	struct {
		sfInput* p0;
		sfKeyCode p1;
		sfBool r;
	} *a = v;
	a->r = sfInput_IsKeyDown(a->p0, a->p1);
}

void
_cgo_sfWindow_GetEvent(void *v)
{
	struct {
		sfWindow* p0;
		sfEvent* p1;
		sfBool r;
	} *a = v;
	a->r = sfWindow_GetEvent(a->p0, a->p1);
}

void
_cgo_sfInput_GetJoystickAxis(void *v)
{
	struct {
		sfInput* p0;
		unsigned int p1;
		sfJoyAxis p2;
		float r;
	} *a = v;
	a->r = sfInput_GetJoystickAxis(a->p0, a->p1, a->p2);
}

void
_cgo_sfWindow_Display(void *v)
{
	struct {
		sfWindow* p0;
	} *a = v;
	sfWindow_Display(a->p0);
}

void
_cgo_sfWindow_SetFramerateLimit(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int p1;
	} *a = v;
	sfWindow_SetFramerateLimit(a->p0, a->p1);
}

void
_cgo_sfWindow_IsOpened(void *v)
{
	struct {
		sfWindow* p0;
		sfBool r;
	} *a = v;
	a->r = sfWindow_IsOpened(a->p0);
}

void
_cgo_sfWindow_EnableKeyRepeat(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
	} *a = v;
	sfWindow_EnableKeyRepeat(a->p0, a->p1);
}

void
_cgo_sfVideoMode_GetMode(void *v)
{
	struct {
		size_t p0;
		sfVideoMode r;
	} *a = v;
	a->r = sfVideoMode_GetMode(a->p0);
}

void
_cgo_sfContext_Destroy(void *v)
{
	struct {
		sfContext* p0;
	} *a = v;
	sfContext_Destroy(a->p0);
}

void
_cgo_sfInput_GetMouseX(void *v)
{
	struct {
		sfInput* p0;
		int r;
	} *a = v;
	a->r = sfInput_GetMouseX(a->p0);
}

void
_cgo_sfWindow_GetWidth(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int r;
	} *a = v;
	a->r = sfWindow_GetWidth(a->p0);
}

void
_cgo_sfWindow_ShowMouseCursor(void *v)
{
	struct {
		sfWindow* p0;
		sfBool p1;
	} *a = v;
	sfWindow_ShowMouseCursor(a->p0, a->p1);
}

void
_cgo_sfInput_IsMouseButtonDown(void *v)
{
	struct {
		sfInput* p0;
		sfMouseButton p1;
		sfBool r;
	} *a = v;
	a->r = sfInput_IsMouseButtonDown(a->p0, a->p1);
}

void
_cgo_sfWindow_SetSize(void *v)
{
	struct {
		sfWindow* p0;
		unsigned int p1;
		unsigned int p2;
	} *a = v;
	sfWindow_SetSize(a->p0, a->p1, a->p2);
}

void
_cgo_sfVideoMode_GetModesCount(void *v)
{
	struct {
		size_t r;
	} *a = v;
	a->r = sfVideoMode_GetModesCount();
}

