# Truelang 开发日志

本项目为阅读《用 Go 语言自制解释器》时，实现的脚本语言解释器，以此来根据实践来学习编译原理。

## 词法分析

第一次要解析的代码

```
    var five = 5;
    var ten = 10;
    var add = fn(x, y) {
        x + y;
    };
    var result = add(five, ten);
```

【进阶】：词法解析器还需要解析 行号、列号、文件名等信息，方面后续的错误处理，提供给用户更多信息。
分析有哪些词法单元？

- 字面量（数字）：5、10
- 变量：five、ten、add、result
- 关键字：var、fn
- 标识符：分号、逗号、括号、花括号、加号、等号

根据以上分析，我们现在需要定义**词法单元 Token** 的数据结构，用来表示词法分析器的输出结果。

```
type TokenType string

// Token 词法单元
type Token struct {
	Type    TokenType
	Literal string
}
```

【进阶】：TokenType 现阶段使用 string 实现没有问题，但是会有性能损耗，后续可以使用 int/byte 类型优化

```go
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
```

【进阶】从`readChar()` 可以看出，当前解释器仅支持 ASCII 字符，支持 UTF-8 字符需要进行优化。

在当前版本的 NextToken() 实现后，已经可以初步解析如下测试用例中 `input` 中的内容，显然作为一个脚本语言的解释器，做的还不够，我们还需要 Lexer 支持更多的词法单元。

```go
// lexer/lexer_test.go
func TestNextToken(t *testing.T) {
    input :=`=+(){},;`
    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.ASSIGN, "="},
        {token.PLUS, "+"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RBRACE, "}"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }
...
}
```

目前我们这种开发模式，可以称之为测试驱动开发（TDD），编写足够多的单元测试，有助于代码的重构和后续维护

```go
// lexer/lexer_test.go
func TestNextToken_2(t *testing.T) {
	input :=
		`
        var five = 5;
        var ten = 10;
        var add = fn(x, y) {
            x + y;
        };
        var result = add(five, ten);
    `
```

在新增的测试用例中，我们需要 Lexer 支持 关键字 var、fn，以及标识符 five、ten、add、result

```go
default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdentifier()
			return tk
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
```

NextToken()中添加 default 分支后，已经支持了关键字和标识符，但是 Lexer 并不能将二者区分开。

将整数视为词法单元
