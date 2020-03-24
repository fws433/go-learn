//利用for range进行通道上的范围循环
package main
import(
	"time"
	"fmt"
)
func main(){
	ch1:=make(chan int)
	go sendData(ch1)
	//for循环的for range形式可用于通道的接收值，直到它关闭为止
	for v:=range ch1{
		time.Sleep(1*time.Second)
		fmt.Println("读取数据：",v)
	}
	fmt.Println("main..over...")
}
func sendData(ch1 chan int){
	for i:=0;i<10;i++{
		ch1<-i
	}
	close(ch1)//通知对方，通道关闭
}