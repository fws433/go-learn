package main
import "fmt"
//顺序查找
func main(){
  //第一种方式
  names:=[4]string{"金毛狮王","白眉","紫杉","青翼蝠王"}
  var heroname=""
  fmt.Println("请输入要查找的人名...")
  fmt.Scanln(&heroname)
  for i:=0;i<len(names);i++{
	  if heroname==names[i]{
		  fmt.Printf("找到%v,下标%v \n",heroname,i)
	  }else if i==(len(names)-1){
		fmt.Printf("没有找到%v \n",heroname)
	  }
  }
  //第二种方式
  index:=-1
  for i:=0;i<len(names);i++{
	  if heroname==names[i]{
		  index=i
		  break
	  }
  }
  if index!=-1{
	  fmt.Printf("找到%v,下标%v\n",heroname,index)

  }else{
	  fmt.Println("没有找到",heroName)


  }


}