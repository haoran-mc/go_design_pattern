package singleton

// singleton 饿汉式单例
type singletonStruct struct{} // 这是一个类（go语言用结构体表示类）

var singleton *singletonStruct // 上面类的全局唯一实例

func init() {
	singleton = &singletonStruct{}
}

// GetSingletonInstance 获取单例实例
func GetSingletonInstance() *singletonStruct {
	return singleton
}

// GetInstance 获取普通实例
func GetInstance() *singletonStruct {
	return &singletonStruct{}
}
