
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

#pragma dynimport _cgo_sfSprite_SetScaleX _cgo_sfSprite_SetScaleX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetScaleX)(void*);

void
·_C_sfSprite_SetScaleX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_SetScaleX, &p);
}

#pragma dynimport _cgo_sfRenderWindow_Create _cgo_sfRenderWindow_Create "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_Create)(void*);

void
·_C_sfRenderWindow_Create(struct{uint8 x[36];}p)
{
	cgocall(_cgo_sfRenderWindow_Create, &p);
}

#pragma dynimport _cgo_sfShape_GetColor _cgo_sfShape_GetColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetColor)(void*);

void
·_C_sfShape_GetColor(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetColor, &p);
}

#pragma dynimport _cgo_sfRenderWindow_SetActive _cgo_sfRenderWindow_SetActive "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_SetActive)(void*);

void
·_C_sfRenderWindow_SetActive(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfRenderWindow_SetActive, &p);
}

#pragma dynimport _cgo_sfString_SetColor _cgo_sfString_SetColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetColor)(void*);

void
·_C_sfString_SetColor(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetColor, &p);
}

#pragma dynimport _cgo_sfShape_SetCenter _cgo_sfShape_SetCenter "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetCenter)(void*);

void
·_C_sfShape_SetCenter(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfShape_SetCenter, &p);
}

#pragma dynimport _cgo_sfSprite_GetCenterX _cgo_sfSprite_GetCenterX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetCenterX)(void*);

void
·_C_sfSprite_GetCenterX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetCenterX, &p);
}

#pragma dynimport _cgo_sfView_Zoom _cgo_sfView_Zoom "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_Zoom)(void*);

void
·_C_sfView_Zoom(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfView_Zoom, &p);
}

#pragma dynimport _cgo_sfRenderWindow_SetView _cgo_sfRenderWindow_SetView "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_SetView)(void*);

void
·_C_sfRenderWindow_SetView(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_SetView, &p);
}

#pragma dynimport _cgo_sfView_SetHalfSize _cgo_sfView_SetHalfSize "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_SetHalfSize)(void*);

void
·_C_sfView_SetHalfSize(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfView_SetHalfSize, &p);
}

#pragma dynimport _cgo_sfShape_SetX _cgo_sfShape_SetX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetX)(void*);

void
·_C_sfShape_SetX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_SetX, &p);
}

#pragma dynimport _cgo_sfRenderWindow_UseVerticalSync _cgo_sfRenderWindow_UseVerticalSync "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_UseVerticalSync)(void*);

void
·_C_sfRenderWindow_UseVerticalSync(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_UseVerticalSync, &p);
}

#pragma dynimport _cgo_sfImage_CreateMaskFromColor _cgo_sfImage_CreateMaskFromColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_CreateMaskFromColor)(void*);

void
·_C_sfImage_CreateMaskFromColor(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfImage_CreateMaskFromColor, &p);
}

#pragma dynimport _cgo_sfShape_CreateCircle _cgo_sfShape_CreateCircle "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_CreateCircle)(void*);

void
·_C_sfShape_CreateCircle(struct{uint8 x[28];}p)
{
	cgocall(_cgo_sfShape_CreateCircle, &p);
}

#pragma dynimport _cgo_sfView_GetRect _cgo_sfView_GetRect "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_GetRect)(void*);

void
·_C_sfView_GetRect(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfView_GetRect, &p);
}

#pragma dynimport _cgo_sfRenderWindow_SetCursorPosition _cgo_sfRenderWindow_SetCursorPosition "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_SetCursorPosition)(void*);

void
·_C_sfRenderWindow_SetCursorPosition(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfRenderWindow_SetCursorPosition, &p);
}

#pragma dynimport _cgo_sfColor_FromRGB _cgo_sfColor_FromRGB "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_FromRGB)(void*);

void
·_C_sfColor_FromRGB(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfColor_FromRGB, &p);
}

#pragma dynimport _cgo_sfShape_CreateRectangle _cgo_sfShape_CreateRectangle "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_CreateRectangle)(void*);

void
·_C_sfShape_CreateRectangle(struct{uint8 x[32];}p)
{
	cgocall(_cgo_sfShape_CreateRectangle, &p);
}

#pragma dynimport _cgo_sfFloatRect_Intersects _cgo_sfFloatRect_Intersects "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFloatRect_Intersects)(void*);

void
·_C_sfFloatRect_Intersects(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfFloatRect_Intersects, &p);
}

#pragma dynimport _cgo_sfShape_SetPointPosition _cgo_sfShape_SetPointPosition "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetPointPosition)(void*);

void
·_C_sfShape_SetPointPosition(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfShape_SetPointPosition, &p);
}

#pragma dynimport _cgo_sfView_CreateFromRect _cgo_sfView_CreateFromRect "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_CreateFromRect)(void*);

void
·_C_sfView_CreateFromRect(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfView_CreateFromRect, &p);
}

#pragma dynimport _cgo_sfUnwraP _cgo_sfUnwraP "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfUnwraP)(void*);

void
·_C_sfUnwraP(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfUnwraP, &p);
}

