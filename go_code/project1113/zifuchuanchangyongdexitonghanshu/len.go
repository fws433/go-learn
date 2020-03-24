package main
import(
	"fmt"
)
//golang的编码统一为utf-8（ascii的字符（字母和数字）占一个字节，汉字占3个字节）
func main(){
	str:="hello被"
	fmt.Println("str len=",len(str))
}