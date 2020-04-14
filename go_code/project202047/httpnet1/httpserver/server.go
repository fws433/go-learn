package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//net/http server
func main(){
	http.HandleFunc("/fws",f1)
	err:=http.ListenAndServe("127.0.0.1:9000",nil)
	if err!=nil{
		fmt.Printf("http server failed,err:%v\n",err)
	}
}

func f1(writer http.ResponseWriter, request *http.Request) {
	//str:="方伍胜"
	b,err:=ioutil.ReadFile("./xx.txt")
	if err!=nil{
		writer.Write([]byte(fmt.Sprintf("%v",err)))
	}
	writer.Write([]byte(b))
}
