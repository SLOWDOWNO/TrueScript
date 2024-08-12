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

【进阶内容】：词法解析器还需要解析 行号、列号、文件名等信息，方面后续的错误处理，提供给用户更多信息。
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

【进阶内容】：TokenType 现阶段使用 string 实现没有问题，但是会有性能损耗，后续可以使用 int/byte 类型优化
