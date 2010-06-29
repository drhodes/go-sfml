
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

#pragma dynimport _cgo_sfWindow_GetSettings _cgo_sfWindow_GetSettings "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_GetSettings)(void*);

void
·_C_sfWindow_GetSettings(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfWindow_GetSettings, &p);
}

#pragma dynimport _cgo_sfVideoMode_GetDesktopMode _cgo_sfVideoMode_GetDesktopMode "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfVideoMode_GetDesktopMode)(void*);

void
·_C_sfVideoMode_GetDesktopMode(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfVideoMode_GetDesktopMode, &p);
}

#pragma dynimport _cgo_sfContext_Create _cgo_sfContext_Create "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfContext_Create)(void*);

void
·_C_sfContext_Create(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfContext_Create, &p);
}

#pragma dynimport _cgo_sfWindow_GetHeight _cgo_sfWindow_GetHeight "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_GetHeight)(void*);

void
·_C_sfWindow_GetHeight(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_GetHeight, &p);
}

#pragma dynimport _cgo_sfContext_SetActive _cgo_sfContext_SetActive "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfContext_SetActive)(void*);

void
·_C_sfContext_SetActive(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfContext_SetActive, &p);
}

#pragma dynimport _cgo_sfWindow_UseVerticalSync _cgo_sfWindow_UseVerticalSync "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_UseVerticalSync)(void*);

void
·_C_sfWindow_UseVerticalSync(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_UseVerticalSync, &p);
}

#pragma dynimport _cgo_sfInput_IsJoystickButtonDown _cgo_sfInput_IsJoystickButtonDown "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfInput_IsJoystickButtonDown)(void*);

void
·_C_sfInput_IsJoystickButtonDown(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfInput_IsJoystickButtonDown, &p);
}

#pragma dynimport _cgo_sfWindow_SetPosition _cgo_sfWindow_SetPosition "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_SetPosition)(void*);

void
·_C_sfWindow_SetPosition(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfWindow_SetPosition, &p);
}

#pragma dynimport _cgo_sfVideoMode_IsValid _cgo_sfVideoMode_IsValid "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfVideoMode_IsValid)(void*);

void
·_C_sfVideoMode_IsValid(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfVideoMode_IsValid, &p);
}

#pragma dynimport _cgo_sfWindow_GetInput _cgo_sfWindow_GetInput "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_GetInput)(void*);

void
·_C_sfWindow_GetInput(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_GetInput, &p);
}

#pragma dynimport _cgo_sfWindow_Show _cgo_sfWindow_Show "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_Show)(void*);

void
·_C_sfWindow_Show(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_Show, &p);
}

#pragma dynimport _cgo_sfWindow_GetFrameTime _cgo_sfWindow_GetFrameTime "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_GetFrameTime)(void*);

void
·_C_sfWindow_GetFrameTime(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_GetFrameTime, &p);
}

#pragma dynimport _cgo_sfWindow_Close _cgo_sfWindow_Close "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_Close)(void*);

void
·_C_sfWindow_Close(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfWindow_Close, &p);
}

#pragma dynimport _cgo_sfWindow_Destroy _cgo_sfWindow_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_Destroy)(void*);

void
·_C_sfWindow_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfWindow_Destroy, &p);
}

#pragma dynimport _cgo_sfInput_GetMouseY _cgo_sfInput_GetMouseY "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfInput_GetMouseY)(void*);

void
·_C_sfInput_GetMouseY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfInput_GetMouseY, &p);
}

#pragma dynimport _cgo_sfWindow_CreateFromHandle _cgo_sfWindow_CreateFromHandle "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_CreateFromHandle)(void*);

void
·_C_sfWindow_CreateFromHandle(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfWindow_CreateFromHandle, &p);
}

#pragma dynimport _cgo_sfWindow_SetCursorPosition _cgo_sfWindow_SetCursorPosition "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_SetCursorPosition)(void*);

void
·_C_sfWindow_SetCursorPosition(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfWindow_SetCursorPosition, &p);
}

#pragma dynimport _cgo_sfWindow_Create _cgo_sfWindow_Create "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_Create)(void*);

void
·_C_sfWindow_Create(struct{uint8 x[36];}p)
{
	cgocall(_cgo_sfWindow_Create, &p);
}

#pragma dynimport _cgo_sfWindow_SetIcon _cgo_sfWindow_SetIcon "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_SetIcon)(void*);

