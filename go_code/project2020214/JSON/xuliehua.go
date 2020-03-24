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

func testStruct(){
	//演示
	monster:=Monster{
		Name:"牛魔王",
		Age:500,
		Birthday:"2020-2-2",
		Sal:8000.0,
		Skill:"牛魔犬",
	}
	//将monster序列化
	data,err:=json.Marshal(&monster)
	if err!=nil{
		fmt.Printf("序列号错误 err=%v\n",err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n",string(data))

}
//将map进行序列化
func  testMap(){
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a=make(map[string]interface{})
	a["name"]="红孩儿"
	a["age"]=30
	a["address"]="洪崖洞"
	//将a这个map进行序列化
	//将monster序列化
	data,err:=json.Marshal(a)
	if err!=nil{
		fmt.Printf("序列化错误 err=%v\n",err)

	}
	//输出序列化的结果
	fmt.Printf("a map 序列化=%v\n",string(data))

}
//演示对切片进行序列化，我们这个切片[]map[string]interface{}
func testSlice(){
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前，需要先make,返回的是引用
	m1=make(map[string]interface{})
	m1["name"]="jack"
	m1["age"]="7"
	m1["address"]="背景"
	slice=append(slice,m1)

	var m2 map[string]interface{}
	m2=make(map[string]interface{})
	m2["name"]="jack1"
	m2["age"]="7"
	m2["address"]=[2]string{"墨西哥","夏威夷"}
	slice=append(slice,m2)
	//将切片进行序列化的结果
	data,err:=json.Marshal(slice)
	if err!=nil{
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	fmt.Printf("slice序列化后=%v\n",string(data))
}
func main(){
	testStruct()
	testMap()
	testSlice()
}

