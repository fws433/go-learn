package main
import "fmt"
type Person struct{
	Name string
}
//给Person类型绑定一个方法
func (person Person)test(){
	fmt.Println("test() name=",person.Name)
}
//给Person 结构体添加speak方法，输出XXX是一个好人
func (person Person)speak(){
	fmt.Println(person.Name,"是一个好人")
}
//给Person结构体添加jisuan方法，可以计算1+..+1000的结果，说明方法体内可以和函数一样，进行各种运算
func (person Person)jisuan(){
	res :=0
	for i:=0;i<1000;i++{
		res +=i
	}
	fmt.Println(person.Name,"计算的结果是=",res)
}
//给Person结构体jisuan2方法，该方法可以接收一个数n,计算1+..+n的结果
func(person Person)jisuan2(n int){
	res:=0
	for i:=1;i<=n;i++{
		res+=i
	}
	fmt.Println(person.Name,"计算的结果=",res)
}

//给Person 结构体添加getSum方法，可以计算两个数的和，并返回结果
func (person Person)getsum(n1 int,n2 int)int{
	return n1+n2
}

func main(){
	var p Person
	p.Name="fws"
	p.test()//调用方法
	p.speak()
	p.jisuan()
	p.jisuan2(10)
	//var res int
	res:=p.getsum(10,20)
	fmt.Println("res=",res)
}