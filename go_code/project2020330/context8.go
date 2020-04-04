package main
import (
	"context"
	"fmt"
	"sync"
	"time"
)

//如何优雅的控制子goroutine退出
//方法一：利用sync.WaitGroup全局变量
//方法二使用channel
//方法三使用context

//方法三
var wg2 sync.WaitGroup
//var exitch=make(chan bool,1)
func main(){
	ctx,cancel:=context.WithCancel(context.Background())
	wg2.Add(1)
	go f3(ctx)
	time.Sleep(time.Second*5)
	cancel()//exitch<-true

}

func f3(ctx context.Context) {
	defer wg2.Done()
LOOP:
	for{
		fmt.Println("方伍胜")
		time.Sleep(time.Millisecond*500)
		select{

		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

