package token

type TokenType string

// Token 词法单元
type Token struct {
	Type    TokenType
	Literal string
}

// Token 的类型
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// 标识符（变量）
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	// 运算符
	ASSIGN = "="
	PLUS   = "+"
	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	// 括号
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// 关键字
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"var": VAR,
}

// lookupIdent 判断输入的 ident 是否是关键字
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
