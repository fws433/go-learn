package main
import(
	"fmt"
	"sort"
	) 
func main(){
	//map的排序
	map1:=make(map[int]int,10)
	map1[10]=100
	map1[1]=4
	map1[3]=56
	map1[6]=90
	fmt.Println(map1)
//如果按照map的key的顺序进行排序
//1先将map的key放入到切片中，2对切片排序，3遍历切片，然后按照key来输出map的值
	var keys []int
	for k,_:=range map1{
		keys=append(keys,k)
	}
// 对切片进行排序
   sort.Ints(keys)
   fmt.Println(keys)
   for _,k:=range keys{
	   fmt.Printf("map1[%v]=%v \n",k,map1[k])
   }
}