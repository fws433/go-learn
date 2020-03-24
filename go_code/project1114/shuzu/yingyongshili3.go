package main
import (
	"fmt"
	"time"
	"math/rand"
)
//随机生成五个数，并将其反转打印
func main(){
	var intArr [5]int
	len:=len(intArr)
	//为了每次生成的随机数不一样，我们需要给一个seed值
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len;i++{
		intArr[i]=rand.Intn(100)  //0<=n<100

	}
	temp:=0
	fmt.Println("交换前=",intArr)
	for i:=0;i<len/2;i++{
		temp=intArr[len-1-i]
		intArr[len-1-i]=intArr[i]
		intArr[i]=temp
	}
	fmt.Println("交换后=",intArr)
}