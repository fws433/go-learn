package main
import "fmt"
func main(){
	//测试
	test()
	fmt.Println("main()下面的代码...")
}
func test(){
	defer func(){
		err:=recover()   //recover()内置函数，可以捕获到异常
		if err!=nil{//说明捕获到错误
			fmt.Println("err=",err)
		}
	}()
	num1:=10
	num2:=0
	res:=num1/num2
	fmt.Println("res=",res)
}