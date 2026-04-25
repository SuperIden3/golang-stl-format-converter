package converter

import (
	"bufio"
	"strings"
)

type Lexer struct {
	scanner *bufio.Scanner
}

// NextLine returns the Token type and any numeric data found on that line
func (l *Lexer) NextLine() (converter.Token, []float64, error) {
	if !l.scanner.Scan() {
		return converter.EOF, nil, nil
	}

	line := strings.TrimSpace(l.scanner.Text())
	words := strings.Fields(strings.ToLower(line))

	if len(words) == 0 {
		return converter.ILLEGAL, nil, nil
	}

	// Identify the token based on the first word(s)
	token := l.identifyToken(words)

	// Extract numbers if it's a data-heavy line
	var data []float64
	if token == converter.VERTEX || token == converter.NORMAL {
		data = l.extractCoords(words)
	}

	return token, data, nil
}

