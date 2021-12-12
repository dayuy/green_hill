package unit_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
	性能测试 Benchmark
	执行命令 go test -bench=.
		    go test -bench=. -benchmem
			结果：
			BenchmarkConcatStringByAdd-4             6176084               190 ns/op
			BenchmarkConcatStringByBytesBuffer-4    10774585               107 ns/op

	func BenchmarkConcatStringByAdd(b *testing.B){
		b.ResetTimer()
		for i:=0;i<b.N;i++{
			// 测试代码
		}
		b.StopTimer()
	}
**/


func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T)  {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems{
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B)  {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems{
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
