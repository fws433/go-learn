package main

import (
	"context"
	"fmt"
)

//withcancel应用案例
func main(){
	ctx,cancel:=context.WithCancel(context.Background())
	defer cancel()//当我们取完需要的整数后调用cancel
	for n:=range gen(ctx){
		fmt.Println(n)
		if n==5{
			break
		}
	}
}

func gen(ctx context.Context)<-chan int {
	dst:=make(chan int)
	n:=1
	go func(){
		for{
			select{
				case<-ctx.Done():
					return //return 结束该goroutine,防止泄漏
				case dst<-n:
					n++
			}
		}
	}()
	return dst
}
