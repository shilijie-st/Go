/*
切片(slice)：动态分配大小的连续空间
	Go语言得内部结构包含地址、大小和容量。
	1. 从数组或切片生成新的切片
		切片默认指向一段连续内存区域，可以是数组也可以是切片本身。从连续内存区域生成切片的格式如下：
		slice[开始位置:结束位置] (slice表示目标切片对象)
			取出元素的数量为：结束位置-开始位置
			取出元素不包含结束位置对应的索引，切片最后一个元素使用slice[len(slice)]获取
			当缺省结束位置时，表示从开始位置到整个连续区域末尾
			两者同时缺省时，与切片本身等效
			两者同时为0时，等效于空切片，一般用于切片复位
			根据索引获取元素时，应当注意取值范围(0~len(slice)-1)
	2. 表示原有切片：生成切片的格式中，起始和结尾的范围都被忽略，那么切片实际上就表示它本身。
	3. 重置切片，清空拥有的元素： 把切片的开始和结束位置都设为0，生成的切片将变空。
		使用 make()函数构造切片，其格式如下：
			make([]T,size,cap)	其中，T表示切片的元素类型；size就是为这个类型分配多少个元素；cap是与分配的元素数量，这个值设定后不影响size，只是能提前分配空间，降低多次分配空间造成的性能影响。
		注意： 使用make()函数生成的切片一定发生了内存分配操作。但是给定开始与结束位置(包括切片复位)的切片只是将新的切片指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作。
		append()函数为切片动态添加元素：当空间不能容纳足够多的元素时，切片就会进行"扩容",而扩容往往发生在append()函数调用时。
		注意：容量的扩展规律按容量的2倍数扩充。
		使用copy()函数，实现复制切片元素到另一个切片：其格式如下：
			copy(destSlice,srcSlice []T)int  其中，srcSlice为数据来源切片；destSlice为复制目标。
		注意：目标切片必须分配过空间且足够承载复制的元素个数。来源和目标的类型一致，copy的返回值表示实际发生复制的元素个数。
		从切片中删除元素：
			Go语言需要使用切片本身的特性来删除元素。Go语言中切片元素的删除过程并没哟提供任何的语法糖或者方法封装。
			Go语言中切片删除元素的本质是：以被删除元素为分界点，将前后两个部分的内存重新连接起来。
*/
package main

import "fmt"

func main() {
	// sliceDemo()
	// appendDemo()
	// copyDemo()
	deleteDemo()
}

func deleteDemo() int {
	seq := []string{"a", "b", "c", "d", "e", "f"}
	index := 2
	fmt.Println(seq[:index], seq[index+1:])
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
	return 0
}
func sliceDemo() int {
	fmt.Println("slice demo:")
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	fmt.Println("range: ", highRiseBuilding[10:15])
	fmt.Println("middle to end: ", highRiseBuilding[20:])
	fmt.Println("begin to middle: ", highRiseBuilding[:20])
	return 0
}

func appendDemo() int {
	fmt.Println("append demo:")
	var car []string
	car = append(car, "Old Driver")            //添加一个元素
	car = append(car, "Ice", "Sniper", "Monk") //添加多个元素
	team := []string{"Pig", "FlyingCake", "Chicken"}
	car = append(car, team...) //添加一个切片  ...表示将team整个添加到car的后面
	fmt.Println(car)
	return 0
}

func copyDemo() int {
	const eleCount = 1000
	srcData := make([]int, eleCount)
	for i := 0; i < eleCount; i++ {
		srcData[i] = i
	}
	refData := srcData //切片会进行复制
	copyData := make([]int, eleCount)
	copy(copyData, srcData)
	srcData[0] = 999
	fmt.Println(refData[0])
	fmt.Println(copyData[0], copyData[eleCount-1])
	fmt.Println(copy(copyData, srcData[4:6]))
	for i := 0; i < 5; i++ {
		fmt.Printf("%d", copyData[i])
	}
	return 0
}
