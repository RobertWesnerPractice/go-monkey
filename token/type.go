package token

type Type int

//goland:noinspection GoCommentStart
const (
	Illegal Type = iota
	EOF
	Identifier
	Number
	Assignment
	Plus
	Comma
	Semicolon
	ParenthesisLeft
	ParenthesisRight
	BraceLeft
	BraceRight

	// keywords
	Function
	Declaration
)

func (t Type) Debug() string {
	switch t {
	case Illegal:
		return "illegal"
	case EOF:
		return "end of file"
	case Identifier:
		return "identifier"
	case Number:
		return "number"
	case Assignment:
		return "="
	case Plus:
		return "+"
	case Comma:
		return ","
	case Semicolon:
		return ";"
	case ParenthesisLeft:
		return "("
	case ParenthesisRight:
		return ")"
	case BraceLeft:
		return "{"
	case BraceRight:
		return "}"
	case Function:
		return "fn"
	case Declaration:
		return "let"
	default:
		return "unknown"
	}
}
