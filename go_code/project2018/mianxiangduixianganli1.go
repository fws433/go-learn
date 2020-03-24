package main
import "fmt"
//1编写一个学生Student结构体
type Student struct{
	name string
	gender string
	age int
	id int
	score float64
}
type Box struct{
	len float64
	width float64
	height float64
}
//声明一个方法获取立方体的体积
func (box *Box)getVolumn()float64{
	return box.len*box.width*box.height
}
//2结构体中声明一个say方法，返回String类型，方法返回信息中包含所有字段值
func(stu *Student)say() string{
	infoStr:=fmt.Sprintf("Student的信息 name=[%v],gender=[%v],age=[%v],id=[%v],score=[%v]",
	  stu.name,stu.gender,stu.age,stu.id,stu.score)
	return infoStr
}
func main(){
	//测试，创建一个Student实例变量
	var stu=Student{
		name:"tom",
		gender:"tom",
		age:18,
		id:1000,
		score:99.98,

	}
	
	var box Box
	box.len=1.1
	box.width=2.0
	box.height=3.0
	volumn:=box.getVolumn()
	fmt.Printf("体积为=%.2f",volumn)
	fmt.Println(stu.say())
}
