package main
import "fmt"
func main(){
	//如果需要修改字符串，可以先将string-[]byte/或者[]rune->修改->重写转成string
	str:="hello@atguigu"
	arr1:=[]byte(str)
	arr1[0]='z'
	str=string(arr1)
	fmt.Println("str=",str)
	fmt.Println()
	arr2:=[]rune(str)
	arr2[0]='北'
	str=string(arr2)
	fmt.Println("str=",str)
}