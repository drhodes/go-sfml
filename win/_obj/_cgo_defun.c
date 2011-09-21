
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

void _cgo_e901e7711f2a_Cfunc_sfWindow_GetSettings(void*);

void
·_Cfunc_sfWindow_GetSettings(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_GetSettings, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfVideoMode_GetDesktopMode(void*);

void
·_Cfunc_sfVideoMode_GetDesktopMode(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfVideoMode_GetDesktopMode, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfContext_Create(void*);

void
·_Cfunc_sfContext_Create(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfContext_Create, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_GetHeight(void*);

void
·_Cfunc_sfWindow_GetHeight(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_GetHeight, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfContext_SetActive(void*);

void
·_Cfunc_sfContext_SetActive(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfContext_SetActive, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_UseVerticalSync(void*);

void
·_Cfunc_sfWindow_UseVerticalSync(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_UseVerticalSync, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfInput_IsJoystickButtonDown(void*);

void
·_Cfunc_sfInput_IsJoystickButtonDown(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfInput_IsJoystickButtonDown, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_SetPosition(void*);

void
·_Cfunc_sfWindow_SetPosition(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_SetPosition, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfVideoMode_IsValid(void*);

void
·_Cfunc_sfVideoMode_IsValid(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfVideoMode_IsValid, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_GetInput(void*);

void
·_Cfunc_sfWindow_GetInput(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_GetInput, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_Show(void*);

void
·_Cfunc_sfWindow_Show(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_Show, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_GetFrameTime(void*);

void
·_Cfunc_sfWindow_GetFrameTime(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_GetFrameTime, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_Close(void*);

void
·_Cfunc_sfWindow_Close(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_Close, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_Destroy(void*);

void
·_Cfunc_sfWindow_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_Destroy, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfInput_GetMouseY(void*);

void
·_Cfunc_sfInput_GetMouseY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfInput_GetMouseY, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_CreateFromHandle(void*);

void
·_Cfunc_sfWindow_CreateFromHandle(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_CreateFromHandle, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_SetCursorPosition(void*);

void
·_Cfunc_sfWindow_SetCursorPosition(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_SetCursorPosition, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_Create(void*);

void
·_Cfunc_sfWindow_Create(struct{uint8 x[36];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_Create, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_SetIcon(void*);

void
·_Cfunc_sfWindow_SetIcon(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_SetIcon, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_SetJoystickThreshold(void*);

void
·_Cfunc_sfWindow_SetJoystickThreshold(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_SetJoystickThreshold, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_SetActive(void*);

void
·_Cfunc_sfWindow_SetActive(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_SetActive, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfInput_IsKeyDown(void*);

void
·_Cfunc_sfInput_IsKeyDown(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfInput_IsKeyDown, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_GetEvent(void*);

void
·_Cfunc_sfWindow_GetEvent(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_GetEvent, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfInput_GetJoystickAxis(void*);

void
·_Cfunc_sfInput_GetJoystickAxis(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfInput_GetJoystickAxis, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_Display(void*);

void
·_Cfunc_sfWindow_Display(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_Display, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_SetFramerateLimit(void*);

void
·_Cfunc_sfWindow_SetFramerateLimit(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_SetFramerateLimit, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_IsOpened(void*);

void
·_Cfunc_sfWindow_IsOpened(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_IsOpened, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_EnableKeyRepeat(void*);

void
·_Cfunc_sfWindow_EnableKeyRepeat(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_EnableKeyRepeat, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfVideoMode_GetMode(void*);

void
·_Cfunc_sfVideoMode_GetMode(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfVideoMode_GetMode, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfContext_Destroy(void*);

void
·_Cfunc_sfContext_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfContext_Destroy, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfInput_GetMouseX(void*);

void
·_Cfunc_sfInput_GetMouseX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfInput_GetMouseX, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_GetWidth(void*);

void
·_Cfunc_sfWindow_GetWidth(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_GetWidth, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_ShowMouseCursor(void*);

void
·_Cfunc_sfWindow_ShowMouseCursor(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_ShowMouseCursor, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfInput_IsMouseButtonDown(void*);

void
·_Cfunc_sfInput_IsMouseButtonDown(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfInput_IsMouseButtonDown, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfWindow_SetSize(void*);

void
·_Cfunc_sfWindow_SetSize(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfWindow_SetSize, &p);
}

void _cgo_e901e7711f2a_Cfunc_sfVideoMode_GetModesCount(void*);

void
·_Cfunc_sfVideoMode_GetModesCount(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_e901e7711f2a_Cfunc_sfVideoMode_GetModesCount, &p);
}

