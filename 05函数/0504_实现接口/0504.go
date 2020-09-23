/*
结构体实现接口，见demo1
函数体实现接口，见demo2
*/
package main

import "fmt"

func main() {
	demo2()
}

//demo2 函数体实现接口
func demo2() {
	//函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体。当类型方法被调用时，还需要调用函数本体。
	//声明接口变量
	var invoker Invoker
	//将匿名函数转为FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function: ", v)
	})
	invoker.Call("hello")
}

//FuncCaller 函数定义为类型
type FuncCaller func(interface{})

//Call 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	f(p)
}

//demo1 结构体实现接口
func demo1() {
	//声明接口变量
	var invoker Invoker
	//实例化结构体
	s := new(Struck) // s:= &Struck
	//将实例化的结构体赋值到接口
	invoker = s
	//使用接口调用实例化结构体的方法Struck.Call
	invoker.Call("hello")
}

//Invoker 调用器接口
type Invoker interface {
	//需要实现一个Call()方法
	Call(interface{}) //这个接口需要实现Call()方法，调用时会传入一个interface{}类型的变量，这种类型的变量表示任意类型的值。
}

//Struck 定义结构体
type Struck struct{}

//Call 实现Invoker的Call
func (s *Struck) Call(p interface{}) {
	fmt.Println("from struck:", p)
}
