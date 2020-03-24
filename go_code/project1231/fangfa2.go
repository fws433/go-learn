package main
import "fmt"
//声明一个结构体
type Circle struct{
	radius float64
}
//给Circle结构体绑定一个area方法，用来计算圆形面积
func (c Circle) area()float64{
	return 3.14*c.radius*c.radius
}
func (c *Circle) area2() float64{
	c.radius=10.0
	return 3.14*c.radius*c.radius
}
func main(){
	//1声明一个结构体Circle,字段为radius
	//2声明一个方法area与Circle绑定，可以返回面积
	var c Circle
	c.radius=4.0
	res:=c.area()
	fmt.Println("面积是=",res)
	res2:=(&c).area2()
	fmt.Println("面积=",res2)
	fmt.Println("c.radius=",c.radius)

}