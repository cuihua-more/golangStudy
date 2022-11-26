package main

import "fmt"

// ---------------抽象层------------------
type Phone interface {
	Show()
}

// 本应定义为接口，但是golang语法限制interface不能继承，只能定义结构体
type Decorator struct {
	phone Phone
}

// func (d *Decorator) Show() {
// }

// ---------------实现层--------------------
type Hawei struct {
}

func (this *Hawei) Show() {
	fmt.Println("秀出了华为手机")
}

type Xiaomi struct {
}

func (this *Xiaomi) Show() {
	fmt.Println("秀出了小米手机")
}

type MoDecorator struct {
	Decorator
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("...贴了膜")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone: phone}}
}

type KeDecoreator struct {
	Decorator
}

func (nk *KeDecoreator) Show() {
	nk.phone.Show()
	fmt.Println("...带了壳")
}

func NewKeDecoreator(phone Phone) Phone {
	return &KeDecoreator{Decorator{phone: phone}}
}

// 业务逻辑层

func main() {
	var xiaomi Phone
	xiaomi = new(Xiaomi)
	xiaomi.Show()

	var huawei Phone
	huawei = new(Hawei)
	huawei.Show()

	fmt.Println("-------------")
	var moXiaomi Phone
	moXiaomi = NewMoDecorator(xiaomi)
	moXiaomi.Show()

	fmt.Println("-------------")
	var kehuawei Phone
	kehuawei = NewKeDecoreator(huawei)
	kehuawei.Show()

	fmt.Println("-------------")
	var mokexiaomi Phone
	mokexiaomi = NewKeDecoreator(moXiaomi)
	mokexiaomi.Show()
}