void
·_C_sfWindow_SetIcon(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfWindow_SetIcon, &p);
}

#pragma dynimport _cgo_sfWindow_SetJoystickThreshold _cgo_sfWindow_SetJoystickThreshold "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_SetJoystickThreshold)(void*);

void
·_C_sfWindow_SetJoystickThreshold(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_SetJoystickThreshold, &p);
}

#pragma dynimport _cgo_sfWindow_SetActive _cgo_sfWindow_SetActive "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_SetActive)(void*);

void
·_C_sfWindow_SetActive(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfWindow_SetActive, &p);
}

#pragma dynimport _cgo_sfInput_IsKeyDown _cgo_sfInput_IsKeyDown "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfInput_IsKeyDown)(void*);

void
·_C_sfInput_IsKeyDown(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfInput_IsKeyDown, &p);
}

#pragma dynimport _cgo_sfWindow_GetEvent _cgo_sfWindow_GetEvent "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_GetEvent)(void*);

void
·_C_sfWindow_GetEvent(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfWindow_GetEvent, &p);
}

#pragma dynimport _cgo_sfInput_GetJoystickAxis _cgo_sfInput_GetJoystickAxis "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfInput_GetJoystickAxis)(void*);

void
·_C_sfInput_GetJoystickAxis(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfInput_GetJoystickAxis, &p);
}

#pragma dynimport _cgo_sfWindow_Display _cgo_sfWindow_Display "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_Display)(void*);

void
·_C_sfWindow_Display(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfWindow_Display, &p);
}

#pragma dynimport _cgo_sfWindow_SetFramerateLimit _cgo_sfWindow_SetFramerateLimit "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_SetFramerateLimit)(void*);

void
·_C_sfWindow_SetFramerateLimit(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_SetFramerateLimit, &p);
}

#pragma dynimport _cgo_sfWindow_IsOpened _cgo_sfWindow_IsOpened "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_IsOpened)(void*);

void
·_C_sfWindow_IsOpened(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_IsOpened, &p);
}

#pragma dynimport _cgo_sfWindow_EnableKeyRepeat _cgo_sfWindow_EnableKeyRepeat "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_EnableKeyRepeat)(void*);

void
·_C_sfWindow_EnableKeyRepeat(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_EnableKeyRepeat, &p);
}

#pragma dynimport _cgo_sfVideoMode_GetMode _cgo_sfVideoMode_GetMode "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfVideoMode_GetMode)(void*);

void
·_C_sfVideoMode_GetMode(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfVideoMode_GetMode, &p);
}

#pragma dynimport _cgo_sfContext_Destroy _cgo_sfContext_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfContext_Destroy)(void*);

void
·_C_sfContext_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfContext_Destroy, &p);
}

#pragma dynimport _cgo_sfInput_GetMouseX _cgo_sfInput_GetMouseX "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfInput_GetMouseX)(void*);

void
·_C_sfInput_GetMouseX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfInput_GetMouseX, &p);
}

#pragma dynimport _cgo_sfWindow_GetWidth _cgo_sfWindow_GetWidth "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_GetWidth)(void*);

void
·_C_sfWindow_GetWidth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_GetWidth, &p);
}

#pragma dynimport _cgo_sfWindow_ShowMouseCursor _cgo_sfWindow_ShowMouseCursor "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_ShowMouseCursor)(void*);

void
·_C_sfWindow_ShowMouseCursor(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfWindow_ShowMouseCursor, &p);
}

#pragma dynimport _cgo_sfInput_IsMouseButtonDown _cgo_sfInput_IsMouseButtonDown "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfInput_IsMouseButtonDown)(void*);

void
·_C_sfInput_IsMouseButtonDown(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfInput_IsMouseButtonDown, &p);
}

#pragma dynimport _cgo_sfWindow_SetSize _cgo_sfWindow_SetSize "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfWindow_SetSize)(void*);

void
·_C_sfWindow_SetSize(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfWindow_SetSize, &p);
}

#pragma dynimport _cgo_sfVideoMode_GetModesCount _cgo_sfVideoMode_GetModesCount "/usr/custom/goroot/pkg/linux_386/sfml/win.so"
void (*_cgo_sfVideoMode_GetModesCount)(void*);

void
·_C_sfVideoMode_GetModesCount(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfVideoMode_GetModesCount, &p);
}

