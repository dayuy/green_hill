package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

/**
	如何获取一个变量的type？ 可以通过反射
**/

func CheckType(v interface{}) {
	t:=reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Printf("Float")
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknow")
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type(), reflect.TypeOf(f).Kind())
}

/**
	format
**/
type Employee struct {
	EmployeeID string
	Name string `format:"normal"` // struct Tag
	Age int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieId string
	Name string
	Age int
}

func TestInvokeByName(t *testing.T) {
	e:=&Employee{"1", "Mike", 30}
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok{
		t.Error("Failed to get Name field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}

/**
	deepEqual
	: slice 和 map 只能跟nil比较，不能相互比较
	便可以使用 deepEqual(a,b) 进行比较
**/
func TestDeepEqual(t *testing.T) {
	a:=map[int]string{1:"one",2:"two",3:"three"}
	b:=map[int]string{1:"one", 2:"two", 4:"three"}
	//t.Log(a==b) // map can only be compared to nil
	t.Log(reflect.DeepEqual(a, b))

	s1:=[...]int{1,2,3}
	s2:=[...]int{1,2,3}
	s3:=[]int{2,3,1}
	//s4:=[]int{2,3,1}
	//t.Log(s3==s4) // slice can only be compared to nil
	t.Log(s1 ==s2, reflect.TypeOf(s3)) // true []int
	t.Logf("%T", s1) // [3]int
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))
}

// 封装方法
func fillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr { // 必须是指针
		// .Elem() 获取指针指向的值
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct { // 必须是struct
			return errors.New("the first param should be a pointer to the struct type")
		}
	}

	if settings == nil{
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok bool
	)

	for k, v := range settings{
		if field,ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok{
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr:=reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	e:=Employee{}
	if err:= fillBySettings(&e, settings);err!=nil{
		t.Fatal(err)
	}
	t.Log(e)
	c:=new(Customer)
	if err:=fillBySettings(c, settings);err!=nil{
		t.Fatal(err)
	}
	t.Log(*c)
}


/**
	unsafe
	不安全行为
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
**/
type MyInt int
// 应用1 可以使用的别名
func TestCovert(t *testing.T) {
	a:=[]int{1,2,3,4}
	b:=*(*[]MyInt)(unsafe.Pointer(&a)) // 如果是转换别名的类型，ok
	t.Log(b)
}

// 得到错误结果
func TestUnsafe(t *testing.T) {
	i:=10
	f:=*(*float64)(unsafe.Pointer(&i))
	t.Log(f) // 5e-323
	t.Log(unsafe.Pointer(&i)) //0xc00001a2b8
}

// 原子类型操作
func TestAtomic(t *testing.T)  {
	var shareBufPtr unsafe.Pointer
	writeDataFn:= func() {
		data := []int{}
		for i:=0;i<100;i++{
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn:= func() {
		data:=atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFn()
	for i:=0;i<10;i++{
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}
}

