package unit_test

import (
	"fmt"
	"testing"
)

func TestErrorInCode(t *testing.T) {
	fmt.Printf("Start")
	t.Error("Error") // 继续执行
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Fatal") // 中止测试
	fmt.Println("End")
}