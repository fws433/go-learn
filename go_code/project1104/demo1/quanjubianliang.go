package main
import "fmt"
var n1=100
var n2=200
var name="jack"
//下面的写法与上面等价，在使用全局变量的时候
var(
	n3=300
	n4=800
	name2="marry"
)
func main(){
	fmt.Println("n1=",n1,"n2=",n2,"name=",name)
	fmt.Println("n3=",n3,"n4=",n4,"name2=",name2)
}