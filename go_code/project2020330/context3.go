package main
//WithCancel使用案例
import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//吃汉堡比赛，每秒吃0-5个，计算吃到10的用时
func main(){
	ctx,cancel:=context.WithCancel(context.Background())
	eatNum:=chihanbao(ctx)
	for n:=range eatNum{
		if n>=10{
			cancel()
			break
		}
	}
	fmt.Println("正在统计结果...")
	time.Sleep(1*time.Second)
}

func chihanbao(ctx context.Context) <-chan int{
	c:=make(chan int)
	//个数
	n:=0
	//时间
	t:=0
	go func(){
		for{
			select{
			case<-ctx.Done():
				fmt.Printf("耗时%d秒，吃了%d个汉堡\n",t,n)
				return
			case c<-n:
				incr:=rand.Intn(5)
				n+=incr
				if n>=10{
					n=10
				}
				t++
				fmt.Printf("我吃了%d个汉堡\n",n)
			}
		}
	}()
	return c

}
