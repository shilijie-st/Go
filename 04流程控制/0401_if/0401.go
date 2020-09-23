/*
流程控制： Go语言常用的流程控制有if和for，而switch和goto主要是为了简化代码，降低重复代码而生的结构，属于扩展类的流程控制。
条件语句 if
	if语言的标准写法如下:
		if 表达式{
			分支1
		}else if{
			分支2
		}else{
			分支3
		}
	特殊写法：在if表达式之前添加一个执行语句，再根据变量值进行判断
		if err:=Connect(); err!=nil{
			fmt.Println(err)
			return
		}
		Connect()是一个带有返回值的函数，err：=Connect()是一个语句，执行Connect后，将错误保存到err变量中。
		err ！= nil才是if的判断表达式，当err不为空时，打印错误并返回。
*/
package main

func main() {

}
