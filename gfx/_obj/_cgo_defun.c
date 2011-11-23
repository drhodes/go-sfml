
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

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetScaleX(void*);

void
·_Cfunc_sfSprite_SetScaleX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetScaleX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_Create(void*);

void
·_Cfunc_sfRenderWindow_Create(struct{uint8 x[36];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_Create, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetColor(void*);

void
·_Cfunc_sfShape_GetColor(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetActive(void*);

void
·_Cfunc_sfRenderWindow_SetActive(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetActive, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetColor(void*);

void
·_Cfunc_sfString_SetColor(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetCenter(void*);

void
·_Cfunc_sfShape_SetCenter(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetCenter, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetCenterX(void*);

void
·_Cfunc_sfSprite_GetCenterX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetCenterX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_Zoom(void*);

void
·_Cfunc_sfView_Zoom(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_Zoom, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetView(void*);

void
·_Cfunc_sfRenderWindow_SetView(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetView, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_SetHalfSize(void*);

void
·_Cfunc_sfView_SetHalfSize(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_SetHalfSize, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetX(void*);

void
·_Cfunc_sfShape_SetX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_UseVerticalSync(void*);

void
·_Cfunc_sfRenderWindow_UseVerticalSync(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_UseVerticalSync, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_CreateMaskFromColor(void*);

void
·_Cfunc_sfImage_CreateMaskFromColor(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_CreateMaskFromColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_CreateCircle(void*);

void
·_Cfunc_sfShape_CreateCircle(struct{uint8 x[28];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_CreateCircle, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_GetRect(void*);

void
·_Cfunc_sfView_GetRect(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_GetRect, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetCursorPosition(void*);

void
·_Cfunc_sfRenderWindow_SetCursorPosition(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetCursorPosition, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfColor_FromRGB(void*);

void
·_Cfunc_sfColor_FromRGB(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfColor_FromRGB, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_CreateRectangle(void*);

void
·_Cfunc_sfShape_CreateRectangle(struct{uint8 x[32];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_CreateRectangle, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfFloatRect_Intersects(void*);

void
·_Cfunc_sfFloatRect_Intersects(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfFloatRect_Intersects, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetPointPosition(void*);

void
·_Cfunc_sfShape_SetPointPosition(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetPointPosition, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_CreateFromRect(void*);

void
·_Cfunc_sfView_CreateFromRect(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_CreateFromRect, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfUnwraP(void*);

void
·_Cfunc_sfUnwraP(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfUnwraP, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetJoystickThreshold(void*);

void
·_Cfunc_sfRenderWindow_SetJoystickThreshold(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetJoystickThreshold, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_Destroy(void*);

void
·_Cfunc_sfRenderWindow_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_Destroy, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetScaleX(void*);

void
·_Cfunc_sfShape_GetScaleX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetScaleX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetRotation(void*);

void
·_Cfunc_sfSprite_SetRotation(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetRotation, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_CreateFromFile(void*);

void
·_Cfunc_sfImage_CreateFromFile(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_CreateFromFile, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetCenterX(void*);

void
·_Cfunc_sfString_GetCenterX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetCenterX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_FlipX(void*);

void
·_Cfunc_sfSprite_FlipX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_FlipX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_SetFromRect(void*);

void
·_Cfunc_sfView_SetFromRect(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_SetFromRect, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_SetSmooth(void*);

void
·_Cfunc_sfImage_SetSmooth(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_SetSmooth, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfPostFX_SetParameter4(void*);

void
·_Cfunc_sfPostFX_SetParameter4(struct{uint8 x[24];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfPostFX_SetParameter4, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetPixel(void*);

void
·_Cfunc_sfSprite_GetPixel(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetPixel, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_GetHalfSizeX(void*);

void
·_Cfunc_sfView_GetHalfSizeX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_GetHalfSizeX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetScaleY(void*);

void
·_Cfunc_sfString_SetScaleY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetScaleY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetY(void*);

void
·_Cfunc_sfShape_GetY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_EnableFill(void*);

void
·_Cfunc_sfShape_EnableFill(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_EnableFill, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_Scale(void*);

void
·_Cfunc_sfSprite_Scale(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_Scale, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfIntRect_Offset(void*);

void
·_Cfunc_sfIntRect_Offset(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfIntRect_Offset, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetImage(void*);

void
·_Cfunc_sfSprite_GetImage(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetImage, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_DrawSprite(void*);

void
·_Cfunc_sfRenderWindow_DrawSprite(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_DrawSprite, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_GetSettings(void*);

void
·_Cfunc_sfRenderWindow_GetSettings(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_GetSettings, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetSubRect(void*);

void
·_Cfunc_sfSprite_SetSubRect(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetSubRect, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_Show(void*);

void
·_Cfunc_sfRenderWindow_Show(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_Show, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_Capture(void*);

void
·_Cfunc_sfRenderWindow_Capture(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_Capture, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_Create(void*);

void
·_Cfunc_sfString_Create(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_Create, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetCenterX(void*);

void
·_Cfunc_sfShape_GetCenterX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetCenterX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_Move(void*);

void
·_Cfunc_sfShape_Move(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_Move, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetUnicodeText(void*);

void
·_Cfunc_sfString_SetUnicodeText(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetUnicodeText, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetScale(void*);

void
·_Cfunc_sfSprite_SetScale(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetScale, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetColor(void*);

void
·_Cfunc_sfSprite_GetColor(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetCenterY(void*);

void
·_Cfunc_sfSprite_GetCenterY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetCenterY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_CreateFromMemory(void*);

void
·_Cfunc_sfImage_CreateFromMemory(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_CreateFromMemory, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetFont(void*);

void
·_Cfunc_sfString_GetFont(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetFont, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetUnicodeText(void*);

void
·_Cfunc_sfString_GetUnicodeText(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetUnicodeText, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetY(void*);

void
·_Cfunc_sfSprite_SetY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfPostFX_CreateFromFile(void*);

void
·_Cfunc_sfPostFX_CreateFromFile(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfPostFX_CreateFromFile, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_TransformToGlobal(void*);

void
·_Cfunc_sfString_TransformToGlobal(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_TransformToGlobal, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetRotation(void*);

void
·_Cfunc_sfShape_GetRotation(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetRotation, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_SetPixel(void*);

void
·_Cfunc_sfImage_SetPixel(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_SetPixel, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_Rotate(void*);

void
·_Cfunc_sfSprite_Rotate(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_Rotate, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_AddPoint(void*);

void
·_Cfunc_sfShape_AddPoint(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_AddPoint, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetColor(void*);

void
·_Cfunc_sfSprite_SetColor(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_GetDefaultView(void*);

void
·_Cfunc_sfRenderWindow_GetDefaultView(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_GetDefaultView, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_Scale(void*);

void
·_Cfunc_sfShape_Scale(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_Scale, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetNbPoints(void*);

void
·_Cfunc_sfShape_GetNbPoints(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetNbPoints, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetBlendMode(void*);

void
·_Cfunc_sfShape_GetBlendMode(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetBlendMode, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetCenterY(void*);

void
·_Cfunc_sfString_GetCenterY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetCenterY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_FlipY(void*);

void
·_Cfunc_sfSprite_FlipY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_FlipY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetPosition(void*);

void
·_Cfunc_sfSprite_SetPosition(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetPosition, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetPointOutlineColor(void*);

void
·_Cfunc_sfShape_GetPointOutlineColor(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetPointOutlineColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetRotation(void*);

void
·_Cfunc_sfSprite_GetRotation(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetRotation, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfPostFX_CanUsePostFX(void*);

void
·_Cfunc_sfPostFX_CanUsePostFX(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfPostFX_CanUsePostFX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetRect(void*);

void
·_Cfunc_sfString_GetRect(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetRect, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetScaleX(void*);

void
·_Cfunc_sfString_SetScaleX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetScaleX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_Resize(void*);

void
·_Cfunc_sfSprite_Resize(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_Resize, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_GetWidth(void*);

void
·_Cfunc_sfRenderWindow_GetWidth(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_GetWidth, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_Close(void*);

void
·_Cfunc_sfRenderWindow_Close(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_Close, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetX(void*);

void
·_Cfunc_sfShape_GetX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfPostFX_SetTexture(void*);

void
·_Cfunc_sfPostFX_SetTexture(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfPostFX_SetTexture, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfColor_Modulate(void*);

void
·_Cfunc_sfColor_Modulate(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfColor_Modulate, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_Move(void*);

void
·_Cfunc_sfView_Move(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_Move, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetSize(void*);

void
·_Cfunc_sfRenderWindow_SetSize(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetSize, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfIntRect_Contains(void*);

void
·_Cfunc_sfIntRect_Contains(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfIntRect_Contains, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfFont_Create(void*);

void
·_Cfunc_sfFont_Create(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfFont_Create, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfPostFX_SetParameter3(void*);

void
·_Cfunc_sfPostFX_SetParameter3(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfPostFX_SetParameter3, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfColor_FromRGB_P(void*);

void
·_Cfunc_sfColor_FromRGB_P(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfColor_FromRGB_P, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_GetHeight(void*);

void
·_Cfunc_sfRenderWindow_GetHeight(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_GetHeight, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetScale(void*);

void
·_Cfunc_sfShape_SetScale(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetScale, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetScale(void*);

void
·_Cfunc_sfString_SetScale(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetScale, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetCenterY(void*);

void
·_Cfunc_sfShape_GetCenterY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetCenterY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfFloatRect_Contains(void*);

void
·_Cfunc_sfFloatRect_Contains(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfFloatRect_Contains, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_Rotate(void*);

void
·_Cfunc_sfShape_Rotate(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_Rotate, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_Copy(void*);

void
·_Cfunc_sfImage_Copy(struct{uint8 x[32];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_Copy, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfPostFX_Destroy(void*);

void
·_Cfunc_sfPostFX_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfPostFX_Destroy, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_Destroy(void*);

void
·_Cfunc_sfView_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_Destroy, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetScaleY(void*);

void
·_Cfunc_sfShape_SetScaleY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetScaleY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetBlendMode(void*);

void
·_Cfunc_sfShape_SetBlendMode(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetBlendMode, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_Destroy(void*);

void
·_Cfunc_sfString_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_Destroy, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetCenter(void*);

void
·_Cfunc_sfSprite_SetCenter(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetCenter, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_DrawShape(void*);

void
·_Cfunc_sfRenderWindow_DrawShape(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_DrawShape, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetX(void*);

void
·_Cfunc_sfSprite_SetX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_Create(void*);

void
·_Cfunc_sfSprite_Create(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_Create, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_EnableOutline(void*);

void
·_Cfunc_sfShape_EnableOutline(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_EnableOutline, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfColor_FromRGBA(void*);

void
·_Cfunc_sfColor_FromRGBA(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfColor_FromRGBA, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetScaleY(void*);

void
·_Cfunc_sfSprite_GetScaleY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetScaleY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_EnableKeyRepeat(void*);

void
·_Cfunc_sfRenderWindow_EnableKeyRepeat(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_EnableKeyRepeat, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_Bind(void*);

void
·_Cfunc_sfImage_Bind(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_Bind, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_GetHeight(void*);

void
·_Cfunc_sfImage_GetHeight(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_GetHeight, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_Move(void*);

void
·_Cfunc_sfSprite_Move(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_Move, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetPosition(void*);

void
·_Cfunc_sfShape_SetPosition(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetPosition, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetOutlineWidth(void*);

void
·_Cfunc_sfShape_GetOutlineWidth(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetOutlineWidth, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_Destroy(void*);

void
·_Cfunc_sfImage_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_Destroy, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_Destroy(void*);

void
·_Cfunc_sfShape_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_Destroy, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetScaleY(void*);

void
·_Cfunc_sfString_GetScaleY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetScaleY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_CreateFromColor(void*);

void
·_Cfunc_sfImage_CreateFromColor(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_CreateFromColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_DrawString(void*);

void
·_Cfunc_sfRenderWindow_DrawString(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_DrawString, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_Create(void*);

void
·_Cfunc_sfShape_Create(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_Create, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetSize(void*);

void
·_Cfunc_sfString_SetSize(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetSize, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetPointColor(void*);

void
·_Cfunc_sfShape_SetPointColor(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetPointColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_IsOpened(void*);

void
·_Cfunc_sfRenderWindow_IsOpened(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_IsOpened, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetImage(void*);

void
·_Cfunc_sfSprite_SetImage(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetImage, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_IsSmooth(void*);

void
·_Cfunc_sfImage_IsSmooth(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_IsSmooth, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetY(void*);

void
·_Cfunc_sfSprite_GetY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_CopyScreen(void*);

void
·_Cfunc_sfImage_CopyScreen(struct{uint8 x[28];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_CopyScreen, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_SetCenter(void*);

void
·_Cfunc_sfView_SetCenter(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_SetCenter, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetPosition(void*);

void
·_Cfunc_sfString_SetPosition(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetPosition, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfPostFX_SetParameter2(void*);

void
·_Cfunc_sfPostFX_SetParameter2(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfPostFX_SetParameter2, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfFont_CreateFromFile(void*);

void
·_Cfunc_sfFont_CreateFromFile(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfFont_CreateFromFile, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_GetCenterX(void*);

void
·_Cfunc_sfView_GetCenterX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_GetCenterX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetCenter(void*);

void
·_Cfunc_sfString_SetCenter(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetCenter, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetY(void*);

void
·_Cfunc_sfString_GetY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_Clear(void*);

void
·_Cfunc_sfRenderWindow_Clear(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_Clear, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_PreserveOpenGLStates(void*);

void
·_Cfunc_sfRenderWindow_PreserveOpenGLStates(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_PreserveOpenGLStates, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetOutlineWidth(void*);

void
·_Cfunc_sfShape_SetOutlineWidth(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetOutlineWidth, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_DrawPostFX(void*);

void
·_Cfunc_sfRenderWindow_DrawPostFX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_DrawPostFX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_ShowMouseCursor(void*);

void
·_Cfunc_sfRenderWindow_ShowMouseCursor(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_ShowMouseCursor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetY(void*);

void
·_Cfunc_sfString_SetY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_GetWidth(void*);

void
·_Cfunc_sfImage_GetWidth(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_GetWidth, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetRotation(void*);

void
·_Cfunc_sfShape_SetRotation(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetRotation, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetScaleX(void*);

void
·_Cfunc_sfShape_SetScaleX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetScaleX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetBlendMode(void*);

void
·_Cfunc_sfString_GetBlendMode(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetBlendMode, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfIntRect_Intersects(void*);

void
·_Cfunc_sfIntRect_Intersects(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfIntRect_Intersects, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetColor(void*);

void
·_Cfunc_sfString_GetColor(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_Create(void*);

void
·_Cfunc_sfImage_Create(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_Create, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_CreateFromPixels(void*);

void
·_Cfunc_sfImage_CreateFromPixels(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_CreateFromPixels, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfFont_GetDefaultFont(void*);

void
·_Cfunc_sfFont_GetDefaultFont(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfFont_GetDefaultFont, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetPosition(void*);

void
·_Cfunc_sfRenderWindow_SetPosition(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetPosition, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetScaleY(void*);

void
·_Cfunc_sfSprite_SetScaleY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetScaleY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetStyle(void*);

void
·_Cfunc_sfString_SetStyle(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetStyle, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetPointOutlineColor(void*);

void
·_Cfunc_sfShape_SetPointOutlineColor(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetPointOutlineColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_Create(void*);

void
·_Cfunc_sfView_Create(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_Create, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_GetPixel(void*);

void
·_Cfunc_sfImage_GetPixel(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_GetPixel, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_CreateLine(void*);

void
·_Cfunc_sfShape_CreateLine(struct{uint8 x[36];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_CreateLine, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetScaleX(void*);

void
·_Cfunc_sfSprite_GetScaleX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetScaleX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfFont_Destroy(void*);

void
·_Cfunc_sfFont_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfFont_Destroy, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetY(void*);

void
·_Cfunc_sfShape_SetY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_Rotate(void*);

void
·_Cfunc_sfString_Rotate(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_Rotate, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetBlendMode(void*);

void
·_Cfunc_sfSprite_GetBlendMode(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetBlendMode, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetPointColor(void*);

void
·_Cfunc_sfShape_GetPointColor(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetPointColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetText(void*);

void
·_Cfunc_sfString_SetText(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetText, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetSize(void*);

void
·_Cfunc_sfString_GetSize(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetSize, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfFloatRect_Offset(void*);

void
·_Cfunc_sfFloatRect_Offset(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfFloatRect_Offset, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetScaleX(void*);

void
·_Cfunc_sfString_GetScaleX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetScaleX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_GetScaleY(void*);

void
·_Cfunc_sfShape_GetScaleY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_GetScaleY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_Scale(void*);

void
·_Cfunc_sfString_Scale(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_Scale, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetHeight(void*);

void
·_Cfunc_sfSprite_GetHeight(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetHeight, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfColor_Add(void*);

void
·_Cfunc_sfColor_Add(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfColor_Add, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetBlendMode(void*);

void
·_Cfunc_sfString_SetBlendMode(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetBlendMode, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfImage_SaveToFile(void*);

void
·_Cfunc_sfImage_SaveToFile(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfImage_SaveToFile, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetX(void*);

void
·_Cfunc_sfSprite_GetX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_Display(void*);

void
·_Cfunc_sfRenderWindow_Display(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_Display, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_TransformToLocal(void*);

void
·_Cfunc_sfString_TransformToLocal(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_TransformToLocal, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetFramerateLimit(void*);

void
·_Cfunc_sfRenderWindow_SetFramerateLimit(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfRenderWindow_SetFramerateLimit, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_GetCenterY(void*);

void
·_Cfunc_sfView_GetCenterY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_GetCenterY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetX(void*);

void
·_Cfunc_sfString_GetX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfView_GetHalfSizeY(void*);

void
·_Cfunc_sfView_GetHalfSizeY(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfView_GetHalfSizeY, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_SetBlendMode(void*);

void
·_Cfunc_sfSprite_SetBlendMode(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_SetBlendMode, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_Destroy(void*);

void
·_Cfunc_sfSprite_Destroy(struct{uint8 x[4];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_Destroy, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetX(void*);

void
·_Cfunc_sfString_SetX(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetX, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfShape_SetColor(void*);

void
·_Cfunc_sfShape_SetColor(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfShape_SetColor, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_GetRotation(void*);

void
·_Cfunc_sfString_GetRotation(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_GetRotation, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfPostFX_SetParameter1(void*);

void
·_Cfunc_sfPostFX_SetParameter1(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfPostFX_SetParameter1, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_Move(void*);

void
·_Cfunc_sfString_Move(struct{uint8 x[12];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_Move, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetWidth(void*);

void
·_Cfunc_sfSprite_GetWidth(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetWidth, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfSprite_GetSubRect(void*);

void
·_Cfunc_sfSprite_GetSubRect(struct{uint8 x[20];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfSprite_GetSubRect, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetFont(void*);

void
·_Cfunc_sfString_SetFont(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetFont, &p);
}

void _cgo_8feb4e74a977_Cfunc_sfString_SetRotation(void*);

void
·_Cfunc_sfString_SetRotation(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_8feb4e74a977_Cfunc_sfString_SetRotation, &p);
}

