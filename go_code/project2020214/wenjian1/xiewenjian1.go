package main
import(
	"fmt"
	"bufio"
	"os"
	"io"
)
func main(){
	//创建一个新文件，写入5句内容为hello，groden
	//打开文件d:/GoCode/abc.txt
	filePath:="d:/GoCode/abc.txt"
	//1写入5句内容为hello，wolrd
	  //file,err:=os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE,0666)
	
	//2将原来的内容覆盖为10句你好，方伍胜
	//file,err:=os.OpenFile(filePath,os.O_WRONLY | os.O_TRUNC,0666)

	//3,4在原先后面追加内容
	file,err:=os.OpenFile(filePath,os.O_WRONLY | os.O_APPEND,0666)
	  if err!=nil{
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//4先取出原来文件的内容，并显示到终端
	   reader:=bufio.NewReader(file)
	   for{
		   str,err:=reader.ReadString('\n')
		   if err==io.EOF{//如果读取到文件的末尾
			  break
		   }
		   //将读到的内容显示到终端
		   fmt.Print(str)
	   }

	//1准备写入5句""hello weold"
	   //str:="hello world\n"//表示换行

	   //2准备写10句你好方伍胜
	   //str:="你好，方伍胜！\r\n"//\r\n表示换行

	   //3在后面追加5句，hello,linlin
	  // str:="hello linlin\r\n"

	  //4准备写入5ju,nihao
	  str:="nihao\r\n"

	//写入时，使用带缓存的*Writer
	Writer:=bufio.NewWriter(file)
	//3追加的5行
	for i:=0;i<5;i++{
		Writer.WriteString(str)

	}
	//writer是带缓存，因此在调用WriterString方法是，内容先写到缓存的，所以需要调用Flush方法
	Writer.Flush()
}