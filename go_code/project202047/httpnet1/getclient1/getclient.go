package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
	apiuri:="http://127.0.0.1:9091/get"

	data:=url.Values{}//url values
	data.Set("name","小王子")
	data.Set("age","28")
	//
	u,err:=url.ParseRequestURI(apiuri)
	if err!=nil{
		fmt.Printf("parse url requestUrl failed,err:%v\n",err)

	}
	//url encode之后的url
	u.RawQuery=data.Encode()
	fmt.Println(u.String())
	resp,err:=http.Get(u.String())
	if err!=nil{
		fmt.Println("post failed,err:%v\n",err)
		return
	}
	defer resp.Body.Close()
	b,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("get resp failed,err:%v\n",err)
		return
	}
	fmt.Println(string(b))
}
