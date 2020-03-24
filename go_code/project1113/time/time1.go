package main
import(
	"fmt"
	"time"
	"strconv"
)
func main(){
	//在执行test03之前，先获取到当前的Unix时间戳
	start:=time.Now().Unix()
	test03()
	end:=time.Now().Unix()
	fmt.Printf("执行test03()耗费时间为%v秒\n",end-start)
}
func test03(){
	str:=""
	for i:=0;i<100000;i++{
		str+="hello"+strconv.Itoa(i)           //Itoa数字转字符串
	}
}