package token

var keywords = map[string]Type{
	"true":   True,
	"false":  False,
	"fn":     Fn,
	"let":    Let,
	"if":     If,
	"else":   Else,
	"return": Return,
}

// LookupIdentifier returns a keyword Type if found, otherwise returns Identifier
func LookupIdentifier(identifier string) Type {
	if token, ok := keywords[identifier]; ok {
		return token
	}

	return Identifier
}
