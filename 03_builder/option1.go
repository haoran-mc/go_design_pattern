package builder

import "fmt"

type ResourcePoolConfigOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOptFunc func(option *ResourcePoolConfigOption)

// 传入函数切片，使用函数进行初始化
func NewResourcePoolConfig(name string, opts ...ResourcePoolConfigOptFunc) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}

	option := &ResourcePoolConfigOption{
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
	}

	for _, opt := range opts {
		opt(option) // opt 函数处理所有的可配置项 option
	}

	if option.maxTotal < 0 || option.maxIdle < 0 || option.minIdle < 0 {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	if option.maxTotal < option.maxIdle || option.minIdle > option.maxIdle {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	return &ResourcePoolConfig{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}
