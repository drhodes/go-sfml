
typedef struct { char *p; int n; } _GoString_;
_GoString_ GoString(char *p);
char *CString(_GoString_);
#include "color_patch.c"
#include <SFML/Graphics/Font.h>
#include <SFML/Graphics/Image.h>
#include <SFML/Graphics/String.h>
#include <SFML/Graphics/View.h>
#include <SFML/Graphics/Shape.h>
#include <SFML/Graphics/Sprite.h>
#include <SFML/Graphics/PostFX.h>
#include <SFML/Graphics/RenderWindow.h>



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
_cgo_sfSprite_SetScaleX(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
	} *a = v;
	sfSprite_SetScaleX(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_Create(void *v)
{
	struct {
		sfVideoMode p0;
		char* p1;
		long unsigned int p2;
		sfWindowSettings p3;
		sfRenderWindow* r;
	} *a = v;
	a->r = sfRenderWindow_Create(a->p0, a->p1, a->p2, a->p3);
}

void
_cgo_sfShape_GetColor(void *v)
{
	struct {
		sfShape* p0;
		sfColor r;
	} *a = v;
	a->r = sfShape_GetColor(a->p0);
}

void
_cgo_sfRenderWindow_SetActive(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfBool p1;
		sfBool r;
	} *a = v;
	a->r = sfRenderWindow_SetActive(a->p0, a->p1);
}

void
_cgo_sfString_SetColor(void *v)
{
	struct {
		sfString* p0;
		sfColor p1;
	} *a = v;
	sfString_SetColor(a->p0, a->p1);
}

void
_cgo_sfShape_SetCenter(void *v)
{
	struct {
		sfShape* p0;
		float p1;
		float p2;
	} *a = v;
	sfShape_SetCenter(a->p0, a->p1, a->p2);
}

void
_cgo_sfSprite_GetCenterX(void *v)
{
	struct {
		sfSprite* p0;
		float r;
	} *a = v;
	a->r = sfSprite_GetCenterX(a->p0);
}