#pragma dynimport _cgo_sfRenderWindow_SetJoystickThreshold _cgo_sfRenderWindow_SetJoystickThreshold "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_SetJoystickThreshold)(void*);

void
·_C_sfRenderWindow_SetJoystickThreshold(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_SetJoystickThreshold, &p);
}

#pragma dynimport _cgo_sfRenderWindow_Destroy _cgo_sfRenderWindow_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_Destroy)(void*);

void
·_C_sfRenderWindow_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfRenderWindow_Destroy, &p);
}

#pragma dynimport _cgo_sfShape_GetScaleX _cgo_sfShape_GetScaleX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetScaleX)(void*);

void
·_C_sfShape_GetScaleX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetScaleX, &p);
}

#pragma dynimport _cgo_sfSprite_SetRotation _cgo_sfSprite_SetRotation "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetRotation)(void*);

void
·_C_sfSprite_SetRotation(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_SetRotation, &p);
}

#pragma dynimport _cgo_sfImage_CreateFromFile _cgo_sfImage_CreateFromFile "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_CreateFromFile)(void*);

void
·_C_sfImage_CreateFromFile(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfImage_CreateFromFile, &p);
}

#pragma dynimport _cgo_sfString_GetCenterX _cgo_sfString_GetCenterX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetCenterX)(void*);

void
·_C_sfString_GetCenterX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetCenterX, &p);
}

#pragma dynimport _cgo_sfSprite_FlipX _cgo_sfSprite_FlipX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_FlipX)(void*);

void
·_C_sfSprite_FlipX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_FlipX, &p);
}

#pragma dynimport _cgo_sfView_SetFromRect _cgo_sfView_SetFromRect "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_SetFromRect)(void*);

void
·_C_sfView_SetFromRect(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfView_SetFromRect, &p);
}

#pragma dynimport _cgo_sfImage_SetSmooth _cgo_sfImage_SetSmooth "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_SetSmooth)(void*);

void
·_C_sfImage_SetSmooth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfImage_SetSmooth, &p);
}

#pragma dynimport _cgo_sfPostFX_SetParameter4 _cgo_sfPostFX_SetParameter4 "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfPostFX_SetParameter4)(void*);

void
·_C_sfPostFX_SetParameter4(struct{uint8 x[24];}p)
{
	cgocall(_cgo_sfPostFX_SetParameter4, &p);
}

#pragma dynimport _cgo_sfSprite_GetPixel _cgo_sfSprite_GetPixel "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetPixel)(void*);

void
·_C_sfSprite_GetPixel(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfSprite_GetPixel, &p);
}

#pragma dynimport _cgo_sfView_GetHalfSizeX _cgo_sfView_GetHalfSizeX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_GetHalfSizeX)(void*);

void
·_C_sfView_GetHalfSizeX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfView_GetHalfSizeX, &p);
}

#pragma dynimport _cgo_sfString_SetScaleY _cgo_sfString_SetScaleY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetScaleY)(void*);

void
·_C_sfString_SetScaleY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetScaleY, &p);
}

#pragma dynimport _cgo_sfShape_GetY _cgo_sfShape_GetY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetY)(void*);

void
·_C_sfShape_GetY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetY, &p);
}

#pragma dynimport _cgo_sfShape_EnableFill _cgo_sfShape_EnableFill "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_EnableFill)(void*);

void
·_C_sfShape_EnableFill(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_EnableFill, &p);
}

#pragma dynimport _cgo_sfSprite_Scale _cgo_sfSprite_Scale "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_Scale)(void*);

void
·_C_sfSprite_Scale(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfSprite_Scale, &p);
}

#pragma dynimport _cgo_sfIntRect_Offset _cgo_sfIntRect_Offset "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfIntRect_Offset)(void*);

void
·_C_sfIntRect_Offset(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfIntRect_Offset, &p);
}

#pragma dynimport _cgo_sfSprite_GetImage _cgo_sfSprite_GetImage "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetImage)(void*);

void
·_C_sfSprite_GetImage(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetImage, &p);
}

#pragma dynimport _cgo_sfRenderWindow_DrawSprite _cgo_sfRenderWindow_DrawSprite "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_DrawSprite)(void*);

void
·_C_sfRenderWindow_DrawSprite(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_DrawSprite, &p);
}

#pragma dynimport _cgo_sfRenderWindow_GetSettings _cgo_sfRenderWindow_GetSettings "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_GetSettings)(void*);

void
·_C_sfRenderWindow_GetSettings(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfRenderWindow_GetSettings, &p);
}

#pragma dynimport _cgo_sfSprite_SetSubRect _cgo_sfSprite_SetSubRect "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetSubRect)(void*);

void
·_C_sfSprite_SetSubRect(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfSprite_SetSubRect, &p);
}

#pragma dynimport _cgo_sfRenderWindow_Show _cgo_sfRenderWindow_Show "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_Show)(void*);

void
·_C_sfRenderWindow_Show(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_Show, &p);
}

#pragma dynimport _cgo_sfRenderWindow_Capture _cgo_sfRenderWindow_Capture "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_Capture)(void*);

