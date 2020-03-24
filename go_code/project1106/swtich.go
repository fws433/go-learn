package main
import "fmt"
func main(){
	//a:=10
	//fmt.Println(a)
	var score int
	fmt.Println("输入一个成绩：")
	fmt.Scanln(&score)
	//如果switch后面没有表达式或者后面直接声明一个或定义一个变量，那么类似if-else,反之case后面表达式类型要与之一致
	//switch score:=89; {
	switch{
		case score==100:
			fmt.Println("非常好")
		case score>=80:
			fmt.Println("良好")
		case score>=60:
			fmt.Println("一般")
		case score>=0:
			fmt.Println("不及格")
	    default:
		    fmt.Println("输入有错误")
	}
}