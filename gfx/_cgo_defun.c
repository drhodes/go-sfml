
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

#pragma dynimport _cgo_sfString_SetColor _cgo_sfString_SetColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetColor)(void*);

void
·_C_sfString_SetColor(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetColor, &p);
}

#pragma dynimport _cgo_sfImage_CreateMaskFromColor _cgo_sfImage_CreateMaskFromColor "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_CreateMaskFromColor)(void*);

void
·_C_sfImage_CreateMaskFromColor(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfImage_CreateMaskFromColor, &p);
}

#pragma dynimport _cgo_sfColor_FromRGB _cgo_sfColor_FromRGB "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_FromRGB)(void*);

void
·_C_sfColor_FromRGB(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfColor_FromRGB, &p);
}

#pragma dynimport _cgo_sfFloatRect_Intersects _cgo_sfFloatRect_Intersects "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFloatRect_Intersects)(void*);

void
·_C_sfFloatRect_Intersects(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfFloatRect_Intersects, &p);
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

#pragma dynimport _cgo_sfImage_SetSmooth _cgo_sfImage_SetSmooth "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_SetSmooth)(void*);

void
·_C_sfImage_SetSmooth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfImage_SetSmooth, &p);
}

#pragma dynimport _cgo_sfString_GetCharacterPos _cgo_sfString_GetCharacterPos "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetCharacterPos)(void*);

void
·_C_sfString_GetCharacterPos(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfString_GetCharacterPos, &p);
}

#pragma dynimport _cgo_sfString_SetScaleY _cgo_sfString_SetScaleY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetScaleY)(void*);

void
·_C_sfString_SetScaleY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetScaleY, &p);
}

#pragma dynimport _cgo_sfIntRect_Offset _cgo_sfIntRect_Offset "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfIntRect_Offset)(void*);

void
·_C_sfIntRect_Offset(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfIntRect_Offset, &p);
}

#pragma dynimport _cgo_sfString_Create _cgo_sfString_Create "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Create)(void*);

void
·_C_sfString_Create(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfString_Create, &p);
}

#pragma dynimport _cgo_sfString_SetUnicodeText _cgo_sfString_SetUnicodeText "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetUnicodeText)(void*);

void
·_C_sfString_SetUnicodeText(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetUnicodeText, &p);
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

#pragma dynimport _cgo_sfString_TransformToGlobal _cgo_sfString_TransformToGlobal "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_TransformToGlobal)(void*);

void
·_C_sfString_TransformToGlobal(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfString_TransformToGlobal, &p);
}

#pragma dynimport _cgo_sfImage_SetPixel _cgo_sfImage_SetPixel "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_SetPixel)(void*);

void
·_C_sfImage_SetPixel(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfImage_SetPixel, &p);
}

#pragma dynimport _cgo_sfString_GetCenterY _cgo_sfString_GetCenterY "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetCenterY)(void*);

void
·_C_sfString_GetCenterY(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetCenterY, &p);
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

#pragma dynimport _cgo_sfColor_Modulate _cgo_sfColor_Modulate "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_Modulate)(void*);

void
·_C_sfColor_Modulate(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfColor_Modulate, &p);
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

#pragma dynimport _cgo_sfString_SetScale _cgo_sfString_SetScale "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetScale)(void*);

void
·_C_sfString_SetScale(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfString_SetScale, &p);
}

#pragma dynimport _cgo_sfFloatRect_Contains _cgo_sfFloatRect_Contains "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFloatRect_Contains)(void*);

void
·_C_sfFloatRect_Contains(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfFloatRect_Contains, &p);
}

#pragma dynimport _cgo_sfImage_Copy _cgo_sfImage_Copy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_Copy)(void*);

void
·_C_sfImage_Copy(struct{uint8 x[32];}p)
{
	cgocall(_cgo_sfImage_Copy, &p);
}

#pragma dynimport _cgo_sfString_Destroy _cgo_sfString_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Destroy)(void*);

void
·_C_sfString_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfString_Destroy, &p);
}

