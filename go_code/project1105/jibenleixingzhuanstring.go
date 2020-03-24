package main
//import "fmt"
import(
	"fmt"
	//"unsafe"
)

//基本数据类型转string
func main(){
	var num1 int=99
	var num2 float64=23.456
	var b bool=true
	var mychar byte='h'
	var str string
	str=fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q\n",str,str)
	str=fmt.Sprintf("%f",num2)
	fmt.Printf("sre type %T str=%q\n",str,str)
	str=fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%q\n",str,str)
	str=fmt.Sprintf("%c",mychar)
	fmt.Printf("str type %T str=%q\n",str,str)
}