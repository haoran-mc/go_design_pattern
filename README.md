# Go 设计模式

| 类型     | 设计模式                                                                             | 是否常用   |
| -------- | ------------------------------------------------------                               | ---------- |
| 创建型   | [单例模式](./01_singleton)（Singleton Design Pattern）                               | ✅         |
|          | [工厂模式](./02_factory)（Factory Design Pattern）                                   | ✅         |
|          | [建造者模式](./03_builder)（Builder Design Pattern）                                 | ✅         |
|          | [原型模式](./04_prototype)（Prototype Design Pattern）                               | ❌         |
| 结构型   | [代理模式](./05_proxy)（Proxy Design Pattern）                                       | ✅         |
|          | [桥接模式](./06_bridge)（Bridge Design Pattern）                                     | ✅         |
|          | [装饰器模式](./07_decorator)（Decorator Design Pattern）                             | ✅         |
|          | [适配器模式](./08_adapter)（Adapter Design Pattern）                                 | ✅         |
|          | [门面模式](./09_facade)（Facade Design Pattern）                                     | ❌         |
|          | [组合模式](./10_composite)（Composite Design Pattern）                               | ❌         |
|          | [享元模式](./11_flyweight)（Flyweight Design Pattern）                               | ❌         |
| 行为型   | [观察者模式](./12_observer)（Observer Design Pattern）                               | ✅         |
|          | [模板模式](./13_template_method)（Template Method Design Pattern）                   | ✅         |
|          | [策略模式](./14_strategy)（Strategy Method Design Pattern）                          | ✅         |
|          | [职责链模式](./15_chain_of_responsibility)（Chain Of Responsibility Design Pattern） | ✅         |
|          | [状态模式](./16_state)（State Design Pattern）                                       | ✅         |
|          | [迭代器模式](./17_iterator)（Iterator Design Pattern）                               | ✅         |
|          | [访问者模式](./18_visitor)（Visitor Design Pattern）                                 | ❌         |
|          | [备忘录模式](./19_memento)（Memento Design Pattern）                                 | ❌         |
|          | [命令模式](./20_command)（Command Design Pattern）                                   | ❌         |
|          | [解释器模式](./21_interpreter)（Interpreter Design Pattern）                         | ❌         |
|          | [中介模式](./22_mediator)（Mediator Design Pattern）                                 | ❌         |
|          | [空对象模式](./23_null_object)（Null Object Design Pattern）                         | ❌         |

## 代理、桥接、装饰、适配的区别

这 4 种模式是比较常用的结构型设计模式，它们的代码结构非常相似，笼统来说，它们都可以称为 wrapper 模式，也就是通过 wrapper 类二次封装原始类。

尽管代码结构相似，但这 4 种设计模式的用意完全不同，也就是说要解决的问题、应用场景不同，这也是它们的主要区别。

- **代理模式**：代理模式在不改变原始类接口的条件下，为原始类定义一个代理类，主要目的是控制访问，而非加强功能，这是它跟装饰器模式最大的不同。
- **桥接模式**：桥接模式的目的是将接口部分和实现部分分离，从而让它们可以较为容易、也相对独立地加以改变。
- **装饰器模式**：装饰者模式在不改变原始类接口的情况下，对原始类功能进行增强，并且支持多个装饰器的嵌套使用。
- **适配器模式**：适配器模式是一种事后的补救策略。适配器提供跟原始类不同的接口，而代理模式、装饰器模式提供的都是跟原始类相同的接口。

## 注意：

> 忌过度设计！设计模式不是银弹，不要拿着 🔨 就觉得哪里都像是钉子，不要过早优化，持续重构才是正道。

## 参考：

- [Graphic Design Patterns](https://design-patterns.readthedocs.io/zh_CN/latest/index.html)
- [图说设计之美](https://time.geekbang.org/column/intro/100039001?tab=catalog)
- [Go设计模式](https://lailin.xyz/post/go-design-pattern.html)
- [All Design Patterns in Go (Golang)](https://golangbyexample.com/all-design-patterns-golang/)

