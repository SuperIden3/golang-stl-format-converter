package converter

// Keep track of the position in the file for error reporting
type Position struct {
	Line   uint
	Column uint
}

type Lexer interface {
	NextFloat32() (float32           , error)
	NextToken()   (NextToken_Return  , error)
	ReadAll()     ([]NextToken_Return, error)
	Close()
	IsClosed()    bool
	GetPosition() Position
}
