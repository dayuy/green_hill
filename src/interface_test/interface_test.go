package interface_test

import "testing"

// 用 接口 定义依赖关系
type Programmer interface {
	WriteHelloWorld() string
}

// 用 struct 定义一个实例
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "Hello World"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer) // 返回指针，类似 &GoProgrammer{}
	t.Log(p.WriteHelloWorld())
}
