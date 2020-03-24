package main
import (
	"runtime"
	"fmt"
)
func main(){
	//获取当前的Gomaxproces
	//println("GOMAXPROCS=",runtime.GOMAXPROCS(0))
	c:=make(chan struct{})
	go func(i chan struct{}){
		sum:=0
		for i:=0;i<10000;i++{
			sum+=i
		}
		fmt.Println(sum)
		c<-struct{}{}
	}(c)
	fmt.Println("NUmgOTOUTIME=",runtime.NumGoroutine())
	//读通道c,通过通道进行同步等待
	<-c
}