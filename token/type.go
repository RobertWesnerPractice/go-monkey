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
	Minus
	Multiplication
	Division
	Comma
	Semicolon
	ParenthesisLeft
	ParenthesisRight
	BraceLeft
	BraceRight
	LessThan
	GreaterThan
	LessOrEqual
	GreaterOrEqual
	Not
	Equal
	NotEqual

	// keywords
	True
	False
	Function
	Declaration
	If
	Else
	Return
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
	case Minus:
		return "-"
	case Multiplication:
		return "*"
	case Division:
		return "/"
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
	case LessThan:
		return "<"
	case GreaterThan:
		return ">"
	case LessOrEqual:
		return "<="
	case GreaterOrEqual:
		return ">="
	case Not:
		return "!"
	case Equal:
		return "=="
	case NotEqual:
		return "!="
	case Function:
		return "fn"
	case Declaration:
		return "let"
	case True:
		return "true"
	case False:
		return "false"
	case If:
		return "if"
	case Else:
		return "else"
	case Return:
		return "return"
	default:
		return "unknown"
	}
}
