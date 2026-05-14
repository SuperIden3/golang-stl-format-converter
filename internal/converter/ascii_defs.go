package converter

type SolidName string
type Number float32

type Vector struct {
	X, Y, Z Number
}

type Facet struct {
	Normal Vector
	Vertices [3]Vector
}