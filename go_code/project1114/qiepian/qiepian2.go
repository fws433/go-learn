package main
import "fmt"
func main(){
	//切片的使用
	var slice []float64=make([]float64,5,10)
	slice[1]=10
	slice[3]=30
	fmt.Println(slice)
	fmt.Println("slice的size=",len(slice))
	fmt.Println("slice的cap=",cap(slice))
}