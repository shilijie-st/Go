/*
匿名函数--没有函数名字的函数
		在需要使用函数时，再定义函数，匿名函数没有函数名，只有函数体，函数可以被作为一种类型被赋值给函数类型的变量，匿名函数也往往以变量方式被传递。
		匿名函数经常被用于实现回调函数、闭包等。
	在定义时调用匿名函数
		func(data int){
			fmt.Println("hello",data)
		}(100) //"(100)"就表示对匿名函数进行调用，传递的参数为100
	将匿名函数赋值给变量
		f:=func(data int){
			fmt.Println("hello",data)
		}
		f(100)//使用f()调用
	匿名函数用作回调函数
		详见demo1
	使用匿名函数实现操作封装
		详见demo2
*/
package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill to perform") //定义命令行参数skill，从命令行输入-skill可以将空格后的字符串传入skillParam指针变量

func main() {
	//命令行输入 go run 05函数\0503_匿名函数\0503.go --skill=fly
	demo2()
}

func demo2() {
	flag.Parse()                   //解析命令行参数，解析完成后，skillParam指针变量指向命令行传入的值
	var skill = map[string]func(){ //定义一个从字符串映射到func()的map，然后填充这个map
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angle fly")
		},
	}
	if f, ok := skill[*skillParam]; ok { //使用*skillParam获取命令行传过来的值，并在map中查找命令行参数指定的字符串的函数
		f()
	} else {
		fmt.Println("skill not found")
	}
}

func demo1() {
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
