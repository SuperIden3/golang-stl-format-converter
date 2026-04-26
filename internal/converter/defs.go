package converter

// --- // ASCII

type SolidName string
type Number float32

type Vector struct {
	X, Y, Z Number
}

type Facet struct {
	Normal Vector
	Vertices [3]Vector
}

// --- // BINARY

// NOTE:
// - Must use little endian: binary.LittleEndian
// - File size must be 84 + (NumTriangles * 50)

const ( // bytes
	HEADER_SIZE uint = 80
	FACET_SIZE = 50
	TRIANGLE_COUNT_SIZE = 4
	ATTRIBUTE_SIZE = 2
)

type Header [HEADER_SIZE]byte
type TriangleCount uint32
type Attribute uint16

type BinaryFacet struct {
	Facet

	Attrib Attribute
}

