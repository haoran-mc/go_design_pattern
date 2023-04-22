package observer

import "fmt"

// interface 商家
type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

// 商品
type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

// 上架商品
func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

// 新的订阅用户
func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

// 订阅用户解除订阅
func (i *item) deregister(o observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

// 发送通知
func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

// 删除不友好的用户
func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// interface 用户
type observer interface {
	update(string)
	getID() string
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
