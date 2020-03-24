package main
import "fmt"
type Person struct{
	Name string
	Age int
}
func main(){
	var p1 Person
	p1.Age=10
	p1.Name="fws"
	var p2 Person=p1
	fmt.Println(p2.Age)
	//p2.Name="TOM"
	//fmt.Printf("p2.Name=%v p1.Name=%v",p2.Name,p1.Name)
}