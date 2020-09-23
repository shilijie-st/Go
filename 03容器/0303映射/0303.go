/*
映射（map） --建立事物关联的容器
Go语言提供的映射关系容器为map。map使用散列表（hash）实现。[大多数语言中映射关系容器使用两种算法：散列表和平衡树]
散列表可以简单描述为一个数组，数组的每个元素是一个列表。散列表的查找复杂度为O(1),与数组一致。最坏的情况为O(n)，n为元素总数。
平衡树类似于有父子关系的一棵数据树，每个元素在放入树时，都要与一些节点进行比较，它的查找复杂度始终为O(log n)
添加关联到map并访问关联和数据
	map的定义： map[KeyType]valueType  其中KeyType是键类型，valueType是键对应的值类型。
遍历map的键值对
	直接使用for range的方法,格式为：for k,v := range map
	如果只遍历值，则将不需要的键改为匿名变量形式：for _,v := map
	如果只遍历键，则将值改为匿名变量形式，忽略值即可： for k := range map
注意：
	遍历输出元素的顺序与填充顺序无关。不能期望map在遍历时返回某种期望顺序的结果。
使用delete()函数从map中删除键值对
	delete)_函数的格式如下： delete(map,键)
清空map中的所有元素
	Go语言并没有提供清空map的方法，清空map的唯一方法就是重新make一个心得map
能够在并发环境中使用的map--sync.Map
	sync.Map与map不同，他不是以语言原生形态提供，而是在sync包下的特殊结构，其特性如下：
	1. 无需初始化，直接声明即可
	2. sync.Map不能使用map的方式进行取值和设置等操作，而是使用sync.Map的方法进行调用。Store表示存储，Load表示获取，Delete表示删除。
	3. 使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。Range参数中的回调函数的返回值功能是：需要继续迭代遍历时，返回true，终止迭代遍历时，返回false。
注意：sync.Map不能使用make创建
sync.Map没有提供获取map数量的方法，替代方法是获取遍历时自行计算数量。
sync.Map为了保证并发并安全有一些性能损失，因此在非并发情况下，使用map相比使用sync.Map会更好一些。
*/
package main
import (
	"fmt"
	"sync"
)
func main() {
// define()
//forRnage()
syncMap()
}
func syncMap()  {
	var sence sync.Map
	//将键值对保存到sync.Map
	sence.Store("greece",97)
	sence.Store("london",100)
	sence.Store("egypt",200)
	//从sync.Map中根据键取值
	fmt.Println(sence.Load("greece"))
	//根据键删除对应的键值对
	sence.Delete("london")
	//遍历所用sync.Map中的键值对  遍历需要提供一个匿名函数，参数为k,v 类型为 interfa{},每次Range()在遍历一个元素时，都会调用这个匿名函数把结果返回。
	sence.Range(func(k,v interface{}) bool{
		fmt.Println("iterate",k,v)
		return true
	})
}
func forRnage()  {
	m:=map[string]string{
		"A":"1",
		"B":"2",
		"C":"3",
		"D":"4",
	}
	for k,v:= range m {
		fmt.Println(k,v)
	}
}
func define()  {
	sence := make(map[string]int)//map是一个内部实现的类型，使用时，需要手动使用make()函数创建
sence["route"] = 66
fmt.Println(sence["route"])
v:=sence["route2"]//需要明确知道查询中某个键值是否在map中存在，可以使用这种特殊的方法来实现。
fmt.Println(v)
/*另一种声明时填充内容的方式*/
m:=map[string]string{
	"A":"1",
	"B":"2",
	"C":"3",
	"D":"4",
}
fmt.Println(m)
}