void
·_C_sfRenderWindow_Capture(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_Capture, &p);
}

#pragma dynimport _cgo_sfString_Create _cgo_sfString_Create "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Create)(void*);

void
·_C_sfString_Create(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfString_Create, &p);
}

#pragma dynimport _cgo_sfShape_GetCenterX _cgo_sfShape_GetCenterX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetCenterX)(void*);

void
·_C_sfShape_GetCenterX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetCenterX, &p);
}

#pragma dynimport _cgo_sfShape_Move _cgo_sfShape_Move "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_Move)(void*);

void
·_C_sfShape_Move(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfShape_Move, &p);
}

#pragma dynimport _cgo_sfString_SetUnicodeText _cgo_sfString_SetUnicodeText "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetUnicodeText)(void*);

void
·_C_sfString_SetUnicodeText(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetUnicodeText, &p);
}

#pragma dynimport _cgo_sfSprite_SetScale _cgo_sfSprite_SetScale "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetScale)(void*);

void
·_C_sfSprite_SetScale(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfSprite_SetScale, &p);
}

#pragma dynimport _cgo_sfSprite_GetColor _cgo_sfSprite_GetColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetColor)(void*);

void
·_C_sfSprite_GetColor(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetColor, &p);
}

#pragma dynimport _cgo_sfSprite_GetCenterY _cgo_sfSprite_GetCenterY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetCenterY)(void*);

void
·_C_sfSprite_GetCenterY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetCenterY, &p);
}

#pragma dynimport _cgo_sfImage_CreateFromMemory _cgo_sfImage_CreateFromMemory "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_CreateFromMemory)(void*);

void
·_C_sfImage_CreateFromMemory(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfImage_CreateFromMemory, &p);
}

#pragma dynimport _cgo_sfString_GetFont _cgo_sfString_GetFont "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetFont)(void*);

void
·_C_sfString_GetFont(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetFont, &p);
}

#pragma dynimport _cgo_sfString_GetUnicodeText _cgo_sfString_GetUnicodeText "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetUnicodeText)(void*);

void
·_C_sfString_GetUnicodeText(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetUnicodeText, &p);
}

#pragma dynimport _cgo_sfSprite_SetY _cgo_sfSprite_SetY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetY)(void*);

void
·_C_sfSprite_SetY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_SetY, &p);
}

#pragma dynimport _cgo_sfPostFX_CreateFromFile _cgo_sfPostFX_CreateFromFile "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfPostFX_CreateFromFile)(void*);

void
·_C_sfPostFX_CreateFromFile(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfPostFX_CreateFromFile, &p);
}

#pragma dynimport _cgo_sfString_TransformToGlobal _cgo_sfString_TransformToGlobal "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_TransformToGlobal)(void*);

void
·_C_sfString_TransformToGlobal(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfString_TransformToGlobal, &p);
}

#pragma dynimport _cgo_sfShape_GetRotation _cgo_sfShape_GetRotation "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetRotation)(void*);

void
·_C_sfShape_GetRotation(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetRotation, &p);
}

#pragma dynimport _cgo_sfImage_SetPixel _cgo_sfImage_SetPixel "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_SetPixel)(void*);

void
·_C_sfImage_SetPixel(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfImage_SetPixel, &p);
}

#pragma dynimport _cgo_sfSprite_Rotate _cgo_sfSprite_Rotate "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_Rotate)(void*);

void
·_C_sfSprite_Rotate(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_Rotate, &p);
}

#pragma dynimport _cgo_sfShape_AddPoint _cgo_sfShape_AddPoint "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_AddPoint)(void*);

void
·_C_sfShape_AddPoint(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfShape_AddPoint, &p);
}

#pragma dynimport _cgo_sfSprite_SetColor _cgo_sfSprite_SetColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetColor)(void*);

void
·_C_sfSprite_SetColor(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_SetColor, &p);
}

#pragma dynimport _cgo_sfRenderWindow_GetDefaultView _cgo_sfRenderWindow_GetDefaultView "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_GetDefaultView)(void*);

void
·_C_sfRenderWindow_GetDefaultView(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_GetDefaultView, &p);
}

#pragma dynimport _cgo_sfShape_Scale _cgo_sfShape_Scale "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_Scale)(void*);

void
·_C_sfShape_Scale(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfShape_Scale, &p);
}

#pragma dynimport _cgo_sfShape_GetNbPoints _cgo_sfShape_GetNbPoints "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetNbPoints)(void*);

void
·_C_sfShape_GetNbPoints(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetNbPoints, &p);
}

#pragma dynimport _cgo_sfShape_GetBlendMode _cgo_sfShape_GetBlendMode "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetBlendMode)(void*);

void
·_C_sfShape_GetBlendMode(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetBlendMode, &p);
}

#pragma dynimport _cgo_sfString_GetCenterY _cgo_sfString_GetCenterY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetCenterY)(void*);

void
·_C_sfString_GetCenterY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetCenterY, &p);
}

#pragma dynimport _cgo_sfSprite_FlipY _cgo_sfSprite_FlipY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_FlipY)(void*);

void
·_C_sfSprite_FlipY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_FlipY, &p);
}

