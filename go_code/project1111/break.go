package main
import(
	"fmt"
	"math/rand"
	"time"
)
func main(){
	var count int=0
	for{
		//生成一个随机数需要一个rand设置一个种子
		//time.Now().Unix():返回一个从1970:01:01的0时0分0秒到现在的秒数
		rand.Seed(time.Now().UnixNano())
		n:=rand.Intn(100)+1
		fmt.Println("n=",n)
		count++
		if(n==99){
			break
		}
		

	}
	fmt.Println("生成99一共使用了",count)	
}