package main
import "fmt"
func main(){
	var a int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&a)
	if a==10{
		fmt.Println("666")
	}else if a>=80&&a<100{
		fmt.Println("良好")
	}else if a>=60&&a<80{
		fmt.Println("一般")
	}else{
		fmt.Println("999")
	}

}