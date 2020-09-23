/*
Go语言支持可变参数特性，函数声明和嗲用是没有固定数量的参数，同时也提供了一套方法进行可变参数的多级传递。其格式如下：
		func 函数名(固定参数列表,v ...T)(返回参数列表){
			函数体
		}
		特性说明：
		1. 可变参数一般被放置在函数列表的末尾，前面是固定参数列表。当没有固定参数时，所有变量九江市可变参数
		2. v 为可变参数变量，类型为[]T，也就是拥有多个T元素的T类型切片，V和T之间由"..."(3个点)组成
		3. T为可变参数的类型，当T为interface{}时，传入的可以是任意类型
	另外，多个可变参数使用"..."就可以在函数之间进行传递，传递的是切片中的元素，而不是可变参数变量本身。
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	demo1()
}

//demo1 遍历可变参数
func demo1() {
	fmt.Println(jsonStrings("pig", "and", "rat"))
}

//jsonStrings 定义一个函数，参数数量为0~n,类型约束为字符串
func jsonStrings(slist ...string) string {
	var b bytes.Buffer
	for _, v := range slist {
		b.WriteString(v)
	}
	return b.String()
}