void
_cgo_sfView_Zoom(void *v)
{
	struct {
		sfView* p0;
		float p1;
	} *a = v;
	sfView_Zoom(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_SetView(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfView* p1;
	} *a = v;
	sfRenderWindow_SetView(a->p0, a->p1);
}

void
_cgo_sfView_SetHalfSize(void *v)
{
	struct {
		sfView* p0;
		float p1;
		float p2;
	} *a = v;
	sfView_SetHalfSize(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_SetX(void *v)
{
	struct {
		sfShape* p0;
		float p1;
	} *a = v;
	sfShape_SetX(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_UseVerticalSync(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfBool p1;
	} *a = v;
	sfRenderWindow_UseVerticalSync(a->p0, a->p1);
}

void
_cgo_sfImage_CreateMaskFromColor(void *v)
{
	struct {
		sfImage* p0;
		sfColor p1;
		sfUint8 p2;
		char __pad0[3];
	} *a = v;
	sfImage_CreateMaskFromColor(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_CreateCircle(void *v)
{
	struct {
		float p0;
		float p1;
		float p2;
		sfColor p3;
		float p4;
		sfColor p5;
		sfShape* r;
	} *a = v;
	a->r = sfShape_CreateCircle(a->p0, a->p1, a->p2, a->p3, a->p4, a->p5);
}

void
_cgo_sfView_GetRect(void *v)
{
	struct {
		sfView* p0;
		sfFloatRect r;
	} *a = v;
	a->r = sfView_GetRect(a->p0);
}

void
_cgo_sfRenderWindow_SetCursorPosition(void *v)
{
	struct {
		sfRenderWindow* p0;
		unsigned int p1;
		unsigned int p2;
	} *a = v;
	sfRenderWindow_SetCursorPosition(a->p0, a->p1, a->p2);
}

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
_cgo_sfShape_CreateRectangle(void *v)
{
	struct {
		float p0;
		float p1;
		float p2;
		float p3;
		sfColor p4;
		float p5;
		sfColor p6;
		sfShape* r;
	} *a = v;
	a->r = sfShape_CreateRectangle(a->p0, a->p1, a->p2, a->p3, a->p4, a->p5, a->p6);
}

void
_cgo_sfFloatRect_Intersects(void *v)
{
	struct {
		sfFloatRect* p0;
		sfFloatRect* p1;
		sfFloatRect* p2;
		sfBool r;
	} *a = v;
	a->r = sfFloatRect_Intersects(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_SetPointPosition(void *v)
{
	struct {
		sfShape* p0;
		unsigned int p1;
		float p2;
		float p3;
	} *a = v;
	sfShape_SetPointPosition(a->p0, a->p1, a->p2, a->p3);
}

void
_cgo_sfView_CreateFromRect(void *v)
{
	struct {
		sfFloatRect p0;
		sfView* r;
	} *a = v;
	a->r = sfView_CreateFromRect(a->p0);
}

void
_cgo_sfUnwraP(void *v)
{
	struct {
		sfColor_P p0;
		sfColor r;
	} *a = v;
	a->r = sfUnwraP(a->p0);
}

void
_cgo_sfRenderWindow_SetJoystickThreshold(void *v)
{
	struct {
		sfRenderWindow* p0;
		float p1;
	} *a = v;
	sfRenderWindow_SetJoystickThreshold(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_Destroy(void *v)
{
	struct {
		sfRenderWindow* p0;
	} *a = v;
	sfRenderWindow_Destroy(a->p0);
}

void
_cgo_sfShape_GetScaleX(void *v)
{
	struct {
		sfShape* p0;
		float r;
	} *a = v;
	a->r = sfShape_GetScaleX(a->p0);
}

void
_cgo_sfSprite_SetRotation(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
	} *a = v;
	sfSprite_SetRotation(a->p0, a->p1);
}

void
_cgo_sfImage_CreateFromFile(void *v)
{
	struct {
		char* p0;
		sfImage* r;
	} *a = v;
	a->r = sfImage_CreateFromFile(a->p0);
}

void
_cgo_sfString_GetCenterX(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetCenterX(a->p0);
}

void
_cgo_sfSprite_FlipX(void *v)
{
	struct {
		sfSprite* p0;
		sfBool p1;
	} *a = v;
	sfSprite_FlipX(a->p0, a->p1);
}

void
_cgo_sfView_SetFromRect(void *v)
{
	struct {
		sfView* p0;
		sfFloatRect p1;
	} *a = v;
	sfView_SetFromRect(a->p0, a->p1);
}

void
_cgo_sfImage_SetSmooth(void *v)
{
	struct {
		sfImage* p0;
		sfBool p1;
	} *a = v;
	sfImage_SetSmooth(a->p0, a->p1);
}

void
_cgo_sfPostFX_SetParameter4(void *v)
{
	struct {
		sfPostFX* p0;
		char* p1;
		float p2;
		float p3;
		float p4;
		float p5;
	} *a = v;
	sfPostFX_SetParameter4(a->p0, a->p1, a->p2, a->p3, a->p4, a->p5);
}

void
_cgo_sfSprite_GetPixel(void *v)
{
	struct {
		sfSprite* p0;
		unsigned int p1;
		unsigned int p2;
		sfColor r;
	} *a = v;
	a->r = sfSprite_GetPixel(a->p0, a->p1, a->p2);
}

void
_cgo_sfView_GetHalfSizeX(void *v)
{
	struct {
		sfView* p0;
		float r;
	} *a = v;
	a->r = sfView_GetHalfSizeX(a->p0);
}

void
_cgo_sfString_SetScaleY(void *v)
{
	struct {
		sfString* p0;
		float p1;
	} *a = v;
	sfString_SetScaleY(a->p0, a->p1);
}

void
_cgo_sfShape_GetY(void *v)
{
	struct {
		sfShape* p0;
		float r;
	} *a = v;
	a->r = sfShape_GetY(a->p0);
}

void
_cgo_sfShape_EnableFill(void *v)
{
	struct {
		sfShape* p0;
		sfBool p1;
	} *a = v;
	sfShape_EnableFill(a->p0, a->p1);
}

void
_cgo_sfSprite_Scale(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
		float p2;
	} *a = v;
	sfSprite_Scale(a->p0, a->p1, a->p2);
}

void
_cgo_sfIntRect_Offset(void *v)
{
	struct {
		sfIntRect* p0;
		int p1;
		int p2;
	} *a = v;
	sfIntRect_Offset(a->p0, a->p1, a->p2);
}

void
_cgo_sfSprite_GetImage(void *v)
{
	struct {
		sfSprite* p0;
		sfImage* r;
	} *a = v;
	a->r = sfSprite_GetImage(a->p0);
}

void
_cgo_sfRenderWindow_DrawSprite(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfSprite* p1;
	} *a = v;
	sfRenderWindow_DrawSprite(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_GetSettings(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfWindowSettings r;
	} *a = v;
	a->r = sfRenderWindow_GetSettings(a->p0);
}

void
_cgo_sfSprite_SetSubRect(void *v)
{
	struct {
		sfSprite* p0;
		sfIntRect p1;
	} *a = v;
	sfSprite_SetSubRect(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_Show(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfBool p1;
	} *a = v;
	sfRenderWindow_Show(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_Capture(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfImage* r;
	} *a = v;
	a->r = sfRenderWindow_Capture(a->p0);
}

void
_cgo_sfString_Create(void *v)
{
	struct {
		sfString* r;
	} *a = v;
	a->r = sfString_Create();
}

void
_cgo_sfShape_GetCenterX(void *v)
{
	struct {
		sfShape* p0;
		float r;
	} *a = v;
	a->r = sfShape_GetCenterX(a->p0);
}

void
_cgo_sfShape_Move(void *v)
{
	struct {
		sfShape* p0;
		float p1;
		float p2;
	} *a = v;
	sfShape_Move(a->p0, a->p1, a->p2);
}

void
_cgo_sfString_SetUnicodeText(void *v)
{
	struct {
		sfString* p0;
		sfUint32* p1;
	} *a = v;
	sfString_SetUnicodeText(a->p0, a->p1);
}

void
_cgo_sfSprite_SetScale(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
		float p2;
	} *a = v;
	sfSprite_SetScale(a->p0, a->p1, a->p2);
}

void
_cgo_sfSprite_GetColor(void *v)
{
	struct {
		sfSprite* p0;
		sfColor r;
	} *a = v;
	a->r = sfSprite_GetColor(a->p0);
}

void
_cgo_sfSprite_GetCenterY(void *v)
{
	struct {
		sfSprite* p0;
		float r;
	} *a = v;
	a->r = sfSprite_GetCenterY(a->p0);
}

void
_cgo_sfImage_CreateFromMemory(void *v)
{
	struct {
		char* p0;
		size_t p1;
		sfImage* r;
	} *a = v;
	a->r = sfImage_CreateFromMemory(a->p0, a->p1);
}

void
_cgo_sfString_GetFont(void *v)
{
	struct {
		sfString* p0;
		sfFont* r;
	} *a = v;
	a->r = sfString_GetFont(a->p0);
}

void
_cgo_sfString_GetUnicodeText(void *v)
{
	struct {
		sfString* p0;
		sfUint32* r;
	} *a = v;
	a->r = sfString_GetUnicodeText(a->p0);
}

void
_cgo_sfSprite_SetY(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
	} *a = v;
	sfSprite_SetY(a->p0, a->p1);
}

void
_cgo_sfPostFX_CreateFromFile(void *v)
{
	struct {
		char* p0;
		sfPostFX* r;
	} *a = v;
	a->r = sfPostFX_CreateFromFile(a->p0);
}

void
_cgo_sfString_TransformToGlobal(void *v)
{
	struct {
		sfString* p0;
		float p1;
		float p2;
		float* p3;
		float* p4;
	} *a = v;
	sfString_TransformToGlobal(a->p0, a->p1, a->p2, a->p3, a->p4);
}

void
_cgo_sfShape_GetRotation(void *v)
{
	struct {
		sfShape* p0;
		float r;
	} *a = v;
	a->r = sfShape_GetRotation(a->p0);
}

void
_cgo_sfImage_SetPixel(void *v)
{
	struct {
		sfImage* p0;
		unsigned int p1;
		unsigned int p2;
		sfColor p3;
	} *a = v;
	sfImage_SetPixel(a->p0, a->p1, a->p2, a->p3);
}

void
_cgo_sfSprite_Rotate(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
	} *a = v;
	sfSprite_Rotate(a->p0, a->p1);
}

void
_cgo_sfShape_AddPoint(void *v)
{
	struct {
		sfShape* p0;
		float p1;
		float p2;
		sfColor p3;
		sfColor p4;
	} *a = v;
	sfShape_AddPoint(a->p0, a->p1, a->p2, a->p3, a->p4);
}

void
_cgo_sfSprite_SetColor(void *v)
{
	struct {
		sfSprite* p0;
		sfColor p1;
	} *a = v;
	sfSprite_SetColor(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_GetDefaultView(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfView* r;
	} *a = v;
	a->r = sfRenderWindow_GetDefaultView(a->p0);
}

void
_cgo_sfShape_Scale(void *v)
{
	struct {
		sfShape* p0;
		float p1;
		float p2;
	} *a = v;
	sfShape_Scale(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_GetNbPoints(void *v)
{
	struct {
		sfShape* p0;
		unsigned int r;
	} *a = v;
	a->r = sfShape_GetNbPoints(a->p0);
}

void
_cgo_sfShape_GetBlendMode(void *v)
{
	struct {
		sfShape* p0;
		sfBlendMode r;
	} *a = v;
	a->r = sfShape_GetBlendMode(a->p0);
}

void
_cgo_sfString_GetCenterY(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetCenterY(a->p0);
}

void
_cgo_sfSprite_FlipY(void *v)
{
	struct {
		sfSprite* p0;
		sfBool p1;
	} *a = v;
	sfSprite_FlipY(a->p0, a->p1);
}

void
_cgo_sfSprite_SetPosition(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
		float p2;
	} *a = v;
	sfSprite_SetPosition(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_GetPointOutlineColor(void *v)
{
	struct {
		sfShape* p0;
		unsigned int p1;
		sfColor r;
	} *a = v;
	a->r = sfShape_GetPointOutlineColor(a->p0, a->p1);
}

void
_cgo_sfSprite_GetRotation(void *v)
{
	struct {
		sfSprite* p0;
		float r;
	} *a = v;
	a->r = sfSprite_GetRotation(a->p0);
}

void
_cgo_sfPostFX_CanUsePostFX(void *v)
{
	struct {
		sfBool r;
	} *a = v;
	a->r = sfPostFX_CanUsePostFX();
}

void
_cgo_sfString_GetRect(void *v)
{
	struct {
		sfString* p0;
		sfFloatRect r;
	} *a = v;
	a->r = sfString_GetRect(a->p0);
}

void
_cgo_sfString_SetScaleX(void *v)
{
	struct {
		sfString* p0;
		float p1;
	} *a = v;
	sfString_SetScaleX(a->p0, a->p1);
}

void
_cgo_sfSprite_Resize(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
		float p2;
	} *a = v;
	sfSprite_Resize(a->p0, a->p1, a->p2);
}

void
_cgo_sfRenderWindow_GetWidth(void *v)
{
	struct {
		sfRenderWindow* p0;
		unsigned int r;
	} *a = v;
	a->r = sfRenderWindow_GetWidth(a->p0);
}

void
_cgo_sfRenderWindow_Close(void *v)
{
	struct {
		sfRenderWindow* p0;
	} *a = v;
	sfRenderWindow_Close(a->p0);
}

void
_cgo_sfShape_GetX(void *v)
{
	struct {
		sfShape* p0;
		float r;
	} *a = v;
	a->r = sfShape_GetX(a->p0);
}

void
_cgo_sfPostFX_SetTexture(void *v)
{
	struct {
		sfPostFX* p0;
		char* p1;
		sfImage* p2;
	} *a = v;
	sfPostFX_SetTexture(a->p0, a->p1, a->p2);
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
_cgo_sfView_Move(void *v)
{
	struct {
		sfView* p0;
		float p1;
		float p2;
	} *a = v;
	sfView_Move(a->p0, a->p1, a->p2);
}

void
_cgo_sfRenderWindow_SetSize(void *v)
{
	struct {
		sfRenderWindow* p0;
		unsigned int p1;
		unsigned int p2;
	} *a = v;
	sfRenderWindow_SetSize(a->p0, a->p1, a->p2);
}

void
_cgo_sfIntRect_Contains(void *v)
{
	struct {
		sfIntRect* p0;
		int p1;
		int p2;
		sfBool r;
	} *a = v;
	a->r = sfIntRect_Contains(a->p0, a->p1, a->p2);
}

void
_cgo_sfFont_Create(void *v)
{
	struct {
		sfFont* r;
	} *a = v;
	a->r = sfFont_Create();
}

void
_cgo_sfPostFX_SetParameter3(void *v)
{
	struct {
		sfPostFX* p0;
		char* p1;
		float p2;
		float p3;
		float p4;
	} *a = v;
	sfPostFX_SetParameter3(a->p0, a->p1, a->p2, a->p3, a->p4);
}

void
_cgo_sfColor_FromRGB_P(void *v)
{
	struct {
		sfUint8 p0;
		sfUint8 p1;
		sfUint8 p2;
		char __pad0[1];
		sfColor_P r;
	} *a = v;
	a->r = sfColor_FromRGB_P(a->p0, a->p1, a->p2);
}

void
_cgo_sfRenderWindow_GetHeight(void *v)
{
	struct {
		sfRenderWindow* p0;
		unsigned int r;
	} *a = v;
	a->r = sfRenderWindow_GetHeight(a->p0);
}

void
_cgo_sfShape_SetScale(void *v)
{
	struct {
		sfShape* p0;
		float p1;
		float p2;
	} *a = v;
	sfShape_SetScale(a->p0, a->p1, a->p2);
}

void
_cgo_sfString_SetScale(void *v)
{
	struct {
		sfString* p0;
		float p1;
		float p2;
	} *a = v;
	sfString_SetScale(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_GetCenterY(void *v)
{
	struct {
		sfShape* p0;
		float r;
	} *a = v;
	a->r = sfShape_GetCenterY(a->p0);
}

void
_cgo_sfFloatRect_Contains(void *v)
{
	struct {
		sfFloatRect* p0;
		float p1;
		float p2;
		sfBool r;
	} *a = v;
	a->r = sfFloatRect_Contains(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_Rotate(void *v)
{
	struct {
		sfShape* p0;
		float p1;
	} *a = v;
	sfShape_Rotate(a->p0, a->p1);
}

void
_cgo_sfImage_Copy(void *v)
{
	struct {
		sfImage* p0;
		sfImage* p1;
		unsigned int p2;
		unsigned int p3;
		sfIntRect p4;
	} *a = v;
	sfImage_Copy(a->p0, a->p1, a->p2, a->p3, a->p4);
}

void
_cgo_sfPostFX_Destroy(void *v)
{
	struct {
		sfPostFX* p0;
	} *a = v;
	sfPostFX_Destroy(a->p0);
}

void
_cgo_sfView_Destroy(void *v)
{
	struct {
		sfView* p0;
	} *a = v;
	sfView_Destroy(a->p0);
}

void
_cgo_sfShape_SetScaleY(void *v)
{
	struct {
		sfShape* p0;
		float p1;
	} *a = v;
	sfShape_SetScaleY(a->p0, a->p1);
}

void
_cgo_sfShape_SetBlendMode(void *v)
{
	struct {
		sfShape* p0;
		sfBlendMode p1;
	} *a = v;
	sfShape_SetBlendMode(a->p0, a->p1);
}

void
_cgo_sfString_Destroy(void *v)
{
	struct {
		sfString* p0;
	} *a = v;
	sfString_Destroy(a->p0);
}

void
_cgo_sfSprite_SetCenter(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
		float p2;
	} *a = v;
	sfSprite_SetCenter(a->p0, a->p1, a->p2);
}

void
_cgo_sfRenderWindow_DrawShape(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfShape* p1;
	} *a = v;
	sfRenderWindow_DrawShape(a->p0, a->p1);
}

void
_cgo_sfSprite_SetX(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
	} *a = v;
	sfSprite_SetX(a->p0, a->p1);
}

void
_cgo_sfSprite_Create(void *v)
{
	struct {
		sfSprite* r;
	} *a = v;
	a->r = sfSprite_Create();
}

void
_cgo_sfShape_EnableOutline(void *v)
{
	struct {
		sfShape* p0;
		sfBool p1;
	} *a = v;
	sfShape_EnableOutline(a->p0, a->p1);
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
_cgo_sfSprite_GetScaleY(void *v)
{
	struct {
		sfSprite* p0;
		float r;
	} *a = v;
	a->r = sfSprite_GetScaleY(a->p0);
}

void
_cgo_sfRenderWindow_EnableKeyRepeat(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfBool p1;
	} *a = v;
	sfRenderWindow_EnableKeyRepeat(a->p0, a->p1);
}

void
_cgo_sfImage_Bind(void *v)
{
	struct {
		sfImage* p0;
	} *a = v;
	sfImage_Bind(a->p0);
}

void
_cgo_sfImage_GetHeight(void *v)
{
	struct {
		sfImage* p0;
		unsigned int r;
	} *a = v;
	a->r = sfImage_GetHeight(a->p0);
}

void
_cgo_sfSprite_Move(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
		float p2;
	} *a = v;
	sfSprite_Move(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_SetPosition(void *v)
{
	struct {
		sfShape* p0;
		float p1;
		float p2;
	} *a = v;
	sfShape_SetPosition(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_GetOutlineWidth(void *v)
{
	struct {
		sfShape* p0;
		float r;
	} *a = v;
	a->r = sfShape_GetOutlineWidth(a->p0);
}

void
_cgo_sfImage_Destroy(void *v)
{
	struct {
		sfImage* p0;
	} *a = v;
	sfImage_Destroy(a->p0);
}

void
_cgo_sfShape_Destroy(void *v)
{
	struct {
		sfShape* p0;
	} *a = v;
	sfShape_Destroy(a->p0);
}

void
_cgo_sfString_GetScaleY(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetScaleY(a->p0);
}

void
_cgo_sfImage_CreateFromColor(void *v)
{
	struct {
		unsigned int p0;
		unsigned int p1;
		sfColor p2;
		sfImage* r;
	} *a = v;
	a->r = sfImage_CreateFromColor(a->p0, a->p1, a->p2);
}

void
_cgo_sfRenderWindow_DrawString(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfString* p1;
	} *a = v;
	sfRenderWindow_DrawString(a->p0, a->p1);
}

void
_cgo_sfShape_Create(void *v)
{
	struct {
		sfShape* r;
	} *a = v;
	a->r = sfShape_Create();
}

void
_cgo_sfString_SetSize(void *v)
{
	struct {
		sfString* p0;
		float p1;
	} *a = v;
	sfString_SetSize(a->p0, a->p1);
}

void
_cgo_sfShape_SetPointColor(void *v)
{
	struct {
		sfShape* p0;
		unsigned int p1;
		sfColor p2;
	} *a = v;
	sfShape_SetPointColor(a->p0, a->p1, a->p2);
}

void
_cgo_sfRenderWindow_IsOpened(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfBool r;
	} *a = v;
	a->r = sfRenderWindow_IsOpened(a->p0);
}

void
_cgo_sfSprite_SetImage(void *v)
{
	struct {
		sfSprite* p0;
		sfImage* p1;
	} *a = v;
	sfSprite_SetImage(a->p0, a->p1);
}

void
_cgo_sfImage_IsSmooth(void *v)
{
	struct {
		sfImage* p0;
		sfBool r;
	} *a = v;
	a->r = sfImage_IsSmooth(a->p0);
}

void
_cgo_sfSprite_GetY(void *v)
{
	struct {
		sfSprite* p0;
		float r;
	} *a = v;
	a->r = sfSprite_GetY(a->p0);
}

void
_cgo_sfImage_CopyScreen(void *v)
{
	struct {
		sfImage* p0;
		sfRenderWindow* p1;
		sfIntRect p2;
		sfBool r;
	} *a = v;
	a->r = sfImage_CopyScreen(a->p0, a->p1, a->p2);
}

void
_cgo_sfView_SetCenter(void *v)
{
	struct {
		sfView* p0;
		float p1;
		float p2;
	} *a = v;
	sfView_SetCenter(a->p0, a->p1, a->p2);
}

void
_cgo_sfString_SetPosition(void *v)
{
	struct {
		sfString* p0;
		float p1;
		float p2;
	} *a = v;
	sfString_SetPosition(a->p0, a->p1, a->p2);
}

void
_cgo_sfPostFX_SetParameter2(void *v)
{
	struct {
		sfPostFX* p0;
		char* p1;
		float p2;
		float p3;
	} *a = v;
	sfPostFX_SetParameter2(a->p0, a->p1, a->p2, a->p3);
}

void
_cgo_sfFont_CreateFromFile(void *v)
{
	struct {
		char* p0;
		unsigned int p1;
		sfUint32* p2;
		sfFont* r;
	} *a = v;
	a->r = sfFont_CreateFromFile(a->p0, a->p1, a->p2);
}

void
_cgo_sfView_GetCenterX(void *v)
{
	struct {
		sfView* p0;
		float r;
	} *a = v;
	a->r = sfView_GetCenterX(a->p0);
}

void
_cgo_sfString_SetCenter(void *v)
{
	struct {
		sfString* p0;
		float p1;
		float p2;
	} *a = v;
	sfString_SetCenter(a->p0, a->p1, a->p2);
}

void
_cgo_sfString_GetY(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetY(a->p0);
}

void
_cgo_sfRenderWindow_Clear(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfColor p1;
	} *a = v;
	sfRenderWindow_Clear(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_PreserveOpenGLStates(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfBool p1;
	} *a = v;
	sfRenderWindow_PreserveOpenGLStates(a->p0, a->p1);
}

void
_cgo_sfShape_SetOutlineWidth(void *v)
{
	struct {
		sfShape* p0;
		float p1;
	} *a = v;
	sfShape_SetOutlineWidth(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_DrawPostFX(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfPostFX* p1;
	} *a = v;
	sfRenderWindow_DrawPostFX(a->p0, a->p1);
}

void
_cgo_sfRenderWindow_ShowMouseCursor(void *v)
{
	struct {
		sfRenderWindow* p0;
		sfBool p1;
	} *a = v;
	sfRenderWindow_ShowMouseCursor(a->p0, a->p1);
}

void
_cgo_sfString_SetY(void *v)
{
	struct {
		sfString* p0;
		float p1;
	} *a = v;
	sfString_SetY(a->p0, a->p1);
}

void
_cgo_sfImage_GetWidth(void *v)
{
	struct {
		sfImage* p0;
		unsigned int r;
	} *a = v;
	a->r = sfImage_GetWidth(a->p0);
}

void
_cgo_sfShape_SetRotation(void *v)
{
	struct {
		sfShape* p0;
		float p1;
	} *a = v;
	sfShape_SetRotation(a->p0, a->p1);
}

void
_cgo_sfShape_SetScaleX(void *v)
{
	struct {
		sfShape* p0;
		float p1;
	} *a = v;
	sfShape_SetScaleX(a->p0, a->p1);
}

void
_cgo_sfString_GetBlendMode(void *v)
{
	struct {
		sfString* p0;
		sfBlendMode r;
	} *a = v;
	a->r = sfString_GetBlendMode(a->p0);
}

void
_cgo_sfIntRect_Intersects(void *v)
{
	struct {
		sfIntRect* p0;
		sfIntRect* p1;
		sfIntRect* p2;
		sfBool r;
	} *a = v;
	a->r = sfIntRect_Intersects(a->p0, a->p1, a->p2);
}

void
_cgo_sfString_GetColor(void *v)
{
	struct {
		sfString* p0;
		sfColor r;
	} *a = v;
	a->r = sfString_GetColor(a->p0);
}

void
_cgo_sfImage_Create(void *v)
{
	struct {
		sfImage* r;
	} *a = v;
	a->r = sfImage_Create();
}

void
_cgo_sfImage_CreateFromPixels(void *v)
{
	struct {
		unsigned int p0;
		unsigned int p1;
		sfUint8* p2;
		sfImage* r;
	} *a = v;
	a->r = sfImage_CreateFromPixels(a->p0, a->p1, a->p2);
}

void
_cgo_sfFont_GetDefaultFont(void *v)
{
	struct {
		sfFont* r;
	} *a = v;
	a->r = sfFont_GetDefaultFont();
}

void
_cgo_sfRenderWindow_SetPosition(void *v)
{
	struct {
		sfRenderWindow* p0;
		int p1;
		int p2;
	} *a = v;
	sfRenderWindow_SetPosition(a->p0, a->p1, a->p2);
}

void
_cgo_sfSprite_SetScaleY(void *v)
{
	struct {
		sfSprite* p0;
		float p1;
	} *a = v;
	sfSprite_SetScaleY(a->p0, a->p1);
}

void
_cgo_sfString_SetStyle(void *v)
{
	struct {
		sfString* p0;
		long unsigned int p1;
	} *a = v;
	sfString_SetStyle(a->p0, a->p1);
}

void
_cgo_sfShape_SetPointOutlineColor(void *v)
{
	struct {
		sfShape* p0;
		unsigned int p1;
		sfColor p2;
	} *a = v;
	sfShape_SetPointOutlineColor(a->p0, a->p1, a->p2);
}

void
_cgo_sfView_Create(void *v)
{
	struct {
		sfView* r;
	} *a = v;
	a->r = sfView_Create();
}

void
_cgo_sfImage_GetPixel(void *v)
{
	struct {
		sfImage* p0;
		unsigned int p1;
		unsigned int p2;
		sfColor r;
	} *a = v;
	a->r = sfImage_GetPixel(a->p0, a->p1, a->p2);
}

void
_cgo_sfShape_CreateLine(void *v)
{
	struct {
		float p0;
		float p1;
		float p2;
		float p3;
		float p4;
		sfColor p5;
		float p6;
		sfColor p7;
		sfShape* r;
	} *a = v;
	a->r = sfShape_CreateLine(a->p0, a->p1, a->p2, a->p3, a->p4, a->p5, a->p6, a->p7);
}

void
_cgo_sfSprite_GetScaleX(void *v)
{
	struct {
		sfSprite* p0;
		float r;
	} *a = v;
	a->r = sfSprite_GetScaleX(a->p0);
}

void
_cgo_sfFont_Destroy(void *v)
{
	struct {
		sfFont* p0;
	} *a = v;
	sfFont_Destroy(a->p0);
}

void
_cgo_sfShape_SetY(void *v)
{
	struct {
		sfShape* p0;
		float p1;
	} *a = v;
	sfShape_SetY(a->p0, a->p1);
}

void
_cgo_sfString_Rotate(void *v)
{
	struct {
		sfString* p0;
		float p1;
	} *a = v;
	sfString_Rotate(a->p0, a->p1);
}

void
_cgo_sfSprite_GetBlendMode(void *v)
{
	struct {
		sfSprite* p0;
		sfBlendMode r;
	} *a = v;
	a->r = sfSprite_GetBlendMode(a->p0);
}

void
_cgo_sfShape_GetPointColor(void *v)
{
	struct {
		sfShape* p0;
		unsigned int p1;
		sfColor r;
	} *a = v;
	a->r = sfShape_GetPointColor(a->p0, a->p1);
}

void
_cgo_sfString_SetText(void *v)
{
	struct {
		sfString* p0;
		char* p1;
	} *a = v;
	sfString_SetText(a->p0, a->p1);
}

void
_cgo_sfString_GetSize(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetSize(a->p0);
}

void
_cgo_sfFloatRect_Offset(void *v)
{
	struct {
		sfFloatRect* p0;
		float p1;
		float p2;
	} *a = v;
	sfFloatRect_Offset(a->p0, a->p1, a->p2);
}

void
_cgo_sfString_GetScaleX(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetScaleX(a->p0);
}

void
_cgo_sfShape_GetScaleY(void *v)
{
	struct {
		sfShape* p0;
		float r;
	} *a = v;
	a->r = sfShape_GetScaleY(a->p0);
}

void
_cgo_sfString_Scale(void *v)
{
	struct {
		sfString* p0;
		float p1;
		float p2;
	} *a = v;
	sfString_Scale(a->p0, a->p1, a->p2);
}

void
_cgo_sfSprite_GetHeight(void *v)
{
	struct {
		sfSprite* p0;
		float r;
	} *a = v;
	a->r = sfSprite_GetHeight(a->p0);
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

void
_cgo_sfString_SetBlendMode(void *v)
{
	struct {
		sfString* p0;
		sfBlendMode p1;
	} *a = v;
	sfString_SetBlendMode(a->p0, a->p1);
}

void
_cgo_sfImage_SaveToFile(void *v)
{
	struct {
		sfImage* p0;
		char* p1;
		sfBool r;
	} *a = v;
	a->r = sfImage_SaveToFile(a->p0, a->p1);
}

void
_cgo_sfSprite_GetX(void *v)
{
	struct {
		sfSprite* p0;
		float r;
	} *a = v;
	a->r = sfSprite_GetX(a->p0);
}

void
_cgo_sfRenderWindow_Display(void *v)
{
	struct {
		sfRenderWindow* p0;
	} *a = v;
	sfRenderWindow_Display(a->p0);
}

void
_cgo_sfString_TransformToLocal(void *v)
{
	struct {
		sfString* p0;
		float p1;
		float p2;
		float* p3;
		float* p4;
	} *a = v;
	sfString_TransformToLocal(a->p0, a->p1, a->p2, a->p3, a->p4);
}

void
_cgo_sfRenderWindow_SetFramerateLimit(void *v)
{
	struct {
		sfRenderWindow* p0;
		unsigned int p1;
	} *a = v;
	sfRenderWindow_SetFramerateLimit(a->p0, a->p1);
}

void
_cgo_sfView_GetCenterY(void *v)
{
	struct {
		sfView* p0;
		float r;
	} *a = v;
	a->r = sfView_GetCenterY(a->p0);
}

void
_cgo_sfString_GetX(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetX(a->p0);
}

void
_cgo_sfView_GetHalfSizeY(void *v)
{
	struct {
		sfView* p0;
		float r;
	} *a = v;
	a->r = sfView_GetHalfSizeY(a->p0);
}

void
_cgo_sfSprite_SetBlendMode(void *v)
{
	struct {
		sfSprite* p0;
		sfBlendMode p1;
	} *a = v;
	sfSprite_SetBlendMode(a->p0, a->p1);
}

void
_cgo_sfSprite_Destroy(void *v)
{
	struct {
		sfSprite* p0;
	} *a = v;
	sfSprite_Destroy(a->p0);
}

void
_cgo_sfString_SetX(void *v)
{
	struct {
		sfString* p0;
		float p1;
	} *a = v;
	sfString_SetX(a->p0, a->p1);
}

void
_cgo_sfShape_SetColor(void *v)
{
	struct {
		sfShape* p0;
		sfColor p1;
	} *a = v;
	sfShape_SetColor(a->p0, a->p1);
}

void
_cgo_sfString_GetRotation(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetRotation(a->p0);
}

void
_cgo_sfPostFX_SetParameter1(void *v)
{
	struct {
		sfPostFX* p0;
		char* p1;
		float p2;
	} *a = v;
	sfPostFX_SetParameter1(a->p0, a->p1, a->p2);
}

void
_cgo_sfString_Move(void *v)
{
	struct {
		sfString* p0;
		float p1;
		float p2;
	} *a = v;
	sfString_Move(a->p0, a->p1, a->p2);
}

void
_cgo_sfSprite_GetWidth(void *v)
{
	struct {
		sfSprite* p0;
		float r;
	} *a = v;
	a->r = sfSprite_GetWidth(a->p0);
}

void
_cgo_sfSprite_GetSubRect(void *v)
{
	struct {
		sfSprite* p0;
		sfIntRect r;
	} *a = v;
	a->r = sfSprite_GetSubRect(a->p0);
}

void
_cgo_sfString_SetFont(void *v)
{
	struct {
		sfString* p0;
		sfFont* p1;
	} *a = v;
	sfString_SetFont(a->p0, a->p1);
}

void
_cgo_sfString_SetRotation(void *v)
{
	struct {
		sfString* p0;
		float p1;
	} *a = v;
	sfString_SetRotation(a->p0, a->p1);
}

