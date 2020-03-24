package main
import "fmt"
func main(){
	var name string
	var age byte
	var sal float32
	var isPass bool

//方式1 Scanln()
	//fmt.Println("请输入姓名：")
	//fmt.Scanln(&name)
	/*fmt.Println("请输入年纪：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)FANG
	fmt.Println("请输入是否通过考试:")
	fmt.Scanln(&isPass)*/
	//fmt.Printf("名字是 %v \n年龄是 %v \n薪水是 %v \n是否通过考试 %v \n",name,age,sal,isPass)
	//fmt.Printf("名字是%v",name)
//方式2 Scanf()
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)
	fmt.Printf("名字是 %v \n年龄是 %v \n薪水是 %v \n是否通过考试 %v \n",name,age,sal,isPass)
}