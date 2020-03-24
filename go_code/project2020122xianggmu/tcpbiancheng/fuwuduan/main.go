package main
import(
	"fmt"
	//做网络开发时，net包含我们所需要的所有的方法和函数
	"net"
	_"io"
)
func process(conn net.Conn){
	//这里我们循环的接收客户端发送 的数据
	defer conn.Close() //关闭conn
	for{
		//创建一个新的切片
		buf:=make([]byte,1024)
		//1等待客户端通过conn发送信息
		//2.如果客户端没有write[发送]，那么携程就阻塞在这里
		fmt.Printf("服务器再等待客户端%s发送信息\n",conn.RemoteAddr().String())
		n,err:=conn.Read(buf)//从conn读取
		if err！=nil{
			fmt.Printf("客户端退出err=%v",err)
			return
		}
		//3显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}
func main(){
	fmt.Println("服务器开始监听...")
	//1.tcp表示使用的网络协议是tcp
	//2.0.0.0.0:8888
	listen,err:=net.Listen("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println("listen err=",err)
		return
	}
	defer listen.Close()//延时关闭linsten
	//循环等待客户端来连接我
	for{
		//等待客户端连接
		fmt.Println("等待客户端来连接...")
		conn,err:=linsten.Accept()
		if err!=nil{
			fmt.Println("Accept() err=",err)

		}else{
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n",conn,conn.RemoteAddr().String())

		}
		//这里准备其一个携程，为客户端服务
		go process(conn)

	}
}