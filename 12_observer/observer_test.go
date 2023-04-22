package observer

import "testing"

func TestObserver(t *testing.T) {
	shirtItem := newItem("Nike Shirt")               // 商品：衬衫
	observerFirst := &customer{id: "abc@gmail.com"}  // 用户1
	observerSecond := &customer{id: "xyz@gmail.com"} // 用户2
	shirtItem.register(observerFirst)                // 订阅
	shirtItem.register(observerSecond)               // 订阅
	shirtItem.updateAvailability()                   // 商品发送通知
}
