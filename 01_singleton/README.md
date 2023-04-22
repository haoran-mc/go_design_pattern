# 1. 饿汉单例

[singleton.go](./singleton.go)

包内私有，然后通过函数获取（就是这么简单）。

# 2. 懒汉单例

[single_lazy.go](./single_lazy.go)

使用的时候才创建，用到 sync.Once{}
