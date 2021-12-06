package groutine_test

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
	使用buffer channel实现对象池
**/
type ReusableObj struct {

}

type ObjPool struct {
	bufChan chan *ReusableObj  // 用于缓冲可重用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i:=0;i<numOfObj;i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret:= <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T)  {
	pool:= NewObjPool(10)
	//if err := pool.ReleaseObj(&ReusableObj{}); err!= nil{ // overflow
	//	t.Error(err)
	//}
	for i:=0;i<11;i++{
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err:=pool.ReleaseObj(v);err!=nil{
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}


/**
	sync.Pool
	尝试从私有对象获取，不存在则从当前Processor的共享池获取，也不存在就会取其他processor的共享池获取，如果都没有则使用new产生新的对象返回
	1. 适合于通过复用，降低复杂对象的创建和GC代价
	2、协程安全，会有锁的开销
	3、生命周期受GC影响，不适合于做连接池等，需自己管理生命周期的资源的池化
**/
func TestSyncPool(t *testing.T) {
	pool:=&sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v:=pool.Get().(int) // 断言
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC() // GC 会清除sync.pool中缓存的对象
	v1, _:=pool.Get().(int)
	fmt.Println(v1)
}



