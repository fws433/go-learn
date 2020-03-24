package main
import "fmt"
//存放三名学生，每名学生三个信息，姓名年龄以及家庭地址
func main(){
	studentMap :=make(map[string] map[string]string)
	studentMap["stu01"]=make(map[string]string,3)
	studentMap["stu01"]["name"]="tom"
	studentMap["stu01"]["sex"]="男"
	studentMap["stu01"]["address"]="北京长安街"
	studentMap["stu01"]["app"]="房屋生命"

	studentMap["stu02"]=make(map[string]string)
	studentMap["stu02"]["name"]="mary"
	studentMap["stu02"]["sex"]="女"
	studentMap["stu02"]["address"]="上海黄浦江"
	fmt.Println(studentMap)
	fmt.Println(studentMap["stu01"])
	fmt.Println(studentMap["stu02"]["address"])
}