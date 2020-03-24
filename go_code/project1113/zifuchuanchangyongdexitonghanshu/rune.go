package main
import "fmt"
//字符串遍历，同时处理右中文的问题 r:=[]rune(str)
func main(){
	str:="hello方伍胜"
	r:=[]rune(str)
	for i:=0;i<len(r);i++{
		fmt.Printf("字符=%c\n",r[i])
	}
}