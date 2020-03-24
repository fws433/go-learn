package main
import "fmt"
type Monster struct{
	Name string
	Age int
}
func main(){
	//不同结构体变量的字段是独立的，互不影响，一个结构体变量字段的更改不影响另外一个，结构体是值类型
	var monster1 Monster
	monster1.Name="方伍胜"
	monster1.Age=24

	monster2 :=monster1  //结构体是值类型，默认是值拷贝
	monster2.Name="伍胜"

	fmt.Println("monster1=",monster1)
	fmt.Println("monster2=",monster2)
}