#pragma dynimport _cgo_sfSprite_SetPosition _cgo_sfSprite_SetPosition "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetPosition)(void*);

void
·_C_sfSprite_SetPosition(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfSprite_SetPosition, &p);
}

#pragma dynimport _cgo_sfShape_GetPointOutlineColor _cgo_sfShape_GetPointOutlineColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetPointOutlineColor)(void*);

void
·_C_sfShape_GetPointOutlineColor(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfShape_GetPointOutlineColor, &p);
}

#pragma dynimport _cgo_sfSprite_GetRotation _cgo_sfSprite_GetRotation "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetRotation)(void*);

void
·_C_sfSprite_GetRotation(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetRotation, &p);
}

#pragma dynimport _cgo_sfPostFX_CanUsePostFX _cgo_sfPostFX_CanUsePostFX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfPostFX_CanUsePostFX)(void*);

void
·_C_sfPostFX_CanUsePostFX(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfPostFX_CanUsePostFX, &p);
}

#pragma dynimport _cgo_sfString_GetRect _cgo_sfString_GetRect "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetRect)(void*);

void
·_C_sfString_GetRect(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfString_GetRect, &p);
}

#pragma dynimport _cgo_sfString_SetScaleX _cgo_sfString_SetScaleX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetScaleX)(void*);

void
·_C_sfString_SetScaleX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetScaleX, &p);
}

#pragma dynimport _cgo_sfSprite_Resize _cgo_sfSprite_Resize "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_Resize)(void*);

void
·_C_sfSprite_Resize(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfSprite_Resize, &p);
}

#pragma dynimport _cgo_sfRenderWindow_GetWidth _cgo_sfRenderWindow_GetWidth "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_GetWidth)(void*);

void
·_C_sfRenderWindow_GetWidth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_GetWidth, &p);
}

#pragma dynimport _cgo_sfRenderWindow_Close _cgo_sfRenderWindow_Close "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_Close)(void*);

void
·_C_sfRenderWindow_Close(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfRenderWindow_Close, &p);
}

#pragma dynimport _cgo_sfShape_GetX _cgo_sfShape_GetX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetX)(void*);

void
·_C_sfShape_GetX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetX, &p);
}

#pragma dynimport _cgo_sfPostFX_SetTexture _cgo_sfPostFX_SetTexture "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfPostFX_SetTexture)(void*);

void
·_C_sfPostFX_SetTexture(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfPostFX_SetTexture, &p);
}

#pragma dynimport _cgo_sfColor_Modulate _cgo_sfColor_Modulate "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_Modulate)(void*);

void
·_C_sfColor_Modulate(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfColor_Modulate, &p);
}

#pragma dynimport _cgo_sfView_Move _cgo_sfView_Move "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_Move)(void*);

void
·_C_sfView_Move(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfView_Move, &p);
}

#pragma dynimport _cgo_sfRenderWindow_SetSize _cgo_sfRenderWindow_SetSize "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_SetSize)(void*);

void
·_C_sfRenderWindow_SetSize(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfRenderWindow_SetSize, &p);
}

#pragma dynimport _cgo_sfIntRect_Contains _cgo_sfIntRect_Contains "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfIntRect_Contains)(void*);

void
·_C_sfIntRect_Contains(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfIntRect_Contains, &p);
}

#pragma dynimport _cgo_sfFont_Create _cgo_sfFont_Create "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFont_Create)(void*);

void
·_C_sfFont_Create(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfFont_Create, &p);
}

#pragma dynimport _cgo_sfPostFX_SetParameter3 _cgo_sfPostFX_SetParameter3 "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfPostFX_SetParameter3)(void*);

void
·_C_sfPostFX_SetParameter3(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfPostFX_SetParameter3, &p);
}

#pragma dynimport _cgo_sfColor_FromRGB_P _cgo_sfColor_FromRGB_P "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_FromRGB_P)(void*);

void
·_C_sfColor_FromRGB_P(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfColor_FromRGB_P, &p);
}

#pragma dynimport _cgo_sfRenderWindow_GetHeight _cgo_sfRenderWindow_GetHeight "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_GetHeight)(void*);

void
·_C_sfRenderWindow_GetHeight(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_GetHeight, &p);
}

#pragma dynimport _cgo_sfShape_SetScale _cgo_sfShape_SetScale "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetScale)(void*);

void
·_C_sfShape_SetScale(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfShape_SetScale, &p);
}

#pragma dynimport _cgo_sfString_SetScale _cgo_sfString_SetScale "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetScale)(void*);

void
·_C_sfString_SetScale(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfString_SetScale, &p);
}

#pragma dynimport _cgo_sfShape_GetCenterY _cgo_sfShape_GetCenterY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetCenterY)(void*);

void
·_C_sfShape_GetCenterY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetCenterY, &p);
}

#pragma dynimport _cgo_sfFloatRect_Contains _cgo_sfFloatRect_Contains "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFloatRect_Contains)(void*);

void
·_C_sfFloatRect_Contains(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfFloatRect_Contains, &p);
}

#pragma dynimport _cgo_sfShape_Rotate _cgo_sfShape_Rotate "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_Rotate)(void*);

