package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//WithDeadline & WithTimeout
//执行一段代码，控制执行到某个时间的时候，整个程序结束
//吃汉堡比赛，奥特曼每秒吃0-5个，用时10秒，可以吃多少个
func main(){
	//ctx,cancel:=context.WithDeadline(context.Background,time.Now().Add(10))
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	chihanbao1(ctx)
	defer cancel()
}
func chihanbao1(ctx context.Context){
	n:=0
	for{
		select{
		case<-ctx.Done():
			fmt.Println("stop\n")
			return
		default:
			incr:=rand.Intn(5)
			n+=incr
			fmt.Printf("我吃了%d个汉堡\n",n)
		}
		time.Sleep(time.Second)

	}
}

