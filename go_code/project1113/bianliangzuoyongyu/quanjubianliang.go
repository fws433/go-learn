package main
import "fmt"
var n int=100
var name string="100"
//在全局变量中不能这样定义：Name:="tom",因为它等价于var name string,name="tom",而赋值语句不能再函数体外
func main(){
	fmt.Println("n=",n)
	fmt.Println("name=",name)
}