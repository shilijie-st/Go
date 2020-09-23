/*
switch:
    基本写法:
    Go语言中switch中的每一个case与case之间都是独立的代码块，不需要通过break语句跳出当前case代码以避免执行到下一行。
    Go语言规定，每个switch只能有一个default分支。
    当多个case要放在一起的时候可以写成如下形式：
        case "apple","bear": 值
    case后不仅仅只是常量，还可以和if一样添加表达式，如： case r>10 && r<20

    fallthrough跨越：
        Go语言中，每一个case都是独立的代码块，执行完毕后不会向C语言那样紧接着下一个case的执行。fallthrough关键字就是来实现这一功能的。
        但是新编写的代码，不建议是用fallthrough这个关键字！
*/
package main

func main() {

}
