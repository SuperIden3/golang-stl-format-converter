package converter

import (
	"os"
	"bufio"
	"strings"
)

type Lexer struct {
	scanner bufio.Scanner
}

func NewLexer(file *os.File) *Lexer {
	return &Lexer{
		scanner: bufio.NewScanner(file),
	}
}

