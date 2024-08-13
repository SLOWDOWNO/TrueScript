package lexer

import "True/token"

// Lexer 词法分析器
type Lexer struct {
	input   string
	pos     int
	readPos int // 始终指向当前字符的下一个字符
	ch      byte
}

// New 初始化词法分析器
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar 读取Lexer::input中的下一个字符
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

// NextToken 根据 l.ch 返回对应的语法单元
func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.skipWhitespace()

	switch l.ch {
	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdentifier()
			tk.Type = token.LookupIdent(tk.Literal)
			return tk
		} else if isDigit(l.ch) {
			tk.Type = token.INT
			tk.Literal = l.readNumber()
			return tk
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
	case '=':
		tk = newToken(token.ASSIGN, l.ch)
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case '(':
		tk = newToken(token.LPAREN, l.ch)
	case ')':
		tk = newToken(token.RPAREN, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF
	}

	l.readChar()
	return tk
}

// readIdentifier 读取标识符
func (l *Lexer) readIdentifier() string {
	pos := l.pos // 记录标识符的起始位置
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

// readNumber 读取数字
func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// newToken 创建一个新的Token,是 NextToken 的辅助函数
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

//  isLetter 判断输入的 ch 是不是字母
func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

// isDigit 判断输入的 ch 是不是数字
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
