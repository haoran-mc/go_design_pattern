package builder

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type Cluster struct {
	// 把所有参数另用一个类存储
	opts options
}

type options struct {
	connectionTimeout time.Duration
	readTimeout       time.Duration
	writeTimeout      time.Duration
	logError          func(ctx context.Context, err error)
}

// 给闭包函数取一个名字
// 闭包：能够捕获其他函数内部变量的函数
type Option func(c *options)

// 将调用时的 d 捕获，然后在构造函数中使用
func ConnectionTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.connectionTimeout = d
	}
}

func ReadTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.readTimeout = d
	}
}

func WriteTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.writeTimeout = d
	}
}

func LogError(f func(ctx context.Context, err error)) Option {
	return func(opts *options) {
		opts.logError = f
	}
}

// NewCluster 构造函数具体实现，传入相关Option，new一个对象并赋值
// 如果参数很多，也不需要传入很多参数，只需要传入 opts ...Option 即可
func NewCluster(opts ...Option) *Cluster {
	clusterOpts := options{} // 会进入所有闭包，一遍又一遍赋值
	for _, opt := range opts {
		opt(&clusterOpts) // 函数指针的赋值调用
	}
	// 在这里处理依赖关系等复杂逻辑
	// 此时 clusterOpts 已经是经过内部检验的 cluster 参数了
	cluster := new(Cluster)
	cluster.opts = clusterOpts // 经过检验的，放心赋值
	return cluster
}

func main() {
	// 储备可配置项
	commonsOpts := []Option{
		ConnectionTimeout(1 * time.Second),
		ReadTimeout(2 * time.Second),
		WriteTimeout(3 * time.Second),
		LogError(func(ctx context.Context, err error) {
		}),
	}

	// 构造函数
	cluster := NewCluster(commonsOpts...)

	// 测试验证
	fmt.Println(cluster.opts.connectionTimeout)
	fmt.Println(cluster.opts.writeTimeout)
}
