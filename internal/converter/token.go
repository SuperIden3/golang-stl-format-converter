package converter

type Token uint

const (
	ILLEGAL Token = iota
	EOF
	SOLID
	NUMBER
	FACET
	NORMAL
	OUTER_LOOP // "outer loop" becomes one token
	VERTEX
	END_LOOP
	END_FACET
	END_SOLID
)

func TokenMap() map[string]Token {
	return map[string]Token {
		"solid": SOLID,
		"facet": FACET,
		"normal": NORMAL,
		"outer": OUTER_LOOP,
		"vertex": VERTEX,
		"endloop": END_LOOP,
		"endfacet": END_FACET,
		"endsolid": END_SOLID,
	}
}

func (t Token) String() string {
	return []string{ "ILLEGAL", "EOF", "SOLID", "NUMBER", "FACET",  "NORMAL", "OUTER_LOOP", "VERTEX", "END_LOOP",  "END_FACET", "END_SOLID", }[t]
}

