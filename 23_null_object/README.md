空对象模式（Null Object Pattern）不属于 GoF 设计模式，但是它作为一种经常出现的模式足以被视为设计模式了。

其具体定义为设计一个空对象取代 NULL 对象实例的检查。

空对象模式的重要组件：

- Entity 定义了子结构必须实现的基本操作的接口
- ConcreteEntity 实现了 Entity 接口
- NullEntity 表示空对象，它也实现 Entity 接口，但是不具有属性


假设我们有一个学院，学院里有许多部门，每个部门有一些教授。

```go
type department interface {
    getNumberOfProfessors() int
    getName() string
}
```

```go
type college struct {
    departments []department
}
```

有一个机构想要计算特定学院特定系的教授总数。

- 机构不在乎大学里有没有特定的部门。如果大学中不存在这个系，则 college 返回空系对象；
- 机构对待 nullDepartment 和 realDepartment 是一样的，因此可以避免 null 检查；