void
·_C_sfShape_Rotate(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_Rotate, &p);
}

#pragma dynimport _cgo_sfImage_Copy _cgo_sfImage_Copy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_Copy)(void*);

void
·_C_sfImage_Copy(struct{uint8 x[32];}p)
{
	cgocall(_cgo_sfImage_Copy, &p);
}

#pragma dynimport _cgo_sfPostFX_Destroy _cgo_sfPostFX_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfPostFX_Destroy)(void*);

void
·_C_sfPostFX_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfPostFX_Destroy, &p);
}

#pragma dynimport _cgo_sfView_Destroy _cgo_sfView_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_Destroy)(void*);

void
·_C_sfView_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfView_Destroy, &p);
}

#pragma dynimport _cgo_sfShape_SetScaleY _cgo_sfShape_SetScaleY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetScaleY)(void*);

void
·_C_sfShape_SetScaleY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_SetScaleY, &p);
}

#pragma dynimport _cgo_sfShape_SetBlendMode _cgo_sfShape_SetBlendMode "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetBlendMode)(void*);

void
·_C_sfShape_SetBlendMode(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_SetBlendMode, &p);
}

#pragma dynimport _cgo_sfString_Destroy _cgo_sfString_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Destroy)(void*);

void
·_C_sfString_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfString_Destroy, &p);
}

#pragma dynimport _cgo_sfSprite_SetCenter _cgo_sfSprite_SetCenter "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetCenter)(void*);

void
·_C_sfSprite_SetCenter(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfSprite_SetCenter, &p);
}

#pragma dynimport _cgo_sfRenderWindow_DrawShape _cgo_sfRenderWindow_DrawShape "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_DrawShape)(void*);

void
·_C_sfRenderWindow_DrawShape(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_DrawShape, &p);
}

#pragma dynimport _cgo_sfSprite_SetX _cgo_sfSprite_SetX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetX)(void*);

void
·_C_sfSprite_SetX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_SetX, &p);
}

#pragma dynimport _cgo_sfSprite_Create _cgo_sfSprite_Create "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_Create)(void*);

void
·_C_sfSprite_Create(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfSprite_Create, &p);
}

#pragma dynimport _cgo_sfShape_EnableOutline _cgo_sfShape_EnableOutline "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_EnableOutline)(void*);

void
·_C_sfShape_EnableOutline(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_EnableOutline, &p);
}

#pragma dynimport _cgo_sfColor_FromRGBA _cgo_sfColor_FromRGBA "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_FromRGBA)(void*);

void
·_C_sfColor_FromRGBA(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfColor_FromRGBA, &p);
}

#pragma dynimport _cgo_sfSprite_GetScaleY _cgo_sfSprite_GetScaleY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetScaleY)(void*);

void
·_C_sfSprite_GetScaleY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetScaleY, &p);
}

#pragma dynimport _cgo_sfRenderWindow_EnableKeyRepeat _cgo_sfRenderWindow_EnableKeyRepeat "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_EnableKeyRepeat)(void*);

void
·_C_sfRenderWindow_EnableKeyRepeat(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_EnableKeyRepeat, &p);
}

#pragma dynimport _cgo_sfImage_Bind _cgo_sfImage_Bind "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_Bind)(void*);

void
·_C_sfImage_Bind(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfImage_Bind, &p);
}

#pragma dynimport _cgo_sfImage_GetHeight _cgo_sfImage_GetHeight "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_GetHeight)(void*);

void
·_C_sfImage_GetHeight(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfImage_GetHeight, &p);
}

#pragma dynimport _cgo_sfSprite_Move _cgo_sfSprite_Move "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_Move)(void*);

void
·_C_sfSprite_Move(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfSprite_Move, &p);
}

#pragma dynimport _cgo_sfShape_SetPosition _cgo_sfShape_SetPosition "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetPosition)(void*);

void
·_C_sfShape_SetPosition(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfShape_SetPosition, &p);
}

#pragma dynimport _cgo_sfShape_GetOutlineWidth _cgo_sfShape_GetOutlineWidth "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetOutlineWidth)(void*);

void
·_C_sfShape_GetOutlineWidth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetOutlineWidth, &p);
}

#pragma dynimport _cgo_sfImage_Destroy _cgo_sfImage_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_Destroy)(void*);

void
·_C_sfImage_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfImage_Destroy, &p);
}

#pragma dynimport _cgo_sfShape_Destroy _cgo_sfShape_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_Destroy)(void*);

void
·_C_sfShape_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfShape_Destroy, &p);
}

#pragma dynimport _cgo_sfString_GetScaleY _cgo_sfString_GetScaleY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetScaleY)(void*);

void
·_C_sfString_GetScaleY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetScaleY, &p);
}

#pragma dynimport _cgo_sfImage_CreateFromColor _cgo_sfImage_CreateFromColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_CreateFromColor)(void*);

void
·_C_sfImage_CreateFromColor(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfImage_CreateFromColor, &p);
}

#pragma dynimport _cgo_sfRenderWindow_DrawString _cgo_sfRenderWindow_DrawString "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_DrawString)(void*);

