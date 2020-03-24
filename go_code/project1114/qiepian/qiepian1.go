package main
import "fmt"
func main(){
	var intArr [5]int=[...]int{1,22,33,44,55}
	//声明和定义一个切片
	ddd:=intArr[1:3]
	fmt.Println("intArr=",intArr)
	fmt.Println("slice的元素是=",ddd)
	fmt.Println("slice的元素个数=",len(ddd))
	fmt.Println("slice的容量=",cap(ddd))
}