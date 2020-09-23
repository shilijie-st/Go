/*
Go语言是静态语言，因此变量是由明确类型的.
	声明变量得一般形式是使用var关键字：
		var name type eg: var a,b *int (定义了a,b两个指针变量)
		var是声明变量得关键字，name是变量名，type是变量的类型
	Go语言的基本类型有：
		bool
		string
		int、int8、int16、int32、int64
		uint、uint8、uint16、uint32、uint64、uintptr
		byte // uint8 的别名
		rune // int32 的别名 代表一个 Unicode 码
		float32、float64
		complex64、complex128
		当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等。所有的内存在 Go 中都是经过初始化的。

	批量定义变量时，Go提供了懒人方式：
	var (
		a int
		b bool
		c string
	)
	还有简短格式： 名称:=表达式，如 x := 10
	这种定义方式需要注意以下三点：
		1.定义变量，同时显示初始化
		2.不能提供数据类型
		3.只能用在函数内部
变量的初始化：
	标准格式： var 变量名 类型 = 表达式
	编译器推导类型的格式： var 变量名 = 表达式  （等号右边的部分在编译原理里被称做右值（rvalue））
	短变量声明并初始化： 使用 := , eg： hp := 100
	使用短变量初始化的时候，需要注意以下几点：
		1. 使用:= 而不是 =，因此使用推导声明写法的左值变量必须是没有被定义过的变量，若定义过，则会发生编译错误。
			// 声明 hp 变量
			var hp int
			// 再次声明并赋值
			hp := 10
		2. 编译报错如下时：no new variables on left side of :=，意思是，在“:=”的左边没有新变量出现，意思就是“:=”的左边变量已经被声明了
		3. 在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名可能是重复声明的，编译器也不会报错
			conn, err := net.Dial("tcp", "127.0.0.1:8080")
			conn2, err := net.Dial("tcp", "127.0.0.1:8080")
多个变量同时赋值：
	赋值原则为：多重赋值时，变量的左值和右值按从左到右的顺序赋值
	var a int = 100
	var b int = 200
	b, a = a, b
	这三行代码实现了a,b之间的数据交换。
Go的匿名变量(在Lua等编译语言中，匿名变量也被称为“哑元变量”)：
	匿名变量的特点是一个下画线“_”，“_”本身就是一个特殊的标识符，被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），
	但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。
		匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用
Go语言中变量的作用域:
	与其他语言一致,但这里需要特别说明的是,Go语言中,全局变量的名字可以与局部变量的相同,但代码执行的时候会优先考虑局部变量

*/
package main

import (
	"fmt"
)

// func main() {
// fmt.Println("变量定义练习")
// x := 100
// a, s := 1, "abc"
// fmt.Println("x:", x, " ,a:", a, " ,s:", s)

// fmt.Println("变量初始化")
// var attack = 40
// var defence = 20
// var damageRate float32 = 0.17
// var damage = float32(attack-defence) * damageRate
// fmt.Println(damage)

// fmt.Println("多个变量同时赋值：")
// var a int = 100
// var b int = 200
// b, a = a, b
// fmt.Println(a, b)

func GetData() (int, int) {
	return 100, 200
}

func main() {
	fmt.Println("匿名变量")
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
	var damageRate float32 = 0.17
}
