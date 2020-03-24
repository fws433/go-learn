package main
import "fmt"
func BubbleSort(arr *[5]int){
	fmt.Println("排序前arr=",(*arr))
	temp:=0
	for i:=0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-1-i;j++{
			if(*arr)[j]>(*arr)[j+1]{
				temp=(*arr)[j]
				(*arr)[j]=(*arr)[j+1]
				(*arr)[j+1]=temp
			}
		}
	}
}
func main(){
	//定义数组
	arr:=[5]int{24,69,80,57,13}
	//将数组传递给一个函数，完成排序
	BubbleSort(&arr)
	fmt.Println("main arr=",arr)
}