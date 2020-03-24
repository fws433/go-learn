package main
import "fmt"

//支持可变参数，但是可变参数不能放在函数的第一个形参变量位置中
func Sum(n1 int,args... int)int{
	sum:=n1
	//遍历args
	for i:=0;i<len(args);i++{
		sum+=args[i]
	}
	return sum
}
func main(){
	res:=Sum(10,0,-1,90,10,100)
	fmt.Println("res=",res)
}