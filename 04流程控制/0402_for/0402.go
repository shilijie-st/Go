/*
for循环：
	基于语句和表达式的基本for循环格式为：
	for 初始语句；条件语句；结束语句{
		循环体代码
	}
	for循环可以通过break、goto、return、panic语句强制退出循环，但此时的结束语句将不会被执行。
for range 键值循环--直接获取对象的索引个数据
	Go语言中可以使用for range遍历数组、切片、字符串、map及通道（channel）。通过for range遍历的返回值有一定的规律：
		1. 数组、切片、字符串返回索引和值(在遍历中key，value分别代表切片得下标以及下标对应得值)
		2. map返回键和值(在遍历中key和value分别代表字符串得索引(base0)和字符串中的每一个字符)
		3. 通道（channel）只返回通道内的值(在遍历中key和value分别代表map的索引值和索引对应的值)[注意对map遍历时，遍历输出的键值是无序的，如果需要有序的键值对输出，需要对结果进行排序]

		遍历管道时，只输出一个值，即管道内的类型对应的数据
		在遍历中，如果不同时需要键值对同时出现，可以采用匿名变量的形式进行遍历，匿名变量的特点如下：
			匿名变量可以理解为一种占位符；
			本身这种变量不会进行空间分配，也不会占用一个变量的名字；
			在for range可以对key使用匿名变量，也可以对value使用匿名变量。
		总结起来就是：
		1. Go语言的for包含初始化语句、条件表达式、结束语句，这三个部分均可缺省
		2. for range支持对数组、切片、字符串、map、通道进行遍历操作
		3. 在需要时，可以使用匿名变量对for range的变量进行选取
*/
package main

import "fmt"

func main() {
	muiltyTable()
}

func muiltyTable() {
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d * %d = %d ", y, x, x*y)
		}
		fmt.Println()
	}
}
