package main

import (
	"fmt"
	"sync"
	"time"
)

//如何优雅的控制子goroutine退出
//方法一：利用sync.WaitGroup全局变量
//方法二使用channel
//方法三使用context

//方法二
var wg1 sync.WaitGroup
var exitch=make(chan bool,1)
func main(){
	wg1.Add(1)
	go f2()
	time.Sleep(time.Second*5)
	exitch<-true

}

func f2() {
	defer wg1.Done()
LOOP:
	for{
		fmt.Println("方伍胜")
		time.Sleep(time.Millisecond*500)
		select{
		case <- exitch:
			break LOOP
		default:
		}
	}
}
