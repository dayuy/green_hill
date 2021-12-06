package main

/**
	OOP面向对象：继承
**/

import "fmt"

type Human struct {
	name string
	sex string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}


//========

type SuperMan struct {
	Human // SuperMan类继承了Human类的方法
	level int
}

// 重定义父类的Eat方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)

}

func main3() {
	h := Human{
		"zhang3",
		"female",
	}

	h.Eat()
	h.Walk()

	// 继承
	s := SuperMan{Human{"zhang3", "female"}, 3}
	s.Eat()
	s.Fly()
	s.Walk()

	var s2 SuperMan
	s2.name = "li4"
	s2.sex = "male"
	s2.level = 88
	s2.Walk()
	s2.Eat()
	s2.Fly()

	s2.Print()
}
