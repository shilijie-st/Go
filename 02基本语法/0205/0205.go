/*
字符串应用:
	len()函数: 可以用来获取切片/字符串/通道(channel)等的长度, len(param).其返回值类型为int,表示字符串的ASCII字符个数或字节长度.
	(注意,Go语言中的字符串是以UTF-8格式保存的,每个中文占3个字节,
		如果需要按照习惯上来获取字符串长度,则需要使用Go语言中UTF-8包提供的RuneCountInString()函数)
	string.Index()
	string.LastIndex()

常量:
	关键字为const, 除此之外与变量的定义方式相同.
枚举: --一组常量值
	Go语言目前阶段没有枚举,可以使用常量配合iota模拟枚举,iota起始值为0,一般情况下也是建议枚举从0开始,让每个枚举类型都有一个空值.
	const(
		Arrow Weapon = iota //开始产生枚举值,默认为0
		Shuriken
		Sniper Rifle
		Rifle
		Blower
	)
	一个const声明内部的每一行常量声明,将会自动套用前面的iota格式,并自动增加.
	iota还可以生成一些高级的枚举常量,如:
	const(
		Flag None = 1 << iota // 每次将上一次的值左移一位
		Flag Red
		Flag Green
		Flag Blue
	)

类型别名:
	类型别名规定:TypeAlias只是Type的别名,本质上他们是同一个类型
	type alisa = sourceType

	非本地类型不能定义方法
*/
package main

import "fmt"

// NewInt 将NewInt定义为int类型,NewInt会形成一种新的类型,它本身依然具有int的特性
type NewInt int

// IntAlias 将int取一个别名叫做IntAlias,使用IntAlias和int是等效的
type IntAlias = int

func main() {
	var a NewInt
	fmt.Printf("a type : %T\n", a)
	var b IntAlias
	fmt.Printf("b type : %T\n", b)
}
