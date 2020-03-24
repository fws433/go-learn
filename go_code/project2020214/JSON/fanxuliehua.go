package main
import(
	"fmt"
	"encoding/json"
)
//定义一个结构体
type Monster struct{
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}
//演示将json字符串反序列化成struct
func unmarshalStruct(){
	//说明str在项目开发中，是通过网络传输获取到
	str:="{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2020-2-2\",\"Sal\":8000.0,\"Skill\":\"牛魔犬\"}"
	//定义一个Monster实例
	var monster Monster
	err:=json.Unmarshal([]byte(str),&monster)
	if err!=nil{
		fmt.Printf("unmarshal en=%v\n",err)
	}
	fmt.Printf("反序列化后 monster=%v\n monster.Name=%v\n",monster,monster.Name)
}
func main(){
	unmarshalStruct()
}