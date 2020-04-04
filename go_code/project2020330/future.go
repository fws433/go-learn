package main

import (
	"fmt"
	"time"
)

type query struct {
	//参数channel
	sql chan string
	//结果channel
	result chan string
}

func main(){
	//初始化Query
	q:=query{make(chan string,1),make(chan string,1)}
	//执行query,注意执行的时候无须准备参数
	go execQuery(q)
	//发送参数
	q.sql<-"select * from table"
	//做其他事情，通过sleep描述
	time.Sleep(1*time.Second)
	//获取结果
	fmt.Println(<-q.result)

}

func execQuery(q query) {
	//启动携程
	go func(){
		//获取输入
		sql:=<-q.sql
		//访问数据库
		//输出结果通道
		q.result<-"result from "+sql
	}()
}

