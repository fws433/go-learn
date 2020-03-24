package main
import "fmt"
//golang中字符串String类型的使用
func main(){
	//反引号，可以输出源码，（键盘最左边波浪号那个键）
	var str string=`
	package main
    import "fmt"
    //演示go中字符类型使用
    func main(){
	 var c1 byte='a'
	 var c2 byte='0'
	 //当我们直接输出byte值，就是输出了对应的字符的码值
	 fmt.Println("c1=",c1)
	 fmt.Println("c2=",c2)
	 //如果我们希望输出对应的字符，需要使用格式化输出
	 fmt.Printf("c1=%c,c2=%c\n",c1,c2)
	 var c3 int='北'
	 fmt.Printf("c3=%c,c3对应的码值=%d",c3,c3)
}
	
	`
	fmt.Println(str)
}