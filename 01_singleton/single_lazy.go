package singleton

import "sync"

// 懒汉式单例
type lazySingletonStruct struct{}

var (
	lazySingleton *lazySingletonStruct
	once          = &sync.Once{}
)

// 获取单例实例
func GetLazySingletonInstance() *lazySingletonStruct {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &lazySingletonStruct{}
		})
	}
	return lazySingleton
}

// 获取普通实例
func GetLazyInstance() *lazySingletonStruct {
	return &lazySingletonStruct{}
}
