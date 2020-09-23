/*
构造函数 -- 结构体和类型的一系列初始化操作的函数封装
	Go语言的类型或是结构体没有构造函数的功能，结构体的初始化过程可以使用函数封装实现。
*/
package main

import "fmt"

//Cat 定义Cat结构体作为"基类"
type Cat struct {
	Color string
	Name  string
}

//BlackCat 定义BlackCat作为子类
type BlackCat struct {
	Cat //嵌入Cat，类似于派生
}

//NewCat "构造基类"
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

//NewBlackCat "构造子类"
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	cat.Name = "花花"
	return cat
}

func main() {
	fmt.Println(NewBlackCat("colorful"))
}
