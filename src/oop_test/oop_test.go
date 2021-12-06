package oop_test

import (
	"fmt"
	"testing"
	"unsafe"
)

// 用 结构 struct 定义一个实例
type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)              // 指针 &{2 Rose 22}
	t.Log(e2.Age)          // 22 指针的访问和变量的一样
	t.Logf("e is %T", e)   // %T 表示type oop_test.Employee
	t.Logf("e2 is %T", e2) // *oop_test.Employee
}

// 实例对应的方法被调用时，实例成员会进行值复制
func (e Employee) String() string {
	fmt.Printf("Address is=== %x\n", unsafe.Pointer(&e.Name)) // c0000524f0
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is=== %x\n", unsafe.Pointer(&e.Name)) // c000052520
	t.Log(e.String())
}

// 通常情况下为了避免内存拷贝 我们使用第二种定义方式，定义在类型的指针上
func (e *Employee) String1() string {
	fmt.Printf("Address is---- %x", unsafe.Pointer(&e.Name)) // c000052580
	return fmt.Sprintf("ID:%s--Name:%s--Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations1(t *testing.T) {
	e := &Employee{"0", "Bob", 20}                             // 指向实例的指针
	fmt.Printf("Address is---- %x\n", unsafe.Pointer(&e.Name)) // c000052580 指针 address相同
	t.Log(e.String1())
}

/**
面向对象的扩展 **继承**
**/
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("  ", host)
}

type Dog struct {
	p *Pet // 直接调用赋值一个引用
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}
func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("Chao")
}

// 继承自一个 Pet
type Dog1 struct {
	Pet
}

func TestDog1(t *testing.T) {
	// var dog2 Pet = new(Dog1) // cannot use new(Dog1) (type *Dog1) as type Pet in assignment
	dog1 := new(Dog1)
	dog1.SpeakTo("hhhhh")
}

/**
  多态
  接口的最佳实现
  1. 最小接口定义：很多接口中只包含一个方法
  2. 较大接口定义：可以由多个小接口定义组合而成
  3. 使用时 只依赖于必要功能的最小接口,也可以是方法的复用更强：func StoreData(reader Reader) err{}
**/
type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "Hello"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "World"
}

func writeFirstProgram(p Programmer) { // interface 参数必须是指针
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}

/**
不一样的类型，一样的多肽
**/
func DoSomething(p interface{}) { // 空的interface
	// if i, ok := p.(int); ok { // 断言
	// 	fmt.Println("Integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }
	// fmt.Println("unkonw")

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
}
