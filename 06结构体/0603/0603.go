/*
方法
	Go语言中的方法是一种作用与特定类型变量的函数，这种特定类型变量叫做--接收器(Receiver)。接收器的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
	接收器--方法作用的目标。其格式如下：
		func (接收器变量 接收器类型) 方法名(参数列表)(返回参数列表){
			函数体
		}
		接收器根据接收器类型可以分为指针接收器和非指针接收器。两种接收器在使用时会产生不同效果。
		1. 指针类型接收器
			指针类型的接收器由一个结构体的指针组成，更接近于面向对象中的this或者self。由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的。
		2. 非指针类型的接收器
			当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份。在非指针类型接收器的方法中可以获取接收器的成员值，但修改后无效。
		在计算机中，小对象由于值赋值的速度较快，所以适合使用非指针接收器。大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针。

		为类型添加方法
			Go语言可以对任何类型添加方法。
			1. 为基本类型添加方法
*/
package main

//Bag 定义Bag结构体
type Bag struct {
	items []int
}

//Insert 定义Insert方法，其中方法名前面的(b *Bag)就称之为接收器
func (b *Bag) Insert(itemid int) {
	b.items = append(b.items, itemid)
}
func main() {

}
