package main

import "fmt"

// 通过interface关键字来达到多态, interface本质是一个指针
// 本质上是定义一个统一的接口，不同的类继承后有不同的实现，叫多态
// 这个AnimalIF是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  //获取动物的种类
}

// 定义一个具体类，不需要手动将AnimalIF写进去继承，只需要实现其函数即可，前提是所有接口全部实现
type Cat struct {
	Color string // 猫的颜色
}

// 实现AnimalIF中的Sleep()方法表示继承于AnimalIF类
func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.Color
}

func (this *Dog) GetType() string {
	return "Cat"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	// 多态第一种使用方式
	var animal AnimalIF // 接口的数据类型，父类指针

	animal = &Cat{"Red"}
	animal.Sleep() // 调用的是Cat的方法

	animal = &Dog{"Yellow"}
	animal.Sleep() // 调用的Dog的方法

	// 多态第二种使用方式
	cat := Cat{"Black"}
	dog := Dog{"Yellow"}
	fmt.Println("---------------------------------------")
	showAnimal(&cat) // 取地址，默认会传继承的AnimalIF类指针
	fmt.Println("---------------------------------------")
	showAnimal(&dog)

}
