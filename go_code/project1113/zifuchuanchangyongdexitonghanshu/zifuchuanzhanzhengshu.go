package main
import(
	"fmt"
	"strconv"
)
func main(){
	//整数转字符串
	str:=strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\n",str,str)
	//字符串转整数
	n,err:=strconv.Atoi("hello")
	if err!=nil{
		fmt.Println("转换错误",err)
	}else{
		fmt.Println("转换的结果是",n)
	}
	//字符串转[]byte
	var bytes=[]byte("hello go")
	fmt.Printf("str=%v\n",bytes)
	//[]byte转字符串
	str=string([]byte{97,98,99})
	fmt.Printf("str=%v\n",str)
}