package main
import (
	"fmt"
	"go_code/project2020112/model"
)
func main(){
	//创建要给Student实例
	var stu=model.Student{
		Name:"tom",
		Score:75.8,
	}
	fmt.Println(stu)
	fmt.Println()
	//结构体首字母小写，使用工厂模式来解决
	var stu1=model.NewStudent("fws",88.8)
	fmt.Println(*stu1)
	fmt.Println("name=",stu1.Name,"score=",stu1.Score)
	fmt.Println()
	var stu2=model.NewStudent1("zh",98.8)
	fmt.Println(*stu2)
	fmt.Println("name=",stu2.Name,"score=",stu2.GetScore())
}
