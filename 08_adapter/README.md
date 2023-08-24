适配器模式

两种类型的电脑：

- macbook pro
- windows laptop

macbook pro 的 usb 接口是方形的，windows laptop 的 usb 接口是圆形的，现在 client 有方形的 usb 线缆，那我们怎么让 client 既能用在 macbook pro 上，又能让 client 用在 windows laptop 上？

这时需要我们掏出适配器（类似手机转接头）接在 windows 上，这样就能在 windows 上正常使用 client 了。

```go
// windows 适配器
type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}
```
