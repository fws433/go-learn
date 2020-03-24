package main
import(
	"fmt"
)
func main(){
	//演示下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan=make(chan int,3)
	//2.看看intChan是什么
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n",intChan,&intChan)
	//3.向管道写入数据
	intChan<-10
	num:=211
	intChan<-num
	intChan<-50

	//4看看管道的长度和cap
	fmt.Printf("channel len=%v cap=%v\n",len(intChan),cap(intChan))
	//5.从管道中读取数据
	var num2 int
	num2=<-intChan
	fmt.Println("num2=",num2)

	num3:=<-intChan
	fmt.Println("num3=",num3)
	num4:=<-intChan
	fmt.Println("num4=",num4)



}