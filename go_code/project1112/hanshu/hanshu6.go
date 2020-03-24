package main
import "fmt"
var age=test()

//为了看到全局变量是先被初始化的，
func test() int  {
	fmt.Println("1:全局变量先被初始化")//1
	return 90
}
//init函数，通常可以在Init函数中完成初始化工作
func init(){
	fmt.Println("2:init()...") //2
}
func main(){
	fmt.Println("3:main()...age=") //3
}