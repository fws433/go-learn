package main
import "fmt"
//请求出一个数组的最大值，并得到对应的下标
//思路：1.假定第一个元素就是最大值，下标为0，2.然后从第二个元素开始循环比较，如果发现有更大，则交换
func main(){
	var intArr [6]int=[...]int{1,-1,8,90,11,100}
	maxVal:=intArr[0]
	lens:=len(intArr)
	maxValIndex:=0
	for i:=1;i<lens;i++{
		if maxVal<intArr[i]{
			maxVal=intArr[i]
			maxValIndex=i
		}
		  
	}
	fmt.Printf("maxVal=%v,maxValIndex=%v",maxVal,maxValIndex)
}