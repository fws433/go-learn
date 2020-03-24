package main
import "fmt"
//创建一个byte类型的26个个元素的数组，分别放置A-Z
func main(){
	var mychars [26]byte
	for i:=0;i<26;i++{
		mychars[i]='A'+byte(i)  //将int类型的i转换为byte

	}
	for i:=0;i<26;i++{
		fmt.Printf("%c",mychars[i])
	}
}