/*
接口
	接口本身就是调用方和实现放均需要遵守的一种协议，大家按照统一的方法明明参数类型和数量来协调逻辑处理的过程。
	Go语言的接口设计是非侵入式的，接口编写者无须知道接口被那些类型实现，而接口实现者只需要知道实现的是什么样子的接口，但无须指明哪一个接口。

	接口的声明
		每个接口类型由数个方法组成，其声明代码如下：
		type 接口类型名interface{
			方法名1(参数列表) 返回值列表
			方法名2(参数列表) 返回值列表
			...
		}
		Go语言的接口在命名时，一般会在单词后面添加er，如Writer,Stringer,Cloder
		方法名：当方法名首字母大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包(package)之外的代码访问。

		开发中常见的接口及写法
		type Writer interface{
			Write(p []byte)(n int,err error)
			Go语言的每个接口中的方法数量不会很多，Go希望通过一个接口精准描述它自己的功能，而通过多个接口的嵌入和组合的方式将简单的接口扩展为复杂的接口。
		}
	实现接口的条件
		1. 接口的方法与实现接口的类型方法格式一致（即方法的签名必须一致）
		2. 接口中的所有方法均被实现

		值得一提的是，Go语言接口的实现是隐式的，无须让实现接口的类写出实现了哪些接口，这种设计被称为侵入式设计。

*/
package main

import "fmt"

//DataWriter A interface
type DataWriter interface {
	WriteDate(data interface{}) error
}

//file difine a struct
type file struct {
}

//WriteDate 实现DataWriter中的WriteDate方法
func (d *file) WriteDate(data interface{}) error {
	fmt.Println("Write Data:", data)
	return nil
}
func main() {
	f := new(file)
	var writer DataWriter
	writer = f
	writer.WriteDate("DATA")
}
