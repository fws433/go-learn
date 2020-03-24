//使用select可以解决从管道中取数据的阻塞问题
package main
import (
	"fmt"
	"time"
)
func main(){
	//1定义一个管道10个数据int
	intChan:=make(chan int,10)
	for i:=0;i<10;i++{
		intChan<-i
	}
	//2定义一个管道5个数据string
	stringChan:=make(chan string,5)
	for i:=0;i<5;i++{
		stringChan<-"hello"+fmt.Sprintf("%d",i)

	}
	//传统的方法在遍历管道时，如果不关闭会阻塞
	//问题，，在实际开发中，可能我们不好确定什么时候关闭该管道
	for{
		select{
			//注意：这里，如果intchan一致没有关闭，则不会出现阻塞
			//会自动到下一个case匹配
			case v:=<-intChan:
				fmt.Printf("从intChan读取的数据%d\n",v)
				time.Sleep(time.Second)
			case v:=<-stringChan:
				fmt.Printf("从stringChan读取的数据%s\n",v)
				time.Sleep(time.Second)
			default:
				fmt.Printf("都取不到了\n")
				time.Sleep(time.Second)
				return
		}
		
	}
}