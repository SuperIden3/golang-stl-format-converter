package converter

import (
	"os"
	"bufio"
	"fmt"
)

type Lexer struct {
	reader *bufio.Reader
}

// New Lexer from a given file object
func NewLexer(file *os.File) *Lexer {
	return &Lexer{
		reader: bufio.NewReaderSize(file, 8),
	}
}

// "Close" the Lexer by nulling the reader
func (l *Lexer) Close() {
	l.reader = nil
}

// Check if the Lexer is closed
func (l *Lexer) IsClosed() bool {
	return nil == l.reader
}

// Return an error if the Lexer is closed
func closedError() error { return fmt.Errorf("Lexer is closed") }

// Return the next word
func (l *Lexer) NextWord() (string, error) {
	if l.IsClosed() { // Check if closed
		return "", closedError()
	}

	// Prepare: word for return, err for error, and currentByte for tracking current byte
	var word string
	var err error
	var currentByte byte

	for {
		currentByte, err = l.reader.ReadByte() // Read byte and catch errors

		if nil != err { return "", err } // Return empty string and error

		if currentByte == ' ' || currentByte == '\n' { break } // If space or newline, break

		word += string(currentByte) // Append byte to word
	}

	return word, nil
}

func (l *Lexer) NextToken() (NextToken_Return, error) {
	if l.IsClosed() {
		return NextToken_Return{}, closedError()
	}

	word, err := l.NextWord() // Get word
	
	if nil != err { return NextToken_Return{}, err } // Return empty struct and error

	mapping, _ := TokenMappings() // Get mapping
	return NextToken_Return{ Token: mapping[word] }, nil // Return token
}

func AddToken(token_arr []NextToken_Return, token NextToken_Return) error {
	new_token_arr = append(token_arr, token)
	if nil != new_token_arr {
		token_arr = new_token_arr
		return nil
	}
	return fmt.Errorf("Failed to add token")
}

func (l *Lexer) ReadAll() ([]NextToken_Return, error) {
	if l.IsClosed() { // Check if closed
		return []NextToken_Return{}, closedError()
	}

	var tokens []NextToken_Return
	var err error

	for {
		token, err := l.NextToken() // Get token

		if nil != err { return []NextToken_Return{}, err } // Return empty struct and error

		if token.Token == EOF { break } // If EOF, break

		tokens, err = AddToken(tokens, token) // Append token to token array
		if nil != err { return []NextToken_Return{}, err }
	}

	return tokens, nil
}
