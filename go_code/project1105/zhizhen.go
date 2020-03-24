package main
import "fmt"
//指针
func main(){
	var i int =10
	fmt.Println(i)
	fmt.Println("i的地址=",&i)

	var ptr *int=&i   //表明ptr是一个指针变量
	fmt.Println("ptr=",ptr)
	fmt.Println("ptr的地址=",&ptr)
	fmt.Println("ptr值指向的值=",*ptr)
}