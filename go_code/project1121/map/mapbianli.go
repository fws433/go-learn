package main
import "fmt"
func main(){
	citis :=make(map[string]string)
	citis["no1"]="北京"
	citis["no2"]="天津"
	citis["no3"]="上海"
//map遍历方式1  for-range
	for k,v :=range citis{
		fmt.Printf("k=%v v=%v\n",k,v)
	}
	fmt.Println()
	//使用for-range遍历一个结构比较复杂的map
	studentmap :=make(map[string] map[string]string)
	studentmap["stu01"]=make(map[string]string,3)
	studentmap["stu01"]["name"]="tom"
	studentmap["stu01"]["sex"]="男"
	studentmap["stu01"]["address"]="北京长安街"

	studentmap["stu02"]=make(map[string]string,3)
	studentmap["stu02"]["name"]="tom1"
	studentmap["stu02"]["sex"]="男1"
	studentmap["stu02"]["address1"]="北京长安街1"

	for k1,v1 :=range studentmap{
		fmt.Println("k1=",k1)
		for k2,v2 :=range v1{
			fmt.Printf("\t k2=%v v2=%v\n",k2,v2)
		}
		fmt.Println()
	}


}