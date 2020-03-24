package main
import "fmt"
func main(){
	//初始化方式1
	var num1 [3]int=[3]int{1,2,3}
	fmt.Println("num1=",num1)
	//初始化方式2
	var num2=[3]int{4,5,6}
	fmt.Println("num2=",num2)
	//初始化方式3
	var num3=[...]int{7,8,9,6}
	fmt.Println("num3=",num3)
	//初始化方式4
	var num4=[...]int{1:100,0:900,2:999}
	fmt.Println("num4=",num4)
	//类型推导
	num5:=[...]string{1:"tom",2:"mary",3:"jack"}
	fmt.Println("num5=",num5)
}