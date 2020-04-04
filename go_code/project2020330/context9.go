package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg3 sync.WaitGroup 
//当子goroutine启动另一个goroutine时
func main(){
	ctx,cancel:=context.WithCancel(context.Background())
	wg3.Add(1)
	go worker(ctx)
	time.Sleep(time.Second*10)
	cancel() //通知子goroutine结束
	wg3.Wait()
	fmt.Println("main over")
}

func worker(ctx context.Context) {
	go worker2(ctx)
	LOOP:
		for{
			fmt.Println("worker")
			time.Sleep(time.Second)
			select{
			case<-ctx.Done()://等待上级通知
				break LOOP
			default:
			}
		}
	wg3.Done()
}

func worker2(ctx context.Context) {
	LOOP:
		for{
			fmt.Println("worker2")
			time.Sleep(time.Second)
			select{
			case<-ctx.Done()://等待上级通知
				break LOOP
			default:
			}
		}
}
