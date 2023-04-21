package adapter

import "fmt"

// 电脑接口
type computer interface {
	insertInSquarePort()
}

// macbook
type mac struct{}

// mac 插入方形usb
func (m *mac) insertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

// windows
type windows struct{}

// windows 插入圆形usb
func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

// -----------------------------

type client struct{}

// client 只能插入 square 类型的 usb
func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}

// -----------------------------

// windows 适配器
type windowsAdapter struct {
	windowMachine *windows
}

// 做了适配，方形的usb也能插入windows上了
func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}
