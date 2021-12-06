package main

import "fmt"

/**
 OOP 面向对象
**/

type Hero struct {
	Name string
	Ad int
	Level int
}

func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this Hero) GetName() string { // this 是一个变量，一个Hero的拷贝
	return this.Name
}

func (this *Hero) SetName(newName string) { // this 是一个引用
	this.Name = newName
}

func main2() {
	hero := Hero{
		Name: "zhang3",
		Ad: 200,
		Level: 1,
	}

	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
