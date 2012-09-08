package sfml
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/PrimitiveType.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/Graphics/Vertex.h>
// #include <SFML/Graphics/VertexArray.h>
import "C"

// typedef struct
// {
//     sfVector2f position;  ///< Position of the vertex
//     sfColor    color;     ///< Color of the vertex
//     sfVector2f texCoords; ///< Coordinates of the texture's pixel to map to the vertex
// } sfVertex;

type Vertex struct {
	Cref *C.sfVertex
}

func (self Vertex) GetPosition() (x, y float32) {
	return float32(self.Cref.position.x), float32(self.Cref.position.y)
}

func (self Vertex) SetPosition(x, y float32) {
	self.Cref.position.x = C.float(x)
	self.Cref.position.y = C.float(y)
}

func (self Vertex) GetTexCoords() (x, y float32) {
	return float32(self.Cref.texCoords.x), float32(self.Cref.texCoords.y)
}

func (self Vertex) SetTexCoords(x, y float32) {
	self.Cref.texCoords.x = C.float(x)
	self.Cref.texCoords.y = C.float(y)
}

func (self Vertex) GetColor() Color {
	return Color{self.Cref.color}
}

func (self Vertex) SetColor(c Color) {
	self.Cref.color = c.Cref
} 