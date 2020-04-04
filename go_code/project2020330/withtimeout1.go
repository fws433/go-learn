package main

import (
	"context"
	"fmt"
	"time"
)

//withtimeout案例
func main(){
	//设置一个50ms的超时
	ctx,cancel:=context.WithTimeout(context.Background(),time.Millisecond*50)
	//wg.Add(1)
	go worker3(ctx)
	time.Sleep(time.Second*5)
	cancel()//通知子goroutine结束
	//wg.Wait()
	fmt.Println("over")
}

func worker3(ctx context.Context) {
	LOOP:
		for{
			fmt.Println("do connection...")
			//假设正常连接数据库需要耗时间10ms
			time.Sleep(time.Millisecond*10)
			select{
			case<-ctx.Done(): //50ms后自动调用
				break LOOP
			default:
			}
			fmt.Println("worker3 done!")
		//	wg.Done()

		}
}
