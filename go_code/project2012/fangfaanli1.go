package main
import "fmt"
//编写结构体
type MethodUtils struct{
  //可以没有字段
}
//1编写Print1()方法，该方法不需要参数，在方法中打印一个10*8的矩形，并在main方法中调用该方法
func (mu MethodUtils)Print1(){
	for i:=1;i<=10;i++{
		for j:=1;j<=8;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
//2编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
func (mu MethodUtils) Print2(m int,n int){
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
//3编写一个方法计算该矩形的面积（可以接收长len,宽width）,将其作为方法返回值，在main方法中调用该方法，接收返回的面积值并打印
func (mu MethodUtils) area(len float64,width float64)(float64){
	return len*width
}
func main(){
	var mu MethodUtils
	mu.Print1()
	fmt.Println()
	mu.Print2(5,20)
	fmt.Println()
	res:=mu.area(2.5,3.5)
	fmt.Println("面积为=",res)
}