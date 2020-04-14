package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
//http client模拟客户端浏览器
func main(){
	//向李文周网站发送一个请求，并得到回应resp
	resp,err:=http.Get("https://www.liwenzhou.com/")
	if err!=nil{
		fmt.Println("get failed,err:",err)
		return
	}
	defer resp.Body.Close()
	///从resp中把服务端返回的数据读出来
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("read from resp.Body failed,err:",err)
		return
	}
	fmt.Print(string(body))
}
