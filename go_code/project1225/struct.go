package main
import "fmt"
type Cat struct{    //定义一个Cat的结构体，将Cat的各个字段/属性信息，放入到Cat结构体进行管理
	Name string
	Age int
	Color string
	Hobby string
}
func main(){
	var Cat1 Cat   //创建一个Cat的变量
	Cat1.Name="小白"
	Cat1.Age=1
	Cat1.Color="白色"
	Cat1.Hobby="吃鱼"
	fmt.Println("Cat1=",Cat1)


}