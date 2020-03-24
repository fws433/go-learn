package main
import "fmt"
type Preson struct{
	Name string
	Age int
}
func main(){
	//创建结构体变量和访问结构体字段,方式1
	p2 :=Preson{"fangwusheng",25}
	fmt.Println(p2)

	//方式2
	var p1 Preson
	p1.Name="fangwushegn"
	p1.Age=25
	fmt.Println(p1)

	//方式3
	var p3 *Preson=new(Preson)
	(*p3).Name="smith"    //跟下面一行等价
	p3.Name="john"

	(*p3).Age=25
	p3.Age=24
	fmt.Println(p3)

	//方式4
	var p4 *Preson=&Preson{}
	(*p4).Name="fws"
	(*p4).Age=88
	fmt.Println(p4)

}