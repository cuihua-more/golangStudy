package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // 表示SuperMan类继承于Human类

	level int
}

// 重定义父类方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类添加新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{name: "Zhang3", sex: "female"}

	h.Eat()
	h.Walk()

	fmt.Println("--------------------------------")
	//s := SuperMan{Human{"Li3", "female"}, 88} //定义除此之外，还有以下方式定义：
	var s SuperMan
	s.name = "Li4"
	s.sex = "female"
	s.level = 88

	s.Walk() // 调用父类的方法
	s.Eat()  // 调用子类的方法
	s.Fly()  // 调用子类的方法
	s.Print()
}
