/*
goto:跳转到指定代码标签
	goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
break:跳出指定循环(跳出多层循环)
	break语句可以结束for、switch和select的代码块。
		当break 用于跳出多层循环时，可以为循环添加标签，用break跳转至指定标签的循环之外。
	continue：继续下一次循环
		continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限于for循环内使用。
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	gotoDemo2()
}

func gotoDemo2() {
	var err bool = check()
	if err {
		goto onExit
    }
    
	err = check()
	if err {
		goto onExit
	}
    fmt.Println("no Err")

    return
onExit:
	fmt.Println("find err: ", err)
}

func check() bool {
    var flag bool
    x := rand.Uint64()
    if x%2 ==0{
        flag = true
    }else{
        flag = false
    }
    return flag
}

func gotoDemo1() {
	//var breakAgain bool
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			breakAgain = true
	// 			break //只能跳出一层循环
	// 		}
	// 	}
	// 	if breakAgain {
	// 		break
	// 	}
	// }
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakHere //跳转到标签
			}
		}
	}
	return //手动返回，避免执行进入标签[标签只能被goto使用，但不影响代码执行流程，此处如果不手动返回，在不满足条件时，也会执行标签内的代码]
	//标签
breakHere:
	fmt.Println("goto done")
}
