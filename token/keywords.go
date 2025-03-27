package token

var keywords = map[string]Type{
	"fn":  Function,
	"let": Declaration,
}

// LookupIdentifier returns a keyword Type if found, otherwise returns Identifier
func LookupIdentifier(identifier string) Type {
	if token, ok := keywords[identifier]; ok {
		return token
	}

	return Identifier
}
