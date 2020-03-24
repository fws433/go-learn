package main
import "fmt"
//defer目的就是为了在函数执行完毕后，能够及时的释放资源
//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈，当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
func main(){
	res:=sum(10,20)
	fmt.Println("res=",res)  //4
}
func sum(n1 int,n2 int)int{
	defer fmt.Println("ok1 n1=",n1)   //3
	defer fmt.Println("ok2 n2=",n2)  //2
	res:=n1+n2
	fmt.Println("ok3 res=",res)    //1
	return res
}