访问者模式

假设你维护一个 lib，其中包含下面三种形状：

- Square
- Circle
- Triangle

这些形状都实现了 shape 这个接口。

现在，有许多团队使用了你的 lib，这些团队希望你能为这些形状添加一个方法 getArea()。

**方法一** ，shape 这个接口添加一个 getArea() 方法，然后这些形状都实现这个方法。

- 作为库的维护者，你不希望通过添加其他行为来更改高度测试的库代码。
- 使用这个库的团队可能会对更多的行为提出更多的请求，例如 getNumSides()、 getMiddleConcerates()。在这种情况下，你更希望其他团队在不对代码进行太多实际修改的情况下扩展你的库。

**方法二** ，让团队自己根据行为本身编写逻辑。

```go
if shape.type == square {
   //Calculate area for squre
} elseif shape.type == circle {
    //Calculate area of triangle 
} elseif shape.type == "triangle" {
    //Calculate area of triangle
} else {
   //Raise error
} 
```

但是这样你将不能充分利用接口的优势，只是做一个显式的类型检查。其次，在运行时获取类型可能会对性能产生影响，在某些语言中甚至没有提供这项功能。

**方法三** ，访问者模式。

1. 定义一个访问者接口

```go
type visitor interface {
   visitForSquare(square)
   visitForCircle(circle)
   visitForTriangle(triangle)
}
```

visitForSquare(square), visitForCircle(circle), visitForTriangle(triangle) 支持我们分别为 square, circle, triangle 添加函数。

<!-- Now the question which comes to mind is why can’t we have a single method visit(shape) in the visitor interface. The reason we don’t have it because GO and also some other languages support method overloading. So a different method for each of the struct. -->

为什么不在 visitor 接口中使用一个方法 visit(shape) 代替呢？因为 Go 语言不支持函数重载。

2. 为接口添加一个 accept 方法，所有形状实现这个 accept 方法

```go
func accept(v visitor)
```

```go
func (obj *squre) accept(v visitor){
    v.visitForSquare(obj)
}
```

这也是我们唯一一次需要修改每个形状结构体的地方，而且只需要修改一次。

3. 添加需要添加的功能 

需要定制功能的团队应该实现 visitor 接口，在具体的实现中完成面积计算的逻辑。

```go
type areaCalculator struct{
    area int
}

func (a *areaCalculator) visitForSquare(s *square){
    // Calculate area for square
}
func (a *areaCalculator) visitForCircle(s *square){
    // Calculate area for circle
}
func (a *areaCalculator) visitForTriangle(s *square){
    // Calculate area for triangle
}
```





