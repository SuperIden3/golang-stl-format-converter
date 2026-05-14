package converter

import (
	"bufio"
	"encoding/binary"
)

type BinaryLexer struct {
	reader   *bufio.Reader
	position Position
}

// New BinaryLexer from a given file object
func NewBinaryLexer(file *os.File) *BinaryLexer {
	l := &BinaryLexer{
		reader:   bufio.NewReader(file),
		position: Position{Line: 1, Column: 1},
	}
	return l
}