package main
import "fmt"
//golang中的方法作用在指定的数据类型上的，因此自定义类型都可以有方法，不仅仅是struct，比如int,float64等都可以有方法
type integer int
type Student struct{
	Name string
	Age int
}
//编写一个print方法，绑定在integer结构体上
func (i integer) print(){
	fmt.Println("i=",i)
}
//编写一个方法，可以在该方法里改变i的值
func (i *integer)change(){
	*i=*i+1
}
//给*student实现方法string()
func (stu *Student) String() string{
	str:=fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
	return str
}
func main(){
	var i integer=10
	i.print()
	i.change()
	fmt.Println("i=",i)
	//定义一个student变量
    stu:=Student{
		Name : "tom",
		Age: 20,
	}
	fmt.Println(&stu)
}