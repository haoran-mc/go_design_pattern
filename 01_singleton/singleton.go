package singleton

// 饿汉式单例
type singletonStruct struct{} // 这是一个类（go语言用结构体表示类）

var singleton *singletonStruct // 上面类的全局唯一实例，设为包内私有

// 全局初始化一次
func init() {
	singleton = &singletonStruct{}
}

// 获取单例实例
func GetSingletonInstance() *singletonStruct {
	return singleton
}

// 获取普通实例
func GetInstance() *singletonStruct {
	return &singletonStruct{}
}
