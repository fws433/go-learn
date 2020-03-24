package main
import (
	"fmt"
)

func test(n1 int){
   n1=n1+1
   fmt.Println("test() n1=",n1)
}

func getSum(n1 int,n2 int)int{
	sum:=n1+n2
	fmt.Println("getSum=",sum)
	return sum   //当函数有return语句时，就是将结果返回给调用者，也就是谁调用我，我就返回给谁，返回到main函数中sun变量中
}

//如果函数return中返回值有多个，则返回值类型就写多个并用大括号扩起来
func getSumandSub(n1 int,n2 int)(int,int){
	sum:=n1+n2
	sub:=n1-n2
	return sum,sub
}
func main(){
	n1:=10
	test(n1)
	sum:=getSum(10,20)
	fmt.Println("main sum=",sum)
	res1,res2:=getSumandSub(1,2)
	fmt.Printf("res1=%v,res2=%v\n",res1,res2)

}