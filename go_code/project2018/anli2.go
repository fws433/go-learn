package main
import "fmt"
type Visitor struct{
	Name string
    Age int
}
func (visitor *Visitor)showPrice(){
	if visitor.Age>=90||visitor.Age<=8{
		fmt.Println("考虑到安全，就不要玩了")
		return
	}
	if visitor.Age>10{
		fmt.Printf("游客的名字为%v 年龄为%v 收费20元\n",visitor.Name,visitor.Age)
	}else{
		fmt.Printf("游客的名字为%v 年龄为%v 免费\n",visitor.Name,visitor.Age)
	}

}
func main(){
	var visitor1 Visitor
	for{
		fmt.Println("请输入你的名字：")
		fmt.Scanln(&visitor1.Name)
		if visitor1.Name=="n"{
			fmt.Println("退出程序...")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&visitor1.Age)
		visitor1.showPrice()
	}
}