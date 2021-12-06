package main

import (
	"fmt"
	"reflect"
)

/**
	变量的结构：pair<type: , value: >
					这里的type是 staticType | concreteType
**/

func reflectNum(arg interface{}) { // arg: pair<type: , value: >
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

func main1() {
	var num float64 = 1.2345
	reflectNum(num)

	user := user{1, "Aceld", 18}
	DoFiledAndMethod(user)

	var re resume
	findTag(&re)
}

type user struct {
	Id int
	Name string
	Age int
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("------input type = ", inputType.Name(), inputType)

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("value = ", inputValue)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface() // field的interface()方法得到对应的value
		fmt.Printf("aaaa--- %s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type 获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("bbbb--- %s: %v\n", m.Name, m.Type)
	}
}

type resume struct {
	Name string `info:"name" do:"my name"`
	Sex string	`info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("info")
		fmt.Println("info: ", tagstring, t.Field(i), t.Field(i).Tag)
	}
}
