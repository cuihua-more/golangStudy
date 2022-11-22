package main

import "fmt"

// 类名首字母大写，其他包能够访问
type Hero struct {
	// 属性名首字符大写，表示公有属性，其他包可以访问；小写，表示私有属性，其他包不能访问
	Name  string
	Ad    int
	Level int
}

// (this Hero) 表示这个Show函数是Hero类的一个方法，this参数表示调用对象的值内存拷贝
func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(Name string) {
	/*
	* 注意，这样并不会就该这个对象的值，因为this传递的是该调用对象的值拷贝
	 */
	this.Name = Name
}

func (this *Hero) SetName1(Name string) {
	/*
	* 注意。这个this是传递的指针，传递的是该调用对象的地址，修改变量可以达到想要效果
	 */
	this.Name = Name
}

func main() {
	// 创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}

	hero.Show()

	fmt.Println("---------------------------")
	hero.SetName("Li4")
	hero.Show() // 注意Name，仍是原来的

	fmt.Println("---------------------------")
	hero.SetName1("55555")
	hero.Show() // 注意Name，在这修改成功
}
