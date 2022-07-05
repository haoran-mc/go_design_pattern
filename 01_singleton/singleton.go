package singleton

// Singleton 饿汉式单例
type Singleton struct{} // 这是一个类（go语言用结构体表示类）

var singleton *Singleton // 上面类的全局唯一实例

func init() {
	singleton = &Singleton{}
}

// GetInstance 获取实例
func GetInstance() *Singleton {
	return singleton
}
