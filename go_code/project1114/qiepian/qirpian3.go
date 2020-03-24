package main
import "fmt"
//切片的使用3，直接指定具体数组，使用原理类似make的方式
func main(){
	var strSlice []string=[]string{"tong","fang","may"}
	fmt.Println("strSlice=",strSlice)
	fmt.Println("strSlice的大小=",len(strSlice))
	fmt.Println("strSlice的容量cap=",cap(strSlice))

}