package singleton

import "sync"

// lazySingletonStruct 懒汉式单例
type lazySingletonStruct struct{}

var (
	lazySingleton *lazySingletonStruct
	once          = &sync.Once{}
)

// GetLazySingletonInstance 获取单例实例
func GetLazySingletonInstance() *lazySingletonStruct {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &lazySingletonStruct{}
		})
	}
	return lazySingleton
}

// GetLazyInstance 获取普通实例
func GetLazyInstance() *lazySingletonStruct {
	return &lazySingletonStruct{}
}
