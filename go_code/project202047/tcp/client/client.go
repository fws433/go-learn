package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp客户端
func main(){
	//与sercer端建立连接
	conn,err:=net.Dial("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("dial 127.0.0.1.20000 failed,err:",err)
		return
	}
	//发送数据
	reader:=bufio.NewReader(os.Stdin)
	var tmp [128]byte
	//客户端发送多条信息给服务端
	for{
		fmt.Print("方伍胜：")
		msg,_:=reader.ReadString('\n')//读到换行
		msg=strings.TrimSpace(msg)
		if msg=="exit"{
			break
		}
		conn.Write([]byte(msg))
		//客户端读取服务端发送过来的数据
		fmt.Print("黄孝顾：")
		n,err:=conn.Read(tmp[:])
		if err!=nil{
			fmt.Println("rea from conn failed,err:",err)
			return
		}
		fmt.Println(string(tmp[:n]))

	}

	conn.Close()
}
