//编写简单的带缓冲的生成器
package main
import(
	"fmt"
	"math/rand"
)
func GenerateInta()chan int{
	ch:=make(chan int,10)
	//启动一个Goroutine用于生成随机数，函数返回一个通道用于获取随机数
	go func(){
		for{
			ch<-rand.Int()
		}
	}()
	return ch
}
func main(){
	ch:=GenerateInta()
	for i:=0;i<100;i++{
		fmt.Println(<-ch)
	}
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}