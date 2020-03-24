package main
import "fmt"
func main(){
	//遍历管道for..range
	intChan2:=make(chan int,100)
	for i:=0;i<100;i++{
		intChan2<-i*2 //放入100个数据到管道中

	}
	//遍历管道不能使用普通的for循环
	//在遍历时,如果channel已经光比，则会正常遍历数据
	close(intChan2)
	for v:=range intChan2{
		fmt.Println("v=",v)
	}
}