package main
import "fmt"
//声明一个接口
type Usb interface{
	Start()
	Stop()
}
type Phone struct{
	name string
}
//让Phone实现Usb接口的方法
func (p Phone)Start(){
	fmt.Println("手机开始工作...")
}
func (p Phone)Stop(){
	fmt.Println("手机停止工作...")
}
func (p Phone)Call(){
	fmt.Println("手机在打电话。。。")
}
type Camera struct{
	name string
}
//让Camera实现Usb接口的方法
func(c Camera)Start(){
	fmt.Println("相机开始工作...")
}
func(c Camera)Stop(){
	fmt.Println("相机停止工作...")
}

type Computer struct{

}
func (computer Computer)Working(usb  Usb){
	usb.Start()
	//如果usb是指向Phone结构体变量，则还需要调用Call方法
	//使用类型断言
	if phone,ok:=usb.(Phone);ok{
		phone.Call()
	}
	usb.Stop()

}

func main(){
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里体现多态数组
	var usbArr [3]Usb
	usbArr[0]=Phone{"vivo"}
	usbArr[1]=Phone{"小米"}
	usbArr[2]=Camera{"尼康"}


//遍历数组
//Phone还有一个特有的方法call(),请遍历Usb,除了遍历Usb接口声明的方法外，还需要调用call方法

     var computer Computer
     for _,v:=range usbArr{
	     computer.Working(v)
	    fmt.Println()
    }


}























