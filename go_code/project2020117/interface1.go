package main
import "fmt"
//声明一个接口
type Usb interface{
	//声明了两个没有实现的方法
	Start()
	Stop()
}
type Phone struct{

}
//让phone实现Usb接口的方法
func (p Phone)Start(){
	fmt.Println("手机开始工作...")
}
func (p Phone)Stop(){
	fmt.Println("让手机停止工作")
}
type Camera struct{}
//让Camera实现USb接口的方法
func(c Camera)Start(){
	fmt.Println("相机开始工作")
}
func (c Camera)Stop(){
	fmt.Println("相机停止工作")
}
type Computer struct{}
//编写一个方法Working方法，接收一个Usb接口类型变量
func (c Computer)Working(usb Usb){//变量会更具传入的实参来判断到底是Phone还是Camera
	//通过usb接口变量来调用start和stop方法
	usb.Start()
	usb.Stop()

}
func main(){
	//先创建结构体变量
	computer:=Computer{}
	phone:=Phone{}
	camera:=Camera{}
	computer.Working(phone)
	computer.Working(camera)
}


















