package main
import(
	"fmt"
	"time"
)
//向intChan放入1-8000个数
func putNum(intChan chan int){
	for i:=1;i<=8000;i++{
		intChan<-i
	}
	//关闭intChan
	close(intChan)
}
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool){
	//使用for循环
	var flag bool
	for{
		time.Sleep(time.Millisecond*10)
		num,ok:=<-intChan
		if !ok{   //intChan取不到
			break    
		}
		flag=true//假设是素数
		//判断num是不是素数
		for i:=2;i<num;i++{
			if num % i==0{
				flag=false
				break
			}
		}
		if flag{
			//将这个数就放入到primeChan中
			primeChan<-num
		}
	}
	//向exitChan写入true
	exitChan<-true
}
func main(){
	intChan:=make(chan int,1000)//存入数据的管道
	primeChan:=make(chan int,2000)//放入结果的管道
	//标识退出的管道
	exitChan:=make(chan bool,4)
	//开启一个携程，向intChan放入1-8000个数
	go putNum(intChan)
	//开启4个携程，从intChan取出数据，并判断是否为素数
	//放入到PrimeChan
	for i:=0;i<4;i++{
		go primeNum(intChan,primeChan,exitChan)
	}
	//我们这里主线程，进行处理
	go func(){
		for i:=0;i<4;i++{
			<-exitChan
		}
		close(primeChan)
	}()
	//遍历我们的primeChan,把结果取出
	for{
		res,ok:=<-primeChan
		if !ok{
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n",res)

	}
	fmt.Println("main线程退出")

}