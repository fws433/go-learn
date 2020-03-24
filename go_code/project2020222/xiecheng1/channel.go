package main
import(
	"fmt"
	"time"
	"sync"
)
//需求：实现要计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中
//最后显示出来，要求使用goroutime完成
//思路：1编写一个函数，来计算各个数的阶乘，并放入到map中
//2.我们启动的携程多个，统计的将结果放入到map中
//3.map应该做出一个全局的
var(
	myMap=make(map[int]int,10)
	//声明一个全局变量的互斥锁
	//lock是一个全局的互斥锁，mutex:是互斥
	lock sync.Mutex
)
//test函数就是计算n!,让这个结果放入到myMap中
func test(n int){
	res:=1
	for i:=1;i<=n;i++{
		res*=i
	}
	//这里我们将res放入到myMap
	//加锁
	lock.Lock()
	myMap[n]=res
	//解锁
	lock.Unlock()
}
func main(){
	//开启多个携程完成这个任务
	for i:=1;i<=200;i++{
		go test(i)
	}
	//休眠10秒钟
	time.Sleep(time.Second*10)
	//这里我们输出结果，变量这个结果,读的时候加锁，主要是因为底层不知道已经解决了
	lock.Lock()
	for i,v:=range myMap{
		fmt.Printf("map[%d]=%d\n",i,v)
	}
	lock.Unlock()
}