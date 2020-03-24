package main
import "fmt"

func getSum(n1 int,n2 int)int {
	return n1+n2
}

//函数既然是一种数据；类型，因此函数可以作为形参，并且调用
//funvar名字可以随便取的
func myFun(funvar func(int,int)int,num1 int,num2 int)int{
	return funvar(num1,num2)   //调用形参（函数）
}

func main(){
	a:=getSum  //函数也是一种数据类型，将其赋值给一个变量，则该变量就是一个函数类型的变量
	fmt.Printf("a的类型%T,getsum类型是%T\n",a,getSum)
	res:=a(10,10)//等价res:=getSum(10,10)
	fmt.Println(res)
	res2:=myFun(getSum,50,60)
	fmt.Println("res2=",res2)

}