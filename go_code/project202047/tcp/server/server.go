package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"net"
)

//tcp server端，例如淘宝网
func main(){
	//1本地端口启动服务
	linstener,err:=net.Listen("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("start tcp sercer on 127.0.0.1:20000 failed,err:",err)
		return
	}
	//2等待别人来跟我建立连接,许多客户端连接，使用for
	for{
		conn,err:=linstener.Accept()
		if err!=nil{
			fmt.Println("accept failed,err:",err)
			return
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()//关闭连接
	//3与客户端通信
	var tmp [128]byte
	reader:=bufio.NewReader(os.Stdin)
	//客户端可以发多条信息给服务端，即服务端可以接收客户端的多条信息，使用for
	for{
		fmt.Print("方伍胜：")
		n,err:=conn.Read(tmp[:])
		if err!=nil{
			fmt.Println("rea from conn failed,err:",err)
			return
		}
		fmt.Println(string(tmp[:n]))

		//服务端回应该客户端
		fmt.Print("黄孝顾：")
		msg,_:=reader.ReadString('\n')//读到换行
		//将最前面和最后面的空格去掉
		msg=strings.TrimSpace(msg)
		if msg=="exit"{
			break
		}
		conn.Write([]byte(msg))
	}

}
