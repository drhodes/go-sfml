
typedef struct { char *p; int n; } _GoString_;
_GoString_ GoString(char *p);
char *CString(_GoString_);
#include <SFML/Graphics/Color.h>
#include <SFML/Graphics/Font.h>
#include <SFML/Graphics/Image.h>
#include <SFML/Graphics/String.h>



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
_cgo_sfString_SetColor(void *v)
{
	struct {
		sfString* p0;
		sfColor p1;
	} *a = v;
	sfString_SetColor(a->p0, a->p1);
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
_cgo_sfImage_SetSmooth(void *v)
{
	struct {
		sfImage* p0;
		sfBool p1;
	} *a = v;
	sfImage_SetSmooth(a->p0, a->p1);
}

void
_cgo_sfString_GetCharacterPos(void *v)
{
	struct {
		sfString* p0;
		size_t p1;
		float* p2;
		float* p3;
	} *a = v;
	sfString_GetCharacterPos(a->p0, a->p1, a->p2, a->p3);
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
_cgo_sfString_Create(void *v)
{
	struct {
		sfString* r;
	} *a = v;
	a->r = sfString_Create();
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
_cgo_sfString_GetCenterY(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetCenterY(a->p0);
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
_cgo_sfString_Destroy(void *v)
{
	struct {
		sfString* p0;
	} *a = v;
	sfString_Destroy(a->p0);
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
_cgo_sfImage_Destroy(void *v)
{
	struct {
		sfImage* p0;
	} *a = v;
	sfImage_Destroy(a->p0);
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
_cgo_sfString_SetSize(void *v)
{
	struct {
		sfString* p0;
		float p1;
	} *a = v;
	sfString_SetSize(a->p0, a->p1);
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
_cgo_sfString_SetStyle(void *v)
{
	struct {
		sfString* p0;
		long unsigned int p1;
	} *a = v;
	sfString_SetStyle(a->p0, a->p1);
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
_cgo_sfFont_Destroy(void *v)
{
	struct {
		sfFont* p0;
	} *a = v;
	sfFont_Destroy(a->p0);
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
_cgo_sfString_GetX(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetX(a->p0);
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
_cgo_sfString_GetRotation(void *v)
{
	struct {
		sfString* p0;
		float r;
	} *a = v;
	a->r = sfString_GetRotation(a->p0);
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

