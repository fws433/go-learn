package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//创建一个context
	ctx:=context.Background()
	//通过cancel派生context:ctxWithCancel
	//返回派生ctxWithCancel和取消函数cancelFuntion,只有创建它的函数才能调用取消函数来取消此 context
	ctxWithCancel,cancelFunction:=context.WithCancel(ctx)
	//推迟取消，以便释放所有资源
	defer func(){
		fmt.Println("main Defer:canceling context")
		cancelFunction()
	}()
	go func(){
		sleepRandom("Main",nil)
		cancelFunction()
		fmt.Println("Main sleep complete. cancel context")
	}()
	//Do work
	doWorkContext(ctxWithCancel)
}
/*
派生一个超时context,
   这个context将取消当：
    	main调用取消函数或超时或doworkContext调用它的取消函数
   启动goroutine传入派生上下文执行一些慢处理
	等待goroutine完成或上下文被main goroutine取消，以优先发生者为准

*/
func doWorkContext(ctx context.Context) {
	//带有取消的上下文派生超时上下文，150ms内超时，由此派生的所有上下文将在150ms内返回
	ctxWithTimeout,cancelFuntion:=context.WithTimeout(ctx,time.Duration(150)*time.Millisecond)
	//功能完成后，取消以释放资源
	defer func(){
		fmt.Println("doWorkContext complete")
		cancelFuntion()
	}()
	//制作频道并调用上下文函数
	//在这种情况下也可以使用等待组
	//由于我们不使用在通道上发送的返回值
	ch:=make(chan bool)
	go sleepRandomContext(ctxWithTimeout,ch)
	select{
	//当传入的上下文通知停止工作时选择此情况
	//在本例中，当主调用cancelFunction时会得到通知
		case<-ctx.Done():
			fmt.Println("doWorkContext: Time to return")
	//在取消上下文之前完成处理时选择此情况
		case<-ch:
			fmt.Println("sleepRandomContext return")
	}
}
/*
开启一个 goroutine 去做些缓慢的处理
等待该 goroutine 完成或，
等待 context 被 main goroutine 取消，操时或它自己的取消函数被调用

*/
//使用上下文执行缓慢处理的函数
//注意上下文是第一个参数
func sleepRandomContext(ctx context.Context, ch chan bool) {
	//清理任务
	//这里没有创建上下文
	//因此，不需要取消
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()
	sleeptimeChan:=make(chan int)
	go sleepRandom("sleepRandomContext",sleeptimeChan)
	//如果上下文过期，请使用select语句退出
	select{
	//如果取消上下文，则选择这种情况
	//如果超时doWorkContext过期或
	// doWorkContext调用cancelFunction或主要调用cancelFunction
	//释放由于中止工作而不再需要的资源
	//给所有应该停止工作的goroutine发出信号（使用通道）
	 //通常，您会在频道上发送一些内容，
	//等待goroutines退出然后返回
	 //或者，使用等待组而不是通道进行同步
		case <-ctx.Done():
			fmt.Println("sleepRandomContext: Time to return")
		case sleeptime:=<-sleeptimeChan:
			fmt.Println("sleep for ",sleeptime,"ms")
	}
}
/*
随机时间休眠
此示例使用休眠来模拟随机处理时间，在实际示例中，您可以使用通道来通知此函数，
以开始清理并在通道上等待它，以确认清理已完成。
*/
func sleepRandom(s string, ch chan int) {
	defer func(){
		fmt.Println(s,"sleepRandom complete")
	}()
	//执行缓慢的任务
	//出于说明目的，
	//在此处睡眠以获取随机毫秒
	seed:=time.Now().UnixNano()
	r:=rand.New(rand.NewSource(seed))
	randomNumber:=r.Intn(100)
	sleeptime:=randomNumber+100
	fmt.Println(s,"start sleep",sleeptime,"ms")
	time.Sleep(time.Duration(sleeptime)*time.Millisecond)
	fmt.Println(s,"waking uo,sleep for",sleeptime,"ms")
	if ch!=nil{
		ch<-sleeptime
	}
}
