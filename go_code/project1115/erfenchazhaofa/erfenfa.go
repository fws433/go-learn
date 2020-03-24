package main
import "fmt"
func main(){
	arr:=[6]int{1,8,10,89,1000,1234}
	BianryFind(&arr,0,len(arr)-1,-6)
}
func BianryFind(arr *[6]int,leftIndex int,rightIndex int,findVal int){
	if leftIndex>rightIndex{
		fmt.Println("找不到")
		return
	}

	middle:=(leftIndex+rightIndex)/2
	if (*arr)[middle]>findVal {
		BianryFind(arr,leftIndex,middle-1,findVal)
	}else if (*arr)[middle]<findVal {
		BianryFind(arr,middle+1,rightIndex,findVal)
	}else{
	    fmt.Printf("找到了，下标为%v\n",middle)
	}
}