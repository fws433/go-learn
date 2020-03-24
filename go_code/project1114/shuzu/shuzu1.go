package main
import "fmt"
//定义一个数组
func main(){
	var hens [7]float64
	hens[0]=3.0
	hens[1]=4.0
	hens[2]=5.0
	hens[3]=6.0
	hens[4]=7.0
	hens[5]=9.0
	hens[6]=10.0
	total:=0.0
	for i:=0;i<len(hens);i++{
		total +=hens[i]
	}
	//求平均值
	avgWeights:=fmt.Sprintf("%.2f",total/float64(len(hens)))
	fmt.Printf("total=%v,avgWeights=%v",total,avgWeights)

}