package main
import (
	"fmt"
	"net/http"
	"sync"

)
var wg sync.WaitGroup
var urls=[]string{
	"http://www.golang.org/",
	"http://www.google.com/",
	"http://www.qq.com/",
}

func main(){
	for_,url:=range urls{
		//每一个url启动一个goroutine,同时给WG加1
		wg.Add(1)
		go func(url string){
			//当前gor outinr结束后给wg技术减1
			defer wg.Done()
			//发送http get请求并打印http返回吗
			resp,err:=http.Get(url)
			if err==nil{
				fmt.Println(resp.Status)
			}
		}(url)
		//等待所有请求结束
		wg.Wait()

	}
}