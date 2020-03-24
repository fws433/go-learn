package main
import "fmt"
type MethodUtils struct{
	//可以没有字段
  }
  //1.编写一个方法，判断是奇数还是偶数
func(mu *MethodUtils) judgeNum(num int){
	if num%2==0{
		fmt.Println(num,"是偶数...")
	}else{
		fmt.Println(num,"是奇数...")
	}
}
//2编写一个方法，根据行，列，字符，打印对应行数和列数的字符
func (mu *MethodUtils)Print3(n int,m int,key string){
	for i:=1;i<=n;i++{
		for j:=1;j<=m;j++{
			fmt.Print(key)
		}
		fmt.Println()
	}
}
func main(){
	var mu MethodUtils
	mu.judgeNum(11)
	fmt.Println()
	mu.Print3(5,6,"#")
}