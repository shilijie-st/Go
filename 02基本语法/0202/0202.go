/*
数据类型：
	Go语言的数据类型非常丰富，基本类型有：整型、浮点型、布尔型、字符串，还有切片、结构体
	函数、map通道（channel）等。
	需要注意的是：切片类型具有指针的便利性，但是比指针更为安全。
整型中的int 与uint
	1. 哪些情况下使用int和uint？
		逻辑对整型范围没有特殊要求；
		使用int 还是uint主要是考虑他们在不同平台上的差异。
浮点型：
	float32 和 float64都遵循IEEE754标准；
正弦函数:
	正弦函数由math包提供，入口为math.Sin，它的参数为float64，返回值也是float64.
字符串类型：
	字符串是不支持换行的，但是可以使用"`"这个符号换行，键盘上1左边的按键（英文状态下输入）
		``之间的代码可以不会呗编译器识别，而只是作为字符串的一部分原样输
字符;
	Go中的字符有两种，一种是uint8类型，或者叫做byte型，代表ASCII码的一个字符
	另一种是rune类型，代表一个UTF-8字符。当需要处理中文、日文或者其他复合字符时，则需要用到rune类型，它实际上是一个int32
		（使用fmt。printf中的%T动词就可以输出变量的实际类型，使用这个方法可以查看byte和rune的本来类型）
切片：
	切片是一个拥有相同类型元素的可变长度的序列，它的声明方式如下：
		var name []T
类型转换：
	类型转换的语法： T(表达式)，其中T 表示要转换的类型，表达式包括变量、复杂算子和函数返回值
	类型转换时，需要考虑两点：两种类型的关系以及范围
	当浮点类型的数据转换为整型时，会将小数部分去掉，只保留整数部分
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	a := make([]int, 3)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
}

func main02() {
	fmt.Println("0202数据类型的学习:")
	fmt.Println("输出正弦函数图像")
	//图片大小
	const size = 300
	//根据给定的大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//遍历每个像素
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pic.SetGray(i, j, color.Gray{155})
		}
	}

	//绘制正弦图像
	//从0到最大（300）像素生成x坐标
	for i := 0; i < size; i++ {
		//让sin的值的范围在0 ~ 2Pi之间
		s := float64(i) * 2 * math.Pi / size
		//sin的幅度为一半的像素，向下偏移一半像素并反转。
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(i, int(y), color.Gray{0})
	}

	//创建文件C:\Users\Lijie.Shi\Desktop\Go\02基本语法\0202
	file, err := os.Create("02基本语法/0202/sin.png")
	if err != nil {
		log.Fatal(err)
	}
	//使用png格式将数据写入文件
	png.Encode(file, pic)
	file.Close()
}
