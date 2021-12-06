package main

import "fmt"

/**
	OOP面向对象：多态
**/

// 本质上是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类 Cat
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类 Dog
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func main4() {
	// 多态用法一：
	var animal AnimalIF // 定义了一个接口数据类型，是一个指针
	animal = &Cat{"green"}
	animal.Sleep()  // 调用的就是Cat的sleep()方法

	animal = &Dog{"Yello"}
	animal.Sleep()

	// 多态用法二：
	cat := Cat{"white"}
	dog := Dog{"Yello"}
	showAnimal(&cat) // ！！！interface 接口一定是一个指针
	showAnimal(&dog)
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() // 多态
	fmt.Println("Color = ", animal.GetColor())
	fmt.Println("Kind = ", animal.GetType())
}


