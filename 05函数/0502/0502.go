/*
函数变量：

*/
package main

import "fmt"

func main() {
	var f func() // 将f声明为func()类型，此时f被俗称为“回调函数”  此时f的值为nil
	f = fire     //将fire()函数名作为值，赋给f变量，此时f的值为fire()函数
	f()          // 使用f变量进行函数调用，实际调用的是fire()函数
}

func fire() {
	fmt.Println("fire")
}
