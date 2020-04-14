package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/get",getHandle)
	err:=http.ListenAndServe(":9091",nil)
	if err!=nil{
		fmt.Printf("http server failed,err:%v\n",err)
		return
	}
}

func getHandle(w http.ResponseWriter, r *http.Request) {
	//对于get请求，参数都放在url上（query param）,请求体中是没有数据的
	defer r.Body.Close()
	//url.query()自动帮我们识别url中的query param
	data:=r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer:=`{"status":"ok"}`
	w.Write([]byte(answer))

}