void
·_C_sfRenderWindow_DrawString(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_DrawString, &p);
}

#pragma dynimport _cgo_sfShape_Create _cgo_sfShape_Create "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_Create)(void*);

void
·_C_sfShape_Create(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfShape_Create, &p);
}

#pragma dynimport _cgo_sfString_SetSize _cgo_sfString_SetSize "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetSize)(void*);

void
·_C_sfString_SetSize(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetSize, &p);
}

#pragma dynimport _cgo_sfShape_SetPointColor _cgo_sfShape_SetPointColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetPointColor)(void*);

void
·_C_sfShape_SetPointColor(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfShape_SetPointColor, &p);
}

#pragma dynimport _cgo_sfRenderWindow_IsOpened _cgo_sfRenderWindow_IsOpened "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_IsOpened)(void*);

void
·_C_sfRenderWindow_IsOpened(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_IsOpened, &p);
}

#pragma dynimport _cgo_sfSprite_SetImage _cgo_sfSprite_SetImage "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetImage)(void*);

void
·_C_sfSprite_SetImage(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_SetImage, &p);
}

#pragma dynimport _cgo_sfImage_IsSmooth _cgo_sfImage_IsSmooth "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_IsSmooth)(void*);

void
·_C_sfImage_IsSmooth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfImage_IsSmooth, &p);
}

#pragma dynimport _cgo_sfSprite_GetY _cgo_sfSprite_GetY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetY)(void*);

void
·_C_sfSprite_GetY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetY, &p);
}

#pragma dynimport _cgo_sfImage_CopyScreen _cgo_sfImage_CopyScreen "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_CopyScreen)(void*);

void
·_C_sfImage_CopyScreen(struct{uint8 x[28];}p)
{
	cgocall(_cgo_sfImage_CopyScreen, &p);
}

#pragma dynimport _cgo_sfView_SetCenter _cgo_sfView_SetCenter "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_SetCenter)(void*);

void
·_C_sfView_SetCenter(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfView_SetCenter, &p);
}

#pragma dynimport _cgo_sfString_SetPosition _cgo_sfString_SetPosition "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetPosition)(void*);

void
·_C_sfString_SetPosition(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfString_SetPosition, &p);
}

#pragma dynimport _cgo_sfPostFX_SetParameter2 _cgo_sfPostFX_SetParameter2 "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfPostFX_SetParameter2)(void*);

void
·_C_sfPostFX_SetParameter2(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfPostFX_SetParameter2, &p);
}

#pragma dynimport _cgo_sfFont_CreateFromFile _cgo_sfFont_CreateFromFile "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFont_CreateFromFile)(void*);

void
·_C_sfFont_CreateFromFile(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfFont_CreateFromFile, &p);
}

#pragma dynimport _cgo_sfView_GetCenterX _cgo_sfView_GetCenterX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_GetCenterX)(void*);

void
·_C_sfView_GetCenterX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfView_GetCenterX, &p);
}

#pragma dynimport _cgo_sfString_SetCenter _cgo_sfString_SetCenter "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetCenter)(void*);

void
·_C_sfString_SetCenter(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfString_SetCenter, &p);
}

#pragma dynimport _cgo_sfString_GetY _cgo_sfString_GetY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetY)(void*);

void
·_C_sfString_GetY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetY, &p);
}

#pragma dynimport _cgo_sfRenderWindow_Clear _cgo_sfRenderWindow_Clear "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_Clear)(void*);

void
·_C_sfRenderWindow_Clear(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_Clear, &p);
}

#pragma dynimport _cgo_sfRenderWindow_PreserveOpenGLStates _cgo_sfRenderWindow_PreserveOpenGLStates "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_PreserveOpenGLStates)(void*);

void
·_C_sfRenderWindow_PreserveOpenGLStates(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_PreserveOpenGLStates, &p);
}

#pragma dynimport _cgo_sfShape_SetOutlineWidth _cgo_sfShape_SetOutlineWidth "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetOutlineWidth)(void*);

void
·_C_sfShape_SetOutlineWidth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_SetOutlineWidth, &p);
}

#pragma dynimport _cgo_sfRenderWindow_DrawPostFX _cgo_sfRenderWindow_DrawPostFX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_DrawPostFX)(void*);

void
·_C_sfRenderWindow_DrawPostFX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_DrawPostFX, &p);
}

#pragma dynimport _cgo_sfRenderWindow_ShowMouseCursor _cgo_sfRenderWindow_ShowMouseCursor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_ShowMouseCursor)(void*);

void
·_C_sfRenderWindow_ShowMouseCursor(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_ShowMouseCursor, &p);
}

#pragma dynimport _cgo_sfString_SetY _cgo_sfString_SetY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetY)(void*);

void
·_C_sfString_SetY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetY, &p);
}

#pragma dynimport _cgo_sfImage_GetWidth _cgo_sfImage_GetWidth "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_GetWidth)(void*);

void
·_C_sfImage_GetWidth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfImage_GetWidth, &p);
}

#pragma dynimport _cgo_sfShape_SetRotation _cgo_sfShape_SetRotation "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetRotation)(void*);

