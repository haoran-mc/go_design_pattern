package builder

import "fmt"

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

// 资源池配置类
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// 用于构建资源池的配置类
type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// 设置资源池配置类的组成部分，这里设置名称（name）
func (b *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name can not be empty")
	}
	b.name = name
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("max tatal cannot < 0, input: %d", minIdle)
	}
	b.minIdle = minIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max tatal cannot < 0, input: %d", maxIdle)
	}
	b.maxIdle = maxIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max tatal cannot <= 0, input: %d", maxTotal)
	}
	b.maxTotal = maxTotal
	return nil
}

// Build 使用构建类构建
// 组装这些部件可能会很麻烦，所以使用建造者在这里统一组装，可能遇到的麻烦：
// 部件存在限制条件、部件的赋值必须按照某个顺序、一个属性没有赋值之前，另一个属性可能无法赋值等
func (b *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if b.name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}

	// 设置默认值
	if b.minIdle == 0 {
		b.minIdle = defaultMinIdle
	}

	if b.maxIdle == 0 {
		b.maxIdle = defaultMaxIdle
	}

	if b.maxTotal == 0 {
		b.maxTotal = defaultMaxTotal
	}

	if b.maxTotal < b.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", b.maxTotal, b.maxIdle)
	}

	if b.minIdle > b.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", b.maxIdle, b.minIdle)
	}

	return &ResourcePoolConfig{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}, nil
}
