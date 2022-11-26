package main

import "fmt"

type Goods struct {
	Kind string
	Fact bool
}

// 接口
type Buy interface {
	BuySomething(*Goods)
}

// 接口实现
type ChinsesBuy struct {
}

func (cb *ChinsesBuy) BuySomething(good *Goods) {
	fmt.Println("去中国买东西: ", good.Kind)
}

type AmericanBuy struct {
}

func (cb *AmericanBuy) BuySomething(good *Goods) {
	fmt.Println("去美国买东西: ", good.Kind)
}

// 代购，对以上接口实现进行功能增强
type OverSeaBuy struct {
	buy Buy
}

func (osb *OverSeaBuy) BuySomething(good *Goods) {
	if osb.distinguish(good) { //buy功能增强一
		osb.buy.BuySomething(good)
		osb.check(good) // buy功能增强二
	}
}

// 代购功能增强一：
func (osb *OverSeaBuy) distinguish(good *Goods) bool {
	fmt.Println("对[", good.Kind, "]进行了辨别真伪.")
	if good.Fact == false {
		fmt.Println("发现假货，[", good.Kind, "]不应该购买")
		return false
	} else {
		return true
	}
}

// 代购功能增强二：
func (osb *OverSeaBuy) check(good *Goods) {
	fmt.Println("对[", good.Kind, "] 进行了安检...")
}

func NewOverSeaBuy(buy Buy) *OverSeaBuy {
	return &OverSeaBuy{buy: buy}
}

// 业务

func main() {
	g1 := Goods{
		Kind: "吃的",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CETS4",
		Fact: false,
	}

	var cbuy Buy
	cbuy = new(ChinsesBuy)
	cbuy.BuySomething(&g1)

	var abuy Buy
	abuy = &AmericanBuy{}
	abuy.BuySomething(&g2)

	var proxyBuy Buy
	proxyBuy = NewOverSeaBuy(cbuy)
	proxyBuy.BuySomething(&g2)
	proxyBuy.BuySomething(&g1)
}
