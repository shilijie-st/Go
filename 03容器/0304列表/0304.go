/*
列表[list] 可以快速增删的非连续空间的容器
列表有多种实现方式，如单链表，双链表等。在Go语言中，将列表使用container/list包来实现，内部的实现原理是双链表。它可以高效地进行任意位置的元素插入和删除操作.
	list的初始化有两种方式:New和声明,效果是一致的.
		1. 通过container/list包的New方法初始化list
			变量名 := list.New()
		2. 通过声明初始化list
			var 变量名 list.List
			列表与切片和map不同得是，列表并没有具体元素类型的限制，因此列表的元素可以是任意类型的。
	在列表中插入元素
		双向列表支持从前方或后方插入元素，分别对应的方法是：PushFront和PushBack.
		(两个方法都会返回一个*list.Element结构。如果在以后的使用中需要删除插入的元素，只能通过*list.Element配合Remove()方法进行删除[这种方法让删除更加效率，也是双链表特性之一])
	从列表中删除元素
		列表中插入函数的返回值会提供一个*list.Element结构，这个结构记录着列表元素的值及和其他系欸但之间的关系等信息。从列表中删除元素时，需要用到这个结构进行快速删除。
	遍历列表--访问列表的每一个元素
		遍历双链表需要配合Front()函数获取头元素，遍历时只要元素不为空就可以继续进行。

*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// var li list.List
	// fmt.Println(li)
	l := list.New()
	l.PushBack("canon") // 尾部添加
	l.PushFront("67")   //头部添加
	//element := l.PushBack("fist")   //尾部添加后，保存元素的句柄
	// l.InsertAfter("high", element)  //在element之后添加high
	// l.InsertBefore("noon", element) // ("noon", element) //在element之前添加noon
	// l.Remove(element)
	fmt.Println("遍历list")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