void
·_C_sfShape_SetRotation(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_SetRotation, &p);
}

#pragma dynimport _cgo_sfShape_SetScaleX _cgo_sfShape_SetScaleX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetScaleX)(void*);

void
·_C_sfShape_SetScaleX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_SetScaleX, &p);
}

#pragma dynimport _cgo_sfString_GetBlendMode _cgo_sfString_GetBlendMode "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetBlendMode)(void*);

void
·_C_sfString_GetBlendMode(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetBlendMode, &p);
}

#pragma dynimport _cgo_sfIntRect_Intersects _cgo_sfIntRect_Intersects "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfIntRect_Intersects)(void*);

void
·_C_sfIntRect_Intersects(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfIntRect_Intersects, &p);
}

#pragma dynimport _cgo_sfString_GetColor _cgo_sfString_GetColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetColor)(void*);

void
·_C_sfString_GetColor(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetColor, &p);
}

#pragma dynimport _cgo_sfImage_Create _cgo_sfImage_Create "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_Create)(void*);

void
·_C_sfImage_Create(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfImage_Create, &p);
}

#pragma dynimport _cgo_sfImage_CreateFromPixels _cgo_sfImage_CreateFromPixels "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_CreateFromPixels)(void*);

void
·_C_sfImage_CreateFromPixels(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfImage_CreateFromPixels, &p);
}

#pragma dynimport _cgo_sfFont_GetDefaultFont _cgo_sfFont_GetDefaultFont "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFont_GetDefaultFont)(void*);

void
·_C_sfFont_GetDefaultFont(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfFont_GetDefaultFont, &p);
}

#pragma dynimport _cgo_sfRenderWindow_SetPosition _cgo_sfRenderWindow_SetPosition "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_SetPosition)(void*);

void
·_C_sfRenderWindow_SetPosition(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfRenderWindow_SetPosition, &p);
}

#pragma dynimport _cgo_sfSprite_SetScaleY _cgo_sfSprite_SetScaleY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetScaleY)(void*);

void
·_C_sfSprite_SetScaleY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_SetScaleY, &p);
}

#pragma dynimport _cgo_sfString_SetStyle _cgo_sfString_SetStyle "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetStyle)(void*);

void
·_C_sfString_SetStyle(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetStyle, &p);
}

#pragma dynimport _cgo_sfShape_SetPointOutlineColor _cgo_sfShape_SetPointOutlineColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetPointOutlineColor)(void*);

void
·_C_sfShape_SetPointOutlineColor(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfShape_SetPointOutlineColor, &p);
}

#pragma dynimport _cgo_sfView_Create _cgo_sfView_Create "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_Create)(void*);

void
·_C_sfView_Create(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfView_Create, &p);
}

#pragma dynimport _cgo_sfImage_GetPixel _cgo_sfImage_GetPixel "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_GetPixel)(void*);

void
·_C_sfImage_GetPixel(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfImage_GetPixel, &p);
}

#pragma dynimport _cgo_sfShape_CreateLine _cgo_sfShape_CreateLine "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_CreateLine)(void*);

void
·_C_sfShape_CreateLine(struct{uint8 x[36];}p)
{
	cgocall(_cgo_sfShape_CreateLine, &p);
}

#pragma dynimport _cgo_sfSprite_GetScaleX _cgo_sfSprite_GetScaleX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetScaleX)(void*);

void
·_C_sfSprite_GetScaleX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetScaleX, &p);
}

#pragma dynimport _cgo_sfFont_Destroy _cgo_sfFont_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFont_Destroy)(void*);

void
·_C_sfFont_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfFont_Destroy, &p);
}

#pragma dynimport _cgo_sfShape_SetY _cgo_sfShape_SetY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetY)(void*);

void
·_C_sfShape_SetY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_SetY, &p);
}

#pragma dynimport _cgo_sfString_Rotate _cgo_sfString_Rotate "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Rotate)(void*);

void
·_C_sfString_Rotate(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_Rotate, &p);
}

#pragma dynimport _cgo_sfSprite_GetBlendMode _cgo_sfSprite_GetBlendMode "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetBlendMode)(void*);

void
·_C_sfSprite_GetBlendMode(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetBlendMode, &p);
}

#pragma dynimport _cgo_sfShape_GetPointColor _cgo_sfShape_GetPointColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetPointColor)(void*);

void
·_C_sfShape_GetPointColor(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfShape_GetPointColor, &p);
}

#pragma dynimport _cgo_sfString_SetText _cgo_sfString_SetText "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetText)(void*);

void
·_C_sfString_SetText(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetText, &p);
}

#pragma dynimport _cgo_sfString_GetSize _cgo_sfString_GetSize "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetSize)(void*);

void
·_C_sfString_GetSize(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetSize, &p);
}

#pragma dynimport _cgo_sfFloatRect_Offset _cgo_sfFloatRect_Offset "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFloatRect_Offset)(void*);

void
·_C_sfFloatRect_Offset(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfFloatRect_Offset, &p);
}

#pragma dynimport _cgo_sfString_GetScaleX _cgo_sfString_GetScaleX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetScaleX)(void*);

