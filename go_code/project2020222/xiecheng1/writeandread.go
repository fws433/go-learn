package main
import(
	"fmt"
)
//writeDate
func writeData(intChan chan int){
	for i:=1;i<=50;i++{
		//放入数据
		intChan<-i
		fmt.Println("writeData",i)
	
	}
	close(intChan)//关闭
}
//read DATA
func readData(intChan chan int,exitChan chan bool){
	for{
		v,ok:=<-intChan
		fmt.Printf("readData 读到数据=%v\n",v)
		if !ok{
			break
		}
	}
	exitChan<-true
	close(exitChan)
}
func main(){
	//创建两个管道
	intChan:=make(chan int,50)
	exitChan:=make(chan bool,1)
	go writeData(intChan)
	go readData(intChan,exitChan)

}