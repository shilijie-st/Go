/*
闭包(Closure)--引用了外部变量的匿名函数
	闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境，也不会被释放或删除，在闭包中可以继续使用这个自由变量。即：
	函数 + 引用环境 = 闭包
	同一个函数与不同引用环境组合，可以形成不同的示例，见图"闭包与引用函数"
	一个函数类型就像结构体一样，可以被实例化，函数本身不存储任何信息，只有与引用环境结合后形成的闭包才具有"记忆性"。函数是编译期静态的概念，而闭包是运行期动态的概念。
	closure在某些编译语言中也被成为Lambda表达式。
*/
package main

import "fmt"

func main() {
	demo3()
}

//demo3 玩家生成器 -- 利用闭包的记忆效应
//闭包的记忆效应类似于设计模式中工厂模式的生成器
func demo3() {
	generator := playerGen("HIGHER")
	name, hp := generator()
	fmt.Println(name, hp)
}

//playerGen 创建一个玩家生成器，输入名称，输出生成器
func playerGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}

//demo2 闭包的记忆效应
//被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。
func demo2() {
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n", accumulator)
	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)
}

//Accumulate 提供一个值，每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	//返回一个闭包
	return func() int {
		value++
		return value
	}
}

//demo1 在闭包内部修改引用的变量
//闭包对它作用域上部变量的引用可以进行修改，修改应用的变量就会对变量进行实际修改
func demo1() string {
	str := "hello world"
	foo := func() {
		str = "hello dude"
	}
	foo()
	return str
}
