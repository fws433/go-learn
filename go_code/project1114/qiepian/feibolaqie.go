//编写一个函数fbn(n int),要求完成
/*
  1.可以接收一个n int
  2.能够将斐波那契的数列放到切片中
  
*/
package main
import "fmt"
func main(){
	fnbslice:=fbn(20)
	fmt.Println("fnbslice=",fnbslice)
}

func fbn(n int)([]uint64){
	//声明一个切片，切片大小为n
	fbslice:=make([]uint64,n)
	fbslice[0]=1
	fbslice[1]=1
	//进行for循环来存放斐波拉契数列
	for i:=2;i<n;i++{
		fbslice[i]=fbslice[i-1]+fbslice[i-2]
	}
	return fbslice
}