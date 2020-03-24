package main
import(
	"fmt"
	"go_code/project2020116/model"
)
func main(){
	p:=model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name,"age=",p.GetAge(),"sal=",p.GetSal())

//创建一个account变量
	account:=model.NewAccount("fws","000",40)
	if account!=nil{
		fmt.Println("创建成功=",account)
	}else{
		fmt.Println("创建失败")
	}
}