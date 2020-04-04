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

//方法一
var wg sync.WaitGroup
var notify bool
func main(){
	wg.Add(1)
	go f()
	time.Sleep(time.Second*5)
	//通知子goroutine退出
	notify=true
	wg.Wait()
}

func f() {
	defer wg.Done()
	for{
		fmt.Println("方伍胜")
		time.Sleep(time.Millisecond*500)
		if notify{
			break
		}
	}
}


