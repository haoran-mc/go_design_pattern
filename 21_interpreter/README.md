解释器模式为某个语言定义它的语法（或者叫文法）表示，并定义一个解释器用来处理这个语法。


假设有一个监控系统，现在需要实现一个告警模块，可以根据输入的告警规则来决定是否触发告警。

告警规则支持 `&&`, `>`, `<` 3种运算符，其中 `>`, `<` 优先级比 `&&` 更高。





