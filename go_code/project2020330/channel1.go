package main

import "fmt"

func main(){
	ch:=make(chan int)
	go func(){
		fmt.Println("hello inline")
		ch<-1
	}()
	go printHell(ch)
	fmt.Println("hello from main")
	i:=<-ch
	fmt.Println("recieed",i)
	<-ch
}

func printHell(ch chan int) {
	fmt.Println("hello from printhello")
	ch<-2
}
