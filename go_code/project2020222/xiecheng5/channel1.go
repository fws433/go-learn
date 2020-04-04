package main

import (
	"fmt"
	"sync"
)

type task struct{
	begin int
	end int
	result chan<-int
}
//执行计算：计算begin到end的和，将执行结果写入到chan result
func (t task) do() {
	sum:=0
	for i:=t.begin;i<=t.end;i++{
		sum+=i
	}
	t.result<-sum
}

//计算100个自然数的和，将其拆分为多个task
func main(){
	//创建多个通道
	taskchan:=make(chan task,10)
	//创建结果通道
	resultchan:=make(chan int,10)
	//wait用于同步等待任务的执行
	wait:=&sync.WaitGroup{}
	//初始化task的goroutine,计算100个自然数之和
	go InitTask(taskchan,resultchan,100)
	//每个task启动一个goroutine进行处理
	go DistributeTask(taskchan,wait,resultchan)
	//通过结果通道获取结果并汇总
	sum:=ProcessResult(resultchan)
	fmt.Println("sum=",sum)


}
//读取通到结果
func ProcessResult(resultchan chan int) int {
	sum:=0
	for r:=range resultchan{
		sum+=r
	}
	return sum
}
//读取taskchan,每个task启动一个worker goroutine进行处理
//并等待每个task运行完，关闭结果通道
func DistributeTask(taskchan <-chan task, wait *sync.WaitGroup, resultchan chan int) {
	for v:=range taskchan{
		wait.Add(1)
		go ProcessTask(v,wait)
	}
	wait.Wait()
	close(resultchan)
}
//处理具体工作，并将结果发送到通道
func ProcessTask(v task, wait *sync.WaitGroup) {
	v.do()
	wait.Done()
}
//构建task并写入task通道
func InitTask(taskchan chan<-task,r chan int,p int){
	qu:=p/10
	mod:=p%10
	high:=qu*10
	for j:=0;j<qu;j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskchan <- tsk
	}
		if mod!=0{
			tsk:=task{
				begin:high+1,
				end:p,
				result:r,
			}
			taskchan<-tsk
		}
		close(taskchan)

}