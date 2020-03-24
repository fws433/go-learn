package main
import "fmt"
type Ainterface interface{
	Say()
}
type Stu struct{
	//Name string
}
type integer int
func (i integer)Say(){
	fmt.Println("integer Say i=",i)
}
func (stu Stu)Say(){
	fmt.Println("Stu Say()")
}
func main(){
	//var stu Stu //结构体变量，实现了Say()
	stu:=Stu{}
	stu.Say()

	var i integer=10
	i.Say()
}