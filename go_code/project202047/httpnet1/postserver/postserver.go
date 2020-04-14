package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	http.HandleFunc("/post",postHandle)
	err:=http.ListenAndServe(":9091",nil)
	if err!=nil{
		fmt.Printf("http server failed,err:%v\n",err)
		return
	}
}

func postHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	/*r.ParseForm()
	fmt.Println(r.PostForm)//打印form数据
	fmt.Println(r.PostForm.Get("name"),r.PostForm.Get("age"))*/


	//2请求类别是appLication/json时从r.Body读取数据
	b,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("read request.Body failed,err:%v\n",err)
		return
	}
	fmt.Println(string(b))
	answer:=`{"status":"ok"}`
	w.Write([]byte(answer))

	//data:=r.URL.Query()

	answer1:=`{"status":"ok"}`
	w.Write([]byte(answer1))

}

