# 1. 建造者模式与工厂模式的区别

工厂模式是用来创建不同但是相关类型的对象（继承同一接口的一组子类），由给定的参数来决定创建哪种类型的对象。建造者模式是用来创建一种类型的复杂对象，通过设置不同的可选参数，“定制化”地创建不同的对象。

# 2. 建造者模式

一个类的可配置项有很多，而且配置项之间的依赖关系复杂，那么一次性传入所有可配置项（代码可读性和易用性变差）和通过 `set()` 函数来设置（无法处理依赖关系）都不是好的选择。

这时我们可以先把所有参数放入构建类，然后通过构建类的 `Build` 方法统一配置所有可配置项。

```go
type ResourcePoolConfig struct {
  // ...
}

type ResourcePoolConfigBuilder struct {
  // ...
}

func (b *ResourcePoolConfigBuilder) SetConf1(conf string) {}
func (b *ResourcePoolConfigBuilder) SetConf2(conf string) {}
func (b *ResourcePoolConfigBuilder) SetConf3(conf string) {}
func (b *ResourcePoolConfigBuilder) SetConf4(conf string) {}

func (b *ResourcePoolConfigBuilder) Build(name string) (*ResourcePoolConfig) {
  // 处理依赖关系等复杂的逻辑
}
```

# 3. Go 语言中对于建造者模式更好的实践

*闭包*

[option1.go](./option1.go)

使用：

```go
ResourcePoolConfig, _ := NewResourcePoolConfig("config_name",
    []ResourcePoolConfigOptFunc{
        func(option *ResourcePoolConfigOption) {
            option.minIdle = 2
        },
        func(option *ResourcePoolConfigOption) {
            option.maxIdle = 3
        },
    }
)
```

# 4. 使用闭包实现 Go 语言中的构造函数

[option2.go](./option2.go)

