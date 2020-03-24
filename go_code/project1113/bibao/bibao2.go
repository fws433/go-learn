package main
import "fmt"
func main(){
   f:=AddUpper()
   fmt.Println(f(1))
   fmt.Println(f(2))
   fmt.Println(f(3))
}
func AddUpper() func(int) int{
	var n int=10
	var str="hello"
	return func(x int)int{
		n=n+x
		str +=string(36)
		fmt.Println("str=",str)
		return n
	}
}