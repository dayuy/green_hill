package interface_test

import "testing"

// 用 接口 定义依赖关系
type Programmer interface {
	WriteHelloWorld() string
}

// 用 struct 定义一个实例
type GoProgrammer struct {
}

// 自定义方法别名（签名）
type IntConv func(a int) int

// 签名要实现接口里的方法
func (g *GoProgrammer) WriteHelloWorld() string {
	return "Hello World"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer) // 返回指针，类似 &GoProgrammer{}
	t.Log(p.WriteHelloWorld())
}
