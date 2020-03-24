package main
import(
	"fmt"
)
var sum=func(a,b int)int{
	return a+b
}
func doinput(f func(int,int)int,a,b int)int{
	return f(a,b)
}
func main(){
	//匿名函数直接被调用
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	sum(1,2)
	//匿名函数作为实参
	doinput(func(x,y int)int{
		return x+y
	},1,2)
}