package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string
var wg4 sync.WaitGroup

//withValue应用案例
func main(){
	//设置一个50ms的超时
	ctx,cancel:=context.WithTimeout(context.Background(),time.Millisecond*50)
	//在系统的入口中设置trace code传递给后续启动的goroutine实现日志聚合
	ctx=context.WithValue(ctx,TraceCode("Trace_code"),"11111111")
	wg4.Add(1)
	go worker4(ctx)
	time.Sleep(time.Second*5)
	cancel()//通知zigoroutine结束
	wg4.Wait()
	fmt.Println("main over")

}

func worker4(ctx context.Context) {
	key:=TraceCode("Trace_code")
	//在子goroutine中获取tracecode
	tracecode,ok:=ctx.Value(key).(string)
	if !ok{
		fmt.Println("invalid trace code")
	}
	LOOP:
		for{
			fmt.Printf("woorker4,trace code:%s\n",tracecode)
			time.Sleep(time.Millisecond*10)
			select{
			case<-ctx.Done()://50ms后自动调用
				break LOOP
			default:
			}
		}
		fmt.Println("worker4 done")
	wg4.Done()
}
