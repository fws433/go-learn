package main
import "fmt"
func main(){
//map的使用方式一
	//map的声明和注意事项
	var a map[string]string
	//在使用make前，需要先make,make的作用就是给map分配数据空间
	a=make(map[string]string,10)
	a["no1"]="宋江"
	a["no2"]="无用"
	a["no3"]="武松"
	a["no4"]="吴用"
	fmt.Println(a)
//map的使用方式二
	cities:=make(map[string]string)
	cities["no1"]="背景"
	cities["no2"]="添加"
	cities["no3"]="上海"
	fmt.Println(cities)
//map的使用方式三
   heros:=map[string]string{
	   "hero1":"wusheng",
	   "hero2":"wusheng1",
	   "heros3":"wusheng3",
   }
   heros["hero4"]="linchong"
   fmt.Println("herss=",heros)
}