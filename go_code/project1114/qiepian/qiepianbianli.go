package main
import "fmt"
func main(){
	//使用常规的for循环遍历切片
	var arr [5]int=[...]int{10,20,30,40,50}
	slice:=arr[1:4]   //20,30,40
	for i:=0;i<len(slice);i++{
		fmt.Printf("slice[%v]=%v",i,slice[i])
	}
	fmt.Println()
	//使用for--range方式遍历切片
	for i,v:=range slice{
		fmt.Printf("i=%v v=%v\n",i,v)
	}
	//切片可以继续切片
	slice2:=slice[1:2]
	slice2[0]=100//由于arr,slice和slice2指向的数据空间是同一个，所以
	fmt.Println("slice2=",slice2)
	fmt.Println("slice=",slice)
	fmt.Println("arr=",arr)
}