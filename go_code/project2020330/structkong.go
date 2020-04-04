package main

import (

	"fmt"
	"time"
)

func main(){
	strchan:=make(chan string,3)
	synChan1:=make(chan struct{},1)//接收同步变量
	synChan2:=make(chan struct{},2)
	//主线程开启了两个goroutine线程,等这两个线程结束后主线程才能结束

	//用于演示接收操作
	go func(){
		<-synChan1//表示可以接收数据，否则等待
		fmt.Println("[receiver] Receiver  a ync signal and  wait 1 second")
		time.Sleep(time.Second)
		for{
			if elem,ok:=<-strchan;ok{
				fmt.Println("[receiver]Receiver:",elem)

			}else{
				break
			}
		}
		fmt.Println("[receiver]Stopped")
		synChan2<-struct{}{}
	}()

	//用于演示发送操作
	go func(){
		for i,elm:=range []string{"a","b","c","d","e","f"}{
			fmt.Println("[sender]Sent：",elm)
			strchan<-elm
			if(i+1)%3==0{
				synChan1<-struct{}{}
				fmt.Println("[sender]Sent a sync singnal 1 second")
				time.Sleep(time.Second)
			}
		}
		fmt.Println("[sender]wait 2 seconds。。")
		time.Sleep(time.Second)
		close(strchan)
		synChan2<-struct{}{}

	}()
	//主线程等待发送线程和接收线程结束后再结束
	fmt.Println("[main]waiting...")
	<-synChan2
	<-synChan2
	fmt.Println("[main]stopped")

}