void
·_C_sfString_GetScaleX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetScaleX, &p);
}

#pragma dynimport _cgo_sfShape_GetScaleY _cgo_sfShape_GetScaleY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_GetScaleY)(void*);

void
·_C_sfShape_GetScaleY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_GetScaleY, &p);
}

#pragma dynimport _cgo_sfString_Scale _cgo_sfString_Scale "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Scale)(void*);

void
·_C_sfString_Scale(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfString_Scale, &p);
}

#pragma dynimport _cgo_sfSprite_GetHeight _cgo_sfSprite_GetHeight "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetHeight)(void*);

void
·_C_sfSprite_GetHeight(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetHeight, &p);
}

#pragma dynimport _cgo_sfColor_Add _cgo_sfColor_Add "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_Add)(void*);

void
·_C_sfColor_Add(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfColor_Add, &p);
}

#pragma dynimport _cgo_sfString_SetBlendMode _cgo_sfString_SetBlendMode "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetBlendMode)(void*);

void
·_C_sfString_SetBlendMode(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetBlendMode, &p);
}

#pragma dynimport _cgo_sfImage_SaveToFile _cgo_sfImage_SaveToFile "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_SaveToFile)(void*);

void
·_C_sfImage_SaveToFile(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfImage_SaveToFile, &p);
}

#pragma dynimport _cgo_sfSprite_GetX _cgo_sfSprite_GetX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetX)(void*);

void
·_C_sfSprite_GetX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetX, &p);
}

#pragma dynimport _cgo_sfRenderWindow_Display _cgo_sfRenderWindow_Display "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_Display)(void*);

void
·_C_sfRenderWindow_Display(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfRenderWindow_Display, &p);
}

#pragma dynimport _cgo_sfString_TransformToLocal _cgo_sfString_TransformToLocal "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_TransformToLocal)(void*);

void
·_C_sfString_TransformToLocal(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfString_TransformToLocal, &p);
}

#pragma dynimport _cgo_sfRenderWindow_SetFramerateLimit _cgo_sfRenderWindow_SetFramerateLimit "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfRenderWindow_SetFramerateLimit)(void*);

void
·_C_sfRenderWindow_SetFramerateLimit(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfRenderWindow_SetFramerateLimit, &p);
}

#pragma dynimport _cgo_sfView_GetCenterY _cgo_sfView_GetCenterY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_GetCenterY)(void*);

void
·_C_sfView_GetCenterY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfView_GetCenterY, &p);
}

#pragma dynimport _cgo_sfString_GetX _cgo_sfString_GetX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetX)(void*);

void
·_C_sfString_GetX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetX, &p);
}

#pragma dynimport _cgo_sfView_GetHalfSizeY _cgo_sfView_GetHalfSizeY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfView_GetHalfSizeY)(void*);

void
·_C_sfView_GetHalfSizeY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfView_GetHalfSizeY, &p);
}

#pragma dynimport _cgo_sfSprite_SetBlendMode _cgo_sfSprite_SetBlendMode "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_SetBlendMode)(void*);

void
·_C_sfSprite_SetBlendMode(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_SetBlendMode, &p);
}

#pragma dynimport _cgo_sfSprite_Destroy _cgo_sfSprite_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_Destroy)(void*);

void
·_C_sfSprite_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfSprite_Destroy, &p);
}

#pragma dynimport _cgo_sfString_SetX _cgo_sfString_SetX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetX)(void*);

void
·_C_sfString_SetX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetX, &p);
}

#pragma dynimport _cgo_sfShape_SetColor _cgo_sfShape_SetColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfShape_SetColor)(void*);

void
·_C_sfShape_SetColor(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfShape_SetColor, &p);
}

#pragma dynimport _cgo_sfString_GetRotation _cgo_sfString_GetRotation "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetRotation)(void*);

void
·_C_sfString_GetRotation(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetRotation, &p);
}

#pragma dynimport _cgo_sfPostFX_SetParameter1 _cgo_sfPostFX_SetParameter1 "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfPostFX_SetParameter1)(void*);

void
·_C_sfPostFX_SetParameter1(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfPostFX_SetParameter1, &p);
}

#pragma dynimport _cgo_sfString_Move _cgo_sfString_Move "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Move)(void*);

void
·_C_sfString_Move(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfString_Move, &p);
}

#pragma dynimport _cgo_sfSprite_GetWidth _cgo_sfSprite_GetWidth "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetWidth)(void*);

void
·_C_sfSprite_GetWidth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfSprite_GetWidth, &p);
}

#pragma dynimport _cgo_sfSprite_GetSubRect _cgo_sfSprite_GetSubRect "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfSprite_GetSubRect)(void*);

void
·_C_sfSprite_GetSubRect(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfSprite_GetSubRect, &p);
}

#pragma dynimport _cgo_sfString_SetFont _cgo_sfString_SetFont "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetFont)(void*);

void
·_C_sfString_SetFont(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetFont, &p);
}

#pragma dynimport _cgo_sfString_SetRotation _cgo_sfString_SetRotation "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetRotation)(void*);

void
·_C_sfString_SetRotation(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetRotation, &p);
}

