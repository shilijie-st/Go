/*
结构体(Struct)
	结构体是类型中带有成员的复合类型，Go语言使用结构体和结构体成员来描述真实世界的实体和实体对应的各种属性。
	结构体可以被实例化，使用new或'&'构造的类型实例的类型是类型的指针。
	结构体成员是由一系列的成员变量(字段)构成，所谓字段，有如下几个特征：
		1. 字段拥有自己的类型和值
		2. 字段名必须唯一
		3. 字段的类型也可以是结构体，甚至是字段所在结构体的类型。
	提示：
		1. Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
		2. Go语言中的结构体与“类”都是复合结构体，但Go语言中结构体的内嵌配合接口比面向对象具有更高的扩展性和灵活性。
		3. Go语言不仅认为结构体能拥有方法，且每种自定义类型也可以拥有自己的方法。

	定义结构体
		结构体的定义格式如下：
		type 类型名struct{
			字段1 类型2
			字段2 类型2
		}
		类型名：标识自定义结构体的名称，在同一个包内不能重复
		struct{}：表示结构体类型
		字段名：表示结构体的字段名必须唯一
		同类型的变量也可以写在一行，如a,b,c byte
		//GO语言的关键字type可以将各种基本类型定义为自定义类型。

	实例化结构体 -- 为结构体分配内存并初始化
		基本的实例化形式：
			var ins T //T为结构体类型，ins为结构体的实例
		创建指针类型的结构体
			ins := new(T) //T为类型，ins： T类型被实例化后保存到ins变量中，ins的类型为*T，属于指针
		Go语言中可以使用“.”来访问结构体的成员变量，它使用了语法糖技术，将ins.name形式转换为(*ins).name
		取结构体的地址实例化
			Go语言中，对结构体进行“&”取地址操作时，视为对该类型进行一次new的实例化操作，其格式如下：
			ins := &T{}
	初始化结构体的成员变量
		使用键值对初始化结构体：
		1. ins := T{
			key1 : value1,
			key2 : value2,
			...
			}
		2. type People struct{
			name string
			child *People
			}
			telation := &People{
				name:"爷爷"，
				child:&People{
					name:"爸爸"，
					child:&People{
						name:"我",
					}，
				}，
			}
			结构体成员中只能包含结构体的指针类型，包含非指针类型会编译引起编译错误
			使用多个值的列表初始化结构体
*/
package main

import "fmt"

//Command 定义一个结构体
type Command struct {
	Name    string //指令名称
	Var     *int   //指令绑定的变量
	Comment string //指令的注释
}

func main() {
	var version int = 1
	cmd := &Command{}
	cmd.Name = "Version"
	cmd.Var = &version
	cmd.Comment = "show version"
	fmt.Println(cmd)
}