#pragma dynimport _cgo_sfColor_FromRGBA _cgo_sfColor_FromRGBA "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfColor_FromRGBA)(void*);

void
·_C_sfColor_FromRGBA(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfColor_FromRGBA, &p);
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

#pragma dynimport _cgo_sfImage_Destroy _cgo_sfImage_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_Destroy)(void*);

void
·_C_sfImage_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfImage_Destroy, &p);
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

#pragma dynimport _cgo_sfString_SetSize _cgo_sfString_SetSize "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetSize)(void*);

void
·_C_sfString_SetSize(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetSize, &p);
}

#pragma dynimport _cgo_sfImage_IsSmooth _cgo_sfImage_IsSmooth "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_IsSmooth)(void*);

void
·_C_sfImage_IsSmooth(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfImage_IsSmooth, &p);
}

#pragma dynimport _cgo_sfImage_CopyScreen _cgo_sfImage_CopyScreen "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_CopyScreen)(void*);

void
·_C_sfImage_CopyScreen(struct{uint8 x[28];}p)
{
	cgocall(_cgo_sfImage_CopyScreen, &p);
}

#pragma dynimport _cgo_sfString_SetPosition _cgo_sfString_SetPosition "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetPosition)(void*);

void
·_C_sfString_SetPosition(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfString_SetPosition, &p);
}

#pragma dynimport _cgo_sfFont_CreateFromFile _cgo_sfFont_CreateFromFile "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFont_CreateFromFile)(void*);

void
·_C_sfFont_CreateFromFile(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfFont_CreateFromFile, &p);
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

#pragma dynimport _cgo_sfString_SetStyle _cgo_sfString_SetStyle "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetStyle)(void*);

void
·_C_sfString_SetStyle(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetStyle, &p);
}

#pragma dynimport _cgo_sfImage_GetPixel _cgo_sfImage_GetPixel "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfImage_GetPixel)(void*);

void
·_C_sfImage_GetPixel(struct{uint8 x[16];}p)
{
	cgocall(_cgo_sfImage_GetPixel, &p);
}

#pragma dynimport _cgo_sfFont_Destroy _cgo_sfFont_Destroy "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfFont_Destroy)(void*);

void
·_C_sfFont_Destroy(struct{uint8 x[4];}p)
{
	cgocall(_cgo_sfFont_Destroy, &p);
}

#pragma dynimport _cgo_sfString_Rotate _cgo_sfString_Rotate "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Rotate)(void*);

void
·_C_sfString_Rotate(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_Rotate, &p);
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

#pragma dynimport _cgo_sfString_Scale _cgo_sfString_Scale "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Scale)(void*);

void
·_C_sfString_Scale(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfString_Scale, &p);
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

#pragma dynimport _cgo_sfString_TransformToLocal _cgo_sfString_TransformToLocal "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_TransformToLocal)(void*);

void
·_C_sfString_TransformToLocal(struct{uint8 x[20];}p)
{
	cgocall(_cgo_sfString_TransformToLocal, &p);
}

#pragma dynimport _cgo_sfString_GetX _cgo_sfString_GetX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetX)(void*);

void
·_C_sfString_GetX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetX, &p);
}

#pragma dynimport _cgo_sfString_SetX _cgo_sfString_SetX "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_SetX)(void*);

void
·_C_sfString_SetX(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_SetX, &p);
}

#pragma dynimport _cgo_sfString_GetRotation _cgo_sfString_GetRotation "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_GetRotation)(void*);

void
·_C_sfString_GetRotation(struct{uint8 x[8];}p)
{
	cgocall(_cgo_sfString_GetRotation, &p);
}

#pragma dynimport _cgo_sfString_Move _cgo_sfString_Move "/usr/custom/goroot/pkg/linux_386/sfml/gfx.so"
void (*_cgo_sfString_Move)(void*);

void
·_C_sfString_Move(struct{uint8 x[12];}p)
{
	cgocall(_cgo_sfString_Move, &p);
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

