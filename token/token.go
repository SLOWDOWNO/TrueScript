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
	// 标识符+字面量
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
	LET      = "LET"
)