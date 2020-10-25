/*
对类型与接口的理解：
	1. 一个类型可以实现多个接口
	2. 多个类型可以实现相同的接口
接口的嵌套组合 --  将多个接口放在一个接口内
	接口与接口嵌套组合形成了新的接口，只要接口的所有方法被实现，则这个接口中的所有嵌套的方法均可以被调用
接口和类型之间的转换
	Go语言中使用接口断言将接口转换成另外一个接口，也可以将接口转换成另外的类型。
	类型断言的基本格式如下：
		t := i.(T)
		i代表接口变量
		T代表转换的目标类型
		t代表转换后的变量
		如果i没有完全实现T接口的方法，这个语句将会触发宕机，为避免宕机，还可以将上面的写法修改为：
		t,ok := i.(T)
		这中写法下，如果发生接口未实现时，将会把ok置为false，t置为T类型的0值。正常实现时，ok是true。这里的ok可以被认为是：
		i接口是否实现T类型的结果。
	接口和其他类型的转换可以在Go语言中自由进行，前提是已经完全实现。
*/
package main

import "fmt"

//Flyer 定义飞行动物接口
type Flyer interface {
	Fly()
}

//Walker 定义行走动物接口
type Walker interface {
	Walk()
}

//Bird 定义鸟类
type Bird struct{}

//Fly 实现飞行动物接口
func (b *Bird) Fly() {
	fmt.Println("bird: fly")
}

//Pig 定义猪
type Pig struct{}

//Walk 实现猪的Walk()
func (p *Pig) Walk() {
	fmt.Println("pig: walk")
}
func main() {
	//创建动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(Bird),
		"pig":  new(Pig),
	}
	//遍历映射
	for name, obj := range animals {
		//判断对象是否为飞行动物
		f, isFlyer := obj.(Flyer)
		//判断对象是否为行走动物
		w, isWalker := obj.(Walker)
		fmt.Printf("name:%s is Flyer: %v is Walker:%v\n", name, isFlyer, isWalker)
		if isFlyer {
			f.Fly()
		}
		if isWalker {
			w.Walk()
		}
	}
}
