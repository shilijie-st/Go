/*
宕机(panic)--程序终止运行
	Go语言中可以在程序中手动触发宕机，让程序崩溃，这样开发者会及时地发现错误，同时减少可能的损失。发生宕机时，会将堆栈和goroutine信息输出到控制台，所以宕机也可以方便地知道发生错误的位置。
	触发宕机的代码如下：
		panic("crash") -- panic()的参数可以是任意类型的，它的定义为： func panic(v interface{})
	当panic()触发的宕机发生时，panic()后面的代码将不会被执行，但是在panic()函数前面已经运行过的defer语句依然会在宕机发生时起作用。
宕机恢复(recover) -- 防止程序崩溃
	无论是代码运行错误由Runtime层抛出的panic还是主动触发的panic，都可以配合defer和recover实现错误捕捉和恢复，让代码在发生崩溃后允许继续运行。

panic 和 recover的关系
	panic 和的肥肉的组合有如下几个特征
		1. 有panic没recover，程序宕机
		2. 有panic也有recover捕获，程序不会宕机。执行完对应的defer后，从宕机点退出当前函数后继续执行。
			在panic触发的defer函数内，可以继续调用panic，进一步将错误外抛，直到程序整体崩溃
			如果想在捕获错误时设置当前函数的返回值，可以对反沪指使用命名返回值方式直接进行设置
*/
package main

import (
	"fmt"
	"runtime"
)

//panicContent 崩溃时需要传递的上下文信息
type panicContent struct {
	function string
}

//ProtectRun 保护方式允许一个函数
func ProtectRun(entry func()) {
	defer func() {
		//发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()
}

func main() {
	fmt.Println("运行前")

	ProtectRun(func() {
		fmt.Println("手动宕机前")
		panic(&panicContent{"手动触发panic"})
		fmt.Println("手动宕机后") //这句代码将不会被执行，所以在此会报警告
	})

	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}
