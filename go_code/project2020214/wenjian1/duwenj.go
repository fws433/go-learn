package main
import(
	"fmt"
	"os"
	"bufio"
	"io"
)
func main(){
	//打开文件
	file,err:=os.Open("d:/test.txt")
	if err!=nil{
		fmt.Println("open file err=",err)
	}
	//当函数退出时，要及时关闭file,否则会有内存泄漏
	defer file.Close()
	//创建一个*Reader，是带缓冲的
	const(
		defaultBufSize=4096 //默认的缓冲区4096

	)
	reader:=bufio.NewReader(file)
	//循环的读取文件的内容
	for{
		str,err:=reader.ReadString('\n')//读到一个换行就结束
		if err=io.EOF{ //表示文件的末尾
			break
		}
		//输出内容
	   fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
	

}