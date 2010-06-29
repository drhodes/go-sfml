#include <SFML/Graphics/Color.h>
#include <SFML/Config.h>


typedef struct
{
    sfUint8 R;
    sfUint8 G;
    sfUint8 B;
    sfUint8 A;
} sfColor_P;


sfColor_P sfColor_FromRGB_P (sfUint8 R, sfUint8 G, sfUint8 B) {
	sfColor c1 = sfColor_FromRGB(R,G,B);
	sfColor_P c2;
	c2.R = c1.r;
	c2.G = c1.g;
	c2.B = c1.b;
	c2.A = c1.a;
	return c2;
}

sfColor sfUnwraP (sfColor_P col) {
	sfColor c1;
	c1.r = col.R;
	c1.g = col.G;
	c1.b = col.B;
	c1.a = col.A;
	return c1;
}

/*

		

////////////////////////////////////////////////////////////
/// Construct a color from its 4 RGBA components
///
/// \param R : Red component   (0 .. 255)
/// \param G : Green component (0 .. 255)
/// \param B : Blue component  (0 .. 255)
/// \param A : Alpha component (0 .. 255)
///
/// \return sfColor constructed from the components
///
////////////////////////////////////////////////////////////
CSFML_API sfColor sfColor_FromRGBA(sfUint8 R, sfUint8 G, sfUint8 B, sfUint8 A);

////////////////////////////////////////////////////////////
/// Add two colors
///
/// \param Color1 : First color
/// \param Color2 : Second color
///
/// \return Component-wise saturated addition of the two colors
///
////////////////////////////////////////////////////////////
CSFML_API sfColor sfColor_Add(sfColor Color1, sfColor Color2);

////////////////////////////////////////////////////////////
/// Modulate two colors
///
/// \param Color1 : First color
/// \param Color2 : Second color
///
/// \return Component-wise multiplication of the two colors
///
////////////////////////////////////////////////////////////
CSFML_API sfColor sfColor_Modulate(sfColor Color1, sfColor Color2);

*/
