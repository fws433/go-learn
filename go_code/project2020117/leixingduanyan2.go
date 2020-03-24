package main
import "fmt"
//写一个函数，循环判断传入参数的类型
func TypeJudge(items... interface{}){
	for index,x:=range items{
		switch x.(type){
			 case bool:
				fmt.Printf("第%v个参数bool类型，值是%v\n",index,x)
			 case float32:
				fmt.Printf("第%v个参数float32类型，值是%%v\n",index,x)
			case float64:
				fmt.Printf("第%v个参数float64类型，值是%%v\n",index,x)
			case int,int32,int64:
				fmt.Printf("第%v个参数整数类型，值是%%v\n",index,x)
			case string:
				fmt.Printf("第%v个参数string类型，值是%%v\n",index,x)
			default:
				fmt.Printf("第%v个参数类型不确定\n",index,x)
		}
	}
}
func main(){
	var n1 float32=1.1
	var n2 float64=2.3
	var n3 int32=30
	var name string="tom"
	address:="背景"
	n4:=300
	TypeJudge(n1,n2,n3,name,address,n4)
}