/*
类型内嵌和结构体内嵌
	结构体允许其成员字段再声明时没有字段名而只有类型，这种形式的字段被称为类型内嵌或匿名字段(demo1)。
	类型内嵌其实仍然拥有自己的字段名，只是字段名就是其类型本身而已，结构体陶丘字段名称必须唯一，因此一个结构体中同类型的匿名字段只能有一个。
	结构体实例化后，如果匿名的字段类型为结构体，那么可以直接访问匿名结构体里的所有成员，这种方式被称为结构体内嵌(demo2)。

	结构内嵌特性：
		1. 内嵌的结构体可以直接访问其成员变量
			嵌入结构的成员，可以通过外部结构体的实例直接访问，如果嵌套了多层，结构体实例再访问任意一级的嵌入结构体成员时，都只用给出字段名而无须一层一层的访问。
		2. 内嵌结构体的字段名是它的类型名
			一个结构体只能嵌入一个同类型的成员，无须担心结构体重名和错误赋值的情况。
*/
package main

import "fmt"

func main() {
	demo2()
}

//BasicColor 定义基础颜色
type BasicColor struct {
	R, G, B float32
}

//Color 定义颜色结构体
type Color struct {
	BasicColor //类型内嵌的写法
	Alpha      float32
}

func demo2() {
	var c Color
	c.R = 1
	c.G = 1
	c.B = 0
	c.Alpha = 1
	// c.Basic.B = 0
	// c.Basic.G = 1
	// c.Basic.R = 1
	// c.Alpha = 0.3

	fmt.Printf("%+v\n", c)
}

//Data 定义内嵌类型
type Data struct {
	int
	float32
	bool
}

func demo1() {
	ins := &Data{
		int:     10,
		float32: 3.14,
		bool:    false,
	}
	fmt.Println(ins)
}
