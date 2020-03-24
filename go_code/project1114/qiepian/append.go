package main
import "fmt"
func main(){
	//用append内置函数，可以对切片进行动态追加
	var slice3 []int=[]int{100,200,300}
	slice3=append(slice3,400,500,600)
	fmt.Println("slice3",slice3)//100,200,300,400,500,600

	//通过append将切片slice3追加到slice3
	slice3=append(slice3,slice3...)
	fmt.Println("slice3",slice3)

	//切片的拷贝操作，使用copy内置函数完成拷贝
	fmt.Println()
	var slice4 []int=[]int{1,2,3,4,5}
	var slice5=make([]int,10)
	copy(slice5,slice4)
	fmt.Println("slice4=",slice4)//1,2,3,4,5
	fmt.Println("slice5=",slice5)//1,2,3,4,5,0,0,0,0,0
}