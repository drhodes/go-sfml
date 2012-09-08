package sfml

type PrimitiveType uint

const (
	// List of individual points
	sfPoints PrimitiveType = iota
	// List of individual lines
	sfLines
	// List of connected lines, a point uses the previous point to form a line
	sfLinesStrip 
	// List of individual triangles
	sfTriangles
	// List of connected triangles, a point uses the two previous points to form a triangle
	sfTrianglesStrip
	// List of connected triangles, a point uses the common center and the previous point to form a triangle
	sfTrianglesFan
	// List of individual quads
	sfQuads
)
