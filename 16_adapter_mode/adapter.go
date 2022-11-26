package main

import "fmt"

// 接口
type V5 interface {
	Use5V()
}

// 被适配的类
type V220 struct {
}

func (this *V220) Use220V() {
	fmt.Println("输出220V电")
}

// 适配器
type Adapter struct {
	v220 *V220
}

func (ad *Adapter) Use5V() {
	ad.v220.Use220V()
	fmt.Println("转化为5V")
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}

// 业务
type Phone struct {
	v V5
}

func (this *Phone) Charge() {
	fmt.Println("Phone 进行了充电")
	this.v.Use5V()
}

func NewPhone(v V5) *Phone {
	phone := &Phone{v: v}

	return phone
}

func main() {
	phone := NewPhone(NewAdapter(new(V220)))
	phone.Charge()
}
