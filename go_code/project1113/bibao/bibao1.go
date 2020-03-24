package main
import "fmt"
//使用闭包实现一个累加器
func AddUpper() func(int)int{
	var m int=30
	return func(x int)int{
		m=m+x
		return m
	}
}

func AddUpper2(n int,n1 int)int{
	var m int
	m=n+n1
	return m
}
func main(){
	f:=AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	//f1:=AddUpper2(30,4)
	/*res1:=f1(30,1)
	res2:=f1(30,2)
	fmt.Println(res1)
	fmt.Println(res2)*/
	fmt.Println(AddUpper2(30,5))

}