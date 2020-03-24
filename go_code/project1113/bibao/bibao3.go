package main
import (
	"fmt"
    "strings"
)

//请编写一个函数，可以接收一个文件后缀名（比如.jpg）,并返回一个闭包
//调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（比如.jpg），则返回文件名，如果已经有jpg后缀，则返回原文件名
//strings.Hasuffix，该函数可以判断某个字符串是否有指定的后缀
func main(){
   f:=makeSuffix(".jpg")
   fmt.Println(f("fws"))
   fmt.Println(f("zh.jpg"))
}
func makeSuffix(suffix string)func (string)string{
	return func(name string)string{
		if ! strings.HasSuffix(name,suffix){
			return name+suffix
		}
		return name
	}
}