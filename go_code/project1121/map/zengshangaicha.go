package main
import "fmt"
func main(){
//增或者改
	//如果key还没有，就是增加，如果key存在就是修改
	citis :=make(map[string]string)
	citis["01"]="背景"
	citis["02"]="huanj"
	citis["03"]="ddd"
	fmt.Println(citis)
	//因为这个key已经存在，因此下面的这句话就是修改
	citis["03"]="ffff"
	fmt.Println(citis)
//删除
   delete(citis,"01")
   fmt.Println(citis)
   //当delete指定的key不存在时，删除不会操作，也不会报错
   delete(citis,"no4")
   fmt.Println(citis)
   citis["01"]="ddddddd"
   citis["05"]="aaaa"
//查找
   val,ok :=citis["01"]
   if ok{
	   fmt.Println("有 01 key值为%v\n",val)
   }else{
	   fmt.Println("没有01 key\n")
   }
}