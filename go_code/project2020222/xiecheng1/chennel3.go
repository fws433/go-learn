package main
import(
	"fmt"
)
func main(){
	intChan:=make(chan int,3)
	intChan<-100
	intChan<-200
	close(intChan)
	fmt.Println("ok")
	n1:=<-intChan
	fmt.Println("n1=",n1)
}