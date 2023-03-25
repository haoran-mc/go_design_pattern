# 需求样例

假设我们需要根据配置文件的文件名来解析一个配置文件，不同格式的文件使用不同的解析器解析。

如：json 格式的文件使用 json 解析器；yaml 格式的文件使用 yaml 解析器...

# 不使用工厂模式

[no_factory.go](./020_no_factory/no_factory.go)

`Load` 函数是我们的加载函数，放入一个配置文件，系统就会找到并加载这个配置文件。

# 简单工厂

[simple_factory.go](./021_simple_factory/simple_factory.go)

使用一个简单工厂模式，将创建解析器的过程单独封装起来。

`NewConfigParser` 用于创建一个解析器，它也是我们所说的工厂。

> 不是很难理解，就是把一个返回接口的地方封装起来。

# 工厂方法

[factory_method.go](./022_factory_method/factory_method.go)

假如有十几种解析器，每种解析器又很复杂，那么可能这个简单工厂要成百上千行，显然不合理。

所以在大工厂里创建多个小工厂，对于每种解析器的创建，封装成一个个单独的小工厂。

> 如果解析器的种类不多，且每种解析器的创建很简单，那就直接使用「简单工厂」就好了，
> 反之使用工厂方法。

# 抽象工厂

[abstract_factory.go](./023_abstract_factory/abstract_factory.go)

不常用，假设这样一种场景，配置文件可能有不同的格式（json, yaml...），而且这个配置文件的服务类型也不一样（汽车的配置文件、电脑的配置文件、台灯的配置文件...）。

## 比较工厂方法与抽象工厂

在这种场景下，比较工厂方法 [factory_method.go](./023_abstract_factory/functory_method/factory_method.go) 与抽象工厂 [abstract_factory.go](./023_abstract_factory/abstract_factory.go)。

> 好像差不多...

## 自顶向下观察抽象工厂

1. 在 `Load` 函数中是怎样使用工厂的？

```go
parserFactory := NewParserFactory(configFilePath) // 获取解析器工厂
parser := parserFactory.CreateCarParser() // 这里我们指定这是一个汽车的配置文件
Parser.ParseCar(configContent) // 解析配置文件的内容
```

2. 首先我们想要获取一个解析器工厂

下面是工厂函数：

```go
// 工厂
func NewConfigParserFactory(fileExt string) ParserFactory {
    switch fileExt {
    case "json":
        return jsonParserFactory{}
    case "yaml":
        return yamlParserFactory{}
    case "toml":
        return tomlParserFactory{}
    case "ini":
        return iniParserFactory{}
    case "conf":
        return confParserFactory{}
    }
    return nil
}
```

通过工厂函数，我们得到一个「解析器工厂」，因为我们还不知道这个配置文件是用来配置什么的，所以还需要这个 jsonParserFactory（也不一定，我们假设配置文件的文件类型是 json）来创建（对滴，我们还没有得到具体的解析器，所以需要创建）一个具体的解析器。

```go
// 工厂返回的是一个接口
type ConfigParserFactory interface {
    CreateCarParser() CarParser
    CreateComputerParser() ComputerParser
    CreateLampParser() LampParser
}

// 接口的不同实现（数目、类型和工厂保持一致）

// json
// yaml
// toml
// ini
// conf

// 我们只举一个 json 的例子
type jsonParserFactory struct{}

func (j jsonParserFactory) CreateCarParser() CarParser {
    return jsonCarParser{}
}

func (j jsonParserFactory) CreateComputerParser() ComputerParser {
    return jsonComputerParser{}
}

func (j jsonParserFactory) CreateLampParser() LampParser {
    return jsonLampParser{}
}
```

`Load` 中看到，我们需要确定是什么物品的解析器，这里是汽车，所以 `Load` 进行了以下的调用：

```go
// 配置文件里是车轮的尺寸等等
parser := parserFactory.CreateCarParser()
```

3. 具体的解析器又是怎么实现的呢？

我们看看汽车解析器吧。

```go
// 汽车解析器
type CarParser interface {
    Parser(data []byte)
}

type jsonCarParser struct{}

func (j jsonCarParser) Parser(data []byte) {
  // 解析过程
}

type yamlCarParser struct{}

func (y yamlCarParser) Parser(data []byte) {
  // 解析过程
}

// toml
// ini
```
