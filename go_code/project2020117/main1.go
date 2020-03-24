package main
import "fmt"
//继承
type A struct{
	Name string
	age int
}
func (a *A)Sayok(){
	fmt.Println("A Sayok",a.Name)
}
func (a *A)hello(){
	fmt.Println("A hello",a.Name)
}
type B struct{
	A   //嵌套
}
func  main(){
	var b B
	b.A.Name="tom"
	b.A.age=19
	b.A.Sayok()
	b.A.hello()
}



