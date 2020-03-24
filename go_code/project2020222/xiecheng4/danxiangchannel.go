package main
import(
	"fmt"
)
func main(){
	ch1:=make(chan int)//双向，读和写
	//ch2:=make(chan<-int)//单向，只写不能读
	//ch3:=make(<-chan int)//单向，只读不能写
	//ch1<-100
	//data:=<-ch1
	go fun1(ch1)
	data:=<-ch1
	fmt.Println("fun1中写出的数据是：",data)
	go fun2(ch1)
	ch1<-200
	fmt.Println("main...over...")
}
//该函数只接收只写的通道
func fun1(ch chan<-int){
	//函数内部，对于ch只能写数据，不能都数据
	ch<-100
	fmt.Println("fun1函数结束。。。")
}
func fun2(ch <-chan int){
	//该函数内部，对于ch只能读数据，不能写数据
	data:=<-ch
	fmt.Println("fun2函数，从ch中读取的数据是：",data)
}