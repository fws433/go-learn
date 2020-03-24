package main
import "fmt"
type Person struct{
	Name string
	Age int
	Scores [5]float64
	ptr *int //指针
	slice []int //切片
	map1 map[string]string  //map
}
func main(){
	//定义结构体变量
	var p1 Person
	fmt.Println(p1)
	if p1.ptr==nil{
		fmt.Println("ok1")
	}
	if p1.slice==nil{
		fmt.Println("ok2")
	}
	if p1.map1==nil{
		fmt.Println("ok3")
	}
	//使用slice,一定要make
	p1.slice=make([]int,10)
	p1.slice[0]=100
	//使用map,一定要先make
	p1.map1=make(map[string]string)
	p1.map1["key1"]="tom~"
	fmt.Println(p1)
}