//将一个文件中的内容写到另一个文件中，这两个文件已经存在
package main
import(
	"fmt"
	"io/ioutil"
)
func main(){
	//将d:/abc.txt文件内容写到 d:/abc1.txt中
	//首先将abc.txt中的内容读取到内存中，将读取到的内容写到abc1.txt
	file1Path:="d:/abc.txt"
	file2Path:="d:/abc1.txt"
	data,err:=ioutil.ReadFile(file1Path)
	if err!=nil{
		//说明读取文件有错误
		fmt.Printf("read file err=%v\n")
		return
	}
	err=ioutil.WriterFile(file2Path,data,0666)
	if err!=nil{
		fmt.Printf("wirte file error=%v\n")
	}
}