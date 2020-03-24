package main
import "fmt"
//定义一个学生结构体
type Stu struct{
	Name string
	Age int
	Address string
}
func main(){
	students:=make(map[string]Stu,10)
	//map的value为结构体，包含学生的名字，年龄和地址
	stu1:=Stu{"tom",18,"背景"}
	stu2:=Stu{"mary",28,"上海"}
	students["no1"]=stu1
	students["no2"]=stu2
	fmt.Println(students)
	for k,v:=range students{
		fmt.Printf("学生的编号是%v \n",k)
		fmt.Printf("学生的名字是%v \n",v.Name)
		fmt.Printf("学生的年龄是%v \n",v.Age)
		fmt.Printf("学生的地址是%v \n",v.Address)
		fmt.Println()
	}
}