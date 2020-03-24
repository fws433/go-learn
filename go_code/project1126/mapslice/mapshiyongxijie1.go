package main
import "fmt"
func main(){
	//map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后会直接修改原来的map
	map1:=make(map[int]int)
	map1[1]=90
	map1[2]=80
	map1[3]=70
	modify(map1)
	fmt.Println(map1)
}
func modify(map1 map[int]int){
	map1[3]=100
}