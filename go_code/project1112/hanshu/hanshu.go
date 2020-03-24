package main
import "fmt"

func test(n1 int){
	n1=n1+1
	fmt.Println("test() n1=",n1)//输出结果=？
}
func main(){
	n1:=10
	//调用test
	test(n1)
	fmt.Println("main() n1=",n1)//输出结果
	

}