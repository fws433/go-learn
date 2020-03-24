package main
import "fmt"
type Usb interface{
	Say()
}
type Stu struct{}
func (this *Stu)Say(){
	fmt.Println("Say()")
}
func main(){
	var stu Stu=Stu{}

	var u Usb=&stu
	u.Say()
	fmt.Println("here",u)
}