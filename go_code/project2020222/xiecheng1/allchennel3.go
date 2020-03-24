package main
import(
	"fmt"
)
type Cat struct{
	Name string
	Age int
}
func main(){
	//定义一个存放任意数据类型的管道 3个数据
	var allChan chan interface{}
	allChan=make(chan interface{},10)
	allChan<-10
	allChan<-"tom jack"
	cat:=Cat{"小花猫",4}
	allChan<-cat
	//我们希望获得到管道中的第三个元素，则先将前两个弹出
	<-allChan
	<-allChan
	newCat:=<-allChan  //从管道中取出的Cat死什么
	fmt.Printf("newCat=%T,newCat=%v\n",newCat,newCat)
	a:=newCat.(Cat)  //使用类型断言
	fmt.Printf("newCat.Name=%v",a.Name)

	
}