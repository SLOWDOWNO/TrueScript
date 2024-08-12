package lexer

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

// readChar 读取下一个字符
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}