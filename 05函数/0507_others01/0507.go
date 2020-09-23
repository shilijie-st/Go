/*
延迟执行语句--defer
	Go语言的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer的逆序进行执行，即first defer last run,last defer first run.
	延迟调用实在defer所在的函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。
	使用defer语句可以方便地处理释放资源的问题！

处理运行时发生的错误
	Go语言的错误处理思想包含以下特征：
	1. 一个可能造成错误的函数，需要返回值中返回一个错误接口(error)。如果调用是成功的，错误接口将返回nil，否则返回错误
	2. 在函数调用后需要检查错误，如果发生错误，进行必要的错误处理
	注意：Go语言中没有异常处理机制。

*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	demo2()
}

//demo2 define a error
func demo2() {
	var e error
	e = newParseError("0507.go", 1)
	fmt.Println(e.Error())
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename: %s Line: %d\n", detail.FileName, detail.Line)
	default:
		fmt.Println("other error")
	}
}

var errDivisionByZero = errors.New("division by zero")

//ParseError 声明一个解析错误
type ParseError struct {
	FileName string
	Line     int
}

//Error 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.FileName, e.Line)
}

//newParseError 创建一下解析错误
func newParseError(fileName string, line int) error {
	return &ParseError{fileName, line}
}

//demo1 defer
func demo1() {
	fmt.Println("defer begin")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("defer end")
}
