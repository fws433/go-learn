package main
import(
	"fmt"
	"os"
)
func main(){
	//打开文件
	//file的叫法：file对象，file指针，file文件句柄
	file,err:=os.Open("d:/test.txt")
	if err!=nil{
		fmt.Println("open file err=",err)
	}
	//输出下文件，看看文件是什么
	fmt.Printf("file=%v",file)
	//关闭文件
	err=file.Close()
	if err!=nil{
		fmt.Println("clsoe file err=",err)
	}
}