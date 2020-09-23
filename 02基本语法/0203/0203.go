/*
指针：
	指针的概念在Go语言中呗拆分为两个核心概念：
		1. 类型指针：允许对这个指针类型的数据进行修改。传递数据使用指针，而无需拷贝数据，
			类型指针不能进行便宜和运算（避免了非法修改关键性数据的问题，同时垃圾回收也比较容易对不会发生偏移指针进行检索和回收）
		2. 切片：有指向起始元素的原始指针、元素数量和容量组成。
	关于指针，我们需要知道以下几个概念：
	指针地址和指针类型：
		变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用“&”操作符放在变量前面对变量进行“取地址”操作，如：
		ptr := &v //v的类型为T  (其中v代表被取地址的变量，被取地址的v使用ptr变量进行接收，ptr的类型就为“*T”，称作T的指针类型，“*”代表指针)
		变量、指针和地址三者的关系是： 每个变量都有地址，指针的值就是地址
	指针取值：
		在对于普通变量使用“&”操作符取地址获得这个变量的指针后，可以对指针使用“*”操作，也就是指针取值。
		指针取值后就是指向变量的值!
	取地址操作符”&“ 和 ”*“是一对互补操作符，& 取出地址， * 根据地址取出地址指向的值
	变量、指针地址、指针变量、取地址、取值的相互关系如下：
		1. 对变量进行取地址（&）操作，可以获得这个变量的指针变量
		2. 指针变量的值就是指针地址
		3. 对指针变量进行取值（*）操作，可以获得指针变量所指向的原变量的值。

	使用指针修改值：
		* 作为右值时。意义就是取指针的值，作为左值时，表示a指向的变量。
		归纳起来就是，* 操作符的根本意义就是操作指针指向的变量，当操作符在右值时，就是取指向变量的值，当操作符在左值时，就是将值设置给指向的变量。
	创建指针的另一种方法： new()函数，其语法为 new (T)
	new()函数可以创建一个对应类型的指针，创建过程会分配内存。被创建的指针指向的值为默认值。
*/
package main

import (
	"flag"
	"fmt"
)

//示例：使用指针变量获取命令行的输入信息
/*
flag.String，定义一个mode变量，这个变量的类型是*string
参数介绍：
参数名称：在给应用输入参数时，使用这个名称
参数值的默认值：与flag所使用的函数创建变量类型对应，String对应字符串、Int对应整型、Bool对应布尔型等
参数说明：使用-help时，会出现在说明中
在下面使用 go run  C:\Users\Lijie.Shi\Desktop\Go\02????\0203\0203.go --mode=fast
*/
var mode = flag.String("mode", "", "process mode")

func main() {
	flag.Parse()
	/*
		由于使用之前，flag.String已经注册了一个mode的命令行参数，flag底层知道怎么解析命令行，并且将值赋给mode *string指针。
		在Parse调用完毕后，无须从flag获取值，而是通过自己注册的mode这个指针，获取到最终的值。
	*/
	fmt.Println(*mode)
}

func func2() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
	a, b := 3, 4
	fmt.Println(&a, &b)
	swap2(&a, &b)
	fmt.Println(a, b)
	fmt.Println(&a, &b)
}
func swap2(a, b *int) {
	a, b = b, a
}

func swap(a, b *int) {
	/*
		需要格外注意的是：
		将a指向的变量赋值给b指向的变量，将b指向的变量赋值给a指向的变量
		（左值）*a的意思是获取a指向的变量
	*/
	*a, *b = *b, *a
}

func func1() {
	fmt.Println("指针地址与指针类型")
	var str string = "banana"
	var cat int = 1
	fmt.Printf("%p %p", &cat, &str)

	fmt.Println("指针取值")
	// 准备一个字符串
	var house string = "Malibu Point 10880,90265"
	//对字符串取地址，ptr类型为 *string
	ptr := &house
	//打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	//打印ptr的指针地址
	fmt.Printf("ptr address: %p\n", ptr)
	//对指针进行取值操作
	value := *ptr
	//取值后的类型
	fmt.Printf("value type: %T\n", value)
	//指针取值后就是指向变量的值
	fmt.Printf("value is: %s\n", value)
}